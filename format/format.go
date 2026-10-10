// Package format formats and parses money amounts with the CLDR data of a locale.
//
// The fulus and currency packages do not import this package or the locale tables,
// so a program that does not format amounts does not include the CLDR data.
package format

import (
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/khatibomar/fulus"
	"github.com/khatibomar/fulus/currency"
	"github.com/khatibomar/fulus/locale"
)

// Info holds the data that formats an amount of one currency in one locale.
type Info struct {
	// Symbol is the currency symbol, for example "$" or "US$".
	Symbol string
	// Pattern is the CLDR currency pattern, for example "¤#,##0.00" or "#,##0.00 ¤;-#,##0.00 ¤".
	Pattern string
	// GroupSeparator separates digit groups, for example ",".
	GroupSeparator string
	// DecimalSeparator separates the fraction, for example ".".
	DecimalSeparator string
	// MinusSign replaces "-" in the pattern.
	MinusSign string
	// MinimumGroupingDigits is the smallest number of digits before the first group separator. Zero means 1.
	MinimumGroupingDigits int
}

// Formatter is a currency with its own format data. A custom currency can implement it.
// For other currencies, the format data comes from CLDR by the currency code.
type Formatter interface {
	FormatInfo(loc locale.Locale) Info
}

// InfoFor returns the format data of c in loc.
// If c implements Formatter, InfoFor returns its data. If not, InfoFor returns the CLDR data for the currency code.
func InfoFor(c currency.Currency, loc locale.Locale) Info {
	if f, ok := c.(Formatter); ok {
		return f.FormatInfo(loc)
	}
	code := c.Code()
	n := loc.CurrencyNumbers(code)
	return Info{
		Symbol:                loc.CurrencySymbol(code),
		Pattern:               n.CurrencyFormat,
		GroupSeparator:        n.GroupSeparator,
		DecimalSeparator:      n.DecimalSeparator,
		MinusSign:             n.MinusSign,
		MinimumGroupingDigits: n.MinimumGroupingDigits,
	}
}

// Money returns m formatted for loc, for example "-$1,234.50" in en or "-1.234,50 $" in de.
// It applies the CLDR pattern of the locale, including the negative subpattern and the group sizes.
// The number of fraction digits is always the minor units of the currency.
// Do not store the result. CLDR updates can change it.
func Money[T currency.Unit](m fulus.Money[T], loc locale.Locale) string {
	c := m.Currency()
	return formatDecimal(m.Decimal(), InfoFor(c, loc))
}

// AnyMoney returns m formatted for loc, like Money. The zero AnyMoney gives the amount only.
func AnyMoney(m fulus.AnyMoney, loc locale.Locale) string {
	c := m.Currency()
	if c == nil {
		return m.String()
	}
	return formatDecimal(m.Decimal(), InfoFor(c, loc))
}

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
func writeAffix(b *strings.Builder, affix string, info Info, prefix bool) {
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
func writeGrouped(b *strings.Builder, digits string, primary, secondary, minimum int, separator string) {
	if primary <= 0 || len(digits) < primary+max(minimum, 1) {
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

// formatDecimal formats a canonical decimal such as "-1234.50" with the given format data.
// The decimal has all the minor units of the currency.
func formatDecimal(decimal string, info Info) string {
	p := parsePattern(info.Pattern)
	digits, negative := strings.CutPrefix(decimal, "-")
	integer, fraction, _ := strings.Cut(digits, ".")

	prefix, suffix := p.posPrefix, p.posSuffix
	if negative {
		prefix, suffix = p.negPrefix, p.negSuffix
	}

	var b strings.Builder
	b.Grow(len(info.Pattern) + 2*len(info.Symbol) + len(currencySpace) + len(info.MinusSign) + len(info.DecimalSeparator) +
		len(digits)*(1+len(info.GroupSeparator)))
	if negative && p.implicitMinus {
		b.WriteString(info.MinusSign)
	}
	writeAffix(&b, prefix, info, true)
	writeGrouped(&b, integer, p.primaryGroup, p.secondaryGroup, info.MinimumGroupingDigits, info.GroupSeparator)
	if fraction != "" {
		b.WriteString(info.DecimalSeparator)
		b.WriteString(fraction)
	}
	writeAffix(&b, suffix, info, false)
	return b.String()
}
