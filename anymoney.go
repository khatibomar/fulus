package fulus

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"math/big"
	"strconv"

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
// Returns ErrCurrencyMismatch if the currency of m does not have the code and the minor units of T.
func As[T currency.Unit](m AnyMoney) (Money[T], error) {
	var c T
	if !m.sameCurrency(c) {
		return Money[T]{}, fmt.Errorf("%w: expected %s, got %s", ErrCurrencyMismatch, describe(c, m.currency), describe(m.currency, c))
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

// IsPositive reports whether m > 0.
func (m AnyMoney) IsPositive() bool {
	return m.amount.sign() > 0
}

// IsNegative reports whether m < 0.
func (m AnyMoney) IsNegative() bool {
	return m.amount.isNeg()
}

// Sign returns -1 if m < 0, 0 if m is zero, and +1 if m > 0.
func (m AnyMoney) Sign() int {
	return m.amount.sign()
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
		return AnyMoney{}, fmt.Errorf("%w: %s and %s", ErrCurrencyMismatch, describe(m.currency, other.currency), describe(other.currency, m.currency))
	}
	result, ok := op(m.amount, other.amount)
	if !ok {
		return AnyMoney{}, ErrOverflow
	}
	return AnyMoney{amount: result, currency: m.currency}, nil
}

// Mul returns m * scale.
// Returns ErrOverflow if the result does not fit.
func (m AnyMoney) Mul(scale int64) (AnyMoney, error) {
	result, ok := m.amount.mulInt64(scale)
	if !ok {
		return AnyMoney{}, ErrOverflow
	}
	return AnyMoney{amount: result, currency: m.currency}, nil
}

// Div returns m / divisor, rounded with mode, like Money.Div.
func (m AnyMoney) Div(divisor int64, mode RoundingMode) (AnyMoney, error) {
	if divisor == 0 {
		return AnyMoney{}, ErrDivisionByZero
	}
	return m.scale(1, divisor, mode)
}

// MulFactor multiplies m by f and rounds the result with mode, like Money.MulFactor.
func (m AnyMoney) MulFactor(f Factor, mode RoundingMode) (AnyMoney, error) {
	if f.den == 0 {
		return AnyMoney{}, ErrInvalidFactor
	}
	return m.scale(f.num, f.den, mode)
}

func (m AnyMoney) scale(numerator, denominator int64, mode RoundingMode) (AnyMoney, error) {
	if !mode.valid() {
		return AnyMoney{}, ErrInvalidRoundingMode
	}
	result, err := mulDivRound(m.amount, numerator, denominator, mode)
	if err != nil {
		return AnyMoney{}, err
	}
	return AnyMoney{amount: result, currency: m.currency}, nil
}

// RoundCash rounds m to the smallest cash amount of its currency, like Money.RoundCash.
// The zero value does not change.
func (m AnyMoney) RoundCash(mode RoundingMode) (AnyMoney, error) {
	result, err := roundCash(m.amount, m.currency, mode)
	if err != nil {
		return AnyMoney{}, err
	}
	return AnyMoney{amount: result, currency: m.currency}, nil
}

// Abs returns the absolute value of m, like Money.Abs.
func (m AnyMoney) Abs() (AnyMoney, error) {
	if !m.amount.isNeg() {
		return m, nil
	}
	return m.Neg()
}

// Neg returns -m, like Money.Neg.
func (m AnyMoney) Neg() (AnyMoney, error) {
	result, ok := m.amount.neg()
	if !ok {
		return AnyMoney{}, ErrOverflow
	}
	return AnyMoney{amount: result, currency: m.currency}, nil
}

// Allocate divides m into parts in proportion to the ratios, like Money.Allocate.
func (m AnyMoney) Allocate(ratios ...int64) ([]AnyMoney, error) {
	amounts, err := allocate[struct{ amount int128 }](m.amount, ratios)
	if err != nil {
		return nil, err
	}
	parts := make([]AnyMoney, len(amounts))
	for i, a := range amounts {
		parts[i] = AnyMoney{amount: a.amount, currency: m.currency}
	}
	return parts, nil
}

// Cmp compares two values in the same currency and returns -1, 0 or +1.
// Returns ErrCurrencyMismatch if the currencies are different.
func (m AnyMoney) Cmp(other AnyMoney) (int, error) {
	if !m.sameCurrency(other.currency) {
		return 0, fmt.Errorf("%w: %s and %s", ErrCurrencyMismatch, describe(m.currency, other.currency), describe(other.currency, m.currency))
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
// The zero value gives null, which UnmarshalJSON reads back as the zero value.
// Returns ErrUnknownCurrency for an amount that is not zero and has no currency.
func (m AnyMoney) MarshalJSON() ([]byte, error) {
	if m.currency == nil {
		if !m.amount.isZero() {
			return nil, fmt.Errorf("%w: amount %s has no currency", ErrUnknownCurrency, m.Decimal())
		}
		return []byte("null"), nil
	}
	return marshalMoneyJSON(m.amount, m.currency.MinorUnits(), m.currency.Code())
}

// UnmarshalJSON implements json.Unmarshaler with the same form as Money. A JSON null does not change m.
// The currency code must be in currency.Default(). For another registry, decode the code and the amount,
// and use NewAnyMoneyFromDecimal.
func (m *AnyMoney) UnmarshalJSON(data []byte) error {
	if isJSONNull(data) {
		return nil
	}
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

// sameCurrency reports whether c has the code and the minor units of the currency of m.
// The minor units must agree, because the amount is in minor units.
func (m AnyMoney) sameCurrency(c currency.Currency) bool {
	return m.currency != nil && c != nil && m.currency.Code() == c.Code() && m.currency.MinorUnits() == c.MinorUnits()
}

// describe returns the code of c for an error message.
// It adds the minor units, for example "USD with 2 minor units", when other has the same code.
func describe(c, other currency.Currency) string {
	if c == nil {
		return "no currency"
	}
	if other != nil && other.Code() == c.Code() {
		return c.Code() + " with " + strconv.Itoa(c.MinorUnits()) + " minor units"
	}
	return c.Code()
}

// ScanColumns returns the sql.Scanner destinations that read m from two columns: a decimal amount column,
// such as NUMERIC, and a currency code column. Pass them to Rows.Scan in the order of the columns.
// The amount column accepts the same values as Money.Scan. The code must be in currency.Default().
// m changes only after both columns are scanned without an error. You can use the destinations again for the next row.
// To write m, use Decimal for the amount column and the Code of Currency for the code column.
func (m *AnyMoney) ScanColumns() (amount, code sql.Scanner) {
	s := &anyMoneyScan{m: m}
	return anyAmountColumn{s}, anyCodeColumn{s}
}

// anyMoneyScan holds the values of the two columns until both are scanned.
type anyMoneyScan struct {
	m                  *AnyMoney
	amount             any
	code               string
	hasAmount, hasCode bool
}

type anyAmountColumn struct{ s *anyMoneyScan }

type anyCodeColumn struct{ s *anyMoneyScan }

// Scan implements sql.Scanner for the amount column.
func (c anyAmountColumn) Scan(value any) error {
	if b, ok := value.([]byte); ok {
		// A driver can use the bytes again after Scan returns.
		value = string(b)
	}
	c.s.amount, c.s.hasAmount = value, true
	return c.s.finish()
}

// Scan implements sql.Scanner for the code column.
func (c anyCodeColumn) Scan(value any) error {
	switch v := value.(type) {
	case string:
		c.s.code = v
	case []byte:
		c.s.code = string(v)
	default:
		c.s.reset()
		return fmt.Errorf("%w: cannot scan %T into a currency code", ErrUnknownCurrency, value)
	}
	c.s.hasCode = true
	return c.s.finish()
}

// finish sets m when both columns are scanned.
func (s *anyMoneyScan) finish() error {
	if !s.hasAmount || !s.hasCode {
		return nil
	}
	defer s.reset()

	if s.amount == nil {
		return fmt.Errorf("%w: cannot scan NULL into AnyMoney", ErrInvalidAmountFormat)
	}
	c, ok := currency.ByCode(s.code)
	if !ok {
		return fmt.Errorf("%w: %q", ErrUnknownCurrency, s.code)
	}
	amount, err := scanDecimal(s.amount, c.MinorUnits())
	if err != nil {
		return err
	}
	*s.m = AnyMoney{amount: amount, currency: c}
	return nil
}

func (s *anyMoneyScan) reset() {
	s.amount, s.code, s.hasAmount, s.hasCode = nil, "", false, false
}
