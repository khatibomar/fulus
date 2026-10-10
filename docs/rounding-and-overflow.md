# Rounding and overflow

This page tells which operations round, how they round, and what happens when a result is too large.

## Amounts

A `Money[T]` value holds one `int64`. The `int64` is the amount in minor units of the currency.
For USD, 1 minor unit is 1 cent, so `NewMoney[currency.USD](1050)` is $10.50.
The minor units come from ISO 4217. `T.MinorUnits()` gives the number of digits after the decimal separator.

## Operations that do not round

These operations give an exact result or an error. They never round.

| Operation | Result |
|-----------|--------|
| `Add`, `Sub`, `Mul`, `Neg`, `Abs` | The exact result, or `ErrOverflow`. |
| `Sum` | The exact sum, or `ErrOverflow`. See [Sum](#sum). |
| `Allocate`, `Distribute` | Parts whose sum is equal to the value. See [Allocation](#allocation). |
| `ParseMoney`, `UnmarshalText` | The exact value. More fraction digits than the minor units give `ErrScaleMismatch`. |
| `ParseFormatted` | The exact value. More fraction digits than the minor units give `ErrScaleMismatch`. |
| `UnmarshalJSON`, `Scan` | The exact amount in minor units. |

The parse functions do not round, because a rounded input hides a data error.
To round an input with more digits, multiply one major unit by the input with `MulDecimal`:

```go
one := fulus.NewMoney[currency.USD](100)                    // $1.00
price, err := one.MulDecimal("19.999", fulus.RoundHalfEven) // $20.00
```

## Operations that round

These operations take a `RoundingMode`. They do not have a default mode.

| Operation | Exact result before rounding |
|-----------|------------------------------|
| `Div(d, mode)` | amount / d |
| `MulFrac(n, d, mode)` | amount × n / d |
| `MulDecimal(f, mode)` | amount × f, where f is a decimal such as `"0.0825"` or a fraction such as `"1/3"` |
| `Convert(m, ratio, mode)` | amount × ratio, with the ratio in major units of each currency |
| `RoundCash(mode)` | amount / increment, rounded, then multiplied by the increment |

Each operation calculates the exact result first, and then it rounds one time.
The product amount × n uses 128 bits, so it cannot overflow before the division.
`MulDecimal` reads the factor as an exact fraction. It does not use floating point.
A factor whose numerator or denominator does not fit in `int64` uses `math/big`.

An unknown mode or the zero mode gives `ErrInvalidRoundingMode`. A zero divisor gives `ErrDivisionByZero` or `ErrZeroDenominator`.

## Rounding modes

The table shows the result in minor units for three exact results.

| Mode | 2.5 | -2.5 | 3.5 | Rule |
|------|-----|------|-----|------|
| `RoundTruncate` | 2 | -2 | 3 | Toward zero. |
| `RoundHalfUp` | 3 | -3 | 4 | To the nearest. A tie goes away from zero. |
| `RoundHalfEven` | 2 | -2 | 4 | To the nearest. A tie goes to the even neighbor. |
| `RoundHalfDown` | 2 | -2 | 3 | To the nearest. A tie goes toward zero. |
| `RoundUp` | 3 | -3 | 4 | Away from zero. |
| `RoundCeiling` | 3 | -2 | 4 | Toward positive infinity. |
| `RoundFloor` | 2 | -3 | 3 | Toward negative infinity. |
| `RoundUnnecessary` | `ErrInexact` | `ErrInexact` | `ErrInexact` | No rounding. An exact result does not change. |

The zero `RoundingMode` is not a valid mode. An operation with an unset mode returns `ErrInvalidRoundingMode`.
Use `RoundUnnecessary` when the result must be exact, for example to check that a total divides into equal parts.

A tie is an exact result that is half way between two minor units.
`RoundHalfEven` is also known as banker's rounding. Use it when many rounded values are added, because it has no bias.
`RoundHalfUp` is the rule that most people learn at school.

## Cash rounding

Some currencies use a larger step for cash than for accounts. For example, CHF cash uses steps of 0.05.
`RoundCash` uses the CLDR cash rules. It rounds to the nearest step with the mode.

```go
chf := fulus.NewMoney[currency.CHF](1003)       // CHF 10.03
cash, err := chf.RoundCash(fulus.RoundHalfUp) // CHF 10.05
```

A currency without a cash rule does not change. A custom currency can implement `currency.CashRounder`.

## Exchange rates

`Ratio` holds an exchange rate as an exact fraction.
`ParseRatioString` reads a decimal string exactly. For example, `"1.07203"` gives 107203/100000.
`ParseRatioFloat64` writes the float as the shortest decimal string that gives the same float, and then reads that string.
So `0.1` gives 1/10, not the binary value of the float.

The ratio is the price of one major unit of the source currency in major units of the target currency, as markets quote it.
For example, a EUR/JPY ratio of 160.25 changes EUR 1.00 to JPY 160. `Convert` adjusts for the different minor units.

`Convert` returns a `ConversionResult`. Its `ActualRate` is the rate after rounding: the result divided by the source amount.
For example, €3.33 at 1/3 with `RoundHalfUp` gives $1.11, and the actual rate is 111/333 = 1/3.
For a zero source amount, `ActualRate` is the requested ratio.

## Allocation

`Allocate` divides a value by ratios. The sum of the parts is always equal to the value.
It uses the largest remainder method:

1. Each part gets its share, rounded toward zero.
2. The minor units that are left go one at a time to the parts with the largest remainders.
3. If two remainders are equal, the part with the lower index gets the unit first.

```go
parts, _ := fulus.NewMoney[currency.USD](100).Allocate(1, 1, 1)
// $0.34 $0.33 $0.33
```

A negative value gives negative parts with the same rule, for example -$0.34, -$0.33 and -$0.33.

`Distribute` divides a value into equal chunks. It returns the size and the count of the smaller and larger chunks.
The two sizes are different by 1 minor unit.

## Overflow

The amount is an `int64`, so the range depends on the minor units of the currency:

| Minor units | Example | Largest amount |
|-------------|---------|----------------|
| 0 | JPY | 9,223,372,036,854,775,807 |
| 2 | USD | 92,233,720,368,547,758.07 |
| 3 | BHD | 9,223,372,036,854,775.807 |
| 4 | CLF | 922,337,203,685,477.5807 |
| 8 | a custom token | 92,233,720,368.54775807 |
| 18 | a custom token | 9.223372036854775807 |

The smallest amount is the negative of the largest amount, minus 1 minor unit.

Every operation checks for overflow. When the result does not fit in `int64`, the operation returns `ErrOverflow` and the zero value.
It never returns a wrong result.

- `MustAdd`, `MustSub` and `MustMul` panic with `ErrOverflow`. Use them only when you know the range of the values.
- `Abs` and `Neg` return `ErrOverflow` for the smallest amount, because its positive value does not fit.
- `RoundCash` returns `ErrOverflow` if the rounded value does not fit.
- `ParseMoney` and `ParseFormatted` return `ErrOverflow` for an amount that does not fit.
- `ParseRatioString` returns `ErrOverflow` if the numerator or the denominator does not fit.
- `AnyMoney.Add` and `AnyMoney.Sub` have the same checks as `Money[T]`.

### Sum

`Sum` adds the values with a 128-bit total. An intermediate total can be larger than `int64`.
`Sum` returns `ErrOverflow` only when the final total does not fit.

```go
big := fulus.NewMoney[currency.USD](1 << 62)
total, err := fulus.Sum(big, big, big.MustMul(-1)) // total is 1 << 62, err is nil
_, err = big.Add(big)                              // err is ErrOverflow
```
