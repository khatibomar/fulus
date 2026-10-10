package fulus

import (
	"errors"
	"fmt"
	"math/big"
	"slices"
	"strconv"

	"github.com/khatibomar/fulus/currency"
	"github.com/khatibomar/fulus/locale"
)

var (
	// ErrValidation is the error returned when money validation fails
	ErrValidation = errors.New("money validation error")

	// ErrOverflow indicates an arithmetic operation would overflow
	ErrOverflow = errors.New("arithmetic operation would overflow")

	// ErrInvalidChunks indicates an invalid number of chunks for distribution
	ErrInvalidChunks = errors.New("number of chunks must be positive")

	// ErrDivisionByZero indicates division by zero
	ErrDivisionByZero = errors.New("division by zero")

	// ErrNoRatios indicates no ratios were provided for allocation
	ErrNoRatios = errors.New("no ratios provided")

	// ErrNegativeOrZeroRatios indicates negative ratios in allocation
	ErrNegativeOrZeroRatios = errors.New("ratios must be positive")

	// ErrInvalidRoundingMode indicates unsupported rounding mode
	ErrInvalidRoundingMode = errors.New("invalid rounding mode")

	// ErrInvalidAmountFormat indicates amount string cannot be parsed
	ErrInvalidAmountFormat = errors.New("invalid amount format")

	// ErrScaleMismatch indicates decimal digits exceed currency minor units
	ErrScaleMismatch = errors.New("amount scale exceeds currency minor units")

	// ErrInvalidExchangeRate indicates an exchange rate that is not positive, cannot be parsed, or is the zero Rate
	ErrInvalidExchangeRate = errors.New("invalid exchange rate")

	// ErrInvalidFactor indicates a factor that cannot be parsed, or the zero Factor
	ErrInvalidFactor = errors.New("invalid factor")

	// ErrCurrencyMismatch indicates a value in a different currency than expected
	ErrCurrencyMismatch = errors.New("currency mismatch")

	// ErrUnknownCurrency indicates a currency code that is not registered
	ErrUnknownCurrency = errors.New("unknown currency")

	// ErrInexact indicates a result that needs rounding when the mode is RoundUnnecessary
	ErrInexact = errors.New("result is not exact")
)

// RoundingMode controls how a result between two minor units is rounded.
// The zero value is not a valid mode, so an operation with an unset mode returns ErrInvalidRoundingMode.
type RoundingMode int

const (
	// RoundTruncate rounds toward zero.
	RoundTruncate RoundingMode = iota + 1
	// RoundHalfUp rounds to nearest, ties away from zero.
	RoundHalfUp
	// RoundHalfEven rounds to nearest, ties to even. This is also known as banker's rounding.
	RoundHalfEven
	// RoundHalfDown rounds to nearest, ties toward zero.
	RoundHalfDown
	// RoundUp rounds away from zero.
	RoundUp
	// RoundCeiling rounds toward positive infinity.
	RoundCeiling
	// RoundFloor rounds toward negative infinity.
	RoundFloor
	// RoundUnnecessary does not round. An operation returns ErrInexact if the exact result needs rounding.
	RoundUnnecessary
)

func (r RoundingMode) valid() bool {
	return r >= RoundTruncate && r <= RoundUnnecessary
}

// String returns the name of the mode, for example "HalfEven".
func (r RoundingMode) String() string {
	switch r {
	case RoundTruncate:
		return "Truncate"
	case RoundHalfUp:
		return "HalfUp"
	case RoundHalfEven:
		return "HalfEven"
	case RoundHalfDown:
		return "HalfDown"
	case RoundUp:
		return "Up"
	case RoundCeiling:
		return "Ceiling"
	case RoundFloor:
		return "Floor"
	case RoundUnnecessary:
		return "Unnecessary"
	default:
		return "RoundingMode(" + strconv.Itoa(int(r)) + ")"
	}
}

// Money is an amount of money in the currency T.
// It holds a signed 128-bit integer in minor units of T, for example cents for USD.
// A Money value is 16 bytes. You can compare it with == and use it as a map key.
// The zero value is zero in T.
type Money[T currency.Unit] struct {
	amount int128
}

