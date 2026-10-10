package format

import (
	"math"
	"math/big"
	"strings"
	"testing"

	"github.com/khatibomar/fulus"
	"github.com/khatibomar/fulus/currency"
	"github.com/khatibomar/fulus/locale"
)

// amount64 returns the amount of m in minor units. It panics if the amount does not fit in int64.
func amount64[T currency.Unit](m fulus.Money[T]) int64 {
	v, ok := m.Int64()
	if !ok {
		panic("amount does not fit in int64")
	}
	return v
}

// canonical returns the canonical decimal of an amount in minor units.
func canonical(amount *big.Int, minorUnits int) string {
	digits := new(big.Int).Abs(amount).String()
	sign := ""
	if amount.Sign() < 0 {
		sign = "-"
	}
	if minorUnits <= 0 {
		return sign + digits
	}
	if len(digits) <= minorUnits {
		digits = strings.Repeat("0", minorUnits-len(digits)+1) + digits
	}
	return sign + digits[:len(digits)-minorUnits] + "." + digits[len(digits)-minorUnits:]
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

func TestFormat(t *testing.T) {
	tests := []struct {
		name     string
		money    fulus.Money[currency.USD]
		locale   locale.Locale
		expected string
	}{
		{
			name:     "positive whole number",
			money:    fulus.NewMoney[currency.USD](1000),
			locale:   locale.EN,
			expected: "$10.00",
		},
		{
			name:     "negative whole number",
			money:    fulus.NewMoney[currency.USD](-1000),
			locale:   locale.EN,
			expected: "-$10.00",
		},
		{
			name:     "minimum int64",
			money:    fulus.NewMoney[currency.USD](math.MinInt64),
			locale:   locale.EN,
			expected: "-$92,233,720,368,547,758.08",
		},
		{
			name:     "zero",
			money:    fulus.NewMoney[currency.USD](0),
			locale:   locale.EN,
			expected: "$0.00",
		},
		{
			name:     "with cents",
			money:    fulus.NewMoney[currency.USD](1234),
			locale:   locale.EN,
			expected: "$12.34",
		},
		{
			name:     "large number with grouping",
			money:    fulus.NewMoney[currency.USD](1234567),
			locale:   locale.EN,
			expected: "$12,345.67",
		},
		{
			name:     "different locale format (fr)",
			money:    fulus.NewMoney[currency.USD](1234567),
			locale:   locale.FR,
			expected: "12\u202f345,67\u00a0$US",
		},
		{
			name:     "different locale format (de)",
			money:    fulus.NewMoney[currency.USD](1234567),
			locale:   locale.DE,
			expected: "12.345,67\u00a0$",
		},
		{
			name:     "Arabic locale format",
			money:    fulus.NewMoney[currency.USD](1234567),
			locale:   locale.AR,
			expected: "\u200f12,345.67\u00a0US$",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Money(tt.money, tt.locale)
			if result != tt.expected {
				t.Errorf("Format() mismatch\nGot:  %+q (len: %d)\nWant: %+q (len: %d)",
					result, len(result),
					tt.expected, len(tt.expected))
			}
		})
	}
}

func TestGeneratedFormatContracts(t *testing.T) {
	tests := []struct {
		name     string
		money    fulus.Money[currency.USD]
		loc      locale.Locale
		expected Info
	}{
		{
			name:     "en positive contract",
			money:    fulus.NewMoney[currency.USD](1234567),
			loc:      locale.EN,
			expected: InfoFor(currency.USD{}, locale.EN),
		},
		{
			name:     "de negative contract",
			money:    fulus.NewMoney[currency.USD](-1234567),
			loc:      locale.DE,
			expected: InfoFor(currency.USD{}, locale.DE),
		},
		{
			name:     "fr zero contract",
			money:    fulus.NewMoney[currency.USD](0),
			loc:      locale.FR,
			expected: InfoFor(currency.USD{}, locale.FR),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			formatted := Money(tt.money, tt.loc)
			if !strings.Contains(formatted, tt.expected.Symbol) {
				t.Fatalf("formatted value %q does not include symbol %q", formatted, tt.expected.Symbol)
			}

			if amount64(tt.money) != 0 && tt.money.Currency().MinorUnits() > 0 &&
				!strings.Contains(formatted, tt.expected.DecimalSeparator) {
				t.Fatalf("formatted value %q does not include decimal separator %q", formatted, tt.expected.DecimalSeparator)
			}

			if amount64(tt.money) < 0 && !strings.Contains(formatted, tt.expected.MinusSign) {
				t.Fatalf("formatted value %q does not include minus sign %q", formatted, tt.expected.MinusSign)
			}

			if absInt64(amount64(tt.money)) >= 1000 && tt.expected.GroupSeparator != "" &&
				!strings.Contains(formatted, tt.expected.GroupSeparator) {
				t.Fatalf("formatted value %q does not include group separator %q", formatted, tt.expected.GroupSeparator)
			}
		})
	}
}

func TestAnyMoneyFormat(t *testing.T) {
	t.Parallel()

	if got := AnyMoney(fulus.NewAnyMoney(-123456, currency.CHF{}), locale.DE_CH); got != "CHF-1'234.56" {
		t.Errorf("AnyMoney() = %q", got)
	}
	if got := AnyMoney(fulus.AnyMoney{}, locale.EN); got != "0" {
		t.Errorf("zero AnyMoney() = %q, want %q", got, "0")
	}
	m, err := ParseAny("CHF-1'234.56", currency.CHF{}, locale.DE_CH)
	if err != nil || m.Decimal() != "-1234.56" || m.Currency().Code() != "CHF" {
		t.Errorf("ParseAny() = %v, %v", m, err)
	}
}

type kanna struct{}

func (kanna) Code() string    { return "KANNA" }
func (kanna) MinorUnits() int { return 2 }
func (kanna) FormatInfo(loc locale.Locale) Info {
	if loc == locale.JA {
		return Info{Symbol: "🐲", Pattern: "#,##0.00 ¤", GroupSeparator: "⚔︎", DecimalSeparator: "🦖", MinusSign: "⛔"}
	}
	return Info{Symbol: "🐉", Pattern: "¤ #,##0.00", GroupSeparator: ",", DecimalSeparator: ".", MinusSign: "-"}
}

func TestCustomFormatter(t *testing.T) {
	t.Parallel()

	m := fulus.NewMoney[kanna](-2000000)
	tests := []struct {
		loc  locale.Locale
		want string
	}{
		{locale.EN, "-🐉 20,000.00"},
		{locale.JA, "⛔20⚔︎000🦖00 🐲"},
	}
	for _, tt := range tests {
		got := Money(m, tt.loc)
		if got != tt.want {
			t.Errorf("Money(%v) = %q, want %q", tt.loc, got, tt.want)
		}
		back, err := Parse[kanna](got, tt.loc)
		if err != nil || back != m {
			t.Errorf("Parse(%q) = %v, %v", got, back, err)
		}
	}

	// A currency without Formatter uses the CLDR data of its code.
	if got := InfoFor(currency.USD{}, locale.EN).Symbol; got != "$" {
		t.Errorf("InfoFor(USD).Symbol = %q", got)
	}
}
