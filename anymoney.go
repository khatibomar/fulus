package fulus

import (
	"cmp"
	"encoding/json"
	"fmt"
	"math/big"
	"strconv"

	"github.com/khatibomar/fulus/currency"
	"github.com/khatibomar/fulus/locale"
)

// AnyMoney is a monetary value with a currency that is known only at run time,
// for example a currency code from a database row or a request.
// Use As to get a type-safe Money[T] before arithmetic in a known currency.
// The zero value has no currency.
type AnyMoney struct {
	amount   int64
	currency currency.Currency
}

// NewAnyMoney creates an AnyMoney from an amount in minor units and a currency.
func NewAnyMoney(amount int64, c currency.Currency) AnyMoney {
	return AnyMoney{amount: amount, currency: c}
}

// NewAnyMoneyFromCode creates an AnyMoney from an amount in minor units and a registered currency code.
// Returns ErrUnknownCurrency if currency.ByCode does not find the code.
func NewAnyMoneyFromCode(amount int64, code string) (AnyMoney, error) {
	c, ok := currency.ByCode(code)
	if !ok {
		return AnyMoney{}, fmt.Errorf("%w: %q", ErrUnknownCurrency, code)
	}
	return AnyMoney{amount: amount, currency: c}, nil
}

// ParseAnyMoney parses a canonical decimal amount such as "12.50" in the currency with the given code.
// The amount format is the same as for ParseMoney.
func ParseAnyMoney(amount, code string) (AnyMoney, error) {
	c, ok := currency.ByCode(code)
	if !ok {
		return AnyMoney{}, fmt.Errorf("%w: %q", ErrUnknownCurrency, code)
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

// Amount returns the amount in minor units.
func (m AnyMoney) Amount() int64 {
	return m.amount
}

// Currency returns the currency, or nil for the zero value.
func (m AnyMoney) Currency() currency.Currency {
	return m.currency
}

// IsZero returns true if the amount is zero.
func (m AnyMoney) IsZero() bool {
	return m.amount == 0
}

// Add returns the sum of two values in the same currency.
// Returns ErrCurrencyMismatch if the currencies are different, and ErrOverflow if the sum does not fit in int64.
func (m AnyMoney) Add(other AnyMoney) (AnyMoney, error) {
	return m.combine(other, (*big.Int).Add)
}

// Sub returns the difference of two values in the same currency.
// Returns ErrCurrencyMismatch if the currencies are different, and ErrOverflow if the result does not fit in int64.
func (m AnyMoney) Sub(other AnyMoney) (AnyMoney, error) {
	return m.combine(other, (*big.Int).Sub)
}

func (m AnyMoney) combine(other AnyMoney, op func(z, x, y *big.Int) *big.Int) (AnyMoney, error) {
	if !m.sameCurrency(other.currency) {
		return AnyMoney{}, fmt.Errorf("%w: %s and %s", ErrCurrencyMismatch, m.code(), other.code())
	}
	result := op(new(big.Int), big.NewInt(m.amount), big.NewInt(other.amount))
	if !result.IsInt64() {
		return AnyMoney{}, ErrOverflow
	}
	return AnyMoney{amount: result.Int64(), currency: m.currency}, nil
}

// Cmp compares two values in the same currency and returns -1, 0 or +1.
// Returns ErrCurrencyMismatch if the currencies are different.
func (m AnyMoney) Cmp(other AnyMoney) (int, error) {
	if !m.sameCurrency(other.currency) {
		return 0, fmt.Errorf("%w: %s and %s", ErrCurrencyMismatch, m.code(), other.code())
	}
	return cmp.Compare(m.amount, other.amount), nil
}

// Format returns the value formatted for the locale. The zero value formats as a plain integer.
func (m AnyMoney) Format(loc locale.Locale) string {
	if m.currency == nil {
		return strconv.FormatInt(m.amount, 10)
	}
	return formatAmount(m.amount, m.currency.MinorUnits(), m.currency.FormatInfo(loc))
}

// String returns the value formatted with DefaultLocale.
func (m AnyMoney) String() string {
	return m.Format(DefaultLocale())
}

// MarshalJSON implements json.Marshaler with the same form as Money.
func (m AnyMoney) MarshalJSON() ([]byte, error) {
	return marshalMoneyJSON(m.amount, m.code())
}

// UnmarshalJSON implements json.Unmarshaler with the same form as Money.
// The currency code must be registered in the currency package.
func (m *AnyMoney) UnmarshalJSON(data []byte) error {
	var temp moneyJSON
	if err := json.Unmarshal(data, &temp); err != nil {
		return fmt.Errorf("failed to unmarshal money: %w", err)
	}
	amount, err := parseIntAmount(temp.Amount)
	if err != nil {
		return err
	}
	parsed, err := NewAnyMoneyFromCode(amount, temp.Currency)
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
