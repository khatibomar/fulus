package fulusproto

import (
	"errors"
	"math"
	"testing"

	"github.com/khatibomar/fulus"
	"github.com/khatibomar/fulus/currency"
	"google.golang.org/genproto/googleapis/type/money"
)

type token struct{}

func (token) Code() string    { return "TOKEN" }
func (token) MinorUnits() int { return 12 }

func TestToMoney(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		got       func() (*money.Money, error)
		wantUnits int64
		wantNanos int32
		wantErr   error
	}{
		{name: "positive", got: func() (*money.Money, error) { return ToMoney(fulus.NewMoney[currency.USD](1050)) }, wantUnits: 10, wantNanos: 500_000_000},
		{name: "negative", got: func() (*money.Money, error) { return ToMoney(fulus.NewMoney[currency.USD](-1050)) }, wantUnits: -10, wantNanos: -500_000_000},
		{name: "negative below one", got: func() (*money.Money, error) { return ToMoney(fulus.NewMoney[currency.USD](-5)) }, wantUnits: 0, wantNanos: -50_000_000},
		{name: "no minor units", got: func() (*money.Money, error) { return ToMoney(fulus.NewMoney[currency.JPY](-1234)) }, wantUnits: -1234},
		{name: "three minor units", got: func() (*money.Money, error) { return ToMoney(fulus.NewMoney[currency.BHD](1)) }, wantNanos: 1_000_000},
		{name: "12 minor units with zeros", got: func() (*money.Money, error) { return ToMoney(fulus.NewMoney[token](1_000)) }, wantNanos: 1},
		{name: "12 minor units inexact", got: func() (*money.Money, error) { return ToMoney(fulus.NewMoney[token](1)) }, wantErr: fulus.ErrInexact},
		{name: "units overflow", got: func() (*money.Money, error) {
			m, _ := fulus.ParseMoney[currency.JPY]("9223372036854775808")
			return ToMoney(m)
		}, wantErr: fulus.ErrOverflow},
		{name: "smallest units", got: func() (*money.Money, error) {
			m, _ := fulus.ParseMoney[currency.JPY]("-9223372036854775808")
			return ToMoney(m)
		}, wantUnits: math.MinInt64},
		{name: "negative units overflow", got: func() (*money.Money, error) {
			m, _ := fulus.ParseMoney[currency.JPY]("-9223372036854775809")
			return ToMoney(m)
		}, wantErr: fulus.ErrOverflow},
		{name: "smallest units with nanos", got: func() (*money.Money, error) {
			m, _ := fulus.ParseMoney[currency.USD]("-9223372036854775808.50")
			return ToMoney(m)
		}, wantUnits: math.MinInt64, wantNanos: -500_000_000},
		{name: "any", got: func() (*money.Money, error) { return ToMoneyAny(fulus.NewAnyMoney(250, currency.EUR{})) }, wantUnits: 2, wantNanos: 500_000_000},
		{name: "any without currency", got: func() (*money.Money, error) { return ToMoneyAny(fulus.AnyMoney{}) }, wantErr: fulus.ErrUnknownCurrency},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := tt.got()
			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("error = %v, want %v", err, tt.wantErr)
			}
			if err == nil && (got.GetUnits() != tt.wantUnits || got.GetNanos() != tt.wantNanos) {
				t.Errorf("got %d units %d nanos, want %d %d", got.GetUnits(), got.GetNanos(), tt.wantUnits, tt.wantNanos)
			}
		})
	}
}

func TestFromMoney(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		in      *money.Money
		want    string
		wantErr error
	}{
		{name: "positive", in: &money.Money{CurrencyCode: "USD", Units: 10, Nanos: 500_000_000}, want: "10.50"},
		{name: "negative", in: &money.Money{CurrencyCode: "USD", Units: -10, Nanos: -500_000_000}, want: "-10.50"},
		{name: "negative below one", in: &money.Money{CurrencyCode: "USD", Nanos: -10_000_000}, want: "-0.01"},
		{name: "smallest units", in: &money.Money{CurrencyCode: "USD", Units: math.MinInt64}, want: "-9223372036854775808.00"},
		{name: "too many digits", in: &money.Money{CurrencyCode: "USD", Nanos: 1}, wantErr: fulus.ErrScaleMismatch},
		{name: "mixed signs", in: &money.Money{CurrencyCode: "USD", Units: 1, Nanos: -1}, wantErr: fulus.ErrInvalidAmountFormat},
		{name: "nanos out of range", in: &money.Money{CurrencyCode: "USD", Nanos: 1_000_000_000}, wantErr: fulus.ErrInvalidAmountFormat},
		{name: "currency mismatch", in: &money.Money{CurrencyCode: "EUR", Units: 1}, wantErr: fulus.ErrCurrencyMismatch},
		{name: "lower case currency", in: &money.Money{CurrencyCode: "usd", Units: 1}, want: "1.00"},
		{name: "nil", in: nil, wantErr: fulus.ErrInvalidAmountFormat},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := FromMoney[currency.USD](tt.in)
			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("FromMoney() error = %v, want %v", err, tt.wantErr)
			}
			if err == nil && got.Decimal() != tt.want {
				t.Errorf("FromMoney() = %s, want %s", got.Decimal(), tt.want)
			}
		})
	}

	got, err := FromMoneyAny(&money.Money{CurrencyCode: "KWD", Units: -1, Nanos: -5_000_000})
	if err != nil || got.String() != "KWD -1.005" {
		t.Errorf("FromMoneyAny() = %v, %v", got, err)
	}
	if _, err := FromMoneyAny(&money.Money{CurrencyCode: "XYZ"}); !errors.Is(err, fulus.ErrUnknownCurrency) {
		t.Errorf("FromMoneyAny() unknown currency error = %v", err)
	}
}

func FuzzRoundTrip(f *testing.F) {
	f.Add(int64(1050))
	f.Add(int64(math.MinInt64))
	f.Add(int64(math.MaxInt64))

	f.Fuzz(func(t *testing.T, amount int64) {
		m := fulus.NewMoney[currency.BHD](amount)
		p, err := ToMoney(m)
		if err != nil {
			t.Fatal(err)
		}
		back, err := FromMoney[currency.BHD](p)
		if err != nil || back != m {
			t.Fatalf("round trip of %s = %s, %v", m, back, err)
		}
	})
}
