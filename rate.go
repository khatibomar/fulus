package fulus

import (
	"fmt"
	"math"
	"math/big"
	"strconv"
	"strings"

	"github.com/khatibomar/fulus/currency"
)

// Rate is an exchange rate: the price of one major unit of Base in major units of Quote.
// For example, a Rate[currency.EUR, currency.USD] of 1.0720 means that EUR 1.00 is USD 1.0720.
// A Rate is an exact positive fraction in lowest terms.
// The zero value is not a valid rate. Convert returns ErrInvalidExchangeRate for it.
type Rate[Base, Quote currency.Unit] struct {
	num, den int64
}

// NewRate returns the rate numerator/denominator.
// Returns ErrInvalidExchangeRate if a term is not positive.
func NewRate[Base, Quote currency.Unit](numerator, denominator int64) (Rate[Base, Quote], error) {
	if numerator <= 0 || denominator <= 0 {
		return Rate[Base, Quote]{}, fmt.Errorf("%w: %d/%d must be positive", ErrInvalidExchangeRate, numerator, denominator)
	}
	g := gcd(uint64(numerator), uint64(denominator))
	return Rate[Base, Quote]{num: numerator / int64(g), den: denominator / int64(g)}, nil
}

// ParseRate parses a decimal rate such as "1.07203" or a fraction such as "1/3". It does not round.
// Returns ErrInvalidExchangeRate if the rate cannot be parsed or is not positive,
// and ErrOverflow if a term of the fraction in lowest terms does not fit in int64.
func ParseRate[Base, Quote currency.Unit](s string) (Rate[Base, Quote], error) {
	// An exponent such as "1e999999999" can make big.Rat allocate a very large number, so it is not accepted.
	r, ok := new(big.Rat).SetString(s)
	if !ok || strings.ContainsAny(s, "eE") {
		return Rate[Base, Quote]{}, fmt.Errorf("%w: %q", ErrInvalidExchangeRate, s)
	}
	if r.Sign() <= 0 {
		return Rate[Base, Quote]{}, fmt.Errorf("%w: %q must be positive", ErrInvalidExchangeRate, s)
	}
	if !r.Num().IsInt64() || !r.Denom().IsInt64() {
		return Rate[Base, Quote]{}, fmt.Errorf("%w: rate %q", ErrOverflow, s)
	}
	return Rate[Base, Quote]{num: r.Num().Int64(), den: r.Denom().Int64()}, nil
}

// MustParseRate is like ParseRate but panics if the rate is not valid.
// Use it only for constant rates, for example in tests.
func MustParseRate[Base, Quote currency.Unit](s string) Rate[Base, Quote] {
	r, err := ParseRate[Base, Quote](s)
	if err != nil {
		panic(err)
	}
	return r
}

// RateFromFloat64 returns the rate with the shortest decimal form that gives f, for example 1/10 for 0.1.
// Prefer ParseRate with the original text. A float64 has only about 15 significant decimal digits.
// Returns ErrInvalidExchangeRate if f is not a positive finite number.
func RateFromFloat64[Base, Quote currency.Unit](f float64) (Rate[Base, Quote], error) {
	if !(f > 0) || math.IsInf(f, 1) {
		return Rate[Base, Quote]{}, fmt.Errorf("%w: %v must be positive and finite", ErrInvalidExchangeRate, f)
	}
	return ParseRate[Base, Quote](strconv.FormatFloat(f, 'f', -1, 64))
}

// Fraction returns the rate as numerator/denominator in lowest terms.
// Both terms are zero for the zero Rate.
func (r Rate[Base, Quote]) Fraction() (numerator, denominator int64) {
	return r.num, r.den
}

// IsValid reports whether r is a valid rate. Only the zero Rate is not valid.
func (r Rate[Base, Quote]) IsValid() bool {
	return r.den > 0
}

// Invert returns the rate in the other direction: the price of one major unit of Quote in Base.
// The inverse is exact. The inverse of the zero Rate is the zero Rate.
func (r Rate[Base, Quote]) Invert() Rate[Quote, Base] {
	return Rate[Quote, Base]{num: r.den, den: r.num}
}

