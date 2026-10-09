package fulus

import (
	"fmt"
	"strings"

	"github.com/khatibomar/fulus/currency"
	"github.com/khatibomar/fulus/locale"
)

// ParseFormatted parses a value in the form that Format writes for the locale, for example "-₹1,23,45,678.90".
// It ignores spaces and bidirectional marks, and it accepts "-" in place of the minus sign of the locale.
// Group separators are optional. If they are present, the group sizes must agree with the pattern.
// A space group separator is ignored and not checked.
// The fraction can have fewer digits than the minor units, but not more.
func ParseFormatted[T currency.Currency](s string, loc locale.Locale) (Money[T], error) {
	var c T
	minor, err := parseFormatted(s, c.MinorUnits(), c.FormatInfo(loc))
	if err != nil {
		return Money[T]{}, err
	}
	return Money[T]{amount: minor}, nil
}

func parseFormatted(s string, minorUnits int, info currency.FormatInfo) (int64, error) {
	p := parsePattern(info.Format)
	affix := func(a string) string {
		var b strings.Builder
		writeAffix(&b, a, info)
		return normalizeFormatted(b.String())
	}
	posPrefix, posSuffix := affix(p.posPrefix), affix(p.posSuffix)
	negPrefix, negSuffix := affix(p.negPrefix), affix(p.negSuffix)
	if p.implicitMinus {
		negPrefix = "-" + negPrefix
	}

	input := normalizeFormatted(s)
	negative := false
	body, ok := cutAffixes(input, negPrefix, negSuffix)
	if ok && (negPrefix != posPrefix || negSuffix != posSuffix) {
		negative = true
	} else if body, ok = cutAffixes(input, posPrefix, posSuffix); !ok {
		return 0, fmt.Errorf("%w: %q does not match pattern %q", ErrInvalidAmountFormat, s, info.Format)
	}

	integer, fraction, hasFraction := strings.Cut(body, normalizeFormatted(info.DecimalSeparator))
	if group := normalizeFormatted(info.GroupSeparator); group != "" && strings.Contains(integer, group) {
		groups := strings.Split(integer, group)
		if !validGroups(groups, p.primaryGroup, p.secondaryGroup) {
			return 0, fmt.Errorf("%w: %q has wrong digit grouping", ErrInvalidAmountFormat, s)
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
	return parseDecimal(canonical, minorUnits)
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
