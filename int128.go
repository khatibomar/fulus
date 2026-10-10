package fulus

import (
	"encoding/binary"
	"math"
	"math/big"
	"math/bits"
	"strconv"
)

// int128 is a signed 128-bit integer in two's complement.
type int128 struct {
	hi, lo uint64
}

// uint128 is an unsigned 128-bit integer.
type uint128 struct {
	hi, lo uint64
}

var (
	maxInt128 = int128{hi: math.MaxInt64, lo: math.MaxUint64}
	minInt128 = int128{hi: 1 << 63}
)

func int128FromInt64(x int64) int128 {
	return int128{hi: uint64(x >> 63), lo: uint64(x)}
}

func (a int128) isNeg() bool  { return int64(a.hi) < 0 }
func (a int128) isZero() bool { return a.hi == 0 && a.lo == 0 }

func (a int128) sign() int {
	switch {
	case a.isNeg():
		return -1
	case a.isZero():
		return 0
	default:
		return 1
	}
}

// int64 returns a as an int64 and reports whether it fits.
func (a int128) int64() (int64, bool) {
	return int64(a.lo), a.hi == uint64(int64(a.lo)>>63)
}

func (a int128) cmp(b int128) int {
	switch {
	case a.hi != b.hi:
		if int64(a.hi) < int64(b.hi) {
			return -1
		}
		return 1
	case a.lo != b.lo:
		if a.lo < b.lo {
			return -1
		}
		return 1
	default:
		return 0
	}
}

// add128 returns a+b and reports whether the sum fits in int128.
func add128(a, b int128) (int128, bool) {
	lo, carry := bits.Add64(a.lo, b.lo, 0)
	hi, _ := bits.Add64(a.hi, b.hi, carry)
	s := int128{hi: hi, lo: lo}
	return s, a.isNeg() != b.isNeg() || s.isNeg() == a.isNeg()
}

// sub128 returns a-b and reports whether the difference fits in int128.
func sub128(a, b int128) (int128, bool) {
	lo, borrow := bits.Sub64(a.lo, b.lo, 0)
	hi, _ := bits.Sub64(a.hi, b.hi, borrow)
	d := int128{hi: hi, lo: lo}
	return d, a.isNeg() == b.isNeg() || d.isNeg() == a.isNeg()
}

// negate returns -a in two's complement. The result for minInt128 is minInt128.
func (a int128) negate() int128 {
	lo, borrow := bits.Sub64(0, a.lo, 0)
	hi, _ := bits.Sub64(0, a.hi, borrow)
	return int128{hi: hi, lo: lo}
}

// neg returns -a and reports whether it fits. Only minInt128 does not fit.
func (a int128) neg() (int128, bool) {
	return a.negate(), a != minInt128
}

// abs returns the magnitude of a. It is correct for minInt128.
func (a int128) abs() uint128 {
	if a.isNeg() {
		return uint128(a.negate())
	}
	return uint128(a)
}

// fromMagnitude128 returns the int128 with magnitude m and the given sign, and reports whether it fits.
func fromMagnitude128(m uint128, negative bool) (int128, bool) {
	limit := uint128{hi: 1 << 63}
	if negative {
		if m.cmp(limit) > 0 {
			return int128{}, false
		}
		return int128(m).negate(), true
	}
	if m.cmp(limit) >= 0 {
		return int128{}, false
	}
	return int128(m), true
}

// mulInt64 returns a*x and reports whether the product fits in int128.
func (a int128) mulInt64(x int64) (int128, bool) {
	m := a.abs()
	h0, lo := bits.Mul64(m.lo, abs64(x))
	h1, l1 := bits.Mul64(m.hi, abs64(x))
	hi, carry := bits.Add64(h0, l1, 0)
	if h1 != 0 || carry != 0 {
		return int128{}, false
	}
	return fromMagnitude128(uint128{hi: hi, lo: lo}, a.isNeg() != (x < 0) && !a.isZero() && x != 0)
}

