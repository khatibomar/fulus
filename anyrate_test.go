package fulus

import (
	"errors"
	"testing"

	"github.com/khatibomar/fulus/currency"
)

func TestAnyRateConstructors(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		create  func() (AnyRate, error)
		wantNum int64
		wantDen int64
		wantErr error
	}{
		{
			name:    "new reduces",
			create:  func() (AnyRate, error) { return NewAnyRate(10, 4, currency.EUR{}, currency.USD{}) },
			wantNum: 5,
			wantDen: 2,
		},
		{
			name:    "new not positive",
			create:  func() (AnyRate, error) { return NewAnyRate(0, 1, currency.EUR{}, currency.USD{}) },
			wantErr: ErrInvalidExchangeRate,
		},
		{
			name:    "new nil base",
			create:  func() (AnyRate, error) { return NewAnyRate(1, 1, nil, currency.USD{}) },
			wantErr: ErrUnknownCurrency,
		},
		{
			name:    "parse decimal",
			create:  func() (AnyRate, error) { return ParseAnyRate("1.07203", currency.EUR{}, currency.USD{}) },
			wantNum: 107203,
			wantDen: 100000,
		},
		{
			name:    "parse fraction",
			create:  func() (AnyRate, error) { return ParseAnyRate("2/6", currency.EUR{}, currency.USD{}) },
			wantNum: 1,
			wantDen: 3,
		},
		{
			name:    "parse negative",
			create:  func() (AnyRate, error) { return ParseAnyRate("-1", currency.EUR{}, currency.USD{}) },
			wantErr: ErrInvalidExchangeRate,
		},
		{
			name: "parse too many digits",
			create: func() (AnyRate, error) {
				return ParseAnyRate("0.0000000000000000000001", currency.EUR{}, currency.USD{})
			},
			wantErr: ErrOverflow,
		},
		{
			name:    "parse nil quote",
			create:  func() (AnyRate, error) { return ParseAnyRate("1", currency.EUR{}, nil) },
			wantErr: ErrUnknownCurrency,
		},
		{
			name:    "from typed rate",
			create:  func() (AnyRate, error) { return MustParseRate[currency.EUR, currency.USD]("1.25").Any(), nil },
			wantNum: 5,
			wantDen: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			r, err := tt.create()
			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("error = %v, expected %v", err, tt.wantErr)
			}
			if err != nil {
				return
			}
			if n, d := r.Fraction(); n != tt.wantNum || d != tt.wantDen {
				t.Errorf("Fraction() = %d/%d, expected %d/%d", n, d, tt.wantNum, tt.wantDen)
			}
			if r.Base().Code() != "EUR" || r.Quote().Code() != "USD" {
				t.Errorf("currencies = %s/%s, expected EUR/USD", r.Base().Code(), r.Quote().Code())
			}
		})
	}
}

func TestAnyRateAccessors(t *testing.T) {
	t.Parallel()

	r, err := ParseAnyRate("1.25", currency.EUR{}, currency.USD{})
	if err != nil {
		t.Fatal(err)
	}
	if got := r.String(); got != "EUR/USD 1.25" {
		t.Errorf("String() = %q, expected EUR/USD 1.25", got)
	}
	inv := r.Invert()
	if n, d := inv.Fraction(); n != 4 || d != 5 || inv.Base().Code() != "USD" || inv.Quote().Code() != "EUR" {
		t.Errorf("Invert() = %v, expected USD/EUR 0.8", inv)
	}

	var zero AnyRate
	if zero.IsValid() || zero.String() != "0" || zero.Base() != nil || zero.Quote() != nil {
		t.Errorf("zero AnyRate = %v, valid %t", zero, zero.IsValid())
	}
	if (Rate[currency.EUR, currency.USD]{}).Any().IsValid() {
		t.Error("Any() of the zero Rate is valid")
	}
}

func TestAsRate(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		rate    AnyRate
		wantErr error
	}{
		{name: "same currencies", rate: MustParseRate[currency.EUR, currency.USD]("1.25").Any()},
		{name: "other base", rate: MustParseRate[currency.GBP, currency.USD]("1.25").Any(), wantErr: ErrCurrencyMismatch},
		{name: "inverted", rate: MustParseRate[currency.USD, currency.EUR]("0.8").Any(), wantErr: ErrCurrencyMismatch},
		{name: "other minor units", rate: mustAnyRate(t, "1.25", currency.EUR{}, fourDigitUSD{}), wantErr: ErrCurrencyMismatch},
		{name: "zero", rate: AnyRate{}, wantErr: ErrInvalidExchangeRate},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			r, err := AsRate[currency.EUR, currency.USD](tt.rate)
			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("AsRate() error = %v, expected %v", err, tt.wantErr)
			}
			if n, d := r.Fraction(); err == nil && (n != 5 || d != 4) {
				t.Errorf("Fraction() = %d/%d, expected 5/4", n, d)
			}
		})
	}
}

