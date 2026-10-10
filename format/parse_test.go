package format

import (
	"errors"
	"math"
	"math/big"
	"testing"

	"github.com/khatibomar/fulus"
	"github.com/khatibomar/fulus/currency"
	"github.com/khatibomar/fulus/locale"
)

func TestParse(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		parse   func() (int64, error)
		want    int64
		wantErr error
	}{
		{
			name: "indian grouping",
			parse: func() (int64, error) {
				m, err := Parse[currency.INR]("-₹1,23,45,678.90", locale.EN_IN)
				return amount64(m), err
			},
			want: -1234567890,
		},
		{
			name: "negative subpattern",
			parse: func() (int64, error) {
				m, err := Parse[currency.CHF]("CHF-1'234.56", locale.DE_CH)
				return amount64(m), err
			},
			want: -123456,
		},
		{
			name: "space instead of no-break space",
			parse: func() (int64, error) {
				m, err := Parse[currency.EUR]("1 234,56 €", locale.FR)
				return amount64(m), err
			},
			want: 123456,
		},
		{
			name: "ascii minus for unicode minus sign",
			parse: func() (int64, error) {
				m, err := Parse[currency.EUR]("-1 234,56 €", locale.SV)
				return amount64(m), err
			},
			want: -123456,
		},
		{
			name: "fewer fraction digits",
			parse: func() (int64, error) {
				m, err := Parse[currency.USD]("$1,234.5", locale.EN)
				return amount64(m), err
			},
			want: 123450,
		},
		{
			name: "no fraction",
			parse: func() (int64, error) {
				m, err := Parse[currency.USD]("$12", locale.EN)
				return amount64(m), err
			},
			want: 1200,
		},
		{
			name: "missing symbol",
			parse: func() (int64, error) {
				m, err := Parse[currency.USD]("1,234.56", locale.EN)
				return amount64(m), err
			},
			wantErr: fulus.ErrInvalidAmountFormat,
		},
		{
			name: "wrong decimal separator",
			parse: func() (int64, error) {
				m, err := Parse[currency.EUR]("1.234.56 €", locale.DE)
				return amount64(m), err
			},
			wantErr: fulus.ErrInvalidAmountFormat,
		},
		{
			name: "wrong group size",
			parse: func() (int64, error) {
				m, err := Parse[currency.USD]("$12,34.00", locale.EN)
				return amount64(m), err
			},
			wantErr: fulus.ErrInvalidAmountFormat,
		},
		{
			name: "wrong indian group size",
			parse: func() (int64, error) {
				m, err := Parse[currency.INR]("₹1,234,567.00", locale.EN_IN)
				return amount64(m), err
			},
			wantErr: fulus.ErrInvalidAmountFormat,
		},
		{
			name: "no group separators",
			parse: func() (int64, error) {
				m, err := Parse[currency.INR]("₹1234567.00", locale.EN_IN)
				return amount64(m), err
			},
			want: 123456700,
		},
		{
			name: "too many fraction digits",
			parse: func() (int64, error) {
				m, err := Parse[currency.USD]("$1.234", locale.EN)
				return amount64(m), err
			},
			wantErr: fulus.ErrScaleMismatch,
		},
		{
			name: "symbol only",
			parse: func() (int64, error) {
				m, err := Parse[currency.USD]("$", locale.EN)
				return amount64(m), err
			},
			wantErr: fulus.ErrInvalidAmountFormat,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := tt.parse()
			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("Parse() error = %v, want %v", err, tt.wantErr)
			}
			if err == nil && got != tt.want {
				t.Errorf("Parse() = %d, want %d", got, tt.want)
			}
		})
	}
}

func TestParseRoundTrip(t *testing.T) {
	t.Parallel()

	locales := []locale.Locale{
		locale.EN, locale.EN_IN, locale.EN_CH, locale.DE, locale.DE_CH, locale.FR, locale.FR_CH,
		locale.NL, locale.SV, locale.FI, locale.PL, locale.ES, locale.PT_PT, locale.IT,
		locale.AR, locale.AR_EG, locale.FA, locale.HE, locale.UR, locale.HI, locale.BN,
		locale.JA, locale.ZH, locale.KO, locale.RU, locale.TR, locale.TOK, locale.SW,
	}
	maxInt128 := new(big.Int).Sub(new(big.Int).Lsh(big.NewInt(1), 127), big.NewInt(1))
	minInt128 := new(big.Int).Neg(new(big.Int).Lsh(big.NewInt(1), 127))
	amounts := []*big.Int{maxInt128, minInt128}
	for _, a := range []int64{0, 1, -1, 5, 1234, -123456, 1234567890, math.MaxInt64, math.MinInt64} {
		amounts = append(amounts, big.NewInt(a))
	}

	for _, c := range currency.All() {
		for _, loc := range locales {
			for _, amount := range amounts {
				m, err := fulus.NewAnyMoneyFromDecimal(canonical(amount, c.MinorUnits()), c)
				if err != nil {
					t.Fatal(err)
				}
				formatted := AnyMoney(m, loc)
				got, err := ParseAny(formatted, c, loc)
				if err != nil || got.Decimal() != m.Decimal() {
					t.Errorf("%s %s: ParseAny(%+q) = %s, %v; want %s", c.Code(), loc, formatted, got.Decimal(), err, m.Decimal())
				}
			}
		}
	}
}
