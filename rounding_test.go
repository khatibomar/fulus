package fulus

import (
	"errors"
	"math"
	"math/big"
	"testing"

	"github.com/khatibomar/fulus/currency"
)

func TestDivideWithRoundingModes(t *testing.T) {
	t.Parallel()

	// Each row divides numerator by 10, so 25 is 2.5 and -25 is -2.5.
	numerators := []int64{25, 24, 26, 35, 20, -25, -24, -26, -35, -20}
	tests := []struct {
		mode RoundingMode
		want []int64
	}{
		{RoundTruncate, []int64{2, 2, 2, 3, 2, -2, -2, -2, -3, -2}},
		{RoundHalfUp, []int64{3, 2, 3, 4, 2, -3, -2, -3, -4, -2}},
		{RoundHalfEven, []int64{2, 2, 3, 4, 2, -2, -2, -3, -4, -2}},
		{RoundHalfDown, []int64{2, 2, 3, 3, 2, -2, -2, -3, -3, -2}},
		{RoundUp, []int64{3, 3, 3, 4, 2, -3, -3, -3, -4, -2}},
		{RoundCeiling, []int64{3, 3, 3, 4, 2, -2, -2, -2, -3, -2}},
		{RoundFloor, []int64{2, 2, 2, 3, 2, -3, -3, -3, -4, -2}},
	}

	for _, tt := range tests {
		for i, n := range numerators {
			got, err := divideWithRounding(big.NewInt(n), big.NewInt(10), tt.mode)
			if err != nil {
				t.Fatalf("mode %d: unexpected error %v", tt.mode, err)
			}
			if got.Int64() != tt.want[i] {
				t.Errorf("mode %d: %d/10 = %d, want %d", tt.mode, n, got.Int64(), tt.want[i])
			}
		}
	}

	if _, err := divideWithRounding(big.NewInt(20), big.NewInt(10), RoundingMode(99)); !errors.Is(err, ErrInvalidRoundingMode) {
		t.Errorf("invalid mode with exact division: error = %v, want %v", err, ErrInvalidRoundingMode)
	}
}

func TestMulFrac(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		amount      int64
		numerator   int64
		denominator int64
		mode        RoundingMode
		want        int64
		wantErr     error
	}{
		{name: "15 percent", amount: 1999, numerator: 15, denominator: 100, mode: RoundHalfEven, want: 300},
		{name: "one third floor", amount: -100, numerator: 1, denominator: 3, mode: RoundFloor, want: -34},
		{name: "one third ceiling", amount: -100, numerator: 1, denominator: 3, mode: RoundCeiling, want: -33},
		{name: "intermediate product larger than int64", amount: math.MaxInt64, numerator: 3, denominator: 4, mode: RoundTruncate, want: 6917529027641081855},
		{name: "zero denominator", amount: 100, numerator: 1, denominator: 0, mode: RoundTruncate, wantErr: ErrDivisionByZero},
		{name: "overflow", amount: math.MaxInt64, numerator: 2, denominator: 1, mode: RoundTruncate, wantErr: ErrOverflow},
		{name: "invalid mode", amount: 100, numerator: 1, denominator: 3, mode: RoundingMode(-1), wantErr: ErrInvalidRoundingMode},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := NewMoney[currency.USD](tt.amount).MulFrac(tt.numerator, tt.denominator, tt.mode)
			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("MulFrac() error = %v, want %v", err, tt.wantErr)
			}
			if err == nil && got.Amount() != tt.want {
				t.Errorf("MulFrac() = %d, want %d", got.Amount(), tt.want)
			}
		})
	}
}

func TestMulDecimal(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		amount  int64
		factor  string
		mode    RoundingMode
		want    int64
		wantErr error
	}{
		{name: "sales tax", amount: 10000, factor: "0.0825", mode: RoundHalfUp, want: 825},
		{name: "tax half up", amount: 1050, factor: "0.0825", mode: RoundHalfUp, want: 87},
		{name: "fraction", amount: 100, factor: "1/3", mode: RoundHalfEven, want: 33},
		{name: "negative factor", amount: 100, factor: "-1.5", mode: RoundTruncate, want: -150},
		{name: "invalid", amount: 100, factor: "abc", mode: RoundTruncate, wantErr: ErrInvalidFactor},
		{name: "overflow", amount: math.MaxInt64, factor: "1.5", mode: RoundTruncate, wantErr: ErrOverflow},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := NewMoney[currency.USD](tt.amount).MulDecimal(tt.factor, tt.mode)
			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("MulDecimal() error = %v, want %v", err, tt.wantErr)
			}
			if err == nil && got.Amount() != tt.want {
				t.Errorf("MulDecimal() = %d, want %d", got.Amount(), tt.want)
			}
		})
	}
}

func TestRoundCash(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		round   func() (int64, error)
		want    int64
		wantErr error
	}{
		{
			name: "CHF up to 0.05",
			round: func() (int64, error) {
				m, err := NewMoney[currency.CHF](1003).RoundCash(RoundHalfUp)
				return m.Amount(), err
			},
			want: 1005,
		},
		{
			name: "CHF down to 0.00",
			round: func() (int64, error) {
				m, err := NewMoney[currency.CHF](1002).RoundCash(RoundHalfUp)
				return m.Amount(), err
			},
			want: 1000,
		},
		{
			name: "CHF negative",
			round: func() (int64, error) {
				m, err := NewMoney[currency.CHF](-1003).RoundCash(RoundHalfUp)
				return m.Amount(), err
			},
			want: -1005,
		},
		{
			name: "CZK whole crowns",
			round: func() (int64, error) {
				m, err := NewMoney[currency.CZK](12350).RoundCash(RoundHalfEven)
				return m.Amount(), err
			},
			want: 12400,
		},
		{
			name: "USD has no cash rounding",
			round: func() (int64, error) {
				m, err := NewMoney[currency.USD](1003).RoundCash(RoundHalfUp)
				return m.Amount(), err
			},
			want: 1003,
		},
		{
			name: "invalid mode",
			round: func() (int64, error) {
				m, err := NewMoney[currency.USD](1003).RoundCash(RoundingMode(42))
				return m.Amount(), err
			},
			wantErr: ErrInvalidRoundingMode,
		},
		{
			name: "overflow",
			round: func() (int64, error) {
				m, err := NewMoney[currency.CHF](math.MaxInt64).RoundCash(RoundUp)
				return m.Amount(), err
			},
			wantErr: ErrOverflow,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := tt.round()
			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("RoundCash() error = %v, want %v", err, tt.wantErr)
			}
			if err == nil && got != tt.want {
				t.Errorf("RoundCash() = %d, want %d", got, tt.want)
			}
		})
	}
}
