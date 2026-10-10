package fulus

import (
	"math"
	"strconv"
	"strings"
	"testing"

	"github.com/khatibomar/fulus/currency"
	"github.com/khatibomar/fulus/locale"
)

func TestFormatGroupingThroughMoneyFormat(t *testing.T) {
	tests := []struct {
		name   string
		amount int64
		locale locale.Locale
		want   string
	}{
		{
			name:   "zero",
			amount: 0,
			locale: locale.EN,
			want:   "$0.00",
		},
		{
			name:   "single digit",
			amount: 500,
			locale: locale.EN,
			want:   "$5.00",
		},
		{
			name:   "four digits with comma",
			amount: 123400,
			locale: locale.EN,
			want:   "$1,234.00",
		},
		{
			name:   "seven digits with comma",
			amount: 123456700,
			locale: locale.EN,
			want:   "$1,234,567.00",
		},
		{
			name:   "ten digits with comma",
			amount: 123456789000,
			locale: locale.EN,
			want:   "$1,234,567,890.00",
		},
		{
			name:   "seven digits with dot",
			amount: 123456700,
			locale: locale.DE,
			want:   "1.234.567,00\u00a0$",
		},
		{
			name:   "seven digits with narrow no-break space",
			amount: 123456700,
			locale: locale.FR,
			want:   "1\u202f234\u202f567,00\u00a0$US",
		},
		{
			name:   "max int64",
			amount: 9223372036854775800,
			locale: locale.EN,
			want:   "$92,233,720,368,547,758.00",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewMoney[currency.USD](tt.amount).Format(tt.locale)
			if got != tt.want {
				t.Errorf("format for %d in locale %v = %q; want %q",
					tt.amount, tt.locale, got, tt.want)
			}
		})
	}
}

func TestFormatCLDRPatterns(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		format func() string
		want   string
	}{
		{
			name:   "indian grouping",
			format: func() string { return NewMoney[currency.INR](-1234567890).Format(locale.EN_IN) },
			want:   "-₹1,23,45,678.90",
		},
		{
			name:   "indian grouping below one group",
			format: func() string { return NewMoney[currency.INR](100).Format(locale.EN_IN) },
			want:   "₹1.00",
		},
		{
			name:   "two digit grouping",
			format: func() string { return NewMoney[currency.AMD](123456789).Format(locale.TOK) },
			want:   "AMD\u00a01\u00a023\u00a045\u00a067,89",
		},
		{
			name:   "currency spacing after alphabetic symbol",
			format: func() string { return NewMoney[currency.BHD](5).Format(locale.EN) },
			want:   "BHD\u00a00.005",
		},
		{
			name:   "currency spacing before alphabetic symbol",
			format: func() string { return NewMoney[currency.XAF](-5).Format(locale.AGQ) },
			want:   "-5\u00a0FCFA",
		},
		{
			name:   "minimum grouping digits",
			format: func() string { return NewMoney[currency.EUR](100000).Format(locale.ES) },
			want:   "1000,00\u00a0€",
		},
		{
			name:   "minimum grouping digits reached",
			format: func() string { return NewMoney[currency.EUR](1000000).Format(locale.ES) },
			want:   "10.000,00\u00a0€",
		},
		{
			name:   "currency group separator",
			format: func() string { return NewMoney[currency.EUR](123456).Format(locale.DE_AT) },
			want:   "€\u00a01.234,56",
		},
		{
			name:   "currency pattern override",
			format: func() string { return NewMoney[currency.EUR](123456).Format(locale.EN_SK) },
			want:   "€1\u00a0234,56",
		},
		{
			name:   "currency decimal separator override",
			format: func() string { return NewMoney[currency.CVE](1234567).Format(locale.PT_CV) },
			want:   "12\u00a0345$67\u00a0\u200b",
		},
		{
			name:   "positive subpattern",
			format: func() string { return NewMoney[currency.CHF](123456).Format(locale.DE_CH) },
			want:   "CHF\u00a01'234.56",
		},
		{
			name:   "negative subpattern",
			format: func() string { return NewMoney[currency.CHF](-123456).Format(locale.DE_CH) },
			want:   "CHF-1'234.56",
		},
		{
			name:   "negative subpattern with minus after symbol",
			format: func() string { return NewMoney[currency.EUR](-123456).Format(locale.NL) },
			want:   "€\u00a0-1.234,56",
		},
		{
			name:   "negative subpattern with bidi marks",
			format: func() string { return NewMoney[currency.EGP](-123456).Format(locale.AR) },
			want:   "\u200f\u200e-1,234.56\u00a0ج.م.\u200f",
		},
		{
			name:   "no minor units with grouping",
			format: func() string { return NewMoney[currency.JPY](-1234).Format(locale.EN) },
			want:   "-¥1,234",
		},
		{
			name:   "no minor units zero",
			format: func() string { return NewMoney[currency.JPY](0).Format(locale.EN) },
			want:   "¥0",
		},
		{
			name:   "minimum int64",
			format: func() string { return NewMoney[currency.USD](math.MinInt64).Format(locale.EN) },
			want:   "-$92,233,720,368,547,758.08",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := tt.format(); got != tt.want {
				t.Errorf("Format() = %+q, want %+q", got, tt.want)
			}
		})
	}
}