// Distribution tells how Distribute splits a value into equal chunks.
// Larger is Smaller plus 1 minor unit, or equal to Smaller if LargerCount is zero.
// The sum of all chunks is equal to the value.
type Distribution[T currency.Unit] struct {
	// Smaller is the value of each smaller chunk.
	Smaller Money[T]
	// SmallerCount is the number of smaller chunks.
	SmallerCount int64
	// Larger is the value of each larger chunk.
	Larger Money[T]
	// LargerCount is the number of larger chunks.
	LargerCount int64
}

// NewMoney returns an amount in minor units of T. For example, NewMoney[currency.USD](1050) is USD 10.50.
// Use ParseMoney or NewMoneyFromBigInt for an amount that does not fit in int64.
func NewMoney[T currency.Unit](amount int64) Money[T] {
	return Money[T]{amount: int128FromInt64(amount)}
}

// NewMoneyFromBigInt returns an amount in minor units of T.
// Returns ErrOverflow if the amount does not fit in 128 bits.
func NewMoneyFromBigInt[T currency.Unit](amount *big.Int) (Money[T], error) {
	a, ok := int128FromBig(amount)
	if !ok {
		return Money[T]{}, ErrOverflow
	}
	return Money[T]{amount: a}, nil
}

// Currency returns the currency of the Money value.
func (m Money[T]) Currency() T {
	var c T
	return c
}

// Int64 returns the amount in minor units, and reports whether it fits in int64.
func (m Money[T]) Int64() (int64, bool) {
	return m.amount.int64()
}

// BigInt returns the amount in minor units.
func (m Money[T]) BigInt() *big.Int {
	return m.amount.big()
}

// Add returns m + other.
// Returns ErrOverflow if the result does not fit.
func (m Money[T]) Add(other Money[T]) (Money[T], error) {
	result, ok := add128(m.amount, other.amount)
	if !ok {
		return Money[T]{}, ErrOverflow
	}
	return Money[T]{amount: result}, nil
}

// Sub returns m - other.
// Returns ErrOverflow if the result does not fit.
func (m Money[T]) Sub(other Money[T]) (Money[T], error) {
	result, ok := sub128(m.amount, other.amount)
	if !ok {
		return Money[T]{}, ErrOverflow
	}
	return Money[T]{amount: result}, nil
}

// Mul returns m * scale.
// Returns ErrOverflow if the result does not fit.
func (m Money[T]) Mul(scale int64) (Money[T], error) {
	result, ok := m.amount.mulInt64(scale)
	if !ok {
		return Money[T]{}, ErrOverflow
	}
	return Money[T]{amount: result}, nil
}

// MustAdd is like Add but panics if the result does not fit.
func (m Money[T]) MustAdd(other Money[T]) Money[T] {
	result, err := m.Add(other)
	if err != nil {
		panic(err)
	}
	return result
}

// MustSub is like Sub but panics if the result does not fit.
func (m Money[T]) MustSub(other Money[T]) Money[T] {
	result, err := m.Sub(other)
	if err != nil {
		panic(err)
	}
	return result
}

// MustMul is like Mul but panics if the result does not fit.
func (m Money[T]) MustMul(scale int64) Money[T] {
	result, err := m.Mul(scale)
	if err != nil {
		panic(err)
	}
	return result
}

// Div returns m / divisor, rounded with mode.
// Returns ErrDivisionByZero if divisor is 0, ErrOverflow if the result does not fit,
// and ErrInvalidRoundingMode if the rounding mode is not valid.
func (m Money[T]) Div(divisor int64, mode RoundingMode) (Money[T], error) {
	if divisor == 0 {
		return Money[T]{}, ErrDivisionByZero
	}
	return m.scale(1, divisor, mode)
}

// MulFactor multiplies the Money value by f and rounds the result with mode.
// For example, m.MulFactor(fulus.Percent(15), fulus.RoundHalfEven) gives 15 percent of m.
// Returns ErrInvalidFactor for the zero Factor and ErrOverflow if the result does not fit.
func (m Money[T]) MulFactor(f Factor, mode RoundingMode) (Money[T], error) {
	if f.den == 0 {
		return Money[T]{}, ErrInvalidFactor
	}
	return m.scale(f.num, f.den, mode)
}

// scale returns m*numerator/denominator rounded with mode. The denominator must not be zero.
func (m Money[T]) scale(numerator, denominator int64, mode RoundingMode) (Money[T], error) {
	if !mode.valid() {
		return Money[T]{}, ErrInvalidRoundingMode
	}
	result, err := mulDivRound(m.amount, numerator, denominator, mode)
	if err != nil {
		return Money[T]{}, err
	}
	return Money[T]{amount: result}, nil
}

