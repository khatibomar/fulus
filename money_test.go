package fulus

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"testing"

	"github.com/khatibomar/fulus/currency"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		name        string
		a, b        int64
		expected    int64
		expectedErr error
	}{
		{
			name:        "simple addition",
			a:           100,
			b:           200,
			expected:    300,
			expectedErr: nil,
		},
		{
			name:        "zero addition",
			a:           100,
			b:           0,
			expected:    100,
			expectedErr: nil,
		},
		{
			name:        "negative addition",
			a:           100,
			b:           -50,
			expected:    50,
			expectedErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m1 := NewMoney[currency.USD](tt.a)
			m2 := NewMoney[currency.USD](tt.b)
			m1, err := m1.Add(m2)

			if err != tt.expectedErr {
				t.Errorf("Add() error = %v, expected error %v", err, tt.expectedErr)
				return
			}

			if tt.expectedErr == nil && m1.amount64() != tt.expected {
				t.Errorf("Add() = %v, expected %v", m1.amount64(), tt.expected)
			}
		})
	}
}

func TestSub(t *testing.T) {
	tests := []struct {
		name        string
		a, b        int64
		expected    int64
		expectedErr error
	}{
		{
			name:        "simple subtraction",
			a:           200,
			b:           100,
			expected:    100,
			expectedErr: nil,
		},
		{
			name:        "zero subtraction",
			a:           100,
			b:           0,
			expected:    100,
			expectedErr: nil,
		},
		{
			name:        "negative subtraction",
			a:           100,
			b:           -50,
			expected:    150,
			expectedErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m1 := NewMoney[currency.USD](tt.a)
			m2 := NewMoney[currency.USD](tt.b)
			m1, err := m1.Sub(m2)

			if err != tt.expectedErr {
				t.Errorf("Sub() error = %v, expected error %v", err, tt.expectedErr)
				return
			}

			if tt.expectedErr == nil && m1.amount64() != tt.expected {
				t.Errorf("Sub() = %v, expected %v", m1.amount64(), tt.expected)
			}
		})
	}
}

func TestMul(t *testing.T) {
	tests := []struct {
		name        string
		amount      int64
		scale       int64
		expected    int64
		expectedErr error
	}{
		{
			name:        "simple multiplication",
			amount:      100,
			scale:       2,
			expected:    200,
			expectedErr: nil,
		},
		{
			name:        "zero multiplication",
			amount:      100,
			scale:       0,
			expected:    0,
			expectedErr: nil,
		},
		{
			name:        "negative multiplication",
			amount:      100,
			scale:       -2,
			expected:    -200,
			expectedErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := NewMoney[currency.USD](tt.amount)
			m, err := m.Mul(tt.scale)

			if err != tt.expectedErr {
				t.Errorf("Mul() error = %v, expected error %v", err, tt.expectedErr)
				return
			}

			if tt.expectedErr == nil && m.amount64() != tt.expected {
				t.Errorf("Mul() = %v, expected %v", m.amount64(), tt.expected)
			}
		})
	}
}

func TestDiv(t *testing.T) {
	tests := []struct {
		name        string
		amount      int64
		divisor     int64
		mode        RoundingMode
		expected    int64
		expectedErr error
	}{
		{
			name:        "exact division",
			amount:      100,
			divisor:     2,
			mode:        RoundHalfUp,
			expected:    50,
			expectedErr: nil,
		},
		{
			name:        "round half up",
			amount:      105,
			divisor:     2,
			mode:        RoundHalfUp,
			expected:    53,
			expectedErr: nil,
		},
		{
			name:        "round half even (even)",
			amount:      105,
			divisor:     2,
			mode:        RoundHalfEven,
			expected:    52,
			expectedErr: nil,
		},
		{
			name:        "round half even (odd)",
			amount:      115, // 115 / 2 = 57.5 -> 58
			divisor:     2,
			mode:        RoundHalfEven,
			expected:    58,
			expectedErr: nil,
		},
		{
			name:        "truncate",
			amount:      105,
			divisor:     2,
			mode:        RoundTruncate,
			expected:    52,
			expectedErr: nil,
		},
		{
			name:        "division by zero",
			amount:      100,
			divisor:     0,
			mode:        RoundHalfUp,
			expected:    0,
			expectedErr: ErrDivisionByZero,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := NewMoney[currency.USD](tt.amount)
			m, err := m.Div(tt.divisor, tt.mode)

			if err != tt.expectedErr {
				t.Errorf("Div() error = %v, expected error %v", err, tt.expectedErr)
				return
			}

			if tt.expectedErr == nil && m.amount64() != tt.expected {
				t.Errorf("Div() = %v, expected %v", m.amount64(), tt.expected)
			}
		})
	}
}

