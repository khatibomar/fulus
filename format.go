package fulus

import (
	"strconv"
	"strings"

	"github.com/khatibomar/fulus/currency"
)

// numberPattern is a parsed CLDR currency pattern such as "¤#,##0.00;¤-#,##0.00".
type numberPattern struct {
	posPrefix, posSuffix string
	negPrefix, negSuffix string
	// primaryGroup is the size of the group nearest to the decimal separator. Zero means no grouping.
	primaryGroup int
	// secondaryGroup is the size of all other groups, for example 2 in "#,##,##0".
	secondaryGroup int
}

// parsePattern parses a CLDR currency pattern.
// When the pattern has no negative subpattern, the negative form is "-" followed by the positive form.
func parsePattern(pattern string) numberPattern {
	pos, neg, hasNeg := strings.Cut(pattern, ";")

	var p numberPattern
	var number string
	p.posPrefix, number, p.posSuffix = splitAffixes(pos)
	p.primaryGroup, p.secondaryGroup = groupingSizes(number)

	if hasNeg {
		p.negPrefix, _, p.negSuffix = splitAffixes(neg)
	} else {
		p.negPrefix, p.negSuffix = "-"+p.posPrefix, p.posSuffix
	}
	return p
}

// splitAffixes splits a subpattern into its prefix, number part and suffix.
func splitAffixes(sub string) (prefix, number, suffix string) {
	start := strings.IndexAny(sub, "#0")
	if start < 0 {
		return sub, "", ""
	}
	end := start
	for end < len(sub) && strings.IndexByte("#0,.", sub[end]) >= 0 {
		end++
	}
	return sub[:start], sub[start:end], sub[end:]
}

// groupingSizes returns the primary and secondary group sizes of a number part such as "#,##,##0.00".
func groupingSizes(number string) (primary, secondary int) {
	integer, _, _ := strings.Cut(number, ".")
	rest, last, ok := strings.CutLast(integer, ",")
	if !ok || len(last) == 0 {
		return 0, 0
	}
	primary = len(last)
	secondary = primary
	if _, mid, ok := strings.CutLast(rest, ","); ok && len(mid) > 0 {
		secondary = len(mid)
	}
	return primary, secondary
}

// writeAffix writes a pattern prefix or suffix. It replaces "¤" with the symbol and "-" with the minus sign.
func writeAffix(b *strings.Builder, affix string, info currency.FormatInfo) {
	for _, r := range affix {
		switch r {
		case '¤':
			b.WriteString(info.Symbol)
		case '-':
			b.WriteString(info.MinusSign)
		default:
			b.WriteRune(r)
		}
	}
}

// writeGrouped writes integer digits with a separator between the groups.
func writeGrouped(b *strings.Builder, digits string, primary, secondary int, separator string) {
	if primary <= 0 || len(digits) <= primary {
		b.WriteString(digits)
		return
	}

	head := len(digits) - primary
	first := head % secondary
	if first == 0 {
		first = secondary
	}
	b.WriteString(digits[:first])
	for i := first; i < head; i += secondary {
		b.WriteString(separator)
		b.WriteString(digits[i : i+secondary])
	}
	b.WriteString(separator)
	b.WriteString(digits[head:])
}

// formatAmount formats an amount in minor units with the given format information.
func formatAmount(amount int64, minorUnits int, info currency.FormatInfo) string {
	p := parsePattern(info.Format)
	minorUnits = max(minorUnits, 0)

	negative := amount < 0
	magnitude := uint64(amount)
	if negative {
		// Unsigned negation also gives the correct magnitude for math.MinInt64.
		magnitude = -magnitude
	}

	digits := strconv.FormatUint(magnitude, 10)
	if len(digits) <= minorUnits {
		digits = strings.Repeat("0", minorUnits-len(digits)+1) + digits
	}
	integer, fraction := digits[:len(digits)-minorUnits], digits[len(digits)-minorUnits:]

	prefix, suffix := p.posPrefix, p.posSuffix
	if negative {
		prefix, suffix = p.negPrefix, p.negSuffix
	}

	var b strings.Builder
	writeAffix(&b, prefix, info)
	writeGrouped(&b, integer, p.primaryGroup, p.secondaryGroup, info.GroupSeparator)
	if minorUnits > 0 {
		b.WriteString(info.DecimalSeparator)
		b.WriteString(fraction)
	}
	writeAffix(&b, suffix, info)
	return b.String()
}
