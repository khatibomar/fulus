package fulus

import (
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"math/big"

	"github.com/khatibomar/fulus/currency"
)

var (
	// ErrOutOfBoundTemplate is the template for out of bounds errors
	ErrOutOfBoundTemplate = "money amount %d%s should be in interval [%d%s, %d%s]"

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
type Ratio struct {
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
type ConversionResult struct {
	// Amount stores the resulting converted monetary value
	Amount int64
	// ActualRate stores the precise conversion rate that was actually applied
	ActualRate Ratio
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
		return fmt.Errorf(
			ErrOutOfBoundTemplate,
			m.amount, m.Currency.MinorUnitSymbol(),
			min, m.Currency.MinorUnitSymbol(),
			max, m.Currency.MinorUnitSymbol(),
		)
	}
	return nil
}

// Amount returns the internal amount value in the currency's smallest unit.
// For example, returns cents for USD or pence for GBP.
func (m Money[T]) Amount() int64 {
	return m.amount
}

// String returns a formatted string representation of the Money value.
// The format includes the currency symbol, sign (if negative), and proper decimal placement.
//
// For currencies with minor units (e.g., USD, EUR):
//   - Positive: "$10.50" (USD), "€10.50" (EUR)
//   - Negative: "-$10.50" (USD), "-€10.50" (EUR)
//   - Zero: "$0" (USD), "€0" (EUR)
//
// For currencies without minor units (e.g., JPY):
//   - Positive: "¥1000"
//   - Negative: "-¥1000"
//   - Zero: "¥0"
//
// The number of decimal places is determined by the currency's MinorUnits() value.
func (m Money[T]) String() string {
	if m.amount == 0 {
		return fmt.Sprintf("%s0", m.Currency.Symbol())
	}

	sign := ""
	amount := m.amount
	if amount < 0 {
		sign = "-"
		amount = -amount
	}

	// If there's no minor unit (like JPY), just return the major amount
	if m.Currency.MinorUnits() == 0 {
		return fmt.Sprintf("%s%s%d", sign, m.Currency.Symbol(), amount)
	}

	// Calculate major and minor parts
	divisor := int64(math.Pow10(int(m.Currency.MinorUnits())))
	major := amount / divisor
	minor := amount % divisor

	format := "%s%s%d.%0" + fmt.Sprintf("%d", m.Currency.MinorUnits()) + "d"
	return fmt.Sprintf(format, sign, m.Currency.Symbol(), major, minor)
}

// Distribute splits the money amount into the specified number of chunks
// Returns a Distribution describing how to split the money
//
// Example:
//
//	money := NewMoney[currency.USD](1000) // $10.00
//	dist, _ := money.Distribute(3)
//	// dist will contain:
//	// SmallerChunkSize: 333 ($3.33) x 2 chunks
//	// LargerChunkSize:  334 ($3.34) x 1 chunk
//	// Total: (333 * 2) + (334 * 1) = 1000
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

	return Distribution{
		SmallerChunkSize: smallerChunkSize,
		SmallerCount:     chunks - remainder,
		LargerChunkSize:  largerChunkSize,
		LargerCount:      remainder,
	}, nil
}

// Convert performs a conversion using a ratio while maintaining accuracy
// The ratio should be provided as (numerator, denominator) representing numerator/denominator
// Returns both the converted Money value and the actual ratio used after rounding
//
// Example:
//
//	// Convert 100 EUR to CHF at rate 1.072032 CHF/EUR
//	eur := NewMoney[currency.EUR](10000) // 100.00 EUR
//	ratio := Ratio{
//		Numerator:   107203,  // 1.07203 represented as 107203/100000
//		Denominator: 100000,
//	}
//	chf, result, _ := fulus.Convert[currency.CHF](eur, ratio)
//	// chf will be 107.20 CHF (10720 cents)
//	// result.ActualRate shows the actual conversion rate used after rounding
func Convert[U, T currency.Currency](m Money[T], ratio Ratio) (Money[U], ConversionResult, error) {
	if ratio.Denominator == 0 {
		return Money[U]{}, ConversionResult{}, ErrZeroDenominator
	}

	theoretical := big.NewInt(m.amount)
	theoretical.Mul(theoretical, big.NewInt(ratio.Numerator))
	theoretical.Div(theoretical, big.NewInt(ratio.Denominator))

	if !theoretical.IsInt64() {
		return Money[U]{}, ConversionResult{}, ErrOverflow
	}

	roundedAmount := theoretical.Int64()

	actualRate := Ratio{
		Numerator:   roundedAmount,
		Denominator: m.amount,
	}

	result := ConversionResult{
		Amount:     roundedAmount,
		ActualRate: actualRate,
	}

	return NewMoney[U](roundedAmount), result, nil
}

// Allocate divides money according to provided ratios
// Example: $100 allocated by ratios [1,1,2] would give [$25,$25,$50]
func (m Money[T]) Allocate(ratios []int64) (Allocation[T], error) {
	// Validate ratios
	if len(ratios) == 0 {
		return Allocation[T]{}, ErrNoRatios
	}

	total := int64(0)
	for _, ratio := range ratios {
		if ratio <= 0 {
			return Allocation[T]{}, ErrNegativeOrZeroRatios
		}
		total += ratio
	}

	parts := make([]Money[T], len(ratios))
	remaining := m.amount

	// Allocate for all parts except the last one
	for i := 0; i < len(ratios)-1; i++ {
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

	var amount int64
	if _, err := fmt.Sscanf(temp.Amount, "%d", &amount); err != nil {
		return fmt.Errorf("invalid amount format: %w", err)
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
