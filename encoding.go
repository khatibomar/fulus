package fulus

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"log/slog"
	"strings"

	"github.com/khatibomar/fulus/currency"
)

// MarshalJSON implements json.Marshaler.
// It writes the amount in minor units as a string, so that a JSON reader does not lose digits.
func (m Money[T]) MarshalJSON() ([]byte, error) {
	return marshalMoneyJSON(m.amount, m.Currency().Code())
}

// marshalMoneyJSON writes the JSON form of an amount without reflection
// when the currency code needs no escaping.
func marshalMoneyJSON(amount int128, code string) ([]byte, error) {
	for i := range len(code) {
		if c := code[i]; c < 0x20 || c > 0x7e || c == '"' || c == '\\' || c == '<' || c == '>' || c == '&' {
			return json.Marshal(moneyJSON{Amount: string(appendInt128(nil, amount)), Currency: code})
		}
	}

	b := make([]byte, 0, len(`{"amount":"","currency":""}`)+40+len(code))
	b = append(b, `{"amount":"`...)
	b = appendInt128(b, amount)
	b = append(b, `","currency":"`...)
	b = append(b, code...)
	return append(b, `"}`...), nil
}

type moneyJSON struct {
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
}

// UnmarshalJSON implements json.Unmarshaler. It reads the form that MarshalJSON writes.
// Returns ErrCurrencyMismatch if the currency code is not the code of T.
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
// It returns the amount in minor units as int64, for use with an integer column such as BIGINT.
// Returns ErrOverflow if the amount does not fit in int64.
func (m Money[T]) Value() (driver.Value, error) {
	v, ok := m.amount.int64()
	if !ok {
		return nil, ErrOverflow
	}
	return v, nil
}

// Scan implements sql.Scanner for database/sql.
// It accepts an int64 amount in minor units, the same amount as integer text,
// or the JSON form that MarshalJSON writes.
func (m *Money[T]) Scan(value any) error {
	switch v := value.(type) {
	case nil:
		return fmt.Errorf("cannot scan NULL into Money")
	case int64:
		m.amount = int128FromInt64(v)
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
	if amount, err := parseIntAmount(s); err == nil {
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