// Cross returns the exact rate from A to C through B. The compiler checks that the middle currencies agree.
// Returns ErrInvalidExchangeRate if a rate is the zero Rate,
// and ErrOverflow if a term of the result does not fit in int64.
func Cross[A, B, C currency.Unit](ab Rate[A, B], bc Rate[B, C]) (Rate[A, C], error) {
	if !ab.IsValid() || !bc.IsValid() {
		return Rate[A, C]{}, ErrInvalidExchangeRate
	}
	g1 := int64(gcd(uint64(ab.num), uint64(bc.den)))
	g2 := int64(gcd(uint64(bc.num), uint64(ab.den)))
	num, okNum := mul64(ab.num/g1, bc.num/g2)
	den, okDen := mul64(ab.den/g2, bc.den/g1)
	if !okNum || !okDen {
		return Rate[A, C]{}, fmt.Errorf("%w: cross rate", ErrOverflow)
	}
	return Rate[A, C]{num: num, den: den}, nil
}

// String returns the rate as a decimal such as "1.07203" if it has at most 18 decimal places,
// and as a fraction such as "1/3" if not.
func (r Rate[Base, Quote]) String() string {
	if !r.IsValid() {
		return "0"
	}
	d := r.den
	places := 0
	for d%10 == 0 {
		d /= 10
		places++
	}
	for _, p := range []int64{2, 5} {
		for d%p == 0 && places <= 18 {
			d /= p
			places++
		}
	}
	if d != 1 || places > 18 {
		return strconv.FormatInt(r.num, 10) + "/" + strconv.FormatInt(r.den, 10)
	}
	return new(big.Rat).SetFrac64(r.num, r.den).FloatString(places)
}

// MarshalText implements encoding.TextMarshaler with the form from String.
func (r Rate[Base, Quote]) MarshalText() ([]byte, error) {
	if !r.IsValid() {
		return nil, ErrInvalidExchangeRate
	}
	return []byte(r.String()), nil
}

// UnmarshalText implements encoding.TextUnmarshaler with ParseRate.
func (r *Rate[Base, Quote]) UnmarshalText(text []byte) error {
	parsed, err := ParseRate[Base, Quote](string(text))
	if err != nil {
		return err
	}
	*r = parsed
	return nil
}

// Convert changes m to the currency Quote at rate r and rounds the result with mode.
// Convert adjusts for the minor units of each currency.
// For example, a EUR/JPY rate of 160.25 changes EUR 1.00 to JPY 160.
// Returns ErrInvalidExchangeRate for the zero Rate and ErrOverflow if the result does not fit.
func Convert[Base, Quote currency.Unit](m Money[Base], r Rate[Base, Quote], mode RoundingMode) (Money[Quote], error) {
	if !r.IsValid() {
		return Money[Quote]{}, ErrInvalidExchangeRate
	}
	if !mode.valid() {
		return Money[Quote]{}, ErrInvalidRoundingMode
	}

	var from Base
	var to Quote
	shift := to.MinorUnits() - from.MinorUnits()

	amount, err := mulDivRoundShift(m.amount, r.num, r.den, shift, mode)
	if err != nil {
		return Money[Quote]{}, err
	}
	return Money[Quote]{amount: amount}, nil
}

// mulDivRoundShift returns a*n*10^shift/d rounded with mode.
func mulDivRoundShift(a int128, n, d int64, shift int, mode RoundingMode) (int128, error) {
	scaledN, scaledD, ok := n, d, true
	switch {
	case shift > 0:
		scaledN, ok = mulPow10(n, shift)
	case shift < 0:
		scaledD, ok = mulPow10(d, -shift)
	}
	if ok {
		return mulDivRound(a, scaledN, scaledD, mode)
	}

	num := new(big.Int).Mul(a.big(), big.NewInt(n))
	den := big.NewInt(d)
	pow := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(max(shift, -shift))), nil)
	if shift > 0 {
		num.Mul(num, pow)
	} else {
		den.Mul(den, pow)
	}
	return roundBig(num, den, mode)
}

// mulPow10 returns x*10^n and reports whether the product fits in int64.
func mulPow10(x int64, n int) (int64, bool) {
	for range n {
		var ok bool
		if x, ok = mul64(x, 10); !ok {
			return 0, false
		}
	}
	return x, true
}
