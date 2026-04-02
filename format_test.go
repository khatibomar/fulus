package fulus

import (
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
			want:   "$0",
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
			for i := 0; i < b.N; i++ {
				bm.money.Format(bm.locale)
			}
		})
	}
}

func BenchmarkGroupingBaselineEN(b *testing.B) {
	m := NewMoney[currency.USD](123456789000)

	b.Run("format", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
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

		majorPart := strings.TrimPrefix(result, "$")
		if n == 0 {
			if majorPart != "0" {
				t.Fatalf("unexpected zero format: %q", result)
			}
		} else {
			if !strings.HasSuffix(majorPart, ".00") {
				t.Fatalf("unexpected EN format shape: %q", result)
			}
			majorPart = strings.TrimSuffix(majorPart, ".00")
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
