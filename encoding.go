package fulus

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"log/slog"
	"strings"

	"github.com/khatibomar/fulus/currency"
)

// MarshalJSON implements json.Marshaler. It writes the form {"amount":"10.50","currency":"USD"}.
// The amount is the canonical decimal from Decimal, in a string, so that a JSON reader does not lose digits.
// The JSON does not depend on the minor units of the reader, so it stays correct if ISO 4217 changes them.
func (m Money[T]) MarshalJSON() ([]byte, error) {
	c := m.Currency()
	return marshalMoneyJSON(m.amount, c.MinorUnits(), c.Code())
}

// marshalMoneyJSON writes the JSON form of an amount without reflection
// when the currency code needs no escaping.
func marshalMoneyJSON(amount int128, minorUnits int, code string) ([]byte, error) {
	for i := range len(code) {
		if c := code[i]; c < 0x20 || c > 0x7e || c == '"' || c == '\\' || c == '<' || c == '>' || c == '&' {
			return json.Marshal(moneyJSON{Amount: string(appendDecimal(nil, amount, minorUnits)), Currency: code})
		}
	}

	b := make([]byte, 0, len(`{"amount":"","currency":""}`)+42+len(code))
	b = append(b, `{"amount":"`...)
	b = appendDecimal(b, amount, minorUnits)
	b = append(b, `","currency":"`...)
	b = append(b, code...)
	return append(b, `"}`...), nil
}

type moneyJSON struct {
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
}

// UnmarshalJSON implements json.Unmarshaler. It reads the form that MarshalJSON writes.
// The amount can have fewer fraction digits than the minor units, but not more.
// Returns ErrCurrencyMismatch if the currency code is not the code of T,
// and ErrScaleMismatch if the amount has more fraction digits than the minor units.
func (m *Money[T]) UnmarshalJSON(data []byte) error {
	var temp moneyJSON
	if err := json.Unmarshal(data, &temp); err != nil {
		return fmt.Errorf("failed to unmarshal money: %w", err)
	}

	c := m.Currency()
	if c.Code() != temp.Currency {
		return fmt.Errorf("%w: expected %s, got %s", ErrCurrencyMismatch, c.Code(), temp.Currency)
	}

	amount, err := parseDecimal(temp.Amount, c.MinorUnits())
	if err != nil {
		return err
	}
	m.amount = amount
	return nil
}

// ParseMoney parses a canonical decimal amount into Money using the currency's minor units.
// Supported format: optional sign (+/-), digits, optional decimal point and digits.
func ParseMoney[T currency.Unit](amount string) (Money[T], error) {
	var c T
	minor, err := parseDecimal(amount, c.MinorUnits())
	if err != nil {
		return Money[T]{}, err
	}
	return Money[T]{amount: minor}, nil
}

// parseDecimal parses a canonical decimal amount into minor units.
func parseDecimal(amount string, minorUnits int) (int128, error) {
	if amount == "" {
		return int128{}, fmt.Errorf("%w: empty amount", ErrInvalidAmountFormat)
	}

	negative := false
	if amount[0] == '+' || amount[0] == '-' {
		negative = amount[0] == '-'
		amount = amount[1:]
		if amount == "" {
			return int128{}, fmt.Errorf("%w: sign without digits", ErrInvalidAmountFormat)
		}
	}

	whole, fractional, hasPoint := strings.Cut(amount, ".")
	if hasPoint && strings.Contains(fractional, ".") {
		return int128{}, fmt.Errorf("%w: multiple decimal separators", ErrInvalidAmountFormat)
	}
	if whole == "" {
		return int128{}, fmt.Errorf("%w: missing whole part", ErrInvalidAmountFormat)
	}
	if !allDigits(whole) {
		return int128{}, fmt.Errorf("%w: invalid whole part", ErrInvalidAmountFormat)
	}
	if hasPoint && fractional == "" {
		return int128{}, fmt.Errorf("%w: missing fractional part", ErrInvalidAmountFormat)
	}
	if !allDigits(fractional) {
		return int128{}, fmt.Errorf("%w: invalid fractional part", ErrInvalidAmountFormat)
	}
	if len(fractional) > minorUnits {
		return int128{}, fmt.Errorf("%w: got %d fractional digits, max is %d", ErrScaleMismatch, len(fractional), minorUnits)
	}

	var magnitude uint128
	push := func(digit byte) bool {
		var ok bool
		magnitude, ok = magnitude.mulAdd10(uint64(digit - '0'))
		return ok
	}
	for i := range len(whole) {
		if !push(whole[i]) {
			return int128{}, ErrOverflow
		}
	}
	for i := range len(fractional) {
		if !push(fractional[i]) {
			return int128{}, ErrOverflow
		}
	}
	for range minorUnits - len(fractional) {
		if !push('0') {
			return int128{}, ErrOverflow
		}
	}

	result, ok := fromMagnitude128(magnitude, negative)
	if !ok {
		return int128{}, ErrOverflow
	}
	return result, nil
}