// RoundCash rounds the Money value to the smallest cash amount of the currency, with mode.
// For example, CHF cash uses steps of 0.05, so 10.03 CHF becomes 10.05 CHF with RoundHalfUp.
// If the currency does not implement currency.CashRounder, RoundCash returns the value unchanged.
// Returns ErrOverflow if the result does not fit.
func (m Money[T]) RoundCash(mode RoundingMode) (Money[T], error) {
	if !mode.valid() {
		return Money[T]{}, ErrInvalidRoundingMode
	}

	rounder, ok := any(m.Currency()).(currency.CashRounder)
	if !ok || rounder.CashIncrement() <= 1 {
		return m, nil
	}

	increment := rounder.CashIncrement()
	steps, err := mulDivRound(m.amount, 1, increment, mode)
	if err != nil {
		return Money[T]{}, err
	}
	result, ok := steps.mulInt64(increment)
	if !ok {
		return Money[T]{}, ErrOverflow
	}
	return Money[T]{amount: result}, nil
}

// Abs returns the absolute value of m.
// Returns ErrOverflow for the smallest amount, because its positive value does not fit.
func (m Money[T]) Abs() (Money[T], error) {
	if !m.amount.isNeg() {
		return m, nil
	}
	return m.Neg()
}

// Neg returns -m.
// Returns ErrOverflow for the smallest amount, because its positive value does not fit.
func (m Money[T]) Neg() (Money[T], error) {
	result, ok := m.amount.neg()
	if !ok {
		return Money[T]{}, ErrOverflow
	}
	return Money[T]{amount: result}, nil
}

// Validate returns ErrValidation if m is not in the closed interval [low, high].
func (m Money[T]) Validate(low, high Money[T]) error {
	if m.LessThan(low) || m.GreaterThan(high) {
		return fmt.Errorf("%w: money amount %s should be in interval [%s, %s]", ErrValidation, m, low, high)
	}
	return nil
}

// Cmp returns -1 if m < other, 0 if m == other, and +1 if m > other.
func (m Money[T]) Cmp(other Money[T]) int {
	return m.amount.cmp(other.amount)
}

// Equal reports whether m == other.
func (m Money[T]) Equal(other Money[T]) bool {
	return m.amount == other.amount
}

// GreaterThan reports whether m > other.
func (m Money[T]) GreaterThan(other Money[T]) bool {
	return m.Cmp(other) > 0
}

// GreaterThanOrEqual reports whether m >= other.
func (m Money[T]) GreaterThanOrEqual(other Money[T]) bool {
	return m.Cmp(other) >= 0
}

// LessThan reports whether m < other.
func (m Money[T]) LessThan(other Money[T]) bool {
	return m.Cmp(other) < 0
}

// LessThanOrEqual reports whether m <= other.
func (m Money[T]) LessThanOrEqual(other Money[T]) bool {
	return m.Cmp(other) <= 0
}

// IsZero reports whether m is zero.
func (m Money[T]) IsZero() bool {
	return m.amount.isZero()
}

// IsPositive reports whether m > 0.
func (m Money[T]) IsPositive() bool {
	return m.amount.sign() > 0
}

// IsNegative reports whether m < 0.
func (m Money[T]) IsNegative() bool {
	return m.amount.isNeg()
}

// Sign returns -1 if m < 0, 0 if m is zero, and +1 if m > 0.
func (m Money[T]) Sign() int {
	return m.amount.sign()
}

// String returns the currency code and the canonical decimal, for example "USD -1234.50".
// The result does not depend on a locale, so it is safe for logs and tests. Use Format for display.
func (m Money[T]) String() string {
	code := m.Currency().Code()
	b := make([]byte, 0, len(code)+42)
	b = append(b, code...)
	b = append(b, ' ')
	return string(appendDecimal(b, m.amount, m.Currency().MinorUnits()))
}

// Format returns a formatted string representation of the Money value for the specified locale.
// It applies the CLDR pattern of the locale, including the negative subpattern and the group sizes.
// The number of fraction digits is always the minor units of the currency.
func (m Money[T]) Format(loc locale.Locale) string {
	c := m.Currency()
	return formatAmount(m.amount, c.MinorUnits(), c.FormatInfo(loc))
}

