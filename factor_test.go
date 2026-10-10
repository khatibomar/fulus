package fulus

import (
	"errors"
	"math"
	"testing"

	"github.com/khatibomar/fulus/currency"
)

func TestParseShortDecimal(t *testing.T) {
	t.Parallel()

	tests := []struct {
		factor string
		num    int64
		den    int64
		ok     bool
	}{
		{factor: "0.0825", num: 825, den: 10000, ok: true},
		{factor: "-1.5", num: -15, den: 10, ok: true},
		{factor: "+2", num: 2, den: 1, ok: true},
		{factor: ".5", num: 5, den: 10, ok: true},
		{factor: "1234567890123456", num: 1234567890123456, den: 1, ok: true},
		{factor: "12345678901234567", ok: false},
		{factor: "1/3", ok: false},
		{factor: "1e3", ok: false},
		{factor: "1.2.3", ok: false},
		{factor: ".", ok: false},
		{factor: "-", ok: false},
		{factor: "", ok: false},
	}
	for _, tt := range tests {
		t.Run(tt.factor, func(t *testing.T) {
			t.Parallel()
			num, den, ok := parseShortDecimal(tt.factor)
			if ok != tt.ok || (ok && (num != tt.num || den != tt.den)) {
				t.Errorf("parseShortDecimal(%q) = %d/%d, %v; want %d/%d, %v", tt.factor, num, den, ok, tt.num, tt.den, tt.ok)
			}
		})
	}
}

func TestParseFactor(t *testing.T) {
	t.Parallel()

	tests := []struct {
		factor  string
		num     int64
		den     int64
		wantErr error
	}{
		{factor: "0.0825", num: 33, den: 400},
		{factor: "8.25%", num: 33, den: 400},
		{factor: "-1.5", num: -3, den: 2},
		{factor: "1/3", num: 1, den: 3},
		{factor: "-2/6", num: -1, den: 3},
		{factor: "1/3%", num: 1, den: 300},
		{factor: "0", num: 0, den: 1},
		{factor: "12345678901234567", num: 12345678901234567, den: 1},
		{factor: "0.00000000000000000001", wantErr: ErrOverflow},
		{factor: "1e3", wantErr: ErrInvalidFactor},
		{factor: "1e999999999", wantErr: ErrInvalidFactor},
		{factor: "abc", wantErr: ErrInvalidFactor},
		{factor: "1/0", wantErr: ErrInvalidFactor},
		{factor: "", wantErr: ErrInvalidFactor},
		{factor: "%", wantErr: ErrInvalidFactor},
	}
	for _, tt := range tests {
		t.Run(tt.factor, func(t *testing.T) {
			t.Parallel()
			f, err := ParseFactor(tt.factor)
			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("ParseFactor(%q) error = %v, want %v", tt.factor, err, tt.wantErr)
			}
			if num, den := f.Fraction(); err == nil && (num != tt.num || den != tt.den) {
				t.Errorf("ParseFactor(%q) = %d/%d, want %d/%d", tt.factor, num, den, tt.num, tt.den)
			}
		})
	}
}

func TestNewFactor(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		num, den int64
		wantNum  int64
		wantDen  int64
		wantErr  error
	}{
		{name: "reduces", num: 15, den: 100, wantNum: 3, wantDen: 20},
		{name: "negative denominator", num: 1, den: -3, wantNum: -1, wantDen: 3},
		{name: "zero", num: 0, den: 5, wantNum: 0, wantDen: 1},
		{name: "zero denominator", num: 1, den: 0, wantErr: ErrDivisionByZero},
		{name: "smallest numerator with negative denominator", num: math.MinInt64, den: -1, wantErr: ErrOverflow},
		{name: "smallest denominator", num: 1, den: math.MinInt64, wantErr: ErrOverflow},
		{name: "smallest numerator", num: math.MinInt64, den: 2, wantNum: math.MinInt64 / 2, wantDen: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			f, err := NewFactor(tt.num, tt.den)
			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("NewFactor() error = %v, want %v", err, tt.wantErr)
			}
			if num, den := f.Fraction(); err == nil && (num != tt.wantNum || den != tt.wantDen) {
				t.Errorf("NewFactor() = %d/%d, want %d/%d", num, den, tt.wantNum, tt.wantDen)
			}
		})
	}
}

func TestMulFactor(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		amount  int64
		factor  Factor
		mode    RoundingMode
		want    int64
		wantErr error
	}{
		{name: "sales tax", amount: 10000, factor: MustParseFactor("0.0825"), mode: RoundHalfUp, want: 825},
		{name: "tax half up", amount: 1050, factor: MustParseFactor("8.25%"), mode: RoundHalfUp, want: 87},
		{name: "15 percent", amount: 1999, factor: Percent(15), mode: RoundHalfEven, want: 300},
		{name: "25 basis points", amount: 100000, factor: Bps(25), mode: RoundHalfEven, want: 250},
		{name: "one third floor", amount: -100, factor: MustParseFactor("1/3"), mode: RoundFloor, want: -34},
		{name: "one third ceiling", amount: -100, factor: MustParseFactor("1/3"), mode: RoundCeiling, want: -33},
		{name: "negative factor", amount: 100, factor: MustParseFactor("-1.5"), mode: RoundTruncate, want: -150},
		{name: "intermediate product larger than int64", amount: math.MaxInt64, factor: MustParseFactor("3/4"), mode: RoundTruncate, want: 6917529027641081855},
		{name: "overflow", amount: math.MaxInt64, factor: MustParseFactor("1.5"), mode: RoundTruncate, wantErr: ErrOverflow},
		{name: "zero factor", amount: 100, factor: Factor{}, mode: RoundTruncate, wantErr: ErrInvalidFactor},
		{name: "invalid mode", amount: 100, factor: MustParseFactor("1/3"), mode: RoundingMode(-1), wantErr: ErrInvalidRoundingMode},
		{name: "unnecessary inexact", amount: 100, factor: MustParseFactor("1/3"), mode: RoundUnnecessary, wantErr: ErrInexact},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := NewMoney[currency.USD](tt.amount).MulFactor(tt.factor, tt.mode)
			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("MulFactor() error = %v, want %v", err, tt.wantErr)
			}
			if err == nil && got.Amount() != tt.want {
				t.Errorf("MulFactor() = %d, want %d", got.Amount(), tt.want)
			}
		})
	}
}

func TestFactorString(t *testing.T) {
	t.Parallel()

	for in, want := range map[string]string{"8.25%": "33/400", "2": "2", "-0.5": "-1/2"} {
		if got := MustParseFactor(in).String(); got != want {
			t.Errorf("String() = %q, want %q", got, want)
		}
	}
}
