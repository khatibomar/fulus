package forex

import (
	"fmt"
	"math"
)

var ErrOutOfBoundTemplate = "money amount %d%s should be in interval [%d%s, %d%s]"
var ErrOverflow = fmt.Errorf("arithmetic operation would overflow")

// Currency represents a monetary unit that follows ISO 4217 standards.
type Currency interface {
	// Code returns the three-letter ISO 4217 currency code
	// Example: "USD", "EUR", "LBP"
	Code() string

	// Symbol returns the currency symbol
	// Example: "$" for USD, "€" for EUR, "ل.ل" for LBP
	Symbol() string

	// MinorUnit returns the number of decimal places used for the currency
	// as defined in ISO 4217. For example:
	// 2 for USD (cents: 100)
	// 3 for BHD (fils: 1000)
	// 0 for JPY (no minor unit)
	MinorUnit() uint8

	// SubUnitSymbol returns the symbol or abbreviation for the minor unit of the currency
	// Example: "¢" for USD cents, "p" for GBP pence
	SubUnitSymbol() string
}

type Money[T Currency] struct {
	amount   int64
	Currency T
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
			m.amount, m.Currency.SubUnitSymbol(),
			min, m.Currency.SubUnitSymbol(),
			max, m.Currency.SubUnitSymbol(),
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
	if m.Currency.MinorUnit() == 0 {
		return fmt.Sprintf("%s%s%d", sign, m.Currency.Symbol(), amount)
	}

	// Calculate major and minor parts
	divisor := int64(math.Pow10(int(m.Currency.MinorUnit())))
	major := amount / divisor
	minor := amount % divisor

	format := "%s%s%d.%0" + fmt.Sprintf("%d", m.Currency.MinorUnit()) + "d"
	return fmt.Sprintf(format, sign, m.Currency.Symbol(), major, minor)
}
