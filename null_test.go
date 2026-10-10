package fulus

import (
	"encoding/json"
	"testing"

	"github.com/khatibomar/fulus/currency"
)

func TestNullMoneySQL(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		value     any
		wantValid bool
		want      int64
		wantErr   bool
	}{
		{name: "NULL", value: nil, wantValid: false},
		{name: "int64", value: int64(1050), wantValid: true, want: 1050},
		{name: "text", value: []byte("-3"), wantValid: true, want: -3},
		{name: "invalid", value: 1.5, wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			n := NullMoney[currency.USD]{Money: NewMoney[currency.USD](99), Valid: true}
			err := n.Scan(tt.value)
			if (err != nil) != tt.wantErr {
				t.Fatalf("Scan() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err != nil {
				if n.Valid {
					t.Error("Valid must be false after a failed Scan")
				}
				return
			}
			if n.Valid != tt.wantValid || n.Money.amount64() != tt.want {
				t.Errorf("Scan() = %+v, want valid %v amount %d", n, tt.wantValid, tt.want)
			}

			v, err := n.Value()
			if err != nil {
				t.Fatalf("Value() error = %v", err)
			}
			if !tt.wantValid && v != nil {
				t.Errorf("Value() = %v, want nil", v)
			}
			if tt.wantValid && v != tt.want {
				t.Errorf("Value() = %v, want %d", v, tt.want)
			}
		})
	}
}

func TestNullMoneyJSON(t *testing.T) {
	t.Parallel()

	type payload struct {
		Price NullMoney[currency.USD] `json:"price"`
	}

	tests := []struct {
		name  string
		input payload
		json  string
	}{
		{name: "null", input: payload{}, json: `{"price":null}`},
		{
			name:  "valid",
			input: payload{Price: NullMoney[currency.USD]{Money: NewMoney[currency.USD](1050), Valid: true}},
			json:  `{"price":{"amount":"1050","currency":"USD"}}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			b, err := json.Marshal(tt.input)
			if err != nil {
				t.Fatalf("Marshal() error = %v", err)
			}
			if string(b) != tt.json {
				t.Fatalf("Marshal() = %s, want %s", b, tt.json)
			}

			var got payload
			got.Price.Valid = true
			if err := json.Unmarshal(b, &got); err != nil {
				t.Fatalf("Unmarshal() error = %v", err)
			}
			if got != tt.input {
				t.Errorf("Unmarshal() = %+v, want %+v", got, tt.input)
			}
		})
	}

	var bad NullMoney[currency.USD]
	if err := json.Unmarshal([]byte(`{"amount":"1","currency":"EUR"}`), &bad); err == nil || bad.Valid {
		t.Errorf("Unmarshal() mismatch = %v, valid %v; want error and not valid", err, bad.Valid)
	}
}
