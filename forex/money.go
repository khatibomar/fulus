package forex

import (
	"fmt"
	"math"
)

var ErrOutOfBoundTemplate = "money amount should be in interval [%d, %d]"
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
	currency T
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
		return fmt.Errorf(ErrOutOfBoundTemplate, min, max)
	}
	return nil
}

func (m *Money[T]) Amount() int64 {
	return m.amount
}

func (m *Money[T]) String() string {
	if m.amount == 0 {
		return fmt.Sprintf("%s0", m.currency.Symbol())
	}

	sign := ""
	amount := m.amount
	if amount < 0 {
		sign = "-"
		amount = -amount
	}

	// If there's no minor unit (like JPY), just return the major amount
	if m.currency.MinorUnit() == 0 {
		return fmt.Sprintf("%s%s%d", sign, m.currency.Symbol(), amount)
	}

	// Calculate major and minor parts
	divisor := int64(math.Pow10(int(m.currency.MinorUnit())))
	major := amount / divisor
	minor := amount % divisor

	format := "%s%s%d.%0" + fmt.Sprintf("%d", m.currency.MinorUnit()) + "d"
	return fmt.Sprintf(format, sign, m.currency.Symbol(), major, minor)
}
