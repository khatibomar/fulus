package fulus

import (
	"cmp"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"math"
	"math/big"
	"math/bits"
	"slices"
	"strconv"
	"strings"
	"sync/atomic"

	"github.com/khatibomar/fulus/currency"
	"github.com/khatibomar/fulus/locale"
)

var defaultLocale atomic.Pointer[locale.Locale]

// DefaultLocale returns the locale that String uses. The initial value is locale.EN.
func DefaultLocale() locale.Locale {
	if l := defaultLocale.Load(); l != nil {
		return *l
	}
	return locale.EN
}

// SetDefaultLocale sets the locale that String uses. It is safe for concurrent use.
func SetDefaultLocale(l locale.Locale) {
	defaultLocale.Store(&l)
}

var (
	// ErrValidation is the error returned when money validation fails
	ErrValidation = errors.New("money validation error")

	// ErrOverflow indicates an arithmetic operation would overflow
	ErrOverflow = errors.New("arithmetic operation would overflow")

	// ErrInvalidChunks indicates an invalid number of chunks for distribution
	ErrInvalidChunks = errors.New("number of chunks must be positive")

	// ErrZeroDenominator indicates division by zero in conversion
	ErrZeroDenominator = errors.New("denominator cannot be zero")

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

	// ErrInvalidExchangeRate indicates an invalid exchange rate format or value
	ErrInvalidExchangeRate = errors.New("invalid exchange rate")

	// ErrInvalidFactor indicates a multiplication factor that cannot be parsed
	ErrInvalidFactor = errors.New("invalid factor")

	// ErrCurrencyMismatch indicates a value in a different currency than expected
	ErrCurrencyMismatch = errors.New("currency mismatch")

	// ErrUnknownCurrency indicates a currency code that is not registered
	ErrUnknownCurrency = errors.New("unknown currency")
)

// RoundingMode controls how a result between two minor units is rounded.
type RoundingMode int

const (
	// RoundTruncate rounds toward zero (default behavior).
	RoundTruncate RoundingMode = iota
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
)

func (r RoundingMode) valid() bool {
	return r >= RoundTruncate && r <= RoundFloor
}

// Money represents a monetary value in a specific currency.
type Money[T currency.Currency] struct {
	// amount stores the monetary value in the currency's smallest unit (e.g., cents for USD)
	amount int64
}

// Distribution represents how to split money into chunks
type Distribution struct {
	// SmallerChunkSize represents the value of the smaller portions in the distribution
	SmallerChunkSize int64
	// SmallerCount represents how many smaller chunks are in the distribution
	SmallerCount int64
	// LargerChunkSize represents the value of the larger portions in the distribution
	LargerChunkSize int64
	// LargerCount represents how many larger chunks are in the distribution
	LargerCount int64
}

// Ratio represents a fraction used for conversion rates
type Ratio[F currency.Currency, T currency.Currency] struct {
	// Numerator is the top number in the fraction (e.g., 107203 for 1.07203)
	Numerator int64
	// Denominator is the bottom number in the fraction (e.g., 100000 for precise decimal representation)
	Denominator int64
}

// ParseRatioString parses a string representation of an exchange rate into a Ratio
func ParseRatioString[F currency.Currency, T currency.Currency](rate string) (Ratio[F, T], error) {
	r, ok := new(big.Rat).SetString(rate)
	if !ok {
		return Ratio[F, T]{}, fmt.Errorf("%w: %s", ErrInvalidExchangeRate, rate)
	}

	if r.Sign() <= 0 {
		return Ratio[F, T]{}, fmt.Errorf("%w: rate must be positive", ErrInvalidExchangeRate)
	}

	if !r.Num().IsInt64() || !r.Denom().IsInt64() {
		return Ratio[F, T]{}, ErrOverflow
	}

	return Ratio[F, T]{
		Numerator:   r.Num().Int64(),
		Denominator: r.Denom().Int64(),
	}, nil
}

// ParseRatioFloat64 parses a float64 representation of an exchange rate into a Ratio.
// It formats the float to a string to avoid floating point precision issues.
func ParseRatioFloat64[F currency.Currency, T currency.Currency](rate float64) (Ratio[F, T], error) {
	if rate <= 0 {
		return Ratio[F, T]{}, fmt.Errorf("%w: rate must be positive", ErrInvalidExchangeRate)
	}
	s := strconv.FormatFloat(rate, 'f', -1, 64)
	return ParseRatioString[F, T](s)
}

