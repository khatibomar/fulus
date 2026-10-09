// Package fulus provides utilities for handling currency conversion and calculations.
//
// Fulus helps developers work with different currencies and monetary operations
// in a type-safe and efficient manner. It includes features for currency parsing,
// formatting, conversion, and basic arithmetic operations while maintaining
// precision and correctness.
//
// The package supports:
//   - type-safe arithmetic across currency types (Add, Sub, Mul, Div, MulFrac, MulDecimal, Abs, Neg)
//   - seven explicit rounding modes and CLDR cash rounding
//   - runtime currencies with AnyMoney and the currency registry
//   - CLDR formatting and parsing of localized strings
//   - JSON, text, log/slog and database/sql interoperability
//
// The amount is an int64 in minor units. For a currency with 2 minor units,
// the range is ±92,233,720,368,547,758.07. Every operation returns ErrOverflow
// instead of a wrong result when the result does not fit.
package fulus
