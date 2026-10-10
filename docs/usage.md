# Usage

This page shows the API of Fulus with short examples.
See [Rounding and overflow](rounding-and-overflow.md) for the exact rules of each operation.

## Arithmetic Operations

Fulus provides basic arithmetic methods with built-in overflow and error checking:

```go
usd10 := fulus.NewMoney[currency.USD](1000)
usd20 := fulus.NewMoney[currency.USD](2000)

// Addition and Subtraction
usd30, err := usd10.Add(usd20)
usd10, err = usd30.Sub(usd20)

// Multiplication
usd20, err = usd10.Mul(2)

// Division with explicit rounding
usd5, err := usd10.Div(2, fulus.RoundHalfUp)

// Multiplication by a Factor, with explicit rounding.
// Parse a Factor one time, for example in a package-level variable.
var salesTax = fulus.MustParseFactor("8.25%")
tip, err := usd10.MulFactor(fulus.Percent(15), fulus.RoundHalfEven) // USD 1.50
tax, err := usd10.MulFactor(salesTax, fulus.RoundHalfUp)            // USD 0.83
fee, err := usd10.MulFactor(fulus.Bps(25), fulus.RoundHalfUp)       // USD 0.03

// Sum, Min and Max
total, err := fulus.Sum(usd10, usd20, usd5)
cheapest := fulus.Min(usd10, usd20, usd5)

// Absolute value and Negation
absUsd, err := fulus.NewMoney[currency.USD](-1000).Abs() // USD 10.00
negUsd, err := usd10.Neg()                               // USD -10.00
```

For ergonomic method chaining when you are confident about bounds (panics on overflow), you can use the `Must` variants:

```go
usd10 := fulus.NewMoney[currency.USD](1000)
usd20 := fulus.NewMoney[currency.USD](2000)
usd30 := fulus.NewMoney[currency.USD](3000)

// Method chaining
result := usd10.MustAdd(usd20).MustSub(usd30).MustMul(2) // USD 0.00
```

## Comparison Operations

Fulus provides standard comparison methods to safely compare money values of the same currency:

```go
usd10 := fulus.NewMoney[currency.USD](1000)
usd20 := fulus.NewMoney[currency.USD](2000)

fmt.Println(usd10.GreaterThan(usd20))      // false
fmt.Println(usd10.LessThan(usd20))         // true
fmt.Println(usd10.Equal(usd20))            // false
fmt.Println(usd10.IsZero())                // false
fmt.Println(usd10.IsPositive())            // true
fmt.Println(usd10.Cmp(usd20))              // -1
fmt.Println(usd10.InRange(usd10, usd20))   // true
```

## Rounding Modes

| Mode | 2.5 | -2.5 | 2.4 |
|------|-----|------|-----|
| `RoundTruncate` | 2 | -2 | 2 |
| `RoundHalfUp` | 3 | -3 | 2 |
| `RoundHalfEven` | 2 | -2 | 2 |
| `RoundHalfDown` | 2 | -2 | 2 |
| `RoundUp` | 3 | -3 | 3 |
| `RoundCeiling` | 3 | -2 | 3 |
| `RoundFloor` | 2 | -3 | 2 |

`RoundUnnecessary` does not round. It returns `ErrInexact` when the exact result needs rounding.
The zero `RoundingMode` is not valid, so an unset mode returns `ErrInvalidRoundingMode`.

Use `Convert` with an explicit rounding mode:

```go
eur := fulus.NewMoney[currency.EUR](5) // €0.05
rate := fulus.MustParseRate[currency.EUR, currency.USD]("1/2")

usdTrunc, _ := fulus.Convert(eur, rate, fulus.RoundTruncate) // USD 0.02
usdHalfUp, _ := fulus.Convert(eur, rate, fulus.RoundHalfUp)  // USD 0.03
usdHalfEven, _ := fulus.Convert(eur, rate, fulus.RoundHalfEven)

fmt.Println(usdTrunc, usdHalfUp, usdHalfEven)
```

## Cash Rounding

Some currencies use a larger step for cash than for accounts, for example 0.05 for CHF.
`RoundCash` uses the CLDR cash rules. Currencies without such a rule do not change.

```go
chf := fulus.NewMoney[currency.CHF](1003)  // CHF 10.03
cash, err := chf.RoundCash(fulus.RoundHalfUp) // CHF 10.05
```

A custom currency can implement `currency.CashRounder` to get cash rounding.

## Runtime Currencies

Use `AnyMoney` when the currency is known only at run time, for example from a database row or a request.
Convert it to a type-safe `Money[T]` with `As` before you calculate with it.

```go
price, err := fulus.ParseAnyMoney("12.50", "eur")
if err != nil {
	panic(err)
}

eur, err := fulus.As[currency.EUR](price) // ErrCurrencyMismatch if price is not in EUR
```

`AnyMoney` also has `Mul`, `Div`, `MulFactor`, `RoundCash`, `Abs`, `Neg` and `Allocate`, with the same rules as `Money[T]`.
`Add`, `Sub` and `Cmp` return `ErrCurrencyMismatch` for two different currencies.
Two currencies are the same only if they have the same code and the same minor units.

`currency.ByCode` and `currency.ByNumber` find a currency by its ISO 4217 code.
`currency.Register` adds a custom currency to the default registry, so that `AnyMoney` can use it.
To keep custom currencies out of the global state, for example for each tenant or in a test, use a `currency.Registry`:

