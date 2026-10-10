package fulus

import (
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/khatibomar/fulus/currency"
)

// numberPattern is a parsed CLDR currency pattern such as "¤#,##0.00;¤-#,##0.00".
type numberPattern struct {
	posPrefix, posSuffix string
	negPrefix, negSuffix string
	// implicitMinus is true if the pattern has no negative subpattern.
	// Then the negative form is "-" followed by the positive form.
	implicitMinus bool
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
		p.negPrefix, p.negSuffix, p.implicitMinus = p.posPrefix, p.posSuffix, true
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

// writeAffix writes a pattern affix with the symbol, the minus sign and the CLDR currency spacing.
func writeAffix(b *strings.Builder, affix string, info currency.FormatInfo, prefix bool) {
	for i, r := range affix {
		switch r {
		case '¤':
			nextToNumber := prefix && i+len("¤") == len(affix) || !prefix && i == 0
			if nextToNumber && !prefix && needsCurrencySpace(info.Symbol, utf8.DecodeRuneInString) {
				b.WriteString(currencySpace)
			}
			b.WriteString(info.Symbol)
			if nextToNumber && prefix && needsCurrencySpace(info.Symbol, utf8.DecodeLastRuneInString) {
				b.WriteString(currencySpace)
			}
		case '-':
			b.WriteString(info.MinusSign)
		default:
			b.WriteRune(r)
		}
	}
}

// currencySpace is the CLDR currency spacing text. It is the same in all CLDR locales.
const currencySpace = "\u00a0"

// needsCurrencySpace reports whether the symbol character next to the number is not a symbol or a separator.
func needsCurrencySpace(symbol string, decode func(string) (rune, int)) bool {
	r, size := decode(symbol)
	return size > 0 && r != utf8.RuneError && !unicode.IsSymbol(r) && !unicode.In(r, unicode.Z)
}

// writeGrouped writes integer digits with group separators, if the first group has at least minimum digits.
func writeGrouped(b *strings.Builder, digits []byte, primary, secondary, minimum int, separator string) {
	if primary <= 0 || len(digits) < primary+max(minimum, 1) {
		b.Write(digits)
		return
	}

	head := len(digits) - primary
	first := head % secondary
	if first == 0 {
		first = secondary
	}
	b.Write(digits[:first])
	for i := first; i < head; i += secondary {
		b.WriteString(separator)
		b.Write(digits[i : i+secondary])
	}
	b.WriteString(separator)
	b.Write(digits[head:])
}

// formatAmount formats an amount in minor units with the given format information.
func formatAmount(amount int128, minorUnits int, info currency.FormatInfo) string {
	p := parsePattern(info.Format)
	minorUnits = max(minorUnits, 0)
	negative := amount.isNeg()

	var buf [40]byte
	digits := amount.abs().appendDecimal(buf[:0])
	integer, fraction, fractionPad := []byte{'0'}, digits, minorUnits-len(digits)
	if len(digits) > minorUnits {
		integer, fraction, fractionPad = digits[:len(digits)-minorUnits], digits[len(digits)-minorUnits:], 0
	}

	prefix, suffix := p.posPrefix, p.posSuffix
	if negative {
		prefix, suffix = p.negPrefix, p.negSuffix
	}

	var b strings.Builder
	b.Grow(len(info.Format) + 2*len(info.Symbol) + len(currencySpace) + len(info.MinusSign) + len(info.DecimalSeparator) +
		len(digits)*(1+len(info.GroupSeparator)) + minorUnits)
	if negative && p.implicitMinus {
		b.WriteString(info.MinusSign)
	}
	writeAffix(&b, prefix, info, true)
	writeGrouped(&b, integer, p.primaryGroup, p.secondaryGroup, info.MinimumGroupingDigits, info.GroupSeparator)
	if minorUnits > 0 {
		b.WriteString(info.DecimalSeparator)
		for range fractionPad {
			b.WriteByte('0')
		}
		b.Write(fraction)
	}
	writeAffix(&b, suffix, info, false)
	return b.String()
}
