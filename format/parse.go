package format

import (
	"fmt"
	"strings"

	"github.com/khatibomar/fulus"
	"github.com/khatibomar/fulus/currency"
	"github.com/khatibomar/fulus/locale"
)

// Parse parses an amount in the form that Money writes for loc, for example "-₹1,23,45,678.90".
// It ignores spaces and bidirectional marks, and it accepts "-" in place of the minus sign of the locale.
// Group separators are optional. If they are present, the group sizes must agree with the pattern.
// A space group separator is ignored and not checked.
// The fraction can have fewer digits than the minor units, but not more.
// Returns fulus.ErrInvalidAmountFormat if s does not agree with the pattern,
// and fulus.ErrScaleMismatch if the fraction has more digits than the minor units.
func Parse[T currency.Unit](s string, loc locale.Locale) (fulus.Money[T], error) {
	var c T
	decimal, err := parseFormatted(s, InfoFor(c, loc))
	if err != nil {
		return fulus.Money[T]{}, err
	}
	return fulus.ParseMoney[T](decimal)
}

// ParseAny parses an amount of the currency c in the form that AnyMoney writes for loc, like Parse.
func ParseAny(s string, c currency.Currency, loc locale.Locale) (fulus.AnyMoney, error) {
	decimal, err := parseFormatted(s, InfoFor(c, loc))
	if err != nil {
		return fulus.AnyMoney{}, err
	}
	return fulus.NewAnyMoneyFromDecimal(decimal, c)
}

// parseFormatted returns the canonical decimal of a formatted amount. It does not check the scale.
func parseFormatted(s string, info Info) (string, error) {
	p := parsePattern(info.Pattern)
	// The affixes use "-" as the minus sign, and the input has the minus sign of the locale replaced with "-".
	dash := info
	dash.MinusSign = "-"
	affix := func(a string, prefix bool) string {
		var b strings.Builder
		writeAffix(&b, a, dash, prefix)
		return normalizeFormatted(b.String())
	}
	posPrefix, posSuffix := affix(p.posPrefix, true), affix(p.posSuffix, false)
	negPrefix, negSuffix := affix(p.negPrefix, true), affix(p.negSuffix, false)
	if p.implicitMinus {
		negPrefix = "-" + negPrefix
	}

	input := normalizeFormatted(s)
	if minus := normalizeFormatted(info.MinusSign); minus != "" && minus != "-" {
		input = strings.ReplaceAll(input, minus, "-")
	}
	negative := false
	body, ok := cutAffixes(input, negPrefix, negSuffix)
	if ok && (negPrefix != posPrefix || negSuffix != posSuffix) {
		negative = true
	} else if body, ok = cutAffixes(input, posPrefix, posSuffix); !ok {
		return "", fmt.Errorf("%w: %q does not match pattern %q", fulus.ErrInvalidAmountFormat, s, info.Pattern)
	}

	integer, fraction, hasFraction := strings.Cut(body, normalizeFormatted(info.DecimalSeparator))
	if group := normalizeFormatted(info.GroupSeparator); group != "" && strings.Contains(integer, group) {
		groups := strings.Split(integer, group)
		if !validGroups(groups, p.primaryGroup, p.secondaryGroup) {
			return "", fmt.Errorf("%w: %q has wrong digit grouping", fulus.ErrInvalidAmountFormat, s)
		}
		integer = strings.Join(groups, "")
	}

	canonical := integer
	if hasFraction {
		canonical += "." + fraction
	}
	if negative {
		canonical = "-" + canonical
	}
	return canonical, nil
}

// validGroups reports whether the integer digit groups agree with the group sizes of the pattern.
func validGroups(groups []string, primary, secondary int) bool {
	if primary <= 0 {
		return false
	}
	last := len(groups) - 1
	for i, g := range groups {
		switch {
		case i == last:
			if len(g) != primary {
				return false
			}
		case i == 0:
			if len(g) == 0 || len(g) > secondary {
				return false
			}
		case len(g) != secondary:
			return false
		}
	}
	return true
}

func cutAffixes(s, prefix, suffix string) (string, bool) {
	body, ok := strings.CutPrefix(s, prefix)
	if !ok {
		return "", false
	}
	body, ok = strings.CutSuffix(body, suffix)
	if !ok || body == "" {
		return "", false
	}
	return body, true
}

// normalizeFormatted removes spaces and bidirectional marks, and it replaces the minus sign U+2212 with "-".
func normalizeFormatted(s string) string {
	return strings.Map(func(r rune) rune {
		switch r {
		case '‎', '‏', '؜':
			return -1
		case '−':
			return '-'
		}
		if r == ' ' || r == ' ' || r == ' ' || r == ' ' {
			return -1
		}
		return r
	}, s)
}
