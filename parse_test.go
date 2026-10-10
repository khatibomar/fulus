package fulus

import (
	"errors"
	"math"
	"testing"

	"github.com/khatibomar/fulus/currency"
	"github.com/khatibomar/fulus/locale"
)

func TestParseFormatted(t *testing.T) {
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
				m, err := ParseFormatted[currency.INR]("-₹1,23,45,678.90", locale.EN_IN)
				return m.amount64(), err
			},
			want: -1234567890,
		},
		{
			name: "negative subpattern",
			parse: func() (int64, error) {
				m, err := ParseFormatted[currency.CHF]("CHF-1'234.56", locale.DE_CH)
				return m.amount64(), err
			},
			want: -123456,
		},
		{
			name: "space instead of no-break space",
			parse: func() (int64, error) {
				m, err := ParseFormatted[currency.EUR]("1 234,56 €", locale.FR)
				return m.amount64(), err
			},
			want: 123456,
		},
		{
			name: "ascii minus for unicode minus sign",
			parse: func() (int64, error) {
				m, err := ParseFormatted[currency.EUR]("-1 234,56 €", locale.SV)
				return m.amount64(), err
			},
			want: -123456,
		},
		{
			name: "fewer fraction digits",
			parse: func() (int64, error) {
				m, err := ParseFormatted[currency.USD]("$1,234.5", locale.EN)
				return m.amount64(), err
			},
			want: 123450,
		},
		{
			name: "no fraction",
			parse: func() (int64, error) {
				m, err := ParseFormatted[currency.USD]("$12", locale.EN)
				return m.amount64(), err
			},
			want: 1200,
		},
		{
			name: "missing symbol",
			parse: func() (int64, error) {
				m, err := ParseFormatted[currency.USD]("1,234.56", locale.EN)
				return m.amount64(), err
			},
			wantErr: ErrInvalidAmountFormat,
		},
		{
			name: "wrong decimal separator",
			parse: func() (int64, error) {
				m, err := ParseFormatted[currency.EUR]("1.234.56 €", locale.DE)
				return m.amount64(), err
			},
			wantErr: ErrInvalidAmountFormat,
		},
		{
			name: "wrong group size",
			parse: func() (int64, error) {
				m, err := ParseFormatted[currency.USD]("$12,34.00", locale.EN)
				return m.amount64(), err
			},
			wantErr: ErrInvalidAmountFormat,
		},
		{
			name: "wrong indian group size",
			parse: func() (int64, error) {
				m, err := ParseFormatted[currency.INR]("₹1,234,567.00", locale.EN_IN)
				return m.amount64(), err
			},
			wantErr: ErrInvalidAmountFormat,
		},
		{
			name: "no group separators",
			parse: func() (int64, error) {
				m, err := ParseFormatted[currency.INR]("₹1234567.00", locale.EN_IN)
				return m.amount64(), err
			},
			want: 123456700,
		},
		{
			name: "too many fraction digits",
			parse: func() (int64, error) {
				m, err := ParseFormatted[currency.USD]("$1.234", locale.EN)
				return m.amount64(), err
			},
			wantErr: ErrScaleMismatch,
		},
		{
			name: "symbol only",
			parse: func() (int64, error) {
				m, err := ParseFormatted[currency.USD]("$", locale.EN)
				return m.amount64(), err
			},
			wantErr: ErrInvalidAmountFormat,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := tt.parse()
			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("ParseFormatted() error = %v, want %v", err, tt.wantErr)
			}
			if err == nil && got != tt.want {
				t.Errorf("ParseFormatted() = %d, want %d", got, tt.want)
			}
		})
	}
}

func TestParseFormattedRoundTrip(t *testing.T) {
	t.Parallel()

	locales := []locale.Locale{
		locale.EN, locale.EN_IN, locale.EN_CH, locale.DE, locale.DE_CH, locale.FR, locale.FR_CH,
		locale.NL, locale.SV, locale.FI, locale.PL, locale.ES, locale.PT_PT, locale.IT,
		locale.AR, locale.AR_EG, locale.FA, locale.HE, locale.UR, locale.HI, locale.BN,
		locale.JA, locale.ZH, locale.KO, locale.RU, locale.TR, locale.TOK, locale.SW,
	}
	amounts := []int128{maxInt128, minInt128}
	for _, a := range []int64{0, 1, -1, 5, 1234, -123456, 1234567890, math.MaxInt64, math.MinInt64} {
		amounts = append(amounts, int128FromInt64(a))
	}

	for _, c := range currency.All() {
		for _, loc := range locales {
			info := c.FormatInfo(loc)
			for _, amount := range amounts {
				formatted := formatAmount(amount, c.MinorUnits(), info)
				got, err := parseFormatted(formatted, c.MinorUnits(), info)
				if err != nil || got != amount {
					t.Errorf("%s %s: parse(%+q) = %s, %v; want %s", c.Code(), loc, formatted, got.big(), err, amount.big())
				}
			}
		}
	}
}
