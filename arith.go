package fulus

import (
	"errors"
	"math/big"
	"math/bits"
	"strings"
)

// mul64 returns a*b and reports whether the product fits in int64.
func mul64(a, b int64) (int64, bool) {
	hi, lo := bits.Mul64(abs64(a), abs64(b))
	if hi != 0 {
		return 0, false
	}
	return fromMagnitude(lo, (a < 0) != (b < 0))
}

// mulDivRound returns a*n/d rounded with mode. The product a*n uses 192 bits, so it cannot overflow.
// The caller must check that d is not zero and that mode is valid.
// It returns ErrOverflow if the result does not fit in int128,
// and ErrInexact if mode is RoundUnnecessary and the result is not exact.
func mulDivRound(a int128, n, d int64, mode RoundingMode) (int128, error) {
	q, r, ok := a.abs().mulDiv64(abs64(n), abs64(d))
	if !ok {
		return int128{}, ErrOverflow
	}
	negative := !a.isZero() && n != 0 && (a.isNeg() != (n < 0)) != (d < 0)
	return roundQuotient(q, r, abs64(d), negative, mode)
}

// roundQuotient rounds the magnitude q with the remainder r of divisor d, and applies the sign.
func roundQuotient(q uint128, r, d uint64, negative bool, mode RoundingMode) (int128, error) {
	if r != 0 {
		if mode == RoundUnnecessary {
			if _, ok := fromMagnitude128(q, negative); !ok {
				return int128{}, ErrOverflow
			}
			return int128{}, ErrInexact
		}
		if roundAwayFromZero(mode, q.lo&1 == 1, r, d, negative) {
			var ok bool
			if q, ok = q.inc(); !ok {
				return int128{}, ErrOverflow
			}
		}
	}
	result, ok := fromMagnitude128(q, negative)
	if !ok {
		return int128{}, ErrOverflow
	}
	return result, nil
}

// roundBig returns num/den rounded with mode, for terms that do not fit in the fast path.
func roundBig(num, den *big.Int, mode RoundingMode) (int128, error) {
	q, err := divideWithRounding(num, den, mode)
	if errors.Is(err, ErrInexact) {
		if _, ok := int128FromBig(new(big.Int).Quo(num, den)); !ok {
			return int128{}, ErrOverflow
		}
	}
	if err != nil {
		return int128{}, err
	}
	result, ok := int128FromBig(q)
	if !ok {
		return int128{}, ErrOverflow
	}
	return result, nil
}

// roundAwayFromZero reports whether a truncated quotient with a non-zero remainder r of divisor d
// must move one unit away from zero. odd tells whether the truncated quotient is odd.
func roundAwayFromZero(mode RoundingMode, odd bool, r, d uint64, negative bool) bool {
	// r < d, so d-r cannot overflow. r > d-r means the remainder is more than half of d.
	switch mode {
	case RoundHalfUp:
		return r >= d-r
	case RoundHalfDown:
		return r > d-r
	case RoundHalfEven:
		return r > d-r || (r == d-r && odd)
	case RoundUp:
		return true
	case RoundCeiling:
		return !negative
	case RoundFloor:
		return negative
	default:
		return false
	}
}

// abs64 returns the magnitude of x. It is correct for math.MinInt64.
func abs64(x int64) uint64 {
	if x < 0 {
		return -uint64(x)
	}
	return uint64(x)
}

// fromMagnitude returns the int64 with magnitude m and the given sign, and reports whether it fits.
func fromMagnitude(m uint64, negative bool) (int64, bool) {
	if negative {
		if m > 1<<63 {
			return 0, false
		}
		return -int64(m), true
	}
	if m > 1<<63-1 {
		return 0, false
	}
	return int64(m), true
}

// gcd returns the greatest common divisor of a and b.
func gcd(a, b uint64) uint64 {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// divideWithRounding returns numerator/denominator rounded with mode, with math/big.
func divideWithRounding(numerator, denominator *big.Int, mode RoundingMode) (*big.Int, error) {
	if !mode.valid() {
		return nil, ErrInvalidRoundingMode
	}

	q := new(big.Int)
	r := new(big.Int)
	q.QuoRem(numerator, denominator, r)

	if r.Sign() == 0 || mode == RoundTruncate {
		return q, nil
	}
	if mode == RoundUnnecessary {
		return nil, ErrInexact
	}

	absRem := new(big.Int).Abs(r)
	absDen := new(big.Int).Abs(denominator)
	twiceRem := new(big.Int).Lsh(absRem, 1)
	cmp := twiceRem.Cmp(absDen)
	negativeResult := (numerator.Sign() < 0) != (denominator.Sign() < 0)

	var awayFromZero bool
	switch mode {
	case RoundHalfUp:
		awayFromZero = cmp >= 0
	case RoundHalfDown:
		awayFromZero = cmp > 0
	case RoundHalfEven:
		awayFromZero = cmp > 0 || (cmp == 0 && q.Bit(0) == 1)
	case RoundUp:
		awayFromZero = true
	case RoundCeiling:
		awayFromZero = !negativeResult
	case RoundFloor:
		awayFromZero = negativeResult
	}

	if awayFromZero {
		if negativeResult {
			q.Sub(q, big.NewInt(1))
		} else {
			q.Add(q, big.NewInt(1))
		}
	}
	return q, nil
}

// parseRat parses a base 10 decimal such as "-1.25" or a fraction such as "1/3".
// Unlike big.Rat.SetString, it does not accept exponents, base prefixes such as "0x" or "010/1", or underscores.
func parseRat(s string) (*big.Rat, bool) {
	sign := ""
	if s != "" && (s[0] == '-' || s[0] == '+') {
		sign, s = s[:1], s[1:]
	}
	if num, den, isFraction := strings.Cut(s, "/"); isFraction {
		if num == "" || den == "" || !allDigits(num) || !allDigits(den) {
			return nil, false
		}
		n, _ := new(big.Int).SetString(sign+num, 10)
		d, _ := new(big.Int).SetString(den, 10)
		if d.Sign() == 0 {
			return nil, false
		}
		return new(big.Rat).SetFrac(n, d), true
	}
	whole, fraction, _ := strings.Cut(s, ".")
	if whole+fraction == "" || !allDigits(whole) || !allDigits(fraction) {
		return nil, false
	}
	return new(big.Rat).SetString(sign + s)
}