func (a int128) big() *big.Int {
	m := a.abs()
	var buf [16]byte
	binary.BigEndian.PutUint64(buf[:8], m.hi)
	binary.BigEndian.PutUint64(buf[8:], m.lo)
	b := new(big.Int).SetBytes(buf[:])
	if a.isNeg() {
		b.Neg(b)
	}
	return b
}

// int128FromBig returns b as an int128 and reports whether it fits.
func int128FromBig(b *big.Int) (int128, bool) {
	if b.BitLen() > 128 {
		return int128{}, false
	}
	var buf [16]byte
	b.FillBytes(buf[:])
	m := uint128{hi: binary.BigEndian.Uint64(buf[:8]), lo: binary.BigEndian.Uint64(buf[8:])}
	return fromMagnitude128(m, b.Sign() < 0)
}

// appendInt128 appends the decimal form of a, with a "-" sign if it is negative.
func appendInt128(b []byte, a int128) []byte {
	if a.isNeg() {
		b = append(b, '-')
	}
	return a.abs().appendDecimal(b)
}

func (m uint128) cmp(n uint128) int {
	switch {
	case m.hi != n.hi:
		if m.hi < n.hi {
			return -1
		}
		return 1
	case m.lo != n.lo:
		if m.lo < n.lo {
			return -1
		}
		return 1
	default:
		return 0
	}
}

// inc returns m+1 and reports whether it fits in uint128.
func (m uint128) inc() (uint128, bool) {
	lo, carry := bits.Add64(m.lo, 1, 0)
	hi, carry := bits.Add64(m.hi, 0, carry)
	return uint128{hi: hi, lo: lo}, carry == 0
}

// mulDiv64 returns the quotient and the remainder of m*n/d.
// The product uses 192 bits, so it cannot overflow. It reports false if the quotient does not fit in uint128.
// The divisor d must not be zero.
func (m uint128) mulDiv64(n, d uint64) (q uint128, r uint64, ok bool) {
	h0, p0 := bits.Mul64(m.lo, n)
	h1, l1 := bits.Mul64(m.hi, n)
	p1, carry := bits.Add64(h0, l1, 0)
	// h1 is at most 2^64-2, so the carry cannot overflow.
	p2 := h1 + carry

	q2, r := bits.Div64(0, p2, d)
	if q2 != 0 {
		return uint128{}, 0, false
	}
	q1, r := bits.Div64(r, p1, d)
	q0, r := bits.Div64(r, p0, d)
	return uint128{hi: q1, lo: q0}, r, true
}

// divRem64 returns the quotient and the remainder of m/d. The divisor d must not be zero.
func (m uint128) divRem64(d uint64) (uint128, uint64) {
	qh, r := bits.Div64(0, m.hi, d)
	ql, r := bits.Div64(r, m.lo, d)
	return uint128{hi: qh, lo: ql}, r
}

// mulAdd10 returns m*10+digit and reports whether it fits in uint128.
func (m uint128) mulAdd10(digit uint64) (uint128, bool) {
	hLo, lo := bits.Mul64(m.lo, 10)
	hHi, hi := bits.Mul64(m.hi, 10)
	if hHi != 0 {
		return uint128{}, false
	}
	hi, carry := bits.Add64(hi, hLo, 0)
	if carry != 0 {
		return uint128{}, false
	}
	lo, carry = bits.Add64(lo, digit, 0)
	hi, carry = bits.Add64(hi, 0, carry)
	return uint128{hi: hi, lo: lo}, carry == 0
}

// appendDecimal appends the decimal digits of m.
func (m uint128) appendDecimal(b []byte) []byte {
	if m.hi == 0 {
		return strconv.AppendUint(b, m.lo, 10)
	}
	const chunk = 10_000_000_000_000_000_000 // 10^19, the largest power of 10 that fits in uint64
	q, r := m.divRem64(chunk)
	b = q.appendDecimal(b)
	var digits [19]byte
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i] = byte('0' + r%10)
		r /= 10
	}
	return append(b, digits[:]...)
}
