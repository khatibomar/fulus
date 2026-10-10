package fulus

import (
	"math"
	"math/big"
	"testing"
)

func TestInt128Conversions(t *testing.T) {
	t.Parallel()

	for _, a := range arithEdges128() {
		b := a.big()
		back, ok := int128FromBig(b)
		if !ok || back != a {
			t.Errorf("int128FromBig(%s) = %v, %v; want %v", b, back, ok, a)
		}
		if got := string(appendInt128(nil, a)); got != b.String() {
			t.Errorf("appendInt128(%v) = %s, want %s", a, got, b)
		}
		v, ok := a.int64()
		if ok != b.IsInt64() || (ok && v != b.Int64()) {
			t.Errorf("int64(%s) = %d, %v", b, v, ok)
		}
		for _, c := range arithEdges128() {
			if got, want := a.cmp(c), b.Cmp(c.big()); got != want {
				t.Errorf("cmp(%s, %s) = %d, want %d", b, c.big(), got, want)
			}
		}
	}

	tooBig := new(big.Int).Lsh(big.NewInt(1), 127)
	if _, ok := int128FromBig(tooBig); ok {
		t.Errorf("int128FromBig(2^127) fits")
	}
	if _, ok := int128FromBig(new(big.Int).Neg(new(big.Int).Add(tooBig, big.NewInt(1)))); ok {
		t.Errorf("int128FromBig(-2^127-1) fits")
	}
}

func TestInt128Neg(t *testing.T) {
	t.Parallel()

	if _, ok := minInt128.neg(); ok {
		t.Error("neg(minInt128) fits")
	}
	if got, ok := maxInt128.neg(); !ok || got != mkInt128(math.MinInt64, 1) {
		t.Errorf("neg(maxInt128) = %v, %v", got, ok)
	}
	if got := minInt128.abs(); got != (uint128{hi: 1 << 63}) {
		t.Errorf("abs(minInt128) = %v", got)
	}
}

func TestUint128MulAdd10(t *testing.T) {
	t.Parallel()

	// 2^128-1 = 340282366920938463463374607431768211455
	digits := "340282366920938463463374607431768211455"
	var m uint128
	for i := range len(digits) {
		var ok bool
		if m, ok = m.mulAdd10(uint64(digits[i] - '0')); !ok {
			t.Fatalf("mulAdd10 overflow at digit %d", i)
		}
	}
	if m != (uint128{hi: math.MaxUint64, lo: math.MaxUint64}) {
		t.Errorf("parsed %v", m)
	}
	if _, ok := m.mulAdd10(0); ok {
		t.Error("mulAdd10 past 2^128 fits")
	}
	if got := string(m.appendDecimal(nil)); got != digits {
		t.Errorf("appendDecimal() = %s", got)
	}
}