```go
reg, err := currency.NewRegistry(currency.Builtin()...)
err = reg.Register(KANNA{})

c, ok := reg.ByCode("KANNA")
price, err := fulus.NewAnyMoneyFromDecimal("12.50", c)
```

`AnyMoney` uses the same JSON form as `Money[T]`: `{"amount":"12.50","currency":"EUR"}`.
The amount is a decimal string, so it does not lose digits in JavaScript and does not depend on the minor units of the reader.
The currency code is not case-sensitive. A JSON `null` does not change the value. Use `NullMoney[T]` to tell `null` from zero.

## Locales

The `format` package formats and parses amounts with CLDR data. The `fulus` and `currency` packages do not import it,
so a program that does not format amounts does not include the CLDR tables (about 0.3 MB).

`format.Money` uses the CLDR pattern of the locale, including the negative pattern and Indian digit grouping.
`String` does not use a locale. It gives the code and the canonical decimal, for example `USD -1234.50`, so logs and tests do not change with the environment.
`locale.Match` finds the best supported locale for a tag such as `en_US.UTF-8` or `fr-CA-u-nu-latn`.
`TestFormatMatchesICU` compares `format` with ICU, an independent CLDR implementation, for more than 25,000 cases.

```go
loc, ok := locale.Match("en_IN.UTF-8")
inr := fulus.NewMoney[currency.INR](-1234567890)
fmt.Println(format.Money(inr, loc)) // -₹1,23,45,678.90

parsed, err := format.Parse[currency.INR]("-₹1,23,45,678.90", loc)
```

## Parse Decimal Strings

Use `ParseMoney` to parse canonical decimal strings safely:

```go
usd, err := fulus.ParseMoney[currency.USD]("123.45")
if err != nil {
	panic(err)
}
fmt.Println(usd) // USD 123.45
```

`ParseMoney` validates fractional scale against the currency minor units and rejects malformed formats.
`Decimal`, `MarshalText` and `UnmarshalText` use the same canonical form.

## Exchange Rates

`Rate[Base, Quote]` is the price of one major unit of `Base` in major units of `Quote`, as markets quote it.
It is an exact fraction. Its fields are not exported, so a rate is always positive.

```go
eurUSD, err := fulus.ParseRate[currency.EUR, currency.USD]("1.07203")
usdJPY, err := fulus.RateFromFloat64[currency.USD, currency.JPY](151.37) // for example from an FX API

usdEUR := eurUSD.Invert()              // Rate[currency.USD, currency.EUR]
eurJPY, err := fulus.Cross(eurUSD, usdJPY) // Rate[currency.EUR, currency.JPY]
// fulus.Cross(usdJPY, eurUSD) does not compile.

yen, err := fulus.Convert(fulus.NewMoney[currency.EUR](100), eurJPY, fulus.RoundHalfEven) // JPY 162
```

`Convert` adjusts for the minor units of each currency.

## SQL Integration

`Money[T]` implements `driver.Valuer` and `sql.Scanner` for a decimal column such as `NUMERIC(38, 2)`.
`Value` writes the canonical decimal, such as `"10.50"`. `Scan` reads the decimal text that drivers return for `NUMERIC`.
`Scan` returns `ErrScaleMismatch` for a value with a digit that is not zero after the minor units. It does not round.
Zeros after the minor units are accepted, so a `NUMERIC(19, 4)` value such as `"10.5000"` reads as USD 10.50.
The type parameter holds the currency, so the column does not store it.

```go
var m fulus.Money[currency.USD]
if err := m.Scan([]byte("10.50")); err != nil {
	panic(err)
}

v, err := m.Value()
fmt.Println(v, err) // 10.50 <nil>
```

For an integer column such as `BIGINT` that holds minor units, use `BigintMoney[T]`.
Fulus does not guess the column type, because the value 1050 is USD 1050.00 in a `NUMERIC` column and USD 10.50 in minor units.

```go
var row fulus.BigintMoney[currency.USD]
err := row.Scan(int64(1050)) // row.Money is $10.50
```

Use `NullMoney[T]` for a decimal column that can be NULL. It also writes and reads JSON `null`.
Use `sql.Null[fulus.BigintMoney[T]]` for an integer column that can be NULL.

`Scan` returns `ErrInvalidAmountFormat` for an `int64`, because an `int64` can be minor units from a `BIGINT` column.

SQLite gives a `float64` for a `NUMERIC` value with a fraction, and an `int64` for a whole value.
`Scan` accepts a `float64` only if it has at most 15 significant digits, because a `float64` always keeps 15 digits.
In SQLite, use a `TEXT` column or `BigintMoney[T]`.

For an `AnyMoney` in two columns, an amount and a currency code, use `ScanColumns`:

```go
var price fulus.AnyMoney
amount, code := price.ScanColumns()
err := row.Scan(amount, code)

_, err = db.Exec("INSERT INTO prices (amount, currency) VALUES ($1, $2)", price.Decimal(), price.Currency().Code())
```

## Protocol Buffers

The `fulusproto` module converts values to and from `google.type.Money`. It is a separate module,
so the `fulus` module has no dependencies.

```go
import "github.com/khatibomar/fulus/fulusproto"

p, err := fulusproto.ToMoney(fulus.NewMoney[currency.USD](1050)) // units 10, nanos 500000000
m, err := fulusproto.FromMoney[currency.USD](p)                 // USD 10.50
```

A conversion does not round. `ToMoney` returns `ErrInexact` for a value with more than 9 fraction digits that are not zero,
and `FromMoney` returns `ErrScaleMismatch` for a value with a digit that is not zero after the minor units.
