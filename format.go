package fulus

import (
	"fmt"
	"strings"
)

// formatMajor formats the major part of a number with group separators
// For example: 1234567 -> "1,234,567" or "1.234.567" depending on separator
func formatMajor(n int64, separator string) string {
	if n == 0 {
		return "0"
	}

	str := fmt.Sprintf("%d", n)

	numSeparators := (len(str) - 1) / 3
	if numSeparators == 0 {
		return str
	}

	var buf strings.Builder
	buf.Grow(len(str) + len(separator)*numSeparators)

	firstGroupLen := len(str) % 3
	if firstGroupLen == 0 {
		firstGroupLen = 3
	}
	buf.WriteString(str[:firstGroupLen])
	for i := firstGroupLen; i < len(str); i += 3 {
		buf.WriteString(separator)
		buf.WriteString(str[i : i+3])
	}

	return buf.String()
}
