package fulus

import (
	"errors"
	"math"
	"math/big"
	"testing"

	"github.com/khatibomar/fulus/currency"
)

var (
	bigMaxInt128 = new(big.Int).Sub(new(big.Int).Lsh(big.NewInt(1), 127), big.NewInt(1))
	bigMinInt128 = new(big.Int).Neg(new(big.Int).Lsh(big.NewInt(1), 127))
)

// fitsInt128 reports whether x is in the range of int128.
func fitsInt128(x *big.Int) bool {
	return x.Cmp(bigMinInt128) >= 0 && x.Cmp(bigMaxInt128) <= 0
}

// mkInt128 builds an int128 from its two halves.
func mkInt128(hi int64, lo uint64) int128 {
	return int128{hi: uint64(hi), lo: lo}
}

func checkAddSubMul(t *testing.T, a int128, b int64) {
	t.Helper()

	bb := int128FromInt64(b)
	cases := []struct {
		name string
		got  func() (int128, bool)
		want *big.Int
	}{
		{"add", func() (int128, bool) { return add128(a, bb) }, new(big.Int).Add(a.big(), bb.big())},
		{"sub", func() (int128, bool) { return sub128(a, bb) }, new(big.Int).Sub(a.big(), bb.big())},
		{"mul", func() (int128, bool) { return a.mulInt64(b) }, new(big.Int).Mul(a.big(), bb.big())},
	}
	for _, c := range cases {
		got, gotOK := c.got()
		wantOK := fitsInt128(c.want)
		if gotOK != wantOK || (gotOK && got.big().Cmp(c.want) != 0) {
			t.Errorf("%s(%s, %d) = %s, %v; want %s, %v", c.name, a.big(), b, got.big(), gotOK, c.want, wantOK)
		}
	}
}

func checkMulDivRound(t *testing.T, a int128, n, d int64) {
	t.Helper()
	if d == 0 {
		return
	}

	product := new(big.Int).Mul(a.big(), big.NewInt(n))
	for mode := RoundTruncate; mode <= RoundUnnecessary; mode++ {
		got, gotErr := mulDivRound(a, n, d, mode)
		want, wantErr := roundBig(product, big.NewInt(d), mode)
		if !errors.Is(gotErr, wantErr) || (gotErr == nil && got != want) {
			t.Errorf("mulDivRound(%s, %d, %d, %v) = %s, %v; want %s, %v", a.big(), n, d, mode, got.big(), gotErr, want.big(), wantErr)
		}
	}
}

func TestRoundBig(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		num, den *big.Int
		mode     RoundingMode
		want     *big.Int
		wantErr  error
	}{
		{name: "exact", num: big.NewInt(10), den: big.NewInt(5), mode: RoundUnnecessary, want: big.NewInt(2)},
		{name: "inexact", num: big.NewInt(10), den: big.NewInt(3), mode: RoundUnnecessary, wantErr: ErrInexact},
		{name: "overflow wins over inexact", num: new(big.Int).Lsh(big.NewInt(1), 140), den: big.NewInt(3), mode: RoundUnnecessary, wantErr: ErrOverflow},
		{name: "overflow", num: new(big.Int).Lsh(big.NewInt(1), 127), den: big.NewInt(1), mode: RoundHalfEven, wantErr: ErrOverflow},
		{name: "smallest", num: new(big.Int).Set(bigMinInt128), den: big.NewInt(1), mode: RoundHalfEven, want: bigMinInt128},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := roundBig(tt.num, tt.den, tt.mode)
			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("roundBig() error = %v, want %v", err, tt.wantErr)
			}
			if err == nil && got.big().Cmp(tt.want) != 0 {
				t.Errorf("roundBig() = %s, want %s", got.big(), tt.want)
			}
		})
	}
}

var arithEdges64 = []int64{0, 1, -1, 2, -2, 3, 7, -7, 10, 100, 1 << 31, -(1 << 31), 1 << 32, math.MaxInt64, math.MinInt64, math.MaxInt64 - 1, math.MinInt64 + 1}

func arithEdges128() []int128 {
	edges := []int128{maxInt128, minInt128, mkInt128(math.MaxInt64, math.MaxUint64-1), mkInt128(math.MinInt64, 1), mkInt128(0, math.MaxUint64), mkInt128(1, 0), mkInt128(-1, 0)}
	for _, e := range arithEdges64 {
		edges = append(edges, int128FromInt64(e))
	}
	return edges
}

func TestArithmeticEdges(t *testing.T) {
	t.Parallel()

	for _, a := range arithEdges128() {
		for _, b := range arithEdges64 {
			checkAddSubMul(t, a, b)
			for _, d := range arithEdges64 {
				checkMulDivRound(t, a, b, d)
			}
		}
	}
}

func FuzzArithmetic(f *testing.F) {
	f.Add(int64(0), uint64(123456789), int64(107203), int64(100000))
	f.Add(int64(math.MinInt64), uint64(0), int64(-1), int64(1))
	f.Add(int64(math.MaxInt64), uint64(math.MaxUint64), int64(3), int64(4))
	f.Add(int64(-1), uint64(math.MaxUint64-24), int64(1), int64(10))

	f.Fuzz(func(t *testing.T, hi int64, lo uint64, n, d int64) {
		a := mkInt128(hi, lo)
		checkAddSubMul(t, a, n)
		checkMulDivRound(t, a, n, d)
	})
}

func FuzzSum(f *testing.F) {
	f.Add(int64(math.MaxInt64), uint64(math.MaxUint64), int64(1), int64(-2))
	f.Add(int64(math.MinInt64), uint64(0), int64(-1), int64(1))

	f.Fuzz(func(t *testing.T, hi int64, lo uint64, b, c int64) {
		a := Money[currency.USD]{amount: mkInt128(hi, lo)}
		want := new(big.Int).Add(a.BigInt(), big.NewInt(b))
		want.Add(want, big.NewInt(c))

		got, err := Sum(a, NewMoney[currency.USD](b), NewMoney[currency.USD](c))
		if fitsInt128(want) != (err == nil) || (err == nil && got.BigInt().Cmp(want) != 0) {
			t.Errorf("Sum(%s, %d, %d) = %s, %v; want %s", a.BigInt(), b, c, got.BigInt(), err, want)
		}
	})
}

func FuzzParseMoneyDecimalRoundTrip(f *testing.F) {
	f.Add(int64(0), uint64(0))
	f.Add(int64(-1), uint64(math.MaxUint64-123449))
	f.Add(int64(math.MinInt64), uint64(0))
	f.Add(int64(math.MaxInt64), uint64(math.MaxUint64))

	f.Fuzz(func(t *testing.T, hi int64, lo uint64) {
		m := Money[currency.BHD]{amount: mkInt128(hi, lo)}
		parsed, err := ParseMoney[currency.BHD](m.Decimal())
		if err != nil || parsed != m {
			t.Errorf("ParseMoney(%q) = %s, %v; want %s", m.Decimal(), parsed.BigInt(), err, m.BigInt())
		}
	})
}