// Distribute splits m into the given number of chunks with sizes that differ by at most 1 minor unit.
// Returns ErrInvalidChunks if chunks is not positive.
func (m Money[T]) Distribute(chunks int64) (Distribution[T], error) {
	if chunks <= 0 {
		return Distribution[T]{}, ErrInvalidChunks
	}

	// Floor division, so that the remainder is never negative.
	negative := m.amount.isNeg()
	q, r := m.amount.abs().divRem64(uint64(chunks))
	remainder := int64(r)
	if negative && r != 0 {
		// q+1 is at most the magnitude of m, so it fits.
		q, _ = q.inc()
		remainder = chunks - remainder
	}
	smaller, _ := fromMagnitude128(q, negative)
	larger := smaller
	if remainder > 0 {
		larger, _ = add128(smaller, int128FromInt64(1))
	}
	return Distribution[T]{
		Smaller:      Money[T]{amount: smaller},
		SmallerCount: chunks - remainder,
		Larger:       Money[T]{amount: larger},
		LargerCount:  remainder,
	}, nil
}

// Allocate divides m into parts in proportion to the ratios. The sum of the parts is equal to m.
// It uses the largest remainder method. See docs/rounding-and-overflow.md.
func (m Money[T]) Allocate(ratios ...int64) ([]Money[T], error) {
	if len(ratios) == 0 {
		return nil, ErrNoRatios
	}

	var total uint64
	for _, ratio := range ratios {
		if ratio <= 0 {
			return nil, ErrNegativeOrZeroRatios
		}
		total += uint64(ratio)
		if total > 1<<63-1 {
			return nil, ErrOverflow
		}
	}

	// Largest remainder method: each part gets its truncated share first.
	// Then the leftover minor units go one by one to the parts with the largest remainders.
	// Ties go to the part with the lower index.
	parts := make([]Money[T], len(ratios))
	remainders := make([]uint64, len(ratios))
	leftover := m.amount
	magnitude, negative := m.amount.abs(), m.amount.isNeg()
	for i, ratio := range ratios {
		// The share is at most m in magnitude, so it always fits.
		q, r, _ := magnitude.mulDiv64(uint64(ratio), total)
		share, _ := fromMagnitude128(q, negative)
		parts[i] = Money[T]{amount: share}
		remainders[i] = r
		leftover, _ = sub128(leftover, share)
	}
	if leftover.isZero() {
		return parts, nil
	}

	// The leftover is smaller than the number of parts, so it fits in int64.
	count, _ := leftover.int64()
	step := int128FromInt64(1)
	if count < 0 {
		step, count = int128FromInt64(-1), -count
	}
	order := make([]int, len(ratios))
	for i := range order {
		order[i] = i
	}
	slices.SortStableFunc(order, func(a, b int) int {
		return cmpUint64(remainders[b], remainders[a])
	})
	for _, i := range order[:count] {
		parts[i].amount, _ = add128(parts[i].amount, step)
	}

	return parts, nil
}

func cmpUint64(a, b uint64) int {
	switch {
	case a < b:
		return -1
	case a > b:
		return 1
	default:
		return 0
	}
}

// Sum returns the sum of the values, or zero if there are no values.
// Returns ErrOverflow if the sum does not fit.
// An intermediate sum can be larger than the range if the final sum fits.
func Sum[T currency.Unit](values ...Money[T]) (Money[T], error) {
	var total int128
	for i, v := range values {
		next, ok := add128(total, v.amount)
		if ok {
			total = next
			continue
		}
		b := total.big()
		for _, w := range values[i:] {
			b.Add(b, w.amount.big())
		}
		return NewMoneyFromBigInt[T](b)
	}
	return Money[T]{amount: total}, nil
}

// Min returns the smallest of the values.
func Min[T currency.Unit](first Money[T], rest ...Money[T]) Money[T] {
	result := first
	for _, v := range rest {
		if v.LessThan(result) {
			result = v
		}
	}
	return result
}

// Max returns the largest of the values.
func Max[T currency.Unit](first Money[T], rest ...Money[T]) Money[T] {
	result := first
	for _, v := range rest {
		if v.GreaterThan(result) {
			result = v
		}
	}
	return result
}
