package fulus

import (
	"fmt"
	"math"
	"math/big"

	"github.com/khatibomar/fulus/currency"
)

var ErrOutOfBoundTemplate = "money amount %d%s should be in interval [%d%s, %d%s]"
var ErrOverflow = fmt.Errorf("arithmetic operation would overflow")

type Money[T currency.Currency] struct {
	amount int64
	currency.Currency
}

// Distribution represents how to split money into chunks
type Distribution struct {
	SmallerChunkSize int64 // Size of the smaller chunks
	SmallerCount     int64 // Number of smaller chunks
	LargerChunkSize  int64 // Size of the larger chunks
	LargerCount      int64 // Number of larger chunks
}

// Ratio represents a fraction used for conversion rates
type Ratio struct {
	Numerator   int64
	Denominator int64
}

// ConversionResult holds both the converted amount and the actual ratio used
type ConversionResult struct {
	Amount     int64
	ActualRate Ratio
}

// NewMoney creates a new Money instance with the given amount and currency.
// The amount should be specified in the currency's smallest sub-unit
// (e.g., cents for USD, pence for GBP). For example:
// USD 10.50 should be passed as 1050
// EUR 5.99 should be passed as 599
func NewMoney[T currency.Currency](amount int64) *Money[T] {
	var c T
	return &Money[T]{
		amount:   amount,
		Currency: c,
	}
}

func (m *Money[T]) Add(other *Money[T]) error {
	if other.amount > 0 && m.amount > math.MaxInt64-other.amount {
		return ErrOverflow
	}
	if other.amount < 0 && m.amount < math.MinInt64-other.amount {
		return ErrOverflow
	}
	m.amount += other.amount
	return nil
}

func (m *Money[T]) Sub(other *Money[T]) error {
	if other.amount > 0 && m.amount < math.MinInt64+other.amount {
		return ErrOverflow
	}
	if other.amount < 0 && m.amount > math.MaxInt64+other.amount {
		return ErrOverflow
	}
	m.amount -= other.amount
	return nil
}

func (m *Money[T]) Mul(scale int64) error {
	if scale == 0 {
		m.amount = 0
		return nil
	}
	// prob not so performant
	if m.amount > math.MaxInt64/scale || m.amount < math.MinInt64/scale {
		return ErrOverflow
	}
	m.amount *= scale
	return nil
}

func (m *Money[T]) Validate(min, max int64) error {
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

func (m *Money[T]) Amount() int64 {
	return m.amount
}

func (m *Money[T]) String() string {
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
func (m *Money[T]) Distribute(chunks int64) (*Distribution, error) {
	if chunks <= 0 {
		return nil, fmt.Errorf("number of chunks must be positive")
	}

	if m.amount < 0 {
		return nil, fmt.Errorf("cannot distribute negative amount")
	}

	if m.amount == 0 {
		return &Distribution{
			SmallerChunkSize: 0,
			SmallerCount:     chunks,
			LargerChunkSize:  0,
			LargerCount:      0,
		}, nil
	}

	smallerSize := m.amount / chunks
	remainder := m.amount % chunks
	largerCount := remainder

	smallerCount := chunks - largerCount
	largerSize := smallerSize + 1

	return &Distribution{
		SmallerChunkSize: smallerSize,
		SmallerCount:     smallerCount,
		LargerChunkSize:  largerSize,
		LargerCount:      largerCount,
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
//	chf, result, _ := eur.Convert[currency.CHF](ratio)
//	// chf will be 107.20 CHF (10720 cents)
//	// result.ActualRate shows the actual conversion rate used after rounding
func Convert[T, U currency.Currency](m *Money[T], ratio Ratio) (*Money[U], *ConversionResult, error) {
	if ratio.Denominator == 0 {
		return nil, nil, fmt.Errorf("denominator cannot be zero")
	}

	theoretical := big.NewInt(m.amount)
	theoretical.Mul(theoretical, big.NewInt(ratio.Numerator))
	theoretical.Div(theoretical, big.NewInt(ratio.Denominator))

	roundedAmount := theoretical.Int64()

	actualRate := Ratio{
		Numerator:   roundedAmount,
		Denominator: m.amount,
	}

	result := &ConversionResult{
		Amount:     roundedAmount,
		ActualRate: actualRate,
	}

	return NewMoney[U](roundedAmount), result, nil
}