func TestAnyMoneyConvert(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		money    AnyMoney
		rate     AnyRate
		mode     RoundingMode
		want     int64
		wantCode string
		wantErr  error
	}{
		{
			name:     "same minor units",
			money:    NewAnyMoney(10000, currency.EUR{}),
			rate:     mustAnyRate(t, "1.07203", currency.EUR{}, currency.USD{}),
			mode:     RoundTruncate,
			want:     10720,
			wantCode: "USD",
		},
		{
			name:     "fewer minor units",
			money:    NewAnyMoney(100, currency.EUR{}),
			rate:     mustAnyRate(t, "160.25", currency.EUR{}, currency.JPY{}),
			mode:     RoundHalfEven,
			want:     160,
			wantCode: "JPY",
		},
		{
			name:     "more minor units",
			money:    NewAnyMoney(1000, currency.USD{}),
			rate:     mustAnyRate(t, "0.376", currency.USD{}, currency.BHD{}),
			mode:     RoundHalfEven,
			want:     3760,
			wantCode: "BHD",
		},
		{
			name:     "half even tie",
			money:    NewAnyMoney(5, currency.EUR{}),
			rate:     mustAnyRate(t, "1/2", currency.EUR{}, currency.USD{}),
			mode:     RoundHalfEven,
			want:     2,
			wantCode: "USD",
		},
		{
			name:    "unnecessary inexact",
			money:   NewAnyMoney(1, currency.EUR{}),
			rate:    mustAnyRate(t, "1/2", currency.EUR{}, currency.USD{}),
			mode:    RoundUnnecessary,
			wantErr: ErrInexact,
		},
		{
			name:    "other base",
			money:   NewAnyMoney(1, currency.GBP{}),
			rate:    mustAnyRate(t, "1.25", currency.EUR{}, currency.USD{}),
			mode:    RoundHalfEven,
			wantErr: ErrCurrencyMismatch,
		},
		{
			name:    "other minor units",
			money:   NewAnyMoney(1, fourDigitUSD{}),
			rate:    mustAnyRate(t, "0.8", currency.USD{}, currency.EUR{}),
			mode:    RoundHalfEven,
			wantErr: ErrCurrencyMismatch,
		},
		{
			name:    "zero money",
			money:   AnyMoney{},
			rate:    mustAnyRate(t, "1.25", currency.EUR{}, currency.USD{}),
			mode:    RoundHalfEven,
			wantErr: ErrCurrencyMismatch,
		},
		{
			name:    "zero rate",
			money:   NewAnyMoney(1, currency.EUR{}),
			mode:    RoundHalfEven,
			wantErr: ErrInvalidExchangeRate,
		},
		{
			name:    "invalid mode",
			money:   NewAnyMoney(1, currency.EUR{}),
			rate:    mustAnyRate(t, "1.25", currency.EUR{}, currency.USD{}),
			mode:    RoundingMode(99),
			wantErr: ErrInvalidRoundingMode,
		},
		{
			name:    "overflow",
			money:   maxMoney[currency.EUR]().Any(),
			rate:    mustAnyRate(t, "2", currency.EUR{}, currency.USD{}),
			mode:    RoundHalfEven,
			wantErr: ErrOverflow,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := tt.money.Convert(tt.rate, tt.mode)
			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("Convert() error = %v, expected %v", err, tt.wantErr)
			}
			if err != nil {
				return
			}
			if got.amount64() != tt.want || got.Currency().Code() != tt.wantCode {
				t.Errorf("Convert() = %d %s, expected %d %s", got.amount64(), got.Currency().Code(), tt.want, tt.wantCode)
			}
		})
	}
}

func TestAnyMoneyConvertMatchesConvert(t *testing.T) {
	t.Parallel()

	rate := MustParseRate[currency.EUR, currency.JPY]("162.337")
	for _, amount := range []int64{-12345, -1, 0, 1, 99, 1000001} {
		m := NewMoney[currency.EUR](amount)
		want, err := Convert(m, rate, RoundHalfUp)
		if err != nil {
			t.Fatal(err)
		}
		got, err := m.Any().Convert(rate.Any(), RoundHalfUp)
		if err != nil {
			t.Fatal(err)
		}
		if typed, err := As[currency.JPY](got); err != nil || typed != want {
			t.Errorf("Convert(%d) = %v, %v, expected %v", amount, got, err, want)
		}
	}
}

func mustAnyRate(t *testing.T, s string, base, quote currency.Currency) AnyRate {
	t.Helper()
	r, err := ParseAnyRate(s, base, quote)
	if err != nil {
		t.Fatal(err)
	}
	return r
}
