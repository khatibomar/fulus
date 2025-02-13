package fulus

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/khatibomar/fulus/currency"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		name     string
		a, b     int64
		want     int64
		wantErr  bool
		errCheck func(error) bool
	}{
		{"simple addition", 100, 200, 300, false, nil},
		{"zero addition", 100, 0, 100, false, nil},
		{"negative addition", 100, -50, 50, false, nil},
		{"overflow", math.MaxInt64, 1, 0, true, func(err error) bool { return err == ErrOverflow }},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m1 := NewMoney[currency.USD](tt.a)
			m2 := NewMoney[currency.USD](tt.b)
			err := m1.Add(m2)

			if tt.wantErr {
				if err == nil || (tt.errCheck != nil && !tt.errCheck(err)) {
					t.Errorf("Add() error = %v, wantErr %v", err, tt.wantErr)
				}
				return
			}

			if err != nil {
				t.Errorf("Add() unexpected error: %v", err)
			}

			if m1.Amount() != tt.want {
				t.Errorf("Add() = %v, want %v", m1.Amount(), tt.want)
			}
		})
	}
}

func TestSub(t *testing.T) {
	tests := []struct {
		name     string
		a, b     int64
		want     int64
		wantErr  bool
		errCheck func(error) bool
	}{
		{"simple subtraction", 200, 100, 100, false, nil},
		{"zero subtraction", 100, 0, 100, false, nil},
		{"negative subtraction", 100, -50, 150, false, nil},
		{"underflow", math.MinInt64, 1, 0, true, func(err error) bool { return err == ErrOverflow }},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m1 := NewMoney[currency.USD](tt.a)
			m2 := NewMoney[currency.USD](tt.b)
			err := m1.Sub(m2)

			if tt.wantErr {
				if err == nil || (tt.errCheck != nil && !tt.errCheck(err)) {
					t.Errorf("Sub() error = %v, wantErr %v", err, tt.wantErr)
				}
				return
			}

			if err != nil {
				t.Errorf("Sub() unexpected error: %v", err)
			}

			if m1.Amount() != tt.want {
				t.Errorf("Sub() = %v, want %v", m1.Amount(), tt.want)
			}
		})
	}
}

func TestMul(t *testing.T) {
	tests := []struct {
		name     string
		amount   int64
		scale    int64
		want     int64
		wantErr  bool
		errCheck func(error) bool
	}{
		{"simple multiplication", 100, 2, 200, false, nil},
		{"zero multiplication", 100, 0, 0, false, nil},
		{"negative multiplication", 100, -2, -200, false, nil},
		{"overflow", math.MaxInt64, 2, 0, true, func(err error) bool { return err == ErrOverflow }},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := NewMoney[currency.USD](tt.amount)
			err := m.Mul(tt.scale)

			if tt.wantErr {
				if err == nil || (tt.errCheck != nil && !tt.errCheck(err)) {
					t.Errorf("Mul() error = %v, wantErr %v", err, tt.wantErr)
				}
				return
			}

			if err != nil {
				t.Errorf("Mul() unexpected error: %v", err)
			}

			if m.Amount() != tt.want {
				t.Errorf("Mul() = %v, want %v", m.Amount(), tt.want)
			}
		})
	}
}