// Allocation represents how money is divided according to ratios
type Allocation[T currency.Currency] struct {
	Parts []Money[T]
	Total Money[T]
}

// ConversionResult holds both the converted amount and the actual ratio used
type ConversionResult[F currency.Currency, T currency.Currency] struct {
	// Amount stores the resulting converted monetary value
	Amount int64
	// ActualRate is the applied rate Amount/source amount, reduced and with a positive denominator.
	// For a zero source amount it is the requested ratio.
	ActualRate Ratio[F, T]
}

// NewMoney creates a new Money instance with the given amount and currency.
// The amount should be specified in the currency's smallest sub-unit
// (e.g., cents for USD, pence for GBP). For example:
// USD 10.50 should be passed as 1050
// EUR 5.99 should be passed as 599
func NewMoney[T currency.Currency](amount int64) Money[T] {
	return Money[T]{amount: amount}
}

// Currency returns the currency of the Money value.
func (m Money[T]) Currency() T {
	var c T
	return c
}

// Add performs addition of two Money values of the same currency.
// Returns ErrOverflow if the operation would overflow int64.
func (m Money[T]) Add(other Money[T]) (Money[T], error) {
	result, ok := add64(m.amount, other.amount)
	if !ok {
		return Money[T]{}, ErrOverflow
	}
	return Money[T]{amount: result}, nil
}

// Sub performs subtraction of two Money values of the same currency.
// Returns ErrOverflow if the operation would overflow int64.
func (m Money[T]) Sub(other Money[T]) (Money[T], error) {
	result, ok := sub64(m.amount, other.amount)
	if !ok {
		return Money[T]{}, ErrOverflow
	}
	return Money[T]{amount: result}, nil
}

// Mul multiplies the Money value by a scalar value.
// Returns ErrOverflow if the operation would overflow int64.
func (m Money[T]) Mul(scale int64) (Money[T], error) {
	result, ok := mul64(m.amount, scale)
	if !ok {
		return Money[T]{}, ErrOverflow
	}
	return Money[T]{amount: result}, nil
}

// MustAdd performs addition of two Money values of the same currency.
// Panics if the operation would overflow int64.
func (m Money[T]) MustAdd(other Money[T]) Money[T] {
	result, err := m.Add(other)
	if err != nil {
		panic(err)
	}
	return result
}

// MustSub performs subtraction of two Money values of the same currency.
// Panics if the operation would overflow int64.
func (m Money[T]) MustSub(other Money[T]) Money[T] {
	result, err := m.Sub(other)
	if err != nil {
		panic(err)
	}
	return result
}

// MustMul multiplies the Money value by a scalar value.
// Panics if the operation would overflow int64.
func (m Money[T]) MustMul(scale int64) Money[T] {
	result, err := m.Mul(scale)
	if err != nil {
		panic(err)
	}
	return result
}

// Div divides the Money value by a scalar value using the specified rounding mode.
// Returns ErrDivisionByZero if divisor is 0.
// Returns ErrOverflow if the operation would overflow int64.
// Returns ErrInvalidRoundingMode if the rounding mode is unsupported.
func (m Money[T]) Div(divisor int64, mode RoundingMode) (Money[T], error) {
	if divisor == 0 {
		return Money[T]{}, ErrDivisionByZero
	}
	return m.scale(1, divisor, mode)
}

// MulFrac multiplies the Money value by numerator/denominator and rounds the result with mode.
// For example, MulFrac(15, 100, RoundHalfEven) gives 15 percent of the value.
// Returns ErrDivisionByZero if denominator is 0.
// Returns ErrOverflow if the result does not fit in int64.
func (m Money[T]) MulFrac(numerator, denominator int64, mode RoundingMode) (Money[T], error) {
	if denominator == 0 {
		return Money[T]{}, ErrDivisionByZero
	}
	return m.scale(numerator, denominator, mode)
}

