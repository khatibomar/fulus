package fulus

import (
	"database/sql/driver"
	"encoding/json"

	"github.com/khatibomar/fulus/currency"
)

// NullMoney is a Money value that can be NULL in a database or null in JSON.
// It works like sql.NullInt64.
type NullMoney[T currency.Unit] struct {
	Money Money[T]
	// Valid is true if Money is not NULL.
	Valid bool
}

// Scan implements sql.Scanner. A NULL value sets Valid to false.
func (n *NullMoney[T]) Scan(value any) error {
	if value == nil {
		*n = NullMoney[T]{}
		return nil
	}
	if err := n.Money.Scan(value); err != nil {
		n.Valid = false
		return err
	}
	n.Valid = true
	return nil
}

// Value implements driver.Valuer. It returns nil if Valid is false.
func (n NullMoney[T]) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return n.Money.Value()
}

// MarshalJSON implements json.Marshaler. It writes null if Valid is false.
func (n NullMoney[T]) MarshalJSON() ([]byte, error) {
	if !n.Valid {
		return []byte("null"), nil
	}
	return n.Money.MarshalJSON()
}

// UnmarshalJSON implements json.Unmarshaler. A JSON null sets Valid to false.
func (n *NullMoney[T]) UnmarshalJSON(data []byte) error {
	if isJSONNull(data) {
		*n = NullMoney[T]{}
		return nil
	}
	if err := json.Unmarshal(data, &n.Money); err != nil {
		n.Valid = false
		return err
	}
	n.Valid = true
	return nil
}
