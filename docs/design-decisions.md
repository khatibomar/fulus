# Design decisions

This page tells why Fulus works as it does, and what each decision costs.

## The amount is a 128-bit integer in minor units

`Money[T]` holds one signed 128-bit integer, the amount in minor units. For USD, the minor unit is the cent.

Why:

- Integer arithmetic is exact. The value 0.1 + 0.2 is exactly 0.3. With `float64` it is not.
- Operations are fast and do not allocate. `Add` takes about 2 ns. See the README for more benchmarks.
- A `Money[T]` value is 16 bytes. You can compare it with `==` and use it as a map key.
- The range is large enough for currencies and tokens with many minor units.
  A token with 18 minor units holds more than 10^20 units. With `int64`, it holds only about 9.2 units.
  For USD, the largest amount is about 1.7 × 10^36.

Cost:

- A value is 16 bytes, not 8. Multiplication and division are about 2 times slower than with `int64`.
- `Int64` reports false for an amount that does not fit in `int64`. Use `BigInt` or `Decimal` for such an amount.
- An amount cannot have more digits than the minor units. A price per unit such as $0.0042 needs another type,
  or a custom currency with more minor units.

Every operation returns `ErrOverflow` when a result does not fit. Fulus never returns a wrong result.

## The currency is a type parameter

The currency is the type parameter `T` of `Money[T]`. It is not a field.

Why:

- The compiler stops you from mixing currencies. `usd.Add(eur)` does not compile.
- The value does not store the currency, so it stays 16 bytes.
- A function can require a currency in its signature, for example `func Charge(m fulus.Money[currency.EUR])`.
- A conversion must name both currencies: `Rate[currency.EUR, currency.USD]`.
  `Cross(Rate[A, B], Rate[B, C])` gives a `Rate[A, C]`, so the compiler checks a triangulation.
- The constraint `currency.Unit` accepts only empty struct types. So `Money[currency.Currency]` and a currency type
  with fields do not compile. The methods of the currency cannot depend on state.
  `TestTypeSafety` type-checks these rules.

Cost:

- The currency must be known at compile time. For a currency that is known only at run time, for example from a
  database row, use `AnyMoney`. `AnyMoney` checks the currency at run time and returns `ErrCurrencyMismatch`.
  Use `As[T]` to change an `AnyMoney` into a `Money[T]`.
- Each currency is a type. The generated `currency` package has one type for each ISO 4217 currency.

## Rounding is always explicit

Each operation that can round takes a `RoundingMode`. There is no default mode.

Why: different domains need different rules. Tax rules often need `RoundHalfUp`. Accounting often needs `RoundHalfEven`.
A hidden default causes errors that are difficult to find. When the mode is in the call, a reviewer can see it.

The operations that cannot round, such as `Add` and `Allocate`, do not take a mode.
The parse functions do not round. They return `ErrScaleMismatch` when the input has too many digits.

## Errors and not panics

Every operation that can fail returns an error. `MustAdd`, `MustSub` and `MustMul` panic instead.
Use them only when you know the range of the values.

## The minor units come from ISO 4217

`MinorUnits` comes from the ISO 4217 list. CLDR has other digits for some currencies, for display.
Fulus uses ISO 4217 because payment systems and banks use it.

ISO 4217 defines no minor units for some codes, for example the metals XAU, XAG, XPT and XPD, and XDR.
For these codes, the generator uses the CLDR digits and writes this in the doc comment of the type.
The generator does not include XTS (the testing code) and XXX (no currency).
`Format` always writes all the minor units, so it does not lose data.

## Formatting uses CLDR data in generated tables

`Format` uses the CLDR currency patterns, symbols and separators.
`generator.go` reads the CLDR JSON data and writes Go tables in `locale/gen_locale.go` and `currency/gen_currencies.go`.

Why:

- Fulus has no dependencies outside the standard library.
- A lookup is an index into a table. It does not parse data at run time.
- `TestFormatMatchesICU` compares the result with ICU, an independent CLDR implementation.

Cost:

- The tables add about 0.6 MB to each binary that imports Fulus.
- Fulus writes Latin digits only. It does not write Arabic-Indic or other digits.
- `locale.Match` does not use the CLDR likely subtags. For example, `zh-TW` gives `zh`, not `zh-Hant`.

## JSON writes the amount as a decimal string

`MarshalJSON` writes `{"amount":"10.50","currency":"USD"}`. The amount is the canonical decimal in a string.

Why:

- JavaScript and many JSON parsers read numbers as `float64`. A `float64` is exact only up to 2^53.
  A string keeps all the digits of a 128-bit amount.
- A decimal does not depend on the minor units of the reader. If ISO 4217 changes the minor units of a currency,
  a payload in minor units changes its value without an error. A decimal payload keeps its value,
  or gives `ErrScaleMismatch` if it has more fraction digits than the reader accepts.
- Other languages and people can read the payload without a table of minor units.
- The currency code lets the reader check the currency.

## SQL stores the minor units only

`Value` writes the amount in minor units as `int64`. The column does not store the currency,
because the type parameter holds it. If a column holds more than one currency, store the currency code in another
column and use `AnyMoney`.