// MulDecimal multiplies the Money value by a decimal factor such as "0.0825" and rounds the result with mode.
// The factor can also be a fraction such as "1/3".
// Returns ErrInvalidFactor if the factor cannot be parsed.
// Returns ErrOverflow if the result does not fit in int64.
func (m Money[T]) MulDecimal(factor string, mode RoundingMode) (Money[T], error) {
	if numerator, denominator, ok := parseFactor(factor); ok {
		return m.scale(numerator, denominator, mode)
	}

	r, ok := new(big.Rat).SetString(factor)
	if !ok {
		return Money[T]{}, fmt.Errorf("%w: %q", ErrInvalidFactor, factor)
	}
	if r.Num().IsInt64() && r.Denom().IsInt64() {
		return m.scale(r.Num().Int64(), r.Denom().Int64(), mode)
	}

	product := new(big.Int).Mul(big.NewInt(m.amount), r.Num())
	result, err := divideWithRounding(product, r.Denom(), mode)
	if err != nil {
		return Money[T]{}, err
	}
	if !result.IsInt64() {
		return Money[T]{}, ErrOverflow
	}
	return Money[T]{amount: result.Int64()}, nil
}

// scale returns m*numerator/denominator rounded with mode. The denominator must not be zero.
func (m Money[T]) scale(numerator, denominator int64, mode RoundingMode) (Money[T], error) {
	if !mode.valid() {
		return Money[T]{}, ErrInvalidRoundingMode
	}
	result, ok := mulDivRound(m.amount, numerator, denominator, mode)
	if !ok {
		return Money[T]{}, ErrOverflow
	}
	return Money[T]{amount: result}, nil
}

// parseFactor parses a short decimal factor such as "-0.0825" into numerator/denominator without allocation.
// It reports false for other forms, which MulDecimal parses with math/big.
func parseFactor(s string) (numerator, denominator int64, ok bool) {
	negative := false
	if s != "" && (s[0] == '-' || s[0] == '+') {
		negative = s[0] == '-'
		s = s[1:]
	}
	if s == "" || len(s) > 18 {
		return 0, 0, false
	}

	denominator = 1
	seenDigit, seenPoint := false, false
	for i := range len(s) {
		switch c := s[i]; {
		case c >= '0' && c <= '9':
			numerator = numerator*10 + int64(c-'0')
			if seenPoint {
				denominator *= 10
			}
			seenDigit = true
		case c == '.' && !seenPoint:
			seenPoint = true
		default:
			return 0, 0, false
		}
	}
	if !seenDigit {
		return 0, 0, false
	}
	if negative {
		numerator = -numerator
	}
	return numerator, denominator, true
}

// RoundCash rounds the Money value to the smallest cash amount of the currency, with mode.
// For example, CHF cash uses steps of 0.05, so 10.03 CHF becomes 10.05 CHF with RoundHalfUp.
// If the currency does not implement currency.CashRounder, RoundCash returns the value unchanged.
// Returns ErrOverflow if the result does not fit in int64.
func (m Money[T]) RoundCash(mode RoundingMode) (Money[T], error) {
	if !mode.valid() {
		return Money[T]{}, ErrInvalidRoundingMode
	}

	rounder, ok := any(m.Currency()).(currency.CashRounder)
	if !ok || rounder.CashIncrement() <= 1 {
		return m, nil
	}

	increment := rounder.CashIncrement()
	steps, ok := mulDivRound(m.amount, 1, increment, mode)
	if !ok {
		return Money[T]{}, ErrOverflow
	}
	result, ok := mul64(steps, increment)
	if !ok {
		return Money[T]{}, ErrOverflow
	}
	return Money[T]{amount: result}, nil
}

// Abs returns the absolute value of the money amount.
// Returns ErrOverflow if the amount is math.MinInt64.
func (m Money[T]) Abs() (Money[T], error) {
	if m.amount == math.MinInt64 {
		return Money[T]{}, ErrOverflow
	}
	if m.amount < 0 {
		return Money[T]{amount: -m.amount}, nil
	}
	return m, nil
}

// Neg returns the negated value of the money amount.
// Returns ErrOverflow if the amount is math.MinInt64.
func (m Money[T]) Neg() (Money[T], error) {
	if m.amount == math.MinInt64 {
		return Money[T]{}, ErrOverflow
	}
	return Money[T]{amount: -m.amount}, nil
}

