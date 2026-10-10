package fulus

import (
	"errors"
	"math"
	"math/big"
	"testing"

	"github.com/khatibomar/fulus/currency"
)

// bigResult returns x as int64 and reports whether it fits.
func bigResult(x *big.Int) (int64, bool) {
	if !x.IsInt64() {
		return 0, false
	}
	return x.Int64(), true
}

func checkAddSubMul(t *testing.T, a, b int64) {
	t.Helper()

	cases := []struct {
		name string
		got  func() (int64, bool)
		want func() (int64, bool)
	}{
		{"add", func() (int64, bool) { return add64(a, b) }, func() (int64, bool) { return bigResult(new(big.Int).Add(big.NewInt(a), big.NewInt(b))) }},
		{"sub", func() (int64, bool) { return sub64(a, b) }, func() (int64, bool) { return bigResult(new(big.Int).Sub(big.NewInt(a), big.NewInt(b))) }},
		{"mul", func() (int64, bool) { return mul64(a, b) }, func() (int64, bool) { return bigResult(new(big.Int).Mul(big.NewInt(a), big.NewInt(b))) }},
	}
	for _, c := range cases {
		got, gotOK := c.got()
		want, wantOK := c.want()
		if gotOK != wantOK || (gotOK && got != want) {
			t.Errorf("%s(%d, %d) = %d, %v; want %d, %v", c.name, a, b, got, gotOK, want, wantOK)
		}
	}
}

func checkMulDivRound(t *testing.T, a, n, d int64) {
	t.Helper()
	if d == 0 {
		return
	}

	product := new(big.Int).Mul(big.NewInt(a), big.NewInt(n))
	for mode := RoundTruncate; mode <= RoundUnnecessary; mode++ {
		got, gotErr := mulDivRound(a, n, d, mode)
		ref, refErr := divideWithRounding(product, big.NewInt(d), mode)
		var want int64
		wantErr := refErr
		if trunc, _ := divideWithRounding(product, big.NewInt(d), RoundTruncate); !trunc.IsInt64() {
			wantErr = ErrOverflow
		} else if refErr == nil {
			var ok bool
			if want, ok = bigResult(ref); !ok {
				wantErr = ErrOverflow
			}
		}
		if !errors.Is(gotErr, wantErr) || (gotErr == nil && got != want) {
			t.Errorf("mulDivRound(%d, %d, %d, %v) = %d, %v; want %d, %v", a, n, d, mode, got, gotErr, want, wantErr)
		}
	}
}

var arithEdges = []int64{0, 1, -1, 2, -2, 3, 7, -7, 10, 100, 1 << 31, -(1 << 31), 1 << 32, math.MaxInt64, math.MinInt64, math.MaxInt64 - 1, math.MinInt64 + 1}

func TestArithmeticEdges(t *testing.T) {
	t.Parallel()

	for _, a := range arithEdges {
		for _, b := range arithEdges {
			checkAddSubMul(t, a, b)
			for _, d := range arithEdges {
				checkMulDivRound(t, a, b, d)
			}
		}
	}
}

func FuzzArithmetic(f *testing.F) {
	f.Add(int64(123456789), int64(107203), int64(100000))
	f.Add(int64(math.MinInt64), int64(-1), int64(1))
	f.Add(int64(math.MaxInt64), int64(3), int64(4))
	f.Add(int64(-25), int64(1), int64(10))

	f.Fuzz(func(t *testing.T, a, n, d int64) {
		checkAddSubMul(t, a, n)
		checkMulDivRound(t, a, n, d)
	})
}

func FuzzSum(f *testing.F) {
	f.Add(int64(math.MaxInt64), int64(1), int64(-2))
	f.Add(int64(math.MinInt64), int64(-1), int64(1))

	f.Fuzz(func(t *testing.T, a, b, c int64) {
		want := new(big.Int).Add(big.NewInt(a), big.NewInt(b))
		want.Add(want, big.NewInt(c))

		got, err := Sum(NewMoney[currency.USD](a), NewMoney[currency.USD](b), NewMoney[currency.USD](c))
		if want.IsInt64() != (err == nil) || (err == nil && got.Amount() != want.Int64()) {
			t.Errorf("Sum(%d, %d, %d) = %d, %v; want %s", a, b, c, got.Amount(), err, want)
		}
	})
}

func FuzzParseMoneyDecimalRoundTrip(f *testing.F) {
	f.Add(int64(0))
	f.Add(int64(-123450))
	f.Add(int64(math.MinInt64))
	f.Add(int64(math.MaxInt64))

	f.Fuzz(func(t *testing.T, amount int64) {
		m := NewMoney[currency.BHD](amount)
		parsed, err := ParseMoney[currency.BHD](m.Decimal())
		if err != nil || parsed != m {
			t.Errorf("ParseMoney(%q) = %d, %v; want %d", m.Decimal(), parsed.Amount(), err, amount)
		}
	})
}
