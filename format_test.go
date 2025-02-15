package fulus

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func TestFormatMajor(t *testing.T) {
	tests := []struct {
		name      string
		input     int64
		separator string
		want      string
	}{
		{
			name:      "zero",
			input:     0,
			separator: ",",
			want:      "0",
		},
		{
			name:      "single digit",
			input:     5,
			separator: ",",
			want:      "5",
		},
		{
			name:      "two digits",
			input:     42,
			separator: ",",
			want:      "42",
		},
		{
			name:      "three digits",
			input:     123,
			separator: ",",
			want:      "123",
		},
		{
			name:      "four digits with comma",
			input:     1234,
			separator: ",",
			want:      "1,234",
		},
		{
			name:      "five digits with comma",
			input:     12345,
			separator: ",",
			want:      "12,345",
		},
		{
			name:      "six digits with comma",
			input:     123456,
			separator: ",",
			want:      "123,456",
		},
		{
			name:      "seven digits with comma",
			input:     1234567,
			separator: ",",
			want:      "1,234,567",
		},
		{
			name:      "eight digits with comma",
			input:     12345678,
			separator: ",",
			want:      "12,345,678",
		},
		{
			name:      "nine digits with comma",
			input:     123456789,
			separator: ",",
			want:      "123,456,789",
		},
		{
			name:      "ten digits with comma",
			input:     1234567890,
			separator: ",",
			want:      "1,234,567,890",
		},
		{
			name:      "seven digits with dot",
			input:     1234567,
			separator: ".",
			want:      "1.234.567",
		},
		{
			name:      "seven digits with arabic separator",
			input:     1234567,
			separator: "،",
			want:      "1،234،567",
		},
		{
			name:      "seven digits with space",
			input:     1234567,
			separator: " ",
			want:      "1 234 567",
		},
		{
			name:      "seven digits with apostrophe",
			input:     1234567,
			separator: "'",
			want:      "1'234'567",
		},
		{
			name:      "max int64",
			input:     9223372036854775807,
			separator: ",",
			want:      "9,223,372,036,854,775,807",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := formatMajor(tt.input, tt.separator)
			if got != tt.want {
				t.Errorf("formatMajor(%d, %q) = %q; want %q",
					tt.input, tt.separator, got, tt.want)
			}
		})
	}
}

func BenchmarkFormatMajor(b *testing.B) {
	benchmarks := []struct {
		name      string
		input     int64
		separator string
	}{
		{"small number", 1234, ","},
		{"medium number", 1234567, ","},
		{"large number", 1234567890, ","},
		{"max int64", 9223372036854775807, ","},
		{"arabic separator", 1234567, "،"},
		{"multi-byte separator", 1234567, "☃"}, // Testing with unicode separator
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				formatMajor(bm.input, bm.separator)
			}
		})
	}
}

func BenchmarkFormatMajorComparison(b *testing.B) {
	input := int64(1234567890)
	separator := ","

	b.Run("iterative", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			formatMajor(input, separator)
		}
	})
}

// TestFormatMajorFuzz adds fuzz testing to catch potential edge cases
func FuzzFormatMajor(f *testing.F) {
	seeds := []int64{0, 1, 12, 123, 1234, 12345, 123456, 1234567}
	for _, seed := range seeds {
		f.Add(seed, ",")
	}

	f.Fuzz(func(t *testing.T, n int64, sep string) {
		if n < 0 {
			return // Skip negative numbers as they're not supported
		}

		result := formatMajor(n, sep)

		// Verify the result doesn't contain unexpected separators
		count := 0
		for i := 0; i < len(result); i++ {
			if i < len(result)-len(sep) && result[i:i+len(sep)] == sep {
				count++
				i += len(sep) - 1
			}
		}

		// Verify the number of separators is correct
		expectedSeps := (len(fmt.Sprint(n)) - 1) / 3
		if count != expectedSeps {
			t.Errorf("Wrong number of separators in %q: got %d, want %d",
				result, count, expectedSeps)
		}

		// Verify the result can be parsed back to the same number
		// after removing separators
		cleaned := result
		if sep != "" {
			cleaned = strings.ReplaceAll(result, sep, "")
		}
		parsed, err := strconv.ParseInt(cleaned, 10, 64)
		if err != nil {
			t.Errorf("Cannot parse result %q back to number: %v", result, err)
		}
		if parsed != n {
			t.Errorf("Parsing result gives wrong number: got %d, want %d",
				parsed, n)
		}
	})
}
