package fulus

import "math/bits"

// add64 returns a+b and reports whether the sum fits in int64.
func add64(a, b int64) (int64, bool) {
	s := a + b
	return s, (a >= 0) != (b >= 0) || (s >= 0) == (a >= 0)
}

// sub64 returns a-b and reports whether the difference fits in int64.
func sub64(a, b int64) (int64, bool) {
	d := a - b
	return d, (a >= 0) == (b >= 0) || (d >= 0) == (a >= 0)
}

// mul64 returns a*b and reports whether the product fits in int64.
func mul64(a, b int64) (int64, bool) {
	hi, lo := bits.Mul64(abs64(a), abs64(b))
	if hi != 0 {
		return 0, false
	}
	return fromMagnitude(lo, (a < 0) != (b < 0))
}

// mulDivRound returns a*n/d rounded with mode. The product a*n uses 128 bits, so it cannot overflow.
// The caller must check that d is not zero and that mode is valid.
// It reports false if the result does not fit in int64.
func mulDivRound(a, n, d int64, mode RoundingMode) (int64, bool) {
	q, r, negative, ok := mulDiv(a, n, d)
	if !ok {
		return 0, false
	}
	if r != 0 && roundAwayFromZero(mode, q, r, abs64(d), negative) {
		q++
		if q == 0 {
			return 0, false
		}
	}
	return fromMagnitude(q, negative)
}

// mulDiv returns the magnitudes of the truncated quotient and the remainder of a*n/d,
// and whether the exact result is negative. It reports false if the quotient does not fit in uint64.
func mulDiv(a, n, d int64) (q, r uint64, negative, ok bool) {
	hi, lo := bits.Mul64(abs64(a), abs64(n))
	dm := abs64(d)
	if hi >= dm {
		return 0, 0, false, false
	}
	q, r = bits.Div64(hi, lo, dm)
	negative = a != 0 && n != 0 && ((a < 0) != (n < 0)) != (d < 0)
	return q, r, negative, true
}

// roundAwayFromZero reports whether a truncated quotient q with a non-zero remainder r of divisor d
// must move one unit away from zero.
func roundAwayFromZero(mode RoundingMode, q, r, d uint64, negative bool) bool {
	// r < d, so d-r cannot overflow. r > d-r means the remainder is more than half of d.
	switch mode {
	case RoundHalfUp:
		return r >= d-r
	case RoundHalfDown:
		return r > d-r
	case RoundHalfEven:
		return r > d-r || (r == d-r && q&1 == 1)
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
