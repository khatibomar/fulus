package fulus

import (
	"strings"
)

func formatMajorString(str, separator string) string {
	if str == "0" {
		return "0"
	}

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
