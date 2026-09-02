package fulus

import (
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"strings"
	"testing"

	"github.com/khatibomar/fulus/currency"
	"github.com/khatibomar/fulus/locale"
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
		{
			name:        "overflow",
			a:           math.MaxInt64,
			b:           1,
			expected:    0,
			expectedErr: ErrOverflow,
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

			if tt.expectedErr == nil && m1.Amount() != tt.expected {
				t.Errorf("Add() = %v, expected %v", m1.Amount(), tt.expected)
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
		{
			name:        "underflow",
			a:           math.MinInt64,
			b:           1,
			expected:    0,
			expectedErr: ErrOverflow,
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

			if tt.expectedErr == nil && m1.Amount() != tt.expected {
				t.Errorf("Sub() = %v, expected %v", m1.Amount(), tt.expected)
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
		{
			name:        "overflow",
			amount:      math.MaxInt64,
			scale:       2,
			expected:    0,
			expectedErr: ErrOverflow,
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

			if tt.expectedErr == nil && m.Amount() != tt.expected {
				t.Errorf("Mul() = %v, expected %v", m.Amount(), tt.expected)
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
		{
			name:        "overflow from minint / -1",
			amount:      math.MinInt64,
			divisor:     -1,
			mode:        RoundHalfUp,
			expected:    0,
			expectedErr: ErrOverflow,
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

			if tt.expectedErr == nil && m.Amount() != tt.expected {
				t.Errorf("Div() = %v, expected %v", m.Amount(), tt.expected)
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
		{
			name:        "overflow",
			amount:      math.MinInt64,
			expected:    0,
			expectedErr: ErrOverflow,
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

			if tt.expectedErr == nil && m.Amount() != tt.expected {
				t.Errorf("Abs() = %v, expected %v", m.Amount(), tt.expected)
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
		{
			name:        "overflow",
			amount:      math.MinInt64,
			expected:    0,
			expectedErr: ErrOverflow,
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

			if tt.expectedErr == nil && m.Amount() != tt.expected {
				t.Errorf("Neg() = %v, expected %v", m.Amount(), tt.expected)
			}
		})
	}
}

func TestFormat(t *testing.T) {
	tests := []struct {
		name     string
		money    Money[currency.USD]
		locale   locale.Locale
		expected string
	}{
		{
			name:     "positive whole number",
			money:    NewMoney[currency.USD](1000),
			locale:   locale.EN,
			expected: "$10.00",
		},
		{
			name:     "negative whole number",
			money:    NewMoney[currency.USD](-1000),
			locale:   locale.EN,
			expected: "-$10.00",
		},
		{
			name:     "minimum int64",
			money:    NewMoney[currency.USD](math.MinInt64),
			locale:   locale.EN,
			expected: "-$92,233,720,368,547,758.08",
		},
		{
			name:     "zero",
			money:    NewMoney[currency.USD](0),
			locale:   locale.EN,
			expected: "$0",
		},
		{
			name:     "with cents",
			money:    NewMoney[currency.USD](1234),
			locale:   locale.EN,
			expected: "$12.34",
		},
		{
			name:     "large number with grouping",
			money:    NewMoney[currency.USD](1234567),
			locale:   locale.EN,
			expected: "$12,345.67",
		},
		{
			name:     "different locale format (fr)",
			money:    NewMoney[currency.USD](1234567),
			locale:   locale.FR,
			expected: "12\u202f345,67\u00a0$US",
		},
		{
			name:     "different locale format (de)",
			money:    NewMoney[currency.USD](1234567),
			locale:   locale.DE,
			expected: "12.345,67\u00a0$",
		},
		{
			name:     "Arabic locale format",
			money:    NewMoney[currency.USD](1234567),
			locale:   locale.AR,
			expected: "\u200f12,345.67\u00a0US$;\u200f-#,##0.00\u00a0\u00a4",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.money.Format(tt.locale)
			if result != tt.expected {
				t.Errorf("Format() mismatch\nGot:  %+q (len: %d)\nWant: %+q (len: %d)",
					result, len(result),
					tt.expected, len(tt.expected))
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
			expected: "$10.50",
			curr:     currency.USD{},
		},
		{
			name:     "negative currency.USD",
			amount:   -1050,
			expected: "-$10.50",
			curr:     currency.USD{},
		},
		{
			name:     "zero currency.USD",
			amount:   0,
			expected: "$0",
			curr:     currency.USD{},
		},
		{
			name:     "JPY no decimals",
			amount:   1000,
			expected: "¥1000",
			curr:     currency.JPY{},
		},
		{
			name:     "EUR positive",
			amount:   1999,
			expected: "€19.99",
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
	tests := []struct {
		name        string
		amount      int64
		chunks      int64
		expected    Distribution
		expectedErr error
	}{
		{
			name:   "even distribution",
			amount: 1000,
			chunks: 4,
			expected: Distribution{
				SmallerChunkSize: 250,
				SmallerCount:     4,
				LargerChunkSize:  250,
				LargerCount:      0,
			},
			expectedErr: nil,
		},
		{
			name:   "uneven distribution",
			amount: 1000,
			chunks: 3,
			expected: Distribution{
				SmallerChunkSize: 333,
				SmallerCount:     2,
				LargerChunkSize:  334,
				LargerCount:      1,
			},
			expectedErr: nil,
		},
		{
			name:        "invalid chunks",
			amount:      1000,
			chunks:      0,
			expected:    Distribution{},
			expectedErr: ErrInvalidChunks,
		},
		{
			name:        "negative chunks",
			amount:      1000,
			chunks:      -1,
			expected:    Distribution{},
			expectedErr: ErrInvalidChunks,
		},
		{
			name:   "negative uneven distribution",
			amount: -1000,
			chunks: 3,
			expected: Distribution{
				SmallerChunkSize: -334,
				SmallerCount:     1,
				LargerChunkSize:  -333,
				LargerCount:      2,
			},
			expectedErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := NewMoney[currency.USD](tt.amount)
			dist, err := m.Distribute(tt.chunks)

			if err != tt.expectedErr {
				t.Errorf("Distribute() error = %v, expected error %v", err, tt.expectedErr)
				return
			}

			if err == nil && dist != tt.expected {
				t.Errorf("Distribute() = %+v, expected %+v", dist, tt.expected)
				return
			}

			if err == nil {
				total := (dist.SmallerChunkSize * dist.SmallerCount) +
					(dist.LargerChunkSize * dist.LargerCount)
				if total != tt.amount {
					t.Errorf("Distribute() total = %d, expected %d", total, tt.amount)
				}
			}
		})
	}
}

func TestConvert(t *testing.T) {
	tests := []struct {
		name        string
		amount      int64
		ratio       Ratio[currency.EUR, currency.USD]
		mode        RoundingMode
		expected    int64
		expectedErr error
	}{
		{
			name:   "simple conversion",
			amount: 10000,
			ratio: Ratio[currency.EUR, currency.USD]{
				Numerator:   107203,
				Denominator: 100000,
			},
			mode:        RoundTruncate,
			expected:    10720,
			expectedErr: nil,
		},
		{
			name:   "zero amount",
			amount: 0,
			ratio: Ratio[currency.EUR, currency.USD]{
				Numerator:   107203,
				Denominator: 100000,
			},
			mode:        RoundTruncate,
			expected:    0,
			expectedErr: nil,
		},
		{
			name:   "invalid ratio",
			amount: 1000,
			ratio: Ratio[currency.EUR, currency.USD]{
				Numerator:   1,
				Denominator: 0,
			},
			mode:        RoundTruncate,
			expected:    0,
			expectedErr: ErrZeroDenominator,
		},
		{
			name:   "overflow case",
			amount: math.MaxInt64,
			ratio: Ratio[currency.EUR, currency.USD]{
				Numerator:   2,
				Denominator: 1,
			},
			mode:        RoundTruncate,
			expected:    0,
			expectedErr: ErrOverflow,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := NewMoney[currency.EUR](tt.amount)
			converted, result, err := Convert(m, tt.ratio, tt.mode)

			if err != tt.expectedErr {
				t.Errorf("Convert() error = %v, expected error %v", err, tt.expectedErr)
				return
			}

			if tt.expectedErr == nil {
				if converted.Amount() != tt.expected {
					t.Errorf("Convert() = %v, expected %v", converted.Amount(), tt.expected)
				}

				if result.Amount != tt.expected {
					t.Errorf("Convert() result amount = %v, expected %v", result.Amount, tt.expected)
				}

				if tt.amount == 0 {
					if result.ActualRate.Denominator != 1 || result.ActualRate.Numerator != 0 {
						t.Errorf("Convert() actual rate for zero amount = %+v, expected numerator=0 denominator=1", result.ActualRate)
					}
				}
			}
		})
	}
}

func TestConvertRoundingModes(t *testing.T) {
	tests := []struct {
		name     string
		amount   int64
		ratio    Ratio[currency.EUR, currency.USD]
		mode     RoundingMode
		expected int64
	}{
		{
			name:   "truncate positive half",
			amount: 1,
			ratio: Ratio[currency.EUR, currency.USD]{
				Numerator:   1,
				Denominator: 2,
			},
			mode:     RoundTruncate,
			expected: 0,
		},
		{
			name:   "half up positive tie",
			amount: 5,
			ratio: Ratio[currency.EUR, currency.USD]{
				Numerator:   1,
				Denominator: 2,
			},
			mode:     RoundHalfUp,
			expected: 3,
		},
		{
			name:   "half even positive tie",
			amount: 5,
			ratio: Ratio[currency.EUR, currency.USD]{
				Numerator:   1,
				Denominator: 2,
			},
			mode:     RoundHalfEven,
			expected: 2,
		},
		{
			name:   "half up negative tie",
			amount: -1,
			ratio: Ratio[currency.EUR, currency.USD]{
				Numerator:   1,
				Denominator: 2,
			},
			mode:     RoundHalfUp,
			expected: -1,
		},
		{
			name:   "half even negative tie",
			amount: -1,
			ratio: Ratio[currency.EUR, currency.USD]{
				Numerator:   1,
				Denominator: 2,
			},
			mode:     RoundHalfEven,
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := NewMoney[currency.EUR](tt.amount)
			converted, _, err := Convert(m, tt.ratio, tt.mode)
			if err != nil {
				t.Fatalf("Convert() unexpected error = %v", err)
			}

			if converted.Amount() != tt.expected {
				t.Errorf("Convert() = %d, expected %d", converted.Amount(), tt.expected)
			}
		})
	}
}

func TestConvertInvalidMode(t *testing.T) {
	m := NewMoney[currency.EUR](1)
	_, _, err := Convert(m, Ratio[currency.EUR, currency.USD]{Numerator: 1, Denominator: 2}, RoundingMode(99))
	if !errors.Is(err, ErrInvalidRoundingMode) {
		t.Fatalf("expected ErrInvalidRoundingMode, got %v", err)
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
				if m.Amount() != tt.expected {
					t.Fatalf("amount = %d, expected %d", m.Amount(), tt.expected)
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
			if m.Amount() != tt.expected {
				t.Fatalf("amount = %d, expected %d", m.Amount(), tt.expected)
			}
		})
	}
}

func TestMoneyValueAndScan(t *testing.T) {
	original := NewMoney[currency.USD](1050)
	v, err := original.Value()
	if err != nil {
		t.Fatalf("Value() error = %v", err)
	}

	var scanned Money[currency.USD]
	if err := scanned.Scan(v); err != nil {
		t.Fatalf("Scan() error = %v", err)
	}
	if scanned.Amount() != original.Amount() {
		t.Fatalf("scanned amount = %d, expected %d", scanned.Amount(), original.Amount())
	}

	if err := scanned.Scan(int64(99)); err != nil {
		t.Fatalf("Scan(int64) error = %v", err)
	}
	if scanned.Amount() != 99 {
		t.Fatalf("scanned int64 amount = %d, expected 99", scanned.Amount())
	}

	if err := scanned.Scan(nil); err == nil {
		t.Fatal("expected error when scanning nil")
	}

	err = scanned.Scan(`{"amount":"100","currency":"EUR"}`)
	if err == nil {
		t.Fatal("expected currency mismatch error")
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
			expected: `{"amount":"1050","currency":"USD"}`,
		},
		{
			name:     "marshal zero",
			money:    NewMoney[currency.USD](0),
			expected: `{"amount":"0","currency":"USD"}`,
		},
		{
			name:     "marshal negative",
			money:    NewMoney[currency.USD](-1050),
			expected: `{"amount":"-1050","currency":"USD"}`,
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
			if unmarshaledMoney.Amount() != tt.money.Amount() {
				t.Errorf("UnmarshalJSON() amount = %v, expected %v", unmarshaledMoney.Amount(), tt.money.Amount())
			}
		})
	}
}

func TestJSONInvalidAmount(t *testing.T) {
	var unmarshaledMoney Money[currency.USD]
	err := json.Unmarshal([]byte(`{"amount":"10oops","currency":"USD"}`), &unmarshaledMoney)
	if err == nil {
		t.Fatal("expected error for invalid amount format")
	}
}

func TestValidate(t *testing.T) {
	tests := []struct {
		name        string
		amount      int64
		min         int64
		max         int64
		expectedErr error
	}{
		{
			name:   "valid amount",
			amount: 500,
			min:    0,
			max:    1000,
		},
		{
			name:   "at minimum",
			amount: 0,
			min:    0,
			max:    1000,
		},
		{
			name:   "at maximum",
			amount: 1000,
			min:    0,
			max:    1000,
		},
		{
			name:        "below minimum",
			amount:      -1,
			min:         0,
			max:         1000,
			expectedErr: ErrValidation,
		},
		{
			name:        "above maximum",
			amount:      1001,
			min:         0,
			max:         1000,
			expectedErr: ErrValidation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := NewMoney[currency.USD](tt.amount)
			err := m.Validate(tt.min, tt.max)

			if tt.expectedErr == nil && err == nil {
				return
			}

			if tt.expectedErr == nil && err != nil {
				t.Errorf("Validate() unexpected error: %v", err)
			}

			if errors.Unwrap(err) != ErrValidation {
				t.Errorf("unwrapped error = %v, want %v", errors.Unwrap(err), ErrValidation)
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
			name:        "zero ratio",
			amount:      100,
			ratios:      []int64{0, 1},
			expectedErr: ErrNegativeOrZeroRatios,
		},
		{
			name:        "negative ratio",
			amount:      100,
			ratios:      []int64{-1, 1},
			expectedErr: ErrNegativeOrZeroRatios,
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
			expected: []int64{333, 333, 334},
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
			allocation, err := money.Allocate(tt.ratios)

			if err != tt.expectedErr {
				t.Errorf("expected error %v, got %v", tt.expectedErr, err)
				return
			}

			if tt.expectedErr == nil {
				for i, expected := range tt.expected {
					if allocation.Parts[i].Amount() != expected {
						t.Errorf("part %d: expected %d, got %d", i, expected, allocation.Parts[i].Amount())
					}
				}

				sum := int64(0)
				for _, part := range allocation.Parts {
					sum += part.Amount()
				}
				if sum != tt.amount {
					t.Errorf("sum of parts (%d) does not equal original amount (%d)", sum, tt.amount)
				}

				if allocation.Total.Amount() != money.Amount() {
					t.Errorf("Total field (%d) does not match original amount (%d)",
						allocation.Total.Amount(), money.Amount())
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
			expected: []string{"$50.00", "$50.00"},
		},
		{
			name:     "split $100 in thirds",
			amount:   10000, // $100.00
			ratios:   []int64{1, 1, 1},
			expected: []string{"$33.33", "$33.33", "$33.34"},
		},
		{
			name:     "split $50.50 by ratio 1:2",
			amount:   5050, // $50.50
			ratios:   []int64{1, 2},
			expected: []string{"$16.83", "$33.67"},
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
			expectedErr: ErrNegativeOrZeroRatios,
		},
		{
			name:        "zero ratio",
			amount:      10000,
			ratios:      []int64{0, 1},
			expectedErr: ErrNegativeOrZeroRatios,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			money := NewMoney[currency.USD](tt.amount)
			allocation, err := money.Allocate(tt.ratios)
			if err != tt.expectedErr {
				t.Errorf("expected error %v, got %v", tt.expectedErr, err)
				return
			}
			if err != nil {
				return
			}

			for i, expected := range tt.expected {
				if allocation.Parts[i].String() != expected {
					t.Errorf("part %d: expected %s, got %s",
						i, expected, allocation.Parts[i].String())
				}
			}

			sum := int64(0)
			for _, part := range allocation.Parts {
				sum += part.Amount()
			}
			if sum != tt.amount {
				t.Errorf("sum of parts (%d) does not equal original amount (%d)",
					sum, tt.amount)
			}
		})
	}
}

func ExampleConvert() {
	eur := NewMoney[currency.EUR](500) // €5.00
	ratio := Ratio[currency.EUR, currency.USD]{
		Numerator:   104565, // 1.04565 represented as 104565/100000
		Denominator: 100000,
	}
	eurInUsd, _, err := Convert(eur, ratio, RoundTruncate)
	if err != nil {
		panic(err)
	}
	fmt.Println(eurInUsd)
	// Output: $5.22
}

func ExampleMoney_Allocate() {
	dollars := NewMoney[currency.USD](10000) // 100.00 USD

	// Allocate in ratio of [1,1,2]
	allocation, err := dollars.Allocate([]int64{1, 1, 2})
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for i, part := range allocation.Parts {
		fmt.Printf("Part %d: %v\n", i+1, part)
	}

	// Output:
	// Part 1: $25.00
	// Part 2: $25.00
	// Part 3: $50.00
}

func ExampleMoney_Distribute() {
	dollars := NewMoney[currency.USD](10000) // 100.00 USD

	// Distribute into 3 chunks
	dist, err := dollars.Distribute(3)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Smaller chunks: %d x $%.2f\n",
		dist.SmallerCount,
		float64(dist.SmallerChunkSize)/100)
	fmt.Printf("Larger chunks: %d x $%.2f\n",
		dist.LargerCount,
		float64(dist.LargerChunkSize)/100)

	// Verify total
	total := (dist.SmallerChunkSize * dist.SmallerCount) +
		(dist.LargerChunkSize * dist.LargerCount)
	fmt.Printf("Total: $%.2f\n", float64(total)/100)

	// Output:
	// Smaller chunks: 2 x $33.33
	// Larger chunks: 1 x $33.34
	// Total: $100.00
}

func TestGeneratedFormatContracts(t *testing.T) {
	tests := []struct {
		name     string
		money    Money[currency.USD]
		loc      locale.Locale
		expected currency.FormatInfo
	}{
		{
			name:     "en positive contract",
			money:    NewMoney[currency.USD](1234567),
			loc:      locale.EN,
			expected: currency.USD{}.FormatInfo(locale.EN),
		},
		{
			name:     "de negative contract",
			money:    NewMoney[currency.USD](-1234567),
			loc:      locale.DE,
			expected: currency.USD{}.FormatInfo(locale.DE),
		},
		{
			name:     "fr zero contract",
			money:    NewMoney[currency.USD](0),
			loc:      locale.FR,
			expected: currency.USD{}.FormatInfo(locale.FR),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			formatted := tt.money.Format(tt.loc)
			if !strings.Contains(formatted, tt.expected.Symbol) {
				t.Fatalf("formatted value %q does not include symbol %q", formatted, tt.expected.Symbol)
			}

			if tt.money.Amount() != 0 && tt.money.Currency.MinorUnits() > 0 &&
				!strings.Contains(formatted, tt.expected.DecimalSeparator) {
				t.Fatalf("formatted value %q does not include decimal separator %q", formatted, tt.expected.DecimalSeparator)
			}

			if tt.money.Amount() < 0 && !strings.Contains(formatted, tt.expected.MinusSign) {
				t.Fatalf("formatted value %q does not include minus sign %q", formatted, tt.expected.MinusSign)
			}

			if absInt64(tt.money.Amount()) >= 1000 && tt.expected.GroupSeparator != "" &&
				!strings.Contains(formatted, tt.expected.GroupSeparator) {
				t.Fatalf("formatted value %q does not include group separator %q", formatted, tt.expected.GroupSeparator)
			}
		})
	}
}

func FuzzMoneyUnmarshalJSON(f *testing.F) {
	f.Add("100", "USD")
	f.Add("-50", "USD")
	f.Add("abc", "USD")

	f.Fuzz(func(t *testing.T, amount, curr string) {
		payload := fmt.Sprintf(`{"amount":%q,"currency":%q}`, amount, curr)
		var m Money[currency.USD]
		_ = json.Unmarshal([]byte(payload), &m)
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

		total := (dist.SmallerChunkSize * dist.SmallerCount) + (dist.LargerChunkSize * dist.LargerCount)
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
		allocation, err := m.Allocate(ratios)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		sum := int64(0)
		for _, part := range allocation.Parts {
			sum += part.Amount()
		}

		if sum != amount {
			t.Fatalf("allocation sum = %d, amount = %d", sum, amount)
		}
	})
}

func FuzzConvert(f *testing.F) {
	f.Add(int64(100), int64(1), int64(3))
	f.Add(int64(-100), int64(1), int64(2))

	f.Fuzz(func(t *testing.T, amount, numerator, denominator int64) {
		if denominator == 0 {
			return
		}

		m := NewMoney[currency.EUR](amount)
		ratio := Ratio[currency.EUR, currency.USD]{
			Numerator:   numerator,
			Denominator: denominator,
		}

		converted, result, err := Convert(m, ratio, RoundHalfEven)
		if errors.Is(err, ErrOverflow) {
			return
		}
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if converted.Amount() != result.Amount {
			t.Fatalf("result mismatch: converted=%d result=%d", converted.Amount(), result.Amount)
		}

		if amount == 0 && result.ActualRate.Denominator != 1 {
			t.Fatalf("zero amount must produce denominator=1, got %d", result.ActualRate.Denominator)
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
	if result.Amount() != 150 {
		t.Errorf("MustAdd = %d, want %d", result.Amount(), 150)
	}

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("MustAdd should have panicked on overflow")
		}
	}()
	mMax := NewMoney[currency.USD](math.MaxInt64)
	_ = mMax.MustAdd(NewMoney[currency.USD](1))
}

func TestMustSub(t *testing.T) {
	m1 := NewMoney[currency.USD](100)
	m2 := NewMoney[currency.USD](50)

	result := m1.MustSub(m2)
	if result.Amount() != 50 {
		t.Errorf("MustSub = %d, want %d", result.Amount(), 50)
	}

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("MustSub should have panicked on overflow")
		}
	}()
	mMin := NewMoney[currency.USD](math.MinInt64)
	_ = mMin.MustSub(NewMoney[currency.USD](1))
}

func TestMustMul(t *testing.T) {
	m := NewMoney[currency.USD](100)

	result := m.MustMul(3)
	if result.Amount() != 300 {
		t.Errorf("MustMul = %d, want %d", result.Amount(), 300)
	}

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("MustMul should have panicked on overflow")
		}
	}()
	mMax := NewMoney[currency.USD](math.MaxInt64)
	_ = mMax.MustMul(2)
}