func TestString(t *testing.T) {
	tests := []struct {
		name   string
		amount int64
		want   string
		curr   currency.Currency
	}{
		{"positive currency.USD", 1050, "$10.50", currency.USD{}},
		{"negative currency.USD", -1050, "-$10.50", currency.USD{}},
		{"zero currency.USD", 0, "$0", currency.USD{}},
		{"JPY no decimals", 1000, "¥1000", currency.JPY{}},
		{"EUR positive", 1999, "€19.99", currency.EUR{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			switch tt.curr.(type) {
			case currency.USD:
				m := NewMoney[currency.USD](tt.amount)
				if got := m.String(); got != tt.want {
					t.Errorf("String() = %v, want %v", got, tt.want)
				}
			case currency.EUR:
				m := NewMoney[currency.EUR](tt.amount)
				if got := m.String(); got != tt.want {
					t.Errorf("String() = %v, want %v", got, tt.want)
				}
			case currency.JPY:
				m := NewMoney[currency.JPY](tt.amount)
				if got := m.String(); got != tt.want {
					t.Errorf("String() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func TestDistribute(t *testing.T) {
	tests := []struct {
		name      string
		amount    int64
		chunks    int64
		wantDist  *Distribution
		wantError bool
	}{
		{
			name:   "even distribution",
			amount: 1000,
			chunks: 4,
			wantDist: &Distribution{
				SmallerChunkSize: 250,
				SmallerCount:     4,
				LargerChunkSize:  250,
				LargerCount:      0,
			},
			wantError: false,
		},
		{
			name:   "uneven distribution",
			amount: 1000,
			chunks: 3,
			wantDist: &Distribution{
				SmallerChunkSize: 333,
				SmallerCount:     2,
				LargerChunkSize:  334,
				LargerCount:      1,
			},
			wantError: false,
		},
		{
			name:      "invalid chunks",
			amount:    1000,
			chunks:    0,
			wantDist:  nil,
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := NewMoney[currency.USD](tt.amount)
			dist, err := m.Distribute(tt.chunks)

			if tt.wantError {
				if err == nil {
					t.Error("Distribute() expected error, got nil")
				}
				return
			}

			if err != nil {
				t.Errorf("Distribute() unexpected error: %v", err)
				return
			}

			if *dist != *tt.wantDist {
				t.Errorf("Distribute() = %+v, want %+v", dist, tt.wantDist)
			}
		})
	}
}

func TestConvert(t *testing.T) {
	tests := []struct {
		name      string
		amount    int64
		ratio     Ratio
		want      int64
		wantError bool
	}{
		{
			name:   "simple conversion",
			amount: 10000,
			ratio: Ratio{
				Numerator:   107203,
				Denominator: 100000,
			},
			want:      10720,
			wantError: false,
		},
		{
			name:   "zero amount",
			amount: 0,
			ratio: Ratio{
				Numerator:   107203,
				Denominator: 100000,
			},
			want:      0,
			wantError: false,
		},
		{
			name:   "invalid ratio",
			amount: 1000,
			ratio: Ratio{
				Numerator:   1,
				Denominator: 0,
			},
			want:      0,
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := NewMoney[currency.EUR](tt.amount)
			converted, result, err := Convert[currency.EUR, currency.USD](m, tt.ratio)

			if tt.wantError {
				if err == nil {
					t.Error("Convert() expected error, got nil")
				}
				return
			}

			if err != nil {
				t.Errorf("Convert() unexpected error: %v", err)
				return
			}

			if converted.Amount() != tt.want {
				t.Errorf("Convert() = %v, want %v", converted.Amount(), tt.want)
			}

			if result.Amount != tt.want {
				t.Errorf("Convert() result amount = %v, want %v", result.Amount, tt.want)
			}
		})
	}
}

func TestJSON(t *testing.T) {
	tests := []struct {
		name      string
		money     *Money[currency.USD]
		wantJSON  string
		wantError bool
	}{
		{
			name:     "marshal simple",
			money:    NewMoney[currency.USD](1050),
			wantJSON: `{"amount":"1050","currency":"USD"}`,
		},
		{
			name:     "marshal zero",
			money:    NewMoney[currency.USD](0),
			wantJSON: `{"amount":"0","currency":"USD"}`,
		},
		{
			name:     "marshal negative",
			money:    NewMoney[currency.USD](-1050),
			wantJSON: `{"amount":"-1050","currency":"USD"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test marshaling
			got, err := json.Marshal(tt.money)
			if err != nil {
				t.Errorf("MarshalJSON() error = %v", err)
				return
			}
			if string(got) != tt.wantJSON {
				t.Errorf("MarshalJSON() = %v, want %v", string(got), tt.wantJSON)
			}

			// Test unmarshaling
			var unmarshaledMoney Money[currency.USD]
			err = json.Unmarshal([]byte(tt.wantJSON), &unmarshaledMoney)
			if err != nil {
				t.Errorf("UnmarshalJSON() error = %v", err)
				return
			}
			if unmarshaledMoney.Amount() != tt.money.Amount() {
				t.Errorf("UnmarshalJSON() amount = %v, want %v", unmarshaledMoney.Amount(), tt.money.Amount())
			}
		})
	}
}

func TestValidate(t *testing.T) {
	tests := []struct {
		name      string
		amount    int64
		min       int64
		max       int64
		wantError bool
	}{
		{"valid amount", 500, 0, 1000, false},
		{"at minimum", 0, 0, 1000, false},
		{"at maximum", 1000, 0, 1000, false},
		{"below minimum", -1, 0, 1000, true},
		{"above maximum", 1001, 0, 1000, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := NewMoney[currency.USD](tt.amount)
			err := m.Validate(tt.min, tt.max)

			if tt.wantError && err == nil {
				t.Error("Validate() expected error, got nil")
			}
			if !tt.wantError && err != nil {
				t.Errorf("Validate() unexpected error: %v", err)
			}
		})
	}
}