func TestAbs(t *testing.T) {
	tests := []struct {
		name        string
		amount      int64
		expected    int64
		expectedErr error
	}{
		{
			name:        "positive",
			amount:      100,
			expected:    100,
			expectedErr: nil,
		},
		{
			name:        "negative",
			amount:      -100,
			expected:    100,
			expectedErr: nil,
		},
		{
			name:        "zero",
			amount:      0,
			expected:    0,
			expectedErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := NewMoney[currency.USD](tt.amount)
			m, err := m.Abs()

			if err != tt.expectedErr {
				t.Errorf("Abs() error = %v, expected error %v", err, tt.expectedErr)
				return
			}

			if tt.expectedErr == nil && m.amount64() != tt.expected {
				t.Errorf("Abs() = %v, expected %v", m.amount64(), tt.expected)
			}
		})
	}
}

func TestNeg(t *testing.T) {
	tests := []struct {
		name        string
		amount      int64
		expected    int64
		expectedErr error
	}{
		{
			name:        "positive",
			amount:      100,
			expected:    -100,
			expectedErr: nil,
		},
		{
			name:        "negative",
			amount:      -100,
			expected:    100,
			expectedErr: nil,
		},
		{
			name:        "zero",
			amount:      0,
			expected:    0,
			expectedErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := NewMoney[currency.USD](tt.amount)
			m, err := m.Neg()

			if err != tt.expectedErr {
				t.Errorf("Neg() error = %v, expected error %v", err, tt.expectedErr)
				return
			}

			if tt.expectedErr == nil && m.amount64() != tt.expected {
				t.Errorf("Neg() = %v, expected %v", m.amount64(), tt.expected)
			}
		})
	}
}

func TestString(t *testing.T) {
	tests := []struct {
		name     string
		amount   int64
		expected string
		curr     currency.Currency
	}{
		{
			name:     "positive currency.USD",
			amount:   1050,
			expected: "USD 10.50",
			curr:     currency.USD{},
		},
		{
			name:     "negative currency.USD",
			amount:   -1050,
			expected: "USD -10.50",
			curr:     currency.USD{},
		},
		{
			name:     "zero currency.USD",
			amount:   0,
			expected: "USD 0.00",
			curr:     currency.USD{},
		},
		{
			name:     "JPY no decimals",
			amount:   1000,
			expected: "JPY 1000",
			curr:     currency.JPY{},
		},
		{
			name:     "EUR positive",
			amount:   1999,
			expected: "EUR 19.99",
			curr:     currency.EUR{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			switch tt.curr.(type) {
			case currency.USD:
				m := NewMoney[currency.USD](tt.amount)
				if got := m.String(); got != tt.expected {
					t.Errorf("String() = %v, expected %v", got, tt.expected)
				}
			case currency.EUR:
				m := NewMoney[currency.EUR](tt.amount)
				if got := m.String(); got != tt.expected {
					t.Errorf("String() = %v, expected %v", got, tt.expected)
				}
			case currency.JPY:
				m := NewMoney[currency.JPY](tt.amount)
				if got := m.String(); got != tt.expected {
					t.Errorf("String() = %v, expected %v", got, tt.expected)
				}
			}
		})
	}
}

