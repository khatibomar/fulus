package fulus

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"math/big"
	"strconv"
	"strings"

	"github.com/khatibomar/fulus/currency"
	"github.com/khatibomar/fulus/locale"
)

// DefaultLocale is the fallback locale when none is specified
var (
	DefaultLocale = locale.EN
)

var (
	// ErrValidation is the error returned when money validation fails
	ErrValidation = errors.New("money validation error")

	// ErrOverflow indicates an arithmetic operation would overflow
	ErrOverflow = errors.New("arithmetic operation would overflow")

	// ErrInvalidChunks indicates an invalid number of chunks for distribution
	ErrInvalidChunks = errors.New("number of chunks must be positive")

	// ErrZeroDenominator indicates division by zero in conversion
	ErrZeroDenominator = errors.New("denominator cannot be zero")

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
)

// RoundingMode controls how division results are rounded in conversion.
type RoundingMode int

const (
	// RoundTruncate rounds toward zero (default behavior).
	RoundTruncate RoundingMode = iota
	// RoundHalfUp rounds to nearest, ties away from zero.
	RoundHalfUp
	// RoundHalfEven rounds to nearest, ties to even.
	RoundHalfEven
)

// Money represents a monetary value in a specific currency.
type Money[T currency.Currency] struct {
	// amount stores the monetary value in the currency's smallest unit (e.g., cents for USD)
	amount int64
	// Currency represents the type of currency for this money value
	currency.Currency
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

// Allocation represents how money is divided according to ratios
type Allocation[T currency.Currency] struct {
	Parts []Money[T]
	Total Money[T]
}

// ConversionResult holds both the converted amount and the actual ratio used
type ConversionResult[F currency.Currency, T currency.Currency] struct {
	// Amount stores the resulting converted monetary value
	Amount int64
	// ActualRate stores the precise conversion rate that was actually applied
	ActualRate Ratio[F, T]
}

// NewMoney creates a new Money instance with the given amount and currency.
// The amount should be specified in the currency's smallest sub-unit
// (e.g., cents for USD, pence for GBP). For example:
// USD 10.50 should be passed as 1050
// EUR 5.99 should be passed as 599
func NewMoney[T currency.Currency](amount int64) Money[T] {
	var c T
	return Money[T]{
		amount:   amount,
		Currency: c,
	}
}

// Add performs addition of two Money values of the same currency.
// Returns ErrOverflow if the operation would overflow int64.
func (m Money[T]) Add(other Money[T]) (Money[T], error) {
	result := big.NewInt(m.amount)
	result.Add(result, big.NewInt(other.amount))

	if !result.IsInt64() {
		return Money[T]{}, ErrOverflow
	}

	return Money[T]{amount: result.Int64(), Currency: m.Currency}, nil
}

// Sub performs subtraction of two Money values of the same currency.
// Returns ErrOverflow if the operation would overflow int64.
func (m Money[T]) Sub(other Money[T]) (Money[T], error) {
	result := big.NewInt(m.amount)
	result.Sub(result, big.NewInt(other.amount))

	if !result.IsInt64() {
		return Money[T]{}, ErrOverflow
	}

	return Money[T]{amount: result.Int64(), Currency: m.Currency}, nil
}

// Mul multiplies the Money value by a scalar value.
// Returns ErrOverflow if the operation would overflow int64.
// If scale is 0, sets the amount to 0 and returns nil.
func (m Money[T]) Mul(scale int64) (Money[T], error) {
	if scale == 0 {
		return Money[T]{amount: 0, Currency: m.Currency}, nil
	}

	// Use math/big to check for overflow
	result := big.NewInt(m.amount)
	result.Mul(result, big.NewInt(scale))

	// Check if result fits in int64
	if !result.IsInt64() {
		return Money[T]{}, ErrOverflow
	}

	return Money[T]{amount: result.Int64(), Currency: m.Currency}, nil
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
	return m.Format(DefaultLocale)
}

// Format returns a formatted string representation of the Money value for the specified locale.
func (m Money[T]) Format(locale locale.Locale) string {
	info := m.Currency.FormatInfo(locale)

	if m.amount == 0 {
		zeroFmt := strings.Replace(info.Format, "#,##0.00", "0", 1)
		return strings.Replace(zeroFmt, "¤", info.Symbol, 1)
	}

	amount := big.NewInt(m.amount)
	negative := amount.Sign() < 0
	if negative {
		amount.Abs(amount)
	}

	divisor := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(m.Currency.MinorUnits())), nil)
	major := new(big.Int).Quo(new(big.Int).Set(amount), divisor)
	minor := new(big.Int).Mod(new(big.Int).Set(amount), divisor)

	majorStr := formatMajorString(major.String(), info.GroupSeparator)

	var result string
	if m.Currency.MinorUnits() > 0 {
		minorStr := minor.String()
		if len(minorStr) < m.Currency.MinorUnits() {
			minorStr = strings.Repeat("0", m.Currency.MinorUnits()-len(minorStr)) + minorStr
		}
		result = strings.Replace(info.Format, "#,##0.00",
			majorStr+info.DecimalSeparator+minorStr,
			1)
	} else {
		result = strings.Replace(info.Format, "#,##0.00", major.String(), 1)
		result = strings.Replace(result, ".00", "", 1)
	}

	result = strings.Replace(result, "¤", info.Symbol, 1)

	if negative {
		if !strings.Contains(result, info.MinusSign) {
			result = info.MinusSign + result
		}
	}

	return result
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

	theoretical := big.NewInt(m.amount)
	theoretical.Mul(theoretical, big.NewInt(ratio.Numerator))
	quotient, err := divideWithRounding(theoretical, big.NewInt(ratio.Denominator), mode)
	if err != nil {
		return Money[T]{}, ConversionResult[F, T]{}, err
	}

	if !quotient.IsInt64() {
		return Money[T]{}, ConversionResult[F, T]{}, ErrOverflow
	}

	roundedAmount := quotient.Int64()
	actualDenominator := m.amount
	if actualDenominator == 0 {
		actualDenominator = 1
	}

	actualRate := Ratio[F, T]{
		Numerator:   roundedAmount,
		Denominator: actualDenominator,
	}

	result := ConversionResult[F, T]{
		Amount:     roundedAmount,
		ActualRate: actualRate,
	}

	return NewMoney[T](roundedAmount), result, nil
}