// Validate checks if the money amount falls within the specified range [min, max].
// Returns an error if the amount is outside the range.
func (m Money[T]) Validate(min, max int64) error {
	if m.amount < min || m.amount > max {
		return fmt.Errorf("%w: money amount %s should be in interval [%s, %s]",
			ErrValidation,
			m,
			NewMoney[T](min),
			NewMoney[T](max),
		)
	}
	return nil
}

// Cmp compares two Money values and returns:
// -1 if m < other
//
//	0 if m == other
//
// +1 if m > other
func (m Money[T]) Cmp(other Money[T]) int {
	if m.amount < other.amount {
		return -1
	}
	if m.amount > other.amount {
		return 1
	}
	return 0
}

// Equal returns true if the two Money values are equal.
func (m Money[T]) Equal(other Money[T]) bool {
	return m.amount == other.amount
}

// GreaterThan returns true if the Money value is greater than the other.
func (m Money[T]) GreaterThan(other Money[T]) bool {
	return m.amount > other.amount
}

// GreaterThanOrEqual returns true if the Money value is greater than or equal to the other.
func (m Money[T]) GreaterThanOrEqual(other Money[T]) bool {
	return m.amount >= other.amount
}

// LessThan returns true if the Money value is less than the other.
func (m Money[T]) LessThan(other Money[T]) bool {
	return m.amount < other.amount
}

// LessThanOrEqual returns true if the Money value is less than or equal to the other.
func (m Money[T]) LessThanOrEqual(other Money[T]) bool {
	return m.amount <= other.amount
}

// IsZero returns true if the Money amount is zero.
func (m Money[T]) IsZero() bool {
	return m.amount == 0
}

// IsPositive returns true if the Money amount is greater than zero.
func (m Money[T]) IsPositive() bool {
	return m.amount > 0
}

// IsNegative returns true if the Money amount is less than zero.
func (m Money[T]) IsNegative() bool {
	return m.amount < 0
}

// Amount returns the internal amount value in the currency's smallest unit.
// For example, returns cents for USD or pence for GBP.
func (m Money[T]) Amount() int64 {
	return m.amount
}

// String returns a formatted string representation of the Money value using the default locale.
// This implements the fmt.Stringer interface.
func (m Money[T]) String() string {
	return m.Format(DefaultLocale())
}

// Format returns a formatted string representation of the Money value for the specified locale.
// It applies the CLDR pattern of the locale, including the negative subpattern and the group sizes.
// The number of fraction digits is always the minor units of the currency.
func (m Money[T]) Format(loc locale.Locale) string {
	c := m.Currency()
	return formatAmount(m.amount, c.MinorUnits(), c.FormatInfo(loc))
}

// Distribute splits the money amount into the specified number of chunks
// Returns a Distribution describing how to split the money
func (m Money[T]) Distribute(chunks int64) (Distribution, error) {
	if chunks <= 0 {
		return Distribution{}, ErrInvalidChunks
	}

	amount := m.Amount()

	// For even distribution
	if amount%chunks == 0 {
		chunkSize := amount / chunks
		return Distribution{
			SmallerChunkSize: chunkSize,
			SmallerCount:     chunks,
			LargerChunkSize:  chunkSize,
			LargerCount:      0,
		}, nil
	}

	// For uneven distribution
	smallerChunkSize := amount / chunks
	largerChunkSize := smallerChunkSize + 1
	remainder := amount % chunks

	if remainder < 0 {
		smallerChunkSize--
		largerChunkSize--
		remainder = -remainder
		return Distribution{
			SmallerChunkSize: smallerChunkSize,
			SmallerCount:     remainder,
			LargerChunkSize:  largerChunkSize,
			LargerCount:      chunks - remainder,
		}, nil
	}

	return Distribution{
		SmallerChunkSize: smallerChunkSize,
		SmallerCount:     chunks - remainder,
		LargerChunkSize:  largerChunkSize,
		LargerCount:      remainder,
	}, nil
}