func TestDistribute(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name         string
		amount       int64
		chunks       int64
		smaller      int64
		smallerCount int64
		larger       int64
		largerCount  int64
		expectedErr  error
	}{
		{name: "even distribution", amount: 1000, chunks: 4, smaller: 250, smallerCount: 4, larger: 250},
		{name: "uneven distribution", amount: 1000, chunks: 3, smaller: 333, smallerCount: 2, larger: 334, largerCount: 1},
		{name: "invalid chunks", amount: 1000, chunks: 0, expectedErr: ErrInvalidChunks},
		{name: "negative chunks", amount: 1000, chunks: -1, expectedErr: ErrInvalidChunks},
		{name: "negative uneven distribution", amount: -1000, chunks: 3, smaller: -334, smallerCount: 1, larger: -333, largerCount: 2},
		{name: "largest amount in one chunk", amount: math.MaxInt64, chunks: 1, smaller: math.MaxInt64, smallerCount: 1, larger: math.MaxInt64},
		{name: "smallest amount in one chunk", amount: math.MinInt64, chunks: 1, smaller: math.MinInt64, smallerCount: 1, larger: math.MinInt64},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			dist, err := NewMoney[currency.USD](tt.amount).Distribute(tt.chunks)
			if !errors.Is(err, tt.expectedErr) {
				t.Fatalf("Distribute() error = %v, expected error %v", err, tt.expectedErr)
			}
			if err != nil {
				return
			}
			want := Distribution[currency.USD]{
				Smaller:      NewMoney[currency.USD](tt.smaller),
				SmallerCount: tt.smallerCount,
				Larger:       NewMoney[currency.USD](tt.larger),
				LargerCount:  tt.largerCount,
			}
			if dist != want {
				t.Errorf("Distribute() = %+v, expected %+v", dist, want)
			}
		})
	}
}

func TestParseMoney(t *testing.T) {
	tests := []struct {
		name        string
		amount      string
		expected    int64
		expectedErr error
		isJPY       bool
	}{
		{name: "usd whole", amount: "10", expected: 1000},
		{name: "usd decimal", amount: "12.34", expected: 1234},
		{name: "usd negative", amount: "-0.99", expected: -99},
		{name: "usd plus sign", amount: "+7.50", expected: 750},
		{name: "jpy whole", amount: "100", expected: 100, isJPY: true},
		{name: "scale mismatch", amount: "1.234", expectedErr: ErrScaleMismatch},
		{name: "zeros after the minor units", amount: "-1.2300", expected: -123},
		{name: "digit after zeros", amount: "1.2301", expectedErr: ErrScaleMismatch},
		{name: "jpy zeros after the point", amount: "100.000", expected: 100, isJPY: true},
		{name: "jpy fraction digit", amount: "100.5", expectedErr: ErrScaleMismatch, isJPY: true},
		{name: "invalid format", amount: "abc", expectedErr: ErrInvalidAmountFormat},
		{name: "empty", amount: "", expectedErr: ErrInvalidAmountFormat},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.isJPY {
				m, err := ParseMoney[currency.JPY](tt.amount)
				if tt.expectedErr != nil {
					if !errors.Is(err, tt.expectedErr) {
						t.Fatalf("expected error %v, got %v", tt.expectedErr, err)
					}
					return
				}
				if err != nil {
					t.Fatalf("unexpected error: %v", err)
				}
				if m.amount64() != tt.expected {
					t.Fatalf("amount = %d, expected %d", m.amount64(), tt.expected)
				}
				return
			}

			m, err := ParseMoney[currency.USD](tt.amount)
			if tt.expectedErr != nil {
				if !errors.Is(err, tt.expectedErr) {
					t.Fatalf("expected error %v, got %v", tt.expectedErr, err)
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if m.amount64() != tt.expected {
				t.Fatalf("amount = %d, expected %d", m.amount64(), tt.expected)
			}
		})
	}
}