func TestFormatAmount(t *testing.T) {
	t.Parallel()

	base := currency.FormatInfo{
		Symbol:           "$",
		GroupSeparator:   ",",
		DecimalSeparator: ".",
		MinusSign:        "-",
	}
	withFormat := func(format string) currency.FormatInfo {
		info := base
		info.Format = format
		return info
	}

	tests := []struct {
		name       string
		amount     int64
		minorUnits int
		info       currency.FormatInfo
		want       string
	}{
		{name: "three minor units", amount: 5, minorUnits: 3, info: withFormat("¤#,##0.00"), want: "$0.005"},
		{name: "four minor units", amount: -123456789, minorUnits: 4, info: withFormat("¤#,##0.00"), want: "-$12,345.6789"},
		{name: "minus suffix", amount: -1234, minorUnits: 2, info: withFormat("¤ #,##0.00;¤ #,##0.00-"), want: "$ 12.34-"},
		{name: "minus suffix positive", amount: 1234, minorUnits: 2, info: withFormat("¤ #,##0.00;¤ #,##0.00-"), want: "$ 12.34"},
		{name: "no grouping", amount: 123456789, minorUnits: 2, info: withFormat("¤0.00"), want: "$1234567.89"},
		{name: "symbol suffix", amount: -123456, minorUnits: 2, info: withFormat("#,##0.00 ¤"), want: "-1,234.56 $"},
		{name: "empty pattern", amount: 123456, minorUnits: 2, info: withFormat(""), want: "1234.56"},
		{name: "negative minor units", amount: 1234, minorUnits: -1, info: withFormat("¤#,##0"), want: "$1,234"},
		{
			name:       "multi rune minus sign",
			amount:     -1,
			minorUnits: 2,
			info: currency.FormatInfo{
				Symbol: "🐉", Format: "#,##0.00 ¤", GroupSeparator: "⚔︎", DecimalSeparator: "🦖", MinusSign: "⛔",
			},
			want: "⛔0🦖01 🐉",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := formatAmount(tt.amount, tt.minorUnits, tt.info); got != tt.want {
				t.Errorf("formatAmount() = %+q, want %+q", got, tt.want)
			}
		})
	}
}

