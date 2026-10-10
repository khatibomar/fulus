# Rounding and overflow

This page tells which operations round, how they round, and what happens when a result is too large.

## Amounts

A `Money[T]` value holds a signed 128-bit integer. The integer is the amount in minor units of the currency.
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
To round an input with more digits, multiply one major unit by the input with `MulFactor`:

```go
one := fulus.NewMoney[currency.USD](100)                    // $1.00
f, err := fulus.ParseFactor("19.999")
price, err := one.MulFactor(f, fulus.RoundHalfEven) // $20.00
```

## Operations that round

These operations take a `RoundingMode`. They do not have a default mode.

| Operation | Exact result before rounding |
|-----------|------------------------------|
| `Div(d, mode)` | amount / d |
| `MulFactor(f, mode)` | amount × f, where f is a `Factor` such as `ParseFactor("0.0825")`, `Percent(15)` or `Bps(25)` |
| `Convert(m, rate, mode)` | amount × rate, with the rate in major units of each currency |
| `RoundCash(mode)` | amount / increment, rounded, then multiplied by the increment |

Each operation calculates the exact result first, and then it rounds one time.
The product amount × n uses 128 bits, so it cannot overflow before the division.
`ParseFactor` reads a decimal, a fraction such as `"1/3"` or a percentage such as `"8.25%"` as an exact fraction.
It does not use floating point. It does not accept an exponent such as `"1e3"`.
`ParseFactor` returns `ErrOverflow` if the numerator or the denominator does not fit in `int64`.

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

`Rate[Base, Quote]` holds an exchange rate as an exact positive fraction in lowest terms.
It is the price of one major unit of `Base` in major units of `Quote`, as markets quote it.
The fields are not exported, so a rate is always valid, except the zero `Rate`. `Convert` returns `ErrInvalidExchangeRate` for the zero `Rate`.

- `ParseRate` reads a decimal string such as `"1.07203"` or a fraction such as `"1/3"` exactly.
- `RateFromFloat64` writes the float as the shortest decimal string that gives the same float, and then reads that string.
  So `0.1` gives 1/10, not the binary value of the float.
- `Invert` returns the exact rate in the other direction, as a `Rate[Quote, Base]`.
- `Cross` returns the exact rate from A to C through B. The compiler checks that the middle currencies agree.

`Convert` adjusts for the minor units of each currency. For example, a EUR/JPY rate of 160.25 changes EUR 1.00 to JPY 160.
It calculates the exact result and rounds one time.

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

The amount is a signed 128-bit integer, so the range depends on the minor units of the currency:

| Minor units | Example | Largest amount |
|-------------|---------|----------------|
| 0 | JPY | 170,141,183,460,469,231,731,687,303,715,884,105,727 |
| 2 | USD | 1,701,411,834,604,692,317,316,873,037,158,841,057.27 |
| 3 | BHD | 170,141,183,460,469,231,731,687,303,715,884,105.727 |
| 4 | CLF | 17,014,118,346,046,923,173,168,730,371,588,410.5727 |
| 8 | a custom token | 1,701,411,834,604,692,317,316,873,037,158.84105727 |
| 18 | a custom token | 170,141,183,460,469,231,731.687303715884105727 |

The smallest amount is the negative of the largest amount, minus 1 minor unit.

Every operation checks for overflow. When the result does not fit in 128 bits, the operation returns `ErrOverflow` and the zero value.
It never returns a wrong result.

- `MustAdd`, `MustSub` and `MustMul` panic with `ErrOverflow`. Use them only when you know the range of the values.
- `Int64` reports false and `Value` returns `ErrOverflow` for an amount that does not fit in `int64`.
- `Abs` and `Neg` return `ErrOverflow` for the smallest amount, because its positive value does not fit.
- `RoundCash` returns `ErrOverflow` if the rounded value does not fit.
- `ParseMoney` and `ParseFormatted` return `ErrOverflow` for an amount that does not fit.
- `ParseRate` and `Cross` return `ErrOverflow` if the numerator or the denominator does not fit in `int64`.
- `AnyMoney.Add` and `AnyMoney.Sub` have the same checks as `Money[T]`.

### Sum

`Sum` adds the values in 128 bits, and changes to `math/big` if an intermediate total does not fit.
So an intermediate total can be larger than the range.
`Sum` returns `ErrOverflow` only when the final total does not fit.

```go
largest, _ := fulus.ParseMoney[currency.USD]("1701411834604692317316873037158841057.27")
total, err := fulus.Sum(largest, largest, largest.MustMul(-1)) // total is largest, err is nil
_, err = largest.Add(largest)                                 // err is ErrOverflow
```