// Convert performs conversion with an explicit rounding strategy.
// The ratio should be provided as (numerator, denominator) representing numerator/denominator.
// Returns both the converted Money value and the actual ratio used after rounding.
func Convert[F, T currency.Currency](m Money[F], ratio Ratio[F, T], mode RoundingMode) (Money[T], ConversionResult[F, T], error) {
	if ratio.Denominator == 0 {
		return Money[T]{}, ConversionResult[F, T]{}, ErrZeroDenominator
	}

	if !mode.valid() {
		return Money[T]{}, ConversionResult[F, T]{}, ErrInvalidRoundingMode
	}

	roundedAmount, ok := mulDivRound(m.amount, ratio.Numerator, ratio.Denominator, mode)
	if !ok {
		return Money[T]{}, ConversionResult[F, T]{}, ErrOverflow
	}

	actualRate := ratio
	if m.amount != 0 {
		actualRate, ok = reducedRatio[F, T](roundedAmount, m.amount)
		if !ok {
			return Money[T]{}, ConversionResult[F, T]{}, ErrOverflow
		}
	}

	result := ConversionResult[F, T]{
		Amount:     roundedAmount,
		ActualRate: actualRate,
	}

	return NewMoney[T](roundedAmount), result, nil
}

// reducedRatio returns numerator/denominator in lowest terms with a positive denominator.
// It reports false if a term does not fit in int64. The denominator must not be zero.
func reducedRatio[F, T currency.Currency](numerator, denominator int64) (Ratio[F, T], bool) {
	g := gcd(abs64(numerator), abs64(denominator))
	negative := (numerator < 0) != (denominator < 0)
	n, okN := fromMagnitude(abs64(numerator)/g, negative)
	d, okD := fromMagnitude(abs64(denominator)/g, false)
	return Ratio[F, T]{Numerator: n, Denominator: d}, okN && okD
}

func divideWithRounding(numerator, denominator *big.Int, mode RoundingMode) (*big.Int, error) {
	if !mode.valid() {
		return nil, ErrInvalidRoundingMode
	}

	q := new(big.Int)
	r := new(big.Int)
	q.QuoRem(numerator, denominator, r)

	if r.Sign() == 0 || mode == RoundTruncate {
		return q, nil
	}

	absRem := new(big.Int).Abs(r)
	absDen := new(big.Int).Abs(denominator)
	twiceRem := new(big.Int).Lsh(absRem, 1)
	cmp := twiceRem.Cmp(absDen)
	negativeResult := (numerator.Sign() < 0) != (denominator.Sign() < 0)

	var awayFromZero bool
	switch mode {
	case RoundHalfUp:
		awayFromZero = cmp >= 0
	case RoundHalfDown:
		awayFromZero = cmp > 0
	case RoundHalfEven:
		awayFromZero = cmp > 0 || (cmp == 0 && q.Bit(0) == 1)
	case RoundUp:
		awayFromZero = true
	case RoundCeiling:
		awayFromZero = !negativeResult
	case RoundFloor:
		awayFromZero = negativeResult
	}

	if awayFromZero {
		if negativeResult {
			q.Sub(q, big.NewInt(1))
		} else {
			q.Add(q, big.NewInt(1))
		}
	}
	return q, nil
}

// Allocate divides money according to provided ratios
func (m Money[T]) Allocate(ratios []int64) (Allocation[T], error) {
	if len(ratios) == 0 {
		return Allocation[T]{}, ErrNoRatios
	}

	total := int64(0)
	for _, ratio := range ratios {
		if ratio <= 0 {
			return Allocation[T]{}, ErrNegativeOrZeroRatios
		}
		if total > math.MaxInt64-ratio {
			return Allocation[T]{}, ErrOverflow
		}
		total += ratio
	}

	// Largest remainder method: each part gets its truncated share first.
	// Then the leftover minor units go one by one to the parts with the largest remainders.
	// Ties go to the part with the lower index.
	parts := make([]Money[T], len(ratios))
	remainders := make([]uint64, len(ratios))
	leftover := m.amount
	for i, ratio := range ratios {
		// The share is at most m.amount in magnitude, so it always fits.
		q, r, negative, _ := mulDiv(m.amount, ratio, total)
		share, _ := fromMagnitude(q, negative)
		parts[i] = Money[T]{amount: share}
		remainders[i] = r
		leftover -= share
	}
	if leftover == 0 {
		return Allocation[T]{Parts: parts, Total: m}, nil
	}

	step := int64(1)
	if leftover < 0 {
		step = -1
	}
	order := make([]int, len(ratios))
	for i := range order {
		order[i] = i
	}
	slices.SortStableFunc(order, func(a, b int) int {
		return cmp.Compare(remainders[b], remainders[a])
	})
	for _, i := range order[:leftover*step] {
		parts[i].amount += step
	}

	return Allocation[T]{Parts: parts, Total: m}, nil
}

