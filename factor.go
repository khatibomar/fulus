package fulus

import (
	"fmt"
	"math"
	"math/big"
	"strconv"
	"strings"
)

// Factor is an exact multiplier, such as a tax rate, a fee or a discount, held as a fraction in lowest terms.
// Parse a Factor one time, for example at start-up, and use it with Money.MulFactor.
// The zero value is not a valid factor. MulFactor returns ErrInvalidFactor for it.
type Factor struct {
	num, den int64
}

// NewFactor returns the factor numerator/denominator.
// Returns ErrDivisionByZero if denominator is zero.
func NewFactor(numerator, denominator int64) (Factor, error) {
	if denominator == 0 {
		return Factor{}, ErrDivisionByZero
	}
	if denominator < 0 {
		if numerator == math.MinInt64 || denominator == math.MinInt64 {
			return Factor{}, fmt.Errorf("%w: factor %d/%d", ErrOverflow, numerator, denominator)
		}
		numerator, denominator = -numerator, -denominator
	}
	g := max(gcd(abs64(numerator), uint64(denominator)), 1)
	return Factor{num: numerator / int64(g), den: denominator / int64(g)}, nil
}

// Percent returns the factor p/100. For example, Percent(15) is 0.15.
func Percent(p int64) Factor {
	f, _ := NewFactor(p, 100)
	return f
}

// Bps returns the factor of n basis points, n/10000. For example, Bps(25) is 0.0025.
func Bps(n int64) Factor {
	f, _ := NewFactor(n, 10000)
	return f
}

// ParseFactor parses a decimal such as "0.0825", a fraction such as "1/3", or a percentage such as "8.25%".
// It does not round. Returns ErrInvalidFactor if s cannot be parsed,
// and ErrOverflow if a term of the fraction in lowest terms does not fit in int64.
func ParseFactor(s string) (Factor, error) {
	text, percent := strings.CutSuffix(s, "%")
	if numerator, denominator, ok := parseShortDecimal(text); ok {
		if percent {
			denominator *= 100
		}
		return NewFactor(numerator, denominator)
	}

	// An exponent such as "1e999999999" can make big.Rat allocate a very large number, so it is not accepted.
	r, ok := new(big.Rat).SetString(text)
	if !ok || strings.ContainsAny(text, "eE") {
		return Factor{}, fmt.Errorf("%w: %q", ErrInvalidFactor, s)
	}
	if percent {
		r.Quo(r, big.NewRat(100, 1))
	}
	if !r.Num().IsInt64() || !r.Denom().IsInt64() {
		return Factor{}, fmt.Errorf("%w: factor %q", ErrOverflow, s)
	}
	return Factor{num: r.Num().Int64(), den: r.Denom().Int64()}, nil
}

// MustParseFactor is like ParseFactor but panics if s is not valid.
// Use it only for constant factors, for example in a package-level variable.
func MustParseFactor(s string) Factor {
	f, err := ParseFactor(s)
	if err != nil {
		panic(err)
	}
	return f
}

// Fraction returns the factor as numerator/denominator in lowest terms, with a positive denominator.
// Both terms are zero for the zero Factor.
func (f Factor) Fraction() (numerator, denominator int64) {
	return f.num, f.den
}

// String returns the factor as a fraction such as "33/400", or as an integer if the denominator is 1.
func (f Factor) String() string {
	if f.den == 1 {
		return strconv.FormatInt(f.num, 10)
	}
	return strconv.FormatInt(f.num, 10) + "/" + strconv.FormatInt(f.den, 10)
}

// parseShortDecimal parses a short decimal such as "-0.0825" into numerator/denominator without allocation.
// It reports false for other forms, which ParseFactor parses with math/big.
func parseShortDecimal(s string) (numerator, denominator int64, ok bool) {
	negative := false
	if s != "" && (s[0] == '-' || s[0] == '+') {
		negative = s[0] == '-'
		s = s[1:]
	}
	// 16 digits and the factor 100 for a percentage fit in int64.
	if s == "" || len(s) > 16 {
		return 0, 0, false
	}

	denominator = 1
	seenDigit, seenPoint := false, false
	for i := range len(s) {
		switch c := s[i]; {
		case c >= '0' && c <= '9':
			numerator = numerator*10 + int64(c-'0')
			if seenPoint {
				denominator *= 10
			}
			seenDigit = true
		case c == '.' && !seenPoint:
			seenPoint = true
		default:
			return 0, 0, false
		}
	}
	if !seenDigit {
		return 0, 0, false
	}
	if negative {
		numerator = -numerator
	}
	return numerator, denominator, true
}
