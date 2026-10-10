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

func TestRoundingModeValidation(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		amount  int64
		mode    RoundingMode
		want    int64
		wantErr error
	}{
		{name: "zero mode", amount: 5, mode: 0, wantErr: ErrInvalidRoundingMode},
		{name: "unnecessary exact", amount: 6, mode: RoundUnnecessary, want: 3},
		{name: "unnecessary inexact", amount: 5, mode: RoundUnnecessary, wantErr: ErrInexact},
		{name: "unnecessary negative inexact", amount: -5, mode: RoundUnnecessary, wantErr: ErrInexact},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := NewMoney[currency.USD](tt.amount).Div(2, tt.mode)
			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("Div() error = %v, expected %v", err, tt.wantErr)
			}
			if err == nil && got.Amount() != tt.want {
				t.Errorf("Div() = %d, expected %d", got.Amount(), tt.want)
			}
		})
	}
}

func TestRoundingModeString(t *testing.T) {
	t.Parallel()

	tests := map[RoundingMode]string{
		RoundTruncate:    "Truncate",
		RoundHalfEven:    "HalfEven",
		RoundUnnecessary: "Unnecessary",
		0:                "RoundingMode(0)",
	}
	for mode, want := range tests {
		if got := mode.String(); got != want {
			t.Errorf("String() = %q, expected %q", got, want)
		}
	}
}
