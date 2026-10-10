// Package fulus calculates with money in a type-safe way.
//
// Money[T] holds an amount in the currency T. The currency is a type parameter, so the compiler
// stops a calculation that mixes currencies. Rate[Base, Quote] converts between two currencies,
// and Cross checks a triangulation at compile time.
//
// The package supports:
//   - arithmetic without floating point (Add, Sub, Mul, Div, MulFactor, Abs, Neg, Sum)
//   - seven explicit rounding modes, RoundUnnecessary and CLDR cash rounding
//   - allocation with the largest remainder method
//   - runtime currencies with AnyMoney and the currency registry
//   - JSON with decimal strings, text, log/slog, and database/sql for NUMERIC and BIGINT columns
//
// The format package formats and parses amounts with CLDR data. This package does not import it.
//
// The amount is a signed 128-bit integer in minor units. For a currency with 2 minor units,
// the range is about ±1.7 × 10^36. Every operation returns ErrOverflow
// instead of a wrong result when the result does not fit.
package fulus