// MarshalJSON implements the json.Marshaler interface
// Serializes the amount as a string to preserve precision
func (m Money[T]) MarshalJSON() ([]byte, error) {
	return marshalMoneyJSON(m.amount, m.Currency().Code())
}

// marshalMoneyJSON writes the JSON form of an amount without reflection
// when the currency code needs no escaping.
func marshalMoneyJSON(amount int64, code string) ([]byte, error) {
	for i := range len(code) {
		if c := code[i]; c < 0x20 || c > 0x7e || c == '"' || c == '\\' || c == '<' || c == '>' || c == '&' {
			return json.Marshal(moneyJSON{Amount: strconv.FormatInt(amount, 10), Currency: code})
		}
	}

	b := make([]byte, 0, len(`{"amount":"","currency":""}`)+20+len(code))
	b = append(b, `{"amount":"`...)
	b = strconv.AppendInt(b, amount, 10)
	b = append(b, `","currency":"`...)
	b = append(b, code...)
	return append(b, `"}`...), nil
}

type moneyJSON struct {
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
}

// UnmarshalJSON implements the json.Unmarshaler interface
// Expects amount as a string to maintain precision
func (m *Money[T]) UnmarshalJSON(data []byte) error {
	var temp moneyJSON
	if err := json.Unmarshal(data, &temp); err != nil {
		return fmt.Errorf("failed to unmarshal money: %w", err)
	}

	amount, err := parseIntAmount(temp.Amount)
	if err != nil {
		return err
	}

	code := m.Currency().Code()
	if code != temp.Currency {
		return fmt.Errorf("%w: expected %s, got %s", ErrCurrencyMismatch, code, temp.Currency)
	}

	m.amount = amount
	return nil
}

// ParseMoney parses a canonical decimal amount into Money using the currency's minor units.
// Supported format: optional sign (+/-), digits, optional decimal point and digits.
func ParseMoney[T currency.Currency](amount string) (Money[T], error) {
	var c T
	minor, err := parseDecimal(amount, c.MinorUnits())
	if err != nil {
		return Money[T]{}, err
	}
	return Money[T]{amount: minor}, nil
}

// parseDecimal parses a canonical decimal amount into minor units.
func parseDecimal(amount string, minorUnits int) (int64, error) {
	if amount == "" {
		return 0, fmt.Errorf("%w: empty amount", ErrInvalidAmountFormat)
	}

	negative := false
	if amount[0] == '+' || amount[0] == '-' {
		negative = amount[0] == '-'
		amount = amount[1:]
		if amount == "" {
			return 0, fmt.Errorf("%w: sign without digits", ErrInvalidAmountFormat)
		}
	}

	whole, fractional, hasPoint := strings.Cut(amount, ".")
	if hasPoint && strings.Contains(fractional, ".") {
		return 0, fmt.Errorf("%w: multiple decimal separators", ErrInvalidAmountFormat)
	}
	if whole == "" {
		return 0, fmt.Errorf("%w: missing whole part", ErrInvalidAmountFormat)
	}
	if !allDigits(whole) {
		return 0, fmt.Errorf("%w: invalid whole part", ErrInvalidAmountFormat)
	}
	if hasPoint && fractional == "" {
		return 0, fmt.Errorf("%w: missing fractional part", ErrInvalidAmountFormat)
	}
	if !allDigits(fractional) {
		return 0, fmt.Errorf("%w: invalid fractional part", ErrInvalidAmountFormat)
	}
	if len(fractional) > minorUnits {
		return 0, fmt.Errorf("%w: got %d fractional digits, max is %d", ErrScaleMismatch, len(fractional), minorUnits)
	}

	var magnitude uint64
	push := func(digit uint64) bool {
		if magnitude > (math.MaxUint64-digit)/10 {
			return false
		}
		magnitude = magnitude*10 + digit
		return true
	}
	for i := range len(whole) {
		if !push(uint64(whole[i] - '0')) {
			return 0, ErrOverflow
		}
	}
	for i := range len(fractional) {
		if !push(uint64(fractional[i] - '0')) {
			return 0, ErrOverflow
		}
	}
	for range minorUnits - len(fractional) {
		if !push(0) {
			return 0, ErrOverflow
		}
	}

	result, ok := fromMagnitude(magnitude, negative)
	if !ok {
		return 0, ErrOverflow
	}
	return result, nil
}

