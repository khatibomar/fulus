package fulus

import (
	"encoding/json"
	"errors"
	"math"
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
			}
		})
	}
}

func TestConvert(t *testing.T) {
	tests := []struct {
		name        string
		amount      int64
		ratio       Ratio
		expected    int64
		expectedErr error
	}{
		{
			name:   "simple conversion",
			amount: 10000,
			ratio: Ratio{
				Numerator:   107203,
				Denominator: 100000,
			},
			expected:    10720,
			expectedErr: nil,
		},
		{
			name:   "zero amount",
			amount: 0,
			ratio: Ratio{
				Numerator:   107203,
				Denominator: 100000,
			},
			expected:    0,
			expectedErr: nil,
		},
		{
			name:   "invalid ratio",
			amount: 1000,
			ratio: Ratio{
				Numerator:   1,
				Denominator: 0,
			},
			expected:    0,
			expectedErr: ErrZeroDenominator,
		},
		{
			name:   "overflow case",
			amount: math.MaxInt64,
			ratio: Ratio{
				Numerator:   2,
				Denominator: 1,
			},
			expected:    0,
			expectedErr: ErrOverflow,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := NewMoney[currency.EUR](tt.amount)
			converted, result, err := Convert[currency.USD](m, tt.ratio)

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
			}
		})
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