// parseIntAmount parses an integer amount in minor units, with an optional sign.
func parseIntAmount(amount string) (int128, error) {
	digits := strings.TrimLeft(amount, "+-")
	if digits == "" || len(amount)-len(digits) > 1 || !allDigits(digits) {
		return int128{}, fmt.Errorf("%w: %q is not an integer", ErrInvalidAmountFormat, amount)
	}
	return parseDecimal(amount, 0)
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
// It returns the canonical decimal from Decimal, such as "10.50", for a decimal column such as NUMERIC(38, 2).
// Use BigintMoney for an integer column that holds minor units.
func (m Money[T]) Value() (driver.Value, error) {
	return m.Decimal(), nil
}

// Scan implements sql.Scanner for database/sql. It reads a decimal column such as NUMERIC.
// It accepts decimal text, as drivers return NUMERIC values, and an int64 as a whole number of major units.
// It does not accept float64, because a float64 is not exact.
// Returns ErrScaleMismatch if the value has more fraction digits than the minor units.
// Use BigintMoney for an integer column that holds minor units.
func (m *Money[T]) Scan(value any) error {
	var c T
	switch v := value.(type) {
	case nil:
		return fmt.Errorf("%w: cannot scan NULL into Money, use NullMoney", ErrInvalidAmountFormat)
	case int64:
		amount, ok := int128FromInt64(v).mulPow10(c.MinorUnits())
		if !ok {
			return ErrOverflow
		}
		m.amount = amount
		return nil
	case []byte:
		return m.scanDecimal(string(v))
	case string:
		return m.scanDecimal(v)
	default:
		return fmt.Errorf("%w: cannot scan %T into Money, use a NUMERIC column or BigintMoney", ErrInvalidAmountFormat, value)
	}
}

func (m *Money[T]) scanDecimal(s string) error {
	var c T
	amount, err := parseDecimal(s, c.MinorUnits())
	if err != nil {
		return err
	}
	m.amount = amount
	return nil
}

// BigintMoney stores a Money value in an integer column, such as BIGINT, as an amount in minor units.
// Use it only for a column that holds minor units. Money itself stores a decimal for a NUMERIC column.
// Use sql.Null[BigintMoney[T]] for a column that can be NULL.
type BigintMoney[T currency.Unit] struct {
	Money Money[T]
}

// Value implements driver.Valuer. It returns the amount in minor units as int64.
// Returns ErrOverflow if the amount does not fit in int64.
func (b BigintMoney[T]) Value() (driver.Value, error) {
	v, ok := b.Money.Int64()
	if !ok {
		return nil, ErrOverflow
	}
	return v, nil
}

// Scan implements sql.Scanner. It accepts an amount in minor units as int64 or as integer text.
func (b *BigintMoney[T]) Scan(value any) error {
	switch v := value.(type) {
	case int64:
		b.Money = Money[T]{amount: int128FromInt64(v)}
		return nil
	case []byte:
		return b.scanText(string(v))
	case string:
		return b.scanText(v)
	case nil:
		return fmt.Errorf("%w: cannot scan NULL into BigintMoney, use sql.Null", ErrInvalidAmountFormat)
	default:
		return fmt.Errorf("%w: cannot scan %T into BigintMoney", ErrInvalidAmountFormat, value)
	}
}

func (b *BigintMoney[T]) scanText(s string) error {
	amount, err := parseIntAmount(s)
	if err != nil {
		return err
	}
	b.Money = Money[T]{amount: amount}
	return nil
}

// Decimal returns the amount as a canonical decimal string such as "-1234.50".
// It has all the minor units, no symbol and no group separator, and ParseMoney accepts it.
func (m Money[T]) Decimal() string {
	return string(appendDecimal(nil, m.amount, m.Currency().MinorUnits()))
}

// appendDecimal appends the canonical decimal form of an amount in minor units, such as "-1234.50".
func appendDecimal(b []byte, amount int128, minorUnits int) []byte {
	if amount.isNeg() {
		b = append(b, '-')
	}
	var buf [40]byte
	digits := amount.abs().appendDecimal(buf[:0])
	if minorUnits <= 0 {
		return append(b, digits...)
	}
	if len(digits) <= minorUnits {
		b = append(b, '0', '.')
		for range minorUnits - len(digits) {
			b = append(b, '0')
		}
		return append(b, digits...)
	}
	split := len(digits) - minorUnits
	b = append(b, digits[:split]...)
	b = append(b, '.')
	return append(b, digits[split:]...)
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