func divideWithRounding(numerator, denominator *big.Int, mode RoundingMode) (*big.Int, error) {
	q := new(big.Int)
	r := new(big.Int)
	q.QuoRem(numerator, denominator, r)

	if r.Sign() == 0 || mode == RoundTruncate {
		return q, nil
	}

	absRem := new(big.Int).Abs(new(big.Int).Set(r))
	absDen := new(big.Int).Abs(new(big.Int).Set(denominator))
	twiceRem := new(big.Int).Lsh(absRem, 1)
	cmp := twiceRem.Cmp(absDen)

	step := big.NewInt(1)
	negativeResult := (numerator.Sign() < 0) != (denominator.Sign() < 0)
	if negativeResult {
		step.Neg(step)
	}

	switch mode {
	case RoundHalfUp:
		if cmp >= 0 {
			q.Add(q, step)
		}
	case RoundHalfEven:
		if cmp > 0 {
			q.Add(q, step)
		} else if cmp == 0 && q.Bit(0) == 1 {
			q.Add(q, step)
		}
	case RoundTruncate:
		// handled above
	default:
		return nil, ErrInvalidRoundingMode
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

	parts := make([]Money[T], len(ratios))
	remaining := m.amount

	// Allocate for all parts except the last one
	for i := range len(ratios) - 1 {
		share := big.NewInt(m.amount)
		share.Mul(share, big.NewInt(ratios[i]))
		share.Quo(share, big.NewInt(total))

		if !share.IsInt64() {
			return Allocation[T]{}, ErrOverflow
		}

		parts[i] = Money[T]{
			amount:   share.Int64(),
			Currency: m.Currency,
		}
		remaining -= parts[i].amount
	}

	// Last part gets the remaining amount to avoid rounding issues
	parts[len(ratios)-1] = Money[T]{
		amount:   remaining,
		Currency: m.Currency,
	}

	return Allocation[T]{
		Parts: parts,
		Total: m,
	}, nil
}

// MarshalJSON implements the json.Marshaler interface
// Serializes the amount as a string to preserve precision
func (m Money[T]) MarshalJSON() ([]byte, error) {
	type MoneyJSON struct {
		Amount   string `json:"amount"`
		Currency string `json:"currency"`
	}

	return json.Marshal(MoneyJSON{
		Amount:   fmt.Sprintf("%d", m.amount),
		Currency: m.Currency.Code(),
	})
}

// UnmarshalJSON implements the json.Unmarshaler interface
// Expects amount as a string to maintain precision
func (m *Money[T]) UnmarshalJSON(data []byte) error {
	var temp struct {
		Amount   string `json:"amount"`
		Currency string `json:"currency"`
	}

	if err := json.Unmarshal(data, &temp); err != nil {
		return fmt.Errorf("failed to unmarshal money: %w", err)
	}

	amount, err := parseIntAmount(temp.Amount)
	if err != nil {
		return err
	}

	var zeroCurrency T
	if zeroCurrency.Code() != temp.Currency {
		return fmt.Errorf(
			"currency mismatch: expected %s, got %s",
			zeroCurrency.Code(),
			temp.Currency,
		)
	}

	m.amount = amount
	m.Currency = zeroCurrency
	return nil
}

// ParseMoney parses a canonical decimal amount into Money using the currency's minor units.
// Supported format: optional sign (+/-), digits, optional decimal point and digits.
func ParseMoney[T currency.Currency](amount string) (Money[T], error) {
	if amount == "" {
		return Money[T]{}, fmt.Errorf("%w: empty amount", ErrInvalidAmountFormat)
	}

	negative := false
	if amount[0] == '+' || amount[0] == '-' {
		negative = amount[0] == '-'
		amount = amount[1:]
		if amount == "" {
			return Money[T]{}, fmt.Errorf("%w: sign without digits", ErrInvalidAmountFormat)
		}
	}

	parts := strings.Split(amount, ".")
	if len(parts) > 2 {
		return Money[T]{}, fmt.Errorf("%w: multiple decimal separators", ErrInvalidAmountFormat)
	}

	whole := parts[0]
	if whole == "" {
		return Money[T]{}, fmt.Errorf("%w: missing whole part", ErrInvalidAmountFormat)
	}
	if !allDigits(whole) {
		return Money[T]{}, fmt.Errorf("%w: invalid whole part", ErrInvalidAmountFormat)
	}

	fractional := ""
	if len(parts) == 2 {
		fractional = parts[1]
		if fractional == "" {
			return Money[T]{}, fmt.Errorf("%w: missing fractional part", ErrInvalidAmountFormat)
		}
		if !allDigits(fractional) {
			return Money[T]{}, fmt.Errorf("%w: invalid fractional part", ErrInvalidAmountFormat)
		}
	}

	var c T
	minorUnits := c.MinorUnits()
	if len(fractional) > minorUnits {
		return Money[T]{}, fmt.Errorf("%w: got %d fractional digits, max is %d", ErrScaleMismatch, len(fractional), minorUnits)
	}

	if len(fractional) < minorUnits {
		fractional = fractional + strings.Repeat("0", minorUnits-len(fractional))
	}

	magnitude := whole + fractional
	if magnitude == "" {
		return Money[T]{}, fmt.Errorf("%w: empty magnitude", ErrInvalidAmountFormat)
	}

	intVal, ok := new(big.Int).SetString(magnitude, 10)
	if !ok {
		return Money[T]{}, fmt.Errorf("%w: cannot parse magnitude", ErrInvalidAmountFormat)
	}
	if negative {
		intVal.Neg(intVal)
	}
	if !intVal.IsInt64() {
		return Money[T]{}, ErrOverflow
	}

	return NewMoney[T](intVal.Int64()), nil
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
func (m Money[T]) Value() (driver.Value, error) {
	b, err := m.MarshalJSON()
	if err != nil {
		return nil, err
	}
	return string(b), nil
}

// Scan implements sql.Scanner for database/sql.
func (m *Money[T]) Scan(value any) error {
	if value == nil {
		return fmt.Errorf("cannot scan NULL into Money")
	}

	switch v := value.(type) {
	case int64:
		var c T
		m.amount = v
		m.Currency = c
		return nil
	case []byte:
		return m.UnmarshalJSON(v)
	case string:
		return m.UnmarshalJSON([]byte(v))
	default:
		return fmt.Errorf("cannot scan type %T into Money", value)
	}
}