func TestMoneyValueAndScan(t *testing.T) {
	t.Parallel()

	original := NewMoney[currency.USD](1050)
	v, err := original.Value()
	if err != nil {
		t.Fatalf("Value() error = %v", err)
	}
	if v != "10.50" {
		t.Fatalf("Value() = %#v, expected \"10.50\"", v)
	}

	tests := []struct {
		name    string
		value   any
		want    int64
		wantErr error
	}{
		{name: "Value output round-trip", value: v, want: 1050},
		{name: "NUMERIC text", value: []byte("10.50"), want: 1050},
		{name: "NUMERIC integer text is major units", value: []byte("1050"), want: 105000},
		{name: "NUMERIC with more scale and zeros", value: "-7.5", want: -750},
		{name: "int64", value: int64(99), wantErr: ErrInvalidAmountFormat},
		{name: "zero int64", value: int64(0), wantErr: ErrInvalidAmountFormat},
		{name: "too many fraction digits", value: "12.505", wantErr: ErrScaleMismatch},
		{name: "NUMERIC(19, 4) text", value: []byte("10.5000"), want: 1050},
		{name: "NUMERIC(19, 4) with a digit after the minor units", value: []byte("10.5001"), wantErr: ErrScaleMismatch},
		{name: "JSON is not accepted", value: `{"amount":"1.00","currency":"USD"}`, wantErr: ErrInvalidAmountFormat},
		{name: "nil", value: nil, wantErr: ErrInvalidAmountFormat},
		{name: "float64", value: 1.5, want: 150},
		{name: "negative float64", value: -0.05, want: -5},
		{name: "float64 with 15 digits", value: 1234567890123.45, want: 123456789012345},
		{name: "float64 with too many digits", value: 12345678901234.56, wantErr: ErrInvalidAmountFormat},
		{name: "large float64", value: 1e20, wantErr: ErrInvalidAmountFormat},
		{name: "float64 sum", value: 0.30000000000000004, wantErr: ErrInvalidAmountFormat},
		{name: "float64 with more fraction digits", value: 1.005, wantErr: ErrScaleMismatch},
		{name: "NaN", value: math.NaN(), wantErr: ErrInvalidAmountFormat},
		{name: "infinity", value: math.Inf(-1), wantErr: ErrInvalidAmountFormat},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var scanned Money[currency.USD]
			err := scanned.Scan(tt.value)
			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("Scan(%#v) error = %v, want %v", tt.value, err, tt.wantErr)
			}
			if err == nil && scanned.amount64() != tt.want {
				t.Fatalf("Scan(%#v) amount = %d, expected %d", tt.value, scanned.amount64(), tt.want)
			}
		})
	}
}

func TestBigintMoney(t *testing.T) {
	t.Parallel()

	v, err := BigintMoney[currency.USD]{Money: NewMoney[currency.USD](1050)}.Value()
	if err != nil || v != int64(1050) {
		t.Fatalf("Value() = %#v, %v", v, err)
	}
	if _, err := (BigintMoney[currency.USD]{Money: maxMoney[currency.USD]()}).Value(); !errors.Is(err, ErrOverflow) {
		t.Errorf("Value() of the largest amount error = %v", err)
	}

	tests := []struct {
		name    string
		value   any
		want    int64
		wantErr error
	}{
		{name: "int64", value: int64(1050), want: 1050},
		{name: "integer text", value: []byte("-7"), want: -7},
		{name: "decimal text", value: "10.50", wantErr: ErrInvalidAmountFormat},
		{name: "NULL", value: nil, wantErr: ErrInvalidAmountFormat},
		{name: "float64", value: 1.5, wantErr: ErrInvalidAmountFormat},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			var b BigintMoney[currency.USD]
			err := b.Scan(tt.value)
			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("Scan() error = %v, want %v", err, tt.wantErr)
			}
			if err == nil && b.Money.amount64() != tt.want {
				t.Errorf("Scan() = %d, want %d", b.Money.amount64(), tt.want)
			}
		})
	}

	var n sql.Null[BigintMoney[currency.USD]]
	if err := n.Scan(int64(5)); err != nil || !n.Valid || n.V.Money.amount64() != 5 {
		t.Errorf("sql.Null Scan() = %+v, %v", n, err)
	}
	if v, err := n.Value(); err != nil || v != int64(5) {
		t.Errorf("sql.Null Value() = %#v, %v", v, err)
	}
	if err := n.Scan(nil); err != nil || n.Valid {
		t.Errorf("sql.Null Scan(nil) = %+v, %v", n, err)
	}
}

func TestZeroValueMoney(t *testing.T) {
	t.Parallel()

	var m Money[currency.USD]
	if got := m.String(); got != "USD 0.00" {
		t.Errorf("String() = %q, expected %q", got, "USD 0.00")
	}
	if got := m.Currency().Code(); got != "USD" {
		t.Errorf("Currency().Code() = %q, expected %q", got, "USD")
	}
	if m != NewMoney[currency.USD](0) {
		t.Error("zero value must equal NewMoney(0)")
	}
	b, err := json.Marshal(m)
	if err != nil {
		t.Fatalf("Marshal() error = %v", err)
	}
	if string(b) != `{"amount":"0.00","currency":"USD"}` {
		t.Errorf("Marshal() = %s", b)
	}
}