func TestParsePattern(t *testing.T) {
	t.Parallel()

	tests := []struct {
		pattern string
		want    numberPattern
	}{
		{
			pattern: "¤#,##0.00",
			want:    numberPattern{posPrefix: "¤", negPrefix: "¤", implicitMinus: true, primaryGroup: 3, secondaryGroup: 3},
		},
		{
			pattern: "¤ #,##,##0.00",
			want:    numberPattern{posPrefix: "¤ ", negPrefix: "¤ ", implicitMinus: true, primaryGroup: 3, secondaryGroup: 2},
		},
		{
			pattern: "#,##0.00 ¤;-#,##0.00 ¤",
			want:    numberPattern{posSuffix: " ¤", negPrefix: "-", negSuffix: " ¤", primaryGroup: 3, secondaryGroup: 3},
		},
		{
			pattern: "¤#,#0.00",
			want:    numberPattern{posPrefix: "¤", negPrefix: "¤", implicitMinus: true, primaryGroup: 2, secondaryGroup: 2},
		},
		{
			pattern: "¤0.00",
			want:    numberPattern{posPrefix: "¤", negPrefix: "¤", implicitMinus: true},
		},
	}

	for _, tt := range tests {
		t.Run(tt.pattern, func(t *testing.T) {
			t.Parallel()

			if got := parsePattern(tt.pattern); got != tt.want {
				t.Errorf("parsePattern(%q) = %+v, want %+v", tt.pattern, got, tt.want)
			}
		})
	}
}

func BenchmarkGroupingThroughMoneyFormat(b *testing.B) {
	benchmarks := []struct {
		name   string
		money  Money[currency.USD]
		locale locale.Locale
	}{
		{"small number en", NewMoney[currency.USD](123400), locale.EN},
		{"medium number en", NewMoney[currency.USD](123456700), locale.EN},
		{"large number en", NewMoney[currency.USD](123456789000), locale.EN},
		{"max major en", NewMoney[currency.USD](9223372036854775800), locale.EN},
		{"dot separator de", NewMoney[currency.USD](123456700), locale.DE},
		{"nnbsp separator fr", NewMoney[currency.USD](123456700), locale.FR},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for b.Loop() {
				bm.money.Format(bm.locale)
			}
		})
	}
}

func BenchmarkGroupingBaselineEN(b *testing.B) {
	m := NewMoney[currency.USD](123456789000)

	b.Run("format", func(b *testing.B) {
		for b.Loop() {
			m.Format(locale.EN)
		}
	})
}

// FuzzFormatGrouping checks grouping behavior via the public formatting API.
func FuzzFormatGrouping(f *testing.F) {
	seeds := []int64{0, 1, 12, 123, 1234, 12345, 123456, 1234567}
	for _, seed := range seeds {
		f.Add(seed)
	}

	f.Fuzz(func(t *testing.T, n int64) {
		if n < 0 || n > 92233720368547758 {
			return // Skip negative numbers as they're not supported
		}

		result := NewMoney[currency.USD](n * 100).Format(locale.EN)
		if !strings.HasPrefix(result, "$") {
			t.Fatalf("unexpected EN format shape: %q", result)
		}

		majorPart, ok := strings.CutSuffix(strings.TrimPrefix(result, "$"), ".00")
		if !ok {
			t.Fatalf("unexpected EN format shape: %q", result)
		}

		// Verify the result doesn't contain unexpected separators
		count := strings.Count(majorPart, ",")

		// Verify the number of separators is correct
		expectedSeps := (len(strconv.FormatInt(n, 10)) - 1) / 3
		if count != expectedSeps {
			t.Errorf("Wrong number of separators in %q: got %d, want %d",
				majorPart, count, expectedSeps)
		}

		// Verify the result can be parsed back to the same number
		// after removing separators
		cleaned := strings.ReplaceAll(majorPart, ",", "")
		parsed, err := strconv.ParseInt(cleaned, 10, 64)
		if err != nil {
			t.Errorf("Cannot parse result %q back to number: %v", majorPart, err)
		}
		if parsed != n {
			t.Errorf("Parsing result gives wrong number: got %d, want %d",
				parsed, n)
		}
	})
}
