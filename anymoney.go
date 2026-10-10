package fulus

import (
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/khatibomar/fulus/currency"
)

// AnyMoney is a monetary value with a currency that is known only at run time,
// for example a currency code from a database row or a request.
// Use As to get a type-safe Money[T] before arithmetic in a known currency.
// The zero value has no currency.
type AnyMoney struct {
	amount   int128
	currency currency.Currency
}

// NewAnyMoney creates an AnyMoney from an amount in minor units and a currency.
func NewAnyMoney(amount int64, c currency.Currency) AnyMoney {
	return AnyMoney{amount: int128FromInt64(amount), currency: c}
}

// NewAnyMoneyFromCode creates an AnyMoney from an amount in minor units and a registered currency code.
// Returns ErrUnknownCurrency if currency.Default() does not have the code.
func NewAnyMoneyFromCode(amount int64, code string) (AnyMoney, error) {
	c, ok := currency.ByCode(code)
	if !ok {
		return AnyMoney{}, fmt.Errorf("%w: %q", ErrUnknownCurrency, code)
	}
	return AnyMoney{amount: int128FromInt64(amount), currency: c}, nil
}

// ParseAnyMoney parses a canonical decimal amount such as "12.50" in the currency with the given code.
// The amount format is the same as for ParseMoney.
func ParseAnyMoney(amount, code string) (AnyMoney, error) {
	c, ok := currency.ByCode(code)
	if !ok {
		return AnyMoney{}, fmt.Errorf("%w: %q", ErrUnknownCurrency, code)
	}
	return NewAnyMoneyFromDecimal(amount, c)
}

// NewAnyMoneyFromDecimal parses a canonical decimal amount such as "12.50" in the currency c.
// The amount format is the same as for ParseMoney.
func NewAnyMoneyFromDecimal(amount string, c currency.Currency) (AnyMoney, error) {
	if c == nil {
		return AnyMoney{}, fmt.Errorf("%w: nil currency", ErrUnknownCurrency)
	}
	minor, err := parseDecimal(amount, c.MinorUnits())
	if err != nil {
		return AnyMoney{}, err
	}
	return AnyMoney{amount: minor, currency: c}, nil
}

// Any returns the Money value as an AnyMoney.
func (m Money[T]) Any() AnyMoney {
	return AnyMoney{amount: m.amount, currency: m.Currency()}
}

// As returns m as a Money[T].
// Returns ErrCurrencyMismatch if the currency code of m is not the code of T.
func As[T currency.Unit](m AnyMoney) (Money[T], error) {
	var c T
	if !m.sameCurrency(c) {
		return Money[T]{}, fmt.Errorf("%w: expected %s, got %s", ErrCurrencyMismatch, c.Code(), m.code())
	}
	return Money[T]{amount: m.amount}, nil
}

// Int64 returns the amount in minor units, and reports whether it fits in int64.
func (m AnyMoney) Int64() (int64, bool) {
	return m.amount.int64()
}

// BigInt returns the amount in minor units.
func (m AnyMoney) BigInt() *big.Int {
	return m.amount.big()
}

// Currency returns the currency, or nil for the zero value.
func (m AnyMoney) Currency() currency.Currency {
	return m.currency
}

// IsZero returns true if the amount is zero.
func (m AnyMoney) IsZero() bool {
	return m.amount.isZero()
}

// Add returns the sum of two values in the same currency.
// Returns ErrCurrencyMismatch if the currencies are different, and ErrOverflow if the sum does not fit.
func (m AnyMoney) Add(other AnyMoney) (AnyMoney, error) {
	return m.combine(other, add128)
}

// Sub returns the difference of two values in the same currency.
// Returns ErrCurrencyMismatch if the currencies are different, and ErrOverflow if the result does not fit.
func (m AnyMoney) Sub(other AnyMoney) (AnyMoney, error) {
	return m.combine(other, sub128)
}

func (m AnyMoney) combine(other AnyMoney, op func(a, b int128) (int128, bool)) (AnyMoney, error) {
	if !m.sameCurrency(other.currency) {
		return AnyMoney{}, fmt.Errorf("%w: %s and %s", ErrCurrencyMismatch, m.code(), other.code())
	}
	result, ok := op(m.amount, other.amount)
	if !ok {
		return AnyMoney{}, ErrOverflow
	}
	return AnyMoney{amount: result, currency: m.currency}, nil
}

// Cmp compares two values in the same currency and returns -1, 0 or +1.
// Returns ErrCurrencyMismatch if the currencies are different.
func (m AnyMoney) Cmp(other AnyMoney) (int, error) {
	if !m.sameCurrency(other.currency) {
		return 0, fmt.Errorf("%w: %s and %s", ErrCurrencyMismatch, m.code(), other.code())
	}
	return m.amount.cmp(other.amount), nil
}

// Decimal returns the amount as a canonical decimal string such as "-1234.50", like Money.Decimal.
// The zero value gives the amount in minor units.
func (m AnyMoney) Decimal() string {
	if m.currency == nil {
		return string(appendInt128(nil, m.amount))
	}
	return string(appendDecimal(nil, m.amount, m.currency.MinorUnits()))
}

// String returns the currency code and the canonical decimal, for example "EUR 12.50".
// The zero value gives the amount only.
func (m AnyMoney) String() string {
	if m.currency == nil {
		return string(appendInt128(nil, m.amount))
	}
	code := m.currency.Code()
	b := make([]byte, 0, len(code)+42)
	b = append(b, code...)
	b = append(b, ' ')
	return string(appendDecimal(b, m.amount, m.currency.MinorUnits()))
}

// MarshalJSON implements json.Marshaler with the same form as Money.
func (m AnyMoney) MarshalJSON() ([]byte, error) {
	if m.currency == nil {
		return marshalMoneyJSON(m.amount, 0, "")
	}
	return marshalMoneyJSON(m.amount, m.currency.MinorUnits(), m.currency.Code())
}

// UnmarshalJSON implements json.Unmarshaler with the same form as Money.
// The currency code must be in currency.Default(). For another registry, decode the code and the amount,
// and use NewAnyMoneyFromDecimal.
func (m *AnyMoney) UnmarshalJSON(data []byte) error {
	var temp moneyJSON
	if err := json.Unmarshal(data, &temp); err != nil {
		return fmt.Errorf("failed to unmarshal money: %w", err)
	}
	parsed, err := ParseAnyMoney(temp.Amount, temp.Currency)
	if err != nil {
		return err
	}
	*m = parsed
	return nil
}

func (m AnyMoney) code() string {
	if m.currency == nil {
		return ""
	}
	return m.currency.Code()
}

func (m AnyMoney) sameCurrency(c currency.Currency) bool {
	return m.currency != nil && c != nil && m.currency.Code() == c.Code()
}