func TestJSON(t *testing.T) {
	tests := []struct {
		name        string
		money       Money[currency.USD]
		expected    string
		expectedErr error
	}{
		{
			name:     "marshal simple",
			money:    NewMoney[currency.USD](1050),
			expected: `{"amount":"10.50","currency":"USD"}`,
		},
		{
			name:     "marshal zero",
			money:    NewMoney[currency.USD](0),
			expected: `{"amount":"0.00","currency":"USD"}`,
		},
		{
			name:     "marshal negative",
			money:    NewMoney[currency.USD](-1050),
			expected: `{"amount":"-10.50","currency":"USD"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := json.Marshal(tt.money)
			if err != tt.expectedErr {
				t.Errorf("MarshalJSON() error = %v, expected error %v", err, tt.expectedErr)
				return
			}
			if string(got) != tt.expected {
				t.Errorf("MarshalJSON() = %v, expected %v", string(got), tt.expected)
			}

			var unmarshaledMoney Money[currency.USD]
			err = json.Unmarshal([]byte(tt.expected), &unmarshaledMoney)
			if err != tt.expectedErr {
				t.Errorf("UnmarshalJSON() error = %v, expected error %v", err, tt.expectedErr)
				return
			}
			if unmarshaledMoney.amount64() != tt.money.amount64() {
				t.Errorf("UnmarshalJSON() amount = %v, expected %v", unmarshaledMoney.amount64(), tt.money.amount64())
			}
		})
	}
}

func TestUnmarshalJSON(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		input   string
		want    int64
		wantErr error
	}{
		{name: "fewer fraction digits", input: `{"amount":"10.5","currency":"USD"}`, want: 1050},
		{name: "integer", input: `{"amount":"10","currency":"USD"}`, want: 1000},
		{name: "invalid amount", input: `{"amount":"10oops","currency":"USD"}`, wantErr: ErrInvalidAmountFormat},
		{name: "too many fraction digits", input: `{"amount":"10.505","currency":"USD"}`, wantErr: ErrScaleMismatch},
		{name: "zeros after the minor units", input: `{"amount":"10.500","currency":"USD"}`, want: 1050},
		{name: "currency mismatch", input: `{"amount":"10.50","currency":"EUR"}`, wantErr: ErrCurrencyMismatch},
		{name: "missing currency", input: `{"amount":"10.50"}`, wantErr: ErrCurrencyMismatch},
		{name: "lower case currency", input: `{"amount":"10.50","currency":"usd"}`, want: 1050},
		{name: "null does not change the value", input: ` null `, want: 7},
		{name: "number amount", input: `{"amount":10.50,"currency":"USD"}`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			m := NewMoney[currency.USD](7)
			err := json.Unmarshal([]byte(tt.input), &m)
			if tt.name == "number amount" {
				if err == nil {
					t.Fatal("Unmarshal() of a JSON number amount succeeded")
				}
				return
			}
			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("Unmarshal() error = %v, want %v", err, tt.wantErr)
			}
			if err == nil && m.amount64() != tt.want {
				t.Errorf("Unmarshal() = %d, want %d", m.amount64(), tt.want)
			}
		})
	}
}

func TestInRange(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		amount int64
		want   bool
	}{
		{name: "inside", amount: 500, want: true},
		{name: "at minimum", amount: 0, want: true},
		{name: "at maximum", amount: 1000, want: true},
		{name: "below minimum", amount: -1, want: false},
		{name: "above maximum", amount: 1001, want: false},
	}

	low, high := NewMoney[currency.USD](0), NewMoney[currency.USD](1000)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := NewMoney[currency.USD](tt.amount).InRange(low, high); got != tt.want {
				t.Errorf("InRange() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAllocate(t *testing.T) {
	tests := []struct {
		name        string
		amount      int64
		ratios      []int64
		expected    []int64
		expectedErr error
	}{
		{
			name:     "simple equal split",
			amount:   100,
			ratios:   []int64{1, 1},
			expected: []int64{50, 50},
		},
		{
			name:     "uneven split",
			amount:   100,
			ratios:   []int64{1, 2},
			expected: []int64{33, 67},
		},
		{
			name:     "three way split",
			amount:   100,
			ratios:   []int64{1, 1, 2},
			expected: []int64{25, 25, 50},
		},
		{
			name:     "complex ratios",
			amount:   1000,
			ratios:   []int64{3, 7},
			expected: []int64{300, 700},
		},
		{
			name:     "zero ratio",
			amount:   100,
			ratios:   []int64{0, 1},
			expected: []int64{0, 100},
		},
		{
			name:     "zero ratio with leftover units",
			amount:   -100,
			ratios:   []int64{1, 0, 1, 1},
			expected: []int64{-34, 0, -33, -33},
		},
		{
			name:        "all ratios zero",
			amount:      100,
			ratios:      []int64{0, 0},
			expectedErr: ErrInvalidRatios,
		},
		{
			name:        "negative ratio",
			amount:      100,
			ratios:      []int64{-1, 1},
			expectedErr: ErrInvalidRatios,
		},
		{
			name:        "empty ratios",
			amount:      100,
			ratios:      []int64{},
			expectedErr: ErrNoRatios,
		},
		{
			name:     "handle remainder",
			amount:   1000,
			ratios:   []int64{1, 1, 1},
			expected: []int64{334, 333, 333},
		},
		{
			name:     "spread leftover units",
			amount:   100,
			ratios:   []int64{1, 1, 1, 1, 1, 1, 1},
			expected: []int64{15, 15, 14, 14, 14, 14, 14},
		},
		{
			name:     "largest remainder gets leftover",
			amount:   100,
			ratios:   []int64{1, 1, 1, 1, 1, 1, 3},
			expected: []int64{11, 11, 11, 11, 11, 11, 34},
		},
		{
			name:     "negative amount",
			amount:   -100,
			ratios:   []int64{1, 1, 1},
			expected: []int64{-34, -33, -33},
		},
		{
			name:     "minimum int64",
			amount:   math.MinInt64,
			ratios:   []int64{1, 2},
			expected: []int64{-3074457345618258603, -6148914691236517205},
		},
		{
			name:     "large amount with multiple ratios",
			amount:   1000000,
			ratios:   []int64{1, 2, 3, 4},
			expected: []int64{100000, 200000, 300000, 400000},
		},
		{
			name:        "ratio total overflow",
			amount:      100,
			ratios:      []int64{math.MaxInt64, 1},
			expectedErr: ErrOverflow,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			money := NewMoney[currency.USD](tt.amount)
			parts, err := money.Allocate(tt.ratios...)

			if err != tt.expectedErr {
				t.Errorf("expected error %v, got %v", tt.expectedErr, err)
				return
			}

			if tt.expectedErr == nil {
				for i, expected := range tt.expected {
					if parts[i].amount64() != expected {
						t.Errorf("part %d: expected %d, got %d", i, expected, parts[i].amount64())
					}
				}

				sum := int64(0)
				for _, part := range parts {
					sum += part.amount64()
				}
				if sum != tt.amount {
					t.Errorf("sum of parts (%d) does not equal original amount (%d)", sum, tt.amount)
				}

			}
		})
	}
}

func TestAllocateRealMoney(t *testing.T) {
	tests := []struct {
		name        string
		amount      int64
		ratios      []int64
		expected    []string
		expectedErr error
	}{
		{
			name:     "split $100 equally",
			amount:   10000, // $100.00
			ratios:   []int64{1, 1},
			expected: []string{"USD 50.00", "USD 50.00"},
		},
		{
			name:     "split $100 in thirds",
			amount:   10000, // $100.00
			ratios:   []int64{1, 1, 1},
			expected: []string{"USD 33.34", "USD 33.33", "USD 33.33"},
		},
		{
			name:     "split $50.50 by ratio 1:2",
			amount:   5050, // $50.50
			ratios:   []int64{1, 2},
			expected: []string{"USD 16.83", "USD 33.67"},
		},
		{
			name:        "empty ratios",
			amount:      10000,
			ratios:      []int64{},
			expectedErr: ErrNoRatios,
		},
		{
			name:        "negative ratio",
			amount:      10000,
			ratios:      []int64{1, -1},
			expectedErr: ErrInvalidRatios,
		},
		{
			name:     "zero ratio",
			amount:   10000,
			ratios:   []int64{0, 1},
			expected: []string{"USD 0.00", "USD 100.00"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			money := NewMoney[currency.USD](tt.amount)
			parts, err := money.Allocate(tt.ratios...)
			if err != tt.expectedErr {
				t.Errorf("expected error %v, got %v", tt.expectedErr, err)
				return
			}
			if err != nil {
				return
			}

			for i, expected := range tt.expected {
				if parts[i].String() != expected {
					t.Errorf("part %d: expected %s, got %s",
						i, expected, parts[i].String())
				}
			}

			sum := int64(0)
			for _, part := range parts {
				sum += part.amount64()
			}
			if sum != tt.amount {
				t.Errorf("sum of parts (%d) does not equal original amount (%d)",
					sum, tt.amount)
			}
		})
	}
}

func ExampleMoney_Allocate() {
	dollars := NewMoney[currency.USD](10000) // 100.00 USD

	// Allocate in ratio of [1,1,2]
	parts, err := dollars.Allocate(1, 1, 2)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for i, part := range parts {
		fmt.Printf("Part %d: %v\n", i+1, part)
	}

	// Output:
	// Part 1: USD 25.00
	// Part 2: USD 25.00
	// Part 3: USD 50.00
}

func ExampleMoney_Distribute() {
	dollars := NewMoney[currency.USD](10000) // 100.00 USD

	dist, err := dollars.Distribute(3)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Smaller chunks: %d x %v\n", dist.SmallerCount, dist.Smaller)
	fmt.Printf("Larger chunks: %d x %v\n", dist.LargerCount, dist.Larger)

	// Output:
	// Smaller chunks: 2 x USD 33.33
	// Larger chunks: 1 x USD 33.34
}

func FuzzMoneyUnmarshalJSON(f *testing.F) {
	f.Add("1.00", "USD")
	f.Add("-0.50", "USD")
	f.Add("abc", "USD")

	f.Fuzz(func(t *testing.T, amount, curr string) {
		payload, err := json.Marshal(moneyJSON{Amount: amount, Currency: curr})
		if err != nil {
			t.Fatal(err)
		}
		var m Money[currency.USD]
		if err := json.Unmarshal(payload, &m); err != nil {
			return
		}
		// A value that unmarshals must marshal to the canonical form and unmarshal to the same value.
		out, err := json.Marshal(m)
		if err != nil {
			t.Fatal(err)
		}
		var back Money[currency.USD]
		if err := json.Unmarshal(out, &back); err != nil || back != m {
			t.Fatalf("round trip of %s: %s, %v", payload, out, err)
		}
	})
}

func FuzzDistributeInvariants(f *testing.F) {
	f.Add(int64(100), int64(3))
	f.Add(int64(-100), int64(3))

	f.Fuzz(func(t *testing.T, amount, chunks int64) {
		if chunks <= 0 {
			return
		}
		m := NewMoney[currency.USD](amount)
		dist, err := m.Distribute(chunks)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if dist.SmallerCount < 0 || dist.LargerCount < 0 {
			t.Fatalf("counts must be non-negative: %+v", dist)
		}

		total := (dist.Smaller.amount64() * dist.SmallerCount) + (dist.Larger.amount64() * dist.LargerCount)
		if total != amount {
			t.Fatalf("distribution total = %d, amount = %d", total, amount)
		}
	})
}

func TestComparisons(t *testing.T) {
	m10 := NewMoney[currency.USD](1000)
	m10_2 := NewMoney[currency.USD](1000)
	m20 := NewMoney[currency.USD](2000)
	m0 := NewMoney[currency.USD](0)
	m_10 := NewMoney[currency.USD](-1000)

	// Test Cmp
	if got := m10.Cmp(m20); got != -1 {
		t.Errorf("Cmp(10, 20) = %d, want -1", got)
	}
	if got := m10.Cmp(m10_2); got != 0 {
		t.Errorf("Cmp(10, 10) = %d, want 0", got)
	}
	if got := m20.Cmp(m10); got != 1 {
		t.Errorf("Cmp(20, 10) = %d, want 1", got)
	}

	// Test Equal
	if !m10.Equal(m10_2) {
		t.Errorf("Equal(10, 10) should be true")
	}
	if m10.Equal(m20) {
		t.Errorf("Equal(10, 20) should be false")
	}

	// Test GreaterThan
	if !m20.GreaterThan(m10) {
		t.Errorf("GreaterThan(20, 10) should be true")
	}
	if m10.GreaterThan(m20) {
		t.Errorf("GreaterThan(10, 20) should be false")
	}

	// Test GreaterThanOrEqual
	if !m20.GreaterThanOrEqual(m10) {
		t.Errorf("GreaterThanOrEqual(20, 10) should be true")
	}
	if !m10.GreaterThanOrEqual(m10_2) {
		t.Errorf("GreaterThanOrEqual(10, 10) should be true")
	}
	if m10.GreaterThanOrEqual(m20) {
		t.Errorf("GreaterThanOrEqual(10, 20) should be false")
	}

	// Test LessThan
	if !m10.LessThan(m20) {
		t.Errorf("LessThan(10, 20) should be true")
	}
	if m20.LessThan(m10) {
		t.Errorf("LessThan(20, 10) should be false")
	}

	// Test LessThanOrEqual
	if !m10.LessThanOrEqual(m20) {
		t.Errorf("LessThanOrEqual(10, 20) should be true")
	}
	if !m10.LessThanOrEqual(m10_2) {
		t.Errorf("LessThanOrEqual(10, 10) should be true")
	}
	if m20.LessThanOrEqual(m10) {
		t.Errorf("LessThanOrEqual(20, 10) should be false")
	}

	// Test Zero, Positive, Negative
	if !m0.IsZero() {
		t.Errorf("IsZero(0) should be true")
	}
	if m10.IsZero() {
		t.Errorf("IsZero(10) should be false")
	}

	if !m10.IsPositive() {
		t.Errorf("IsPositive(10) should be true")
	}
	if m_10.IsPositive() {
		t.Errorf("IsPositive(-10) should be false")
	}

	if !m_10.IsNegative() {
		t.Errorf("IsNegative(-10) should be true")
	}
	if m10.IsNegative() {
		t.Errorf("IsNegative(10) should be false")
	}
}

func FuzzAllocateInvariants(f *testing.F) {
	f.Add(int64(100), int64(1), int64(1), int64(2))

	f.Fuzz(func(t *testing.T, amount, r1, r2, r3 int64) {
		ratios := []int64{boundedPositive(r1), boundedPositive(r2), boundedPositive(r3)}
		m := NewMoney[currency.USD](amount)
		parts, err := m.Allocate(ratios...)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		sum := int64(0)
		for _, part := range parts {
			sum += part.amount64()
		}

		if sum != amount {
			t.Fatalf("allocation sum = %d, amount = %d", sum, amount)
		}
	})
}

func absInt64(v int64) int64 {
	if v == math.MinInt64 {
		return math.MaxInt64
	}
	if v < 0 {
		return -v
	}
	return v
}

func boundedPositive(v int64) int64 {
	v = absInt64(v)
	v = (v % 1000) + 1
	return v
}

func TestMustAdd(t *testing.T) {
	m1 := NewMoney[currency.USD](100)
	m2 := NewMoney[currency.USD](50)

	result := m1.MustAdd(m2)
	if result.amount64() != 150 {
		t.Errorf("MustAdd = %d, want %d", result.amount64(), 150)
	}

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("MustAdd should have panicked on overflow")
		}
	}()
	mMax := maxMoney[currency.USD]()
	_ = mMax.MustAdd(NewMoney[currency.USD](1))
}

func TestMustSub(t *testing.T) {
	m1 := NewMoney[currency.USD](100)
	m2 := NewMoney[currency.USD](50)

	result := m1.MustSub(m2)
	if result.amount64() != 50 {
		t.Errorf("MustSub = %d, want %d", result.amount64(), 50)
	}

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("MustSub should have panicked on overflow")
		}
	}()
	mMin := minMoney[currency.USD]()
	_ = mMin.MustSub(NewMoney[currency.USD](1))
}

func TestMustMul(t *testing.T) {
	m := NewMoney[currency.USD](100)

	result := m.MustMul(3)
	if result.amount64() != 300 {
		t.Errorf("MustMul = %d, want %d", result.amount64(), 300)
	}

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("MustMul should have panicked on overflow")
		}
	}()
	mMax := maxMoney[currency.USD]()
	_ = mMax.MustMul(2)
}