func parseIntAmount(amount string) (int64, error) {
	if amount == "" {
		return 0, fmt.Errorf("%w: empty amount", ErrInvalidAmountFormat)
	}

	parsed, err := strconv.ParseInt(amount, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("%w: %v", ErrInvalidAmountFormat, err)
	}

	return parsed, nil
}

func allDigits(s string) bool {
	for _, ch := range s {
		if ch < '0' || ch > '9' {
			return false
		}
	}
	return true
}

// Value implements driver.Valuer for database/sql.
// It returns the amount in minor units as int64, for use with an integer column such as BIGINT.
func (m Money[T]) Value() (driver.Value, error) {
	return m.amount, nil
}

// Scan implements sql.Scanner for database/sql.
// It accepts an int64 amount in minor units, the same amount as integer text,
// or the JSON form that MarshalJSON writes.
func (m *Money[T]) Scan(value any) error {
	switch v := value.(type) {
	case nil:
		return fmt.Errorf("cannot scan NULL into Money")
	case int64:
		m.amount = v
		return nil
	case []byte:
		return m.scanText(string(v))
	case string:
		return m.scanText(v)
	default:
		return fmt.Errorf("cannot scan type %T into Money", value)
	}
}

func (m *Money[T]) scanText(s string) error {
	if amount, err := strconv.ParseInt(s, 10, 64); err == nil {
		m.amount = amount
		return nil
	}
	return m.UnmarshalJSON([]byte(s))
}

// Decimal returns the amount as a canonical decimal string such as "-1234.50".
// It has no symbol and no group separator, and ParseMoney accepts it.
func (m Money[T]) Decimal() string {
	return formatAmount(m.amount, m.Currency().MinorUnits(), canonicalFormatInfo)
}

var canonicalFormatInfo = currency.FormatInfo{
	Format:           "0.00",
	DecimalSeparator: ".",
	MinusSign:        "-",
}

// MarshalText implements encoding.TextMarshaler. It writes the canonical decimal form from Decimal.
func (m Money[T]) MarshalText() ([]byte, error) {
	return []byte(m.Decimal()), nil
}

// UnmarshalText implements encoding.TextUnmarshaler. It reads the canonical decimal form with ParseMoney.
func (m *Money[T]) UnmarshalText(text []byte) error {
	parsed, err := ParseMoney[T](string(text))
	if err != nil {
		return err
	}
	*m = parsed
	return nil
}

// LogValue implements slog.LogValuer. It logs the decimal amount and the currency code.
func (m Money[T]) LogValue() slog.Value {
	return slog.GroupValue(
		slog.String("amount", m.Decimal()),
		slog.String("currency", m.Currency().Code()),
	)
}

// Sum returns the sum of the values, or zero if there are no values.
// Returns ErrOverflow if the sum does not fit in int64.
// An intermediate sum can be larger than int64 if the final sum fits.
func Sum[T currency.Currency](values ...Money[T]) (Money[T], error) {
	// hi and lo hold a signed 128-bit total, so no intermediate sum can overflow.
	var hi int64
	var lo uint64
	for _, v := range values {
		var carry uint64
		lo, carry = bits.Add64(lo, uint64(v.amount), 0)
		hi += v.amount>>63 + int64(carry)
	}
	if hi != int64(lo)>>63 {
		return Money[T]{}, ErrOverflow
	}
	return Money[T]{amount: int64(lo)}, nil
}

// Min returns the smallest of the values.
func Min[T currency.Currency](first Money[T], rest ...Money[T]) Money[T] {
	result := first
	for _, v := range rest {
		result.amount = min(result.amount, v.amount)
	}
	return result
}

// Max returns the largest of the values.
func Max[T currency.Currency](first Money[T], rest ...Money[T]) Money[T] {
	result := first
	for _, v := range rest {
		result.amount = max(result.amount, v.amount)
	}
	return result
}
