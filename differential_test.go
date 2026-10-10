package fulus

import (
	"errors"
	"math"
	"math/big"
	"testing"

	"github.com/khatibomar/fulus/currency"
)

// refRound rounds num/den with mode. It is a reference that does not share code with the package:
// it compares the exact fraction with the two integers next to it.
func refRound(num, den *big.Int, mode RoundingMode) (*big.Int, error) {
	if mode < RoundTruncate || mode > RoundUnnecessary {
		return nil, ErrInvalidRoundingMode
	}
	if den.Sign() < 0 {
		num, den = new(big.Int).Neg(num), new(big.Int).Neg(den)
	}
	x := new(big.Rat).SetFrac(num, den)
	// For a positive divisor, the Euclidean quotient is the floor.
	floor := new(big.Int).Div(num, den)
	if new(big.Rat).SetInt(floor).Cmp(x) == 0 {
		return floor, nil
	}
	ceil := new(big.Int).Add(floor, big.NewInt(1))

	// diff is x - floor, in (0, 1).
	diff := new(big.Rat).Sub(x, new(big.Rat).SetInt(floor))
	half := big.NewRat(1, 2)
	positive := x.Sign() > 0
	towardZero, awayFromZero := floor, ceil
	if !positive {
		towardZero, awayFromZero = ceil, floor
	}
	nearest := func(tie *big.Int) *big.Int {
		switch diff.Cmp(half) {
		case -1:
			return floor
		case 1:
			return ceil
		default:
			return tie
		}
	}

	switch mode {
	case RoundTruncate:
		return towardZero, nil
	case RoundUp:
		return awayFromZero, nil
	case RoundCeiling:
		return ceil, nil
	case RoundFloor:
		return floor, nil
	case RoundHalfUp:
		return nearest(awayFromZero), nil
	case RoundHalfDown:
		return nearest(towardZero), nil
	case RoundHalfEven:
		if floor.Bit(0) == 0 {
			return nearest(floor), nil
		}
		return nearest(ceil), nil
	case RoundUnnecessary:
		return nil, ErrInexact
	default:
		return nil, ErrInvalidRoundingMode
	}
}

// refResult turns a reference result into the expected value and error of the package.
// An overflow of the truncated result wins over ErrInexact, as in the package.
func refResult(num, den *big.Int, mode RoundingMode) (*big.Int, error) {
	q, err := refRound(num, den, mode)
	if errors.Is(err, ErrInexact) {
		if !fitsInt128(new(big.Int).Quo(num, den)) {
			return nil, ErrOverflow
		}
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	if !fitsInt128(q) {
		return nil, ErrOverflow
	}
	return q, nil
}

func checkResult[T currency.Unit](t *testing.T, op string, got Money[T], gotErr error, want *big.Int, wantErr error) {
	t.Helper()
	if !errors.Is(gotErr, wantErr) || (gotErr == nil && got.BigInt().Cmp(want) != 0) {
		t.Errorf("%s = %s, %v; want %v, %v", op, got.BigInt(), gotErr, want, wantErr)
	}
}

// checkMoneyOps compares the public operations on Money with the math/big reference, for each rounding mode.
func checkMoneyOps(t *testing.T, a int128, n, d int64) {
	t.Helper()
	m := Money[currency.USD]{amount: a}
	ab, nb, db := a.big(), big.NewInt(n), big.NewInt(d)

	exact := func(x *big.Int) (*big.Int, error) {
		if !fitsInt128(x) {
			return nil, ErrOverflow
		}
		return x, nil
	}
	{
		got, err := m.Add(NewMoney[currency.USD](n))
		want, wantErr := exact(new(big.Int).Add(ab, nb))
		checkResult(t, "Add", got, err, want, wantErr)
	}
	{
		got, err := m.Sub(NewMoney[currency.USD](n))
		want, wantErr := exact(new(big.Int).Sub(ab, nb))
		checkResult(t, "Sub", got, err, want, wantErr)
	}
	{
		got, err := m.Mul(n)
		want, wantErr := exact(new(big.Int).Mul(ab, nb))
		checkResult(t, "Mul", got, err, want, wantErr)
	}
	{
		got, err := m.Neg()
		want, wantErr := exact(new(big.Int).Neg(ab))
		checkResult(t, "Neg", got, err, want, wantErr)
	}
	{
		got, err := m.Abs()
		want, wantErr := exact(new(big.Int).Abs(ab))
		checkResult(t, "Abs", got, err, want, wantErr)
	}

	for mode := RoundingMode(0); mode <= RoundUnnecessary+1; mode++ {
		if d != 0 {
			got, err := m.Div(d, mode)
			want, wantErr := refResult(ab, db, mode)
			checkResult(t, "Div("+mode.String()+")", got, err, want, wantErr)

			if f, ferr := NewFactor(n, d); ferr == nil {
				got, err := m.MulFactor(f, mode)
				want, wantErr := refResult(new(big.Int).Mul(ab, nb), db, mode)
				checkResult(t, "MulFactor("+mode.String()+")", got, err, want, wantErr)
			}
		}

		// EUR has 2 minor units and JPY has 0, so the exact result is amount*n/(d*100).
		if r, rerr := NewRate[currency.EUR, currency.JPY](n, d); rerr == nil {
			got, err := Convert(Money[currency.EUR]{amount: a}, r, mode)
			want, wantErr := refResult(new(big.Int).Mul(ab, nb), new(big.Int).Mul(db, big.NewInt(100)), mode)
			checkResult(t, "Convert("+mode.String()+")", got, err, want, wantErr)
		}

		// CHF cash uses steps of 5 minor units.
		{
			got, err := Money[currency.CHF]{amount: a}.RoundCash(mode)
			want, wantErr := refResult(ab, big.NewInt(5), mode)
			if wantErr == nil {
				want, wantErr = exact(want.Mul(want, big.NewInt(5)))
			}
			checkResult(t, "RoundCash("+mode.String()+")", got, err, want, wantErr)
		}
	}
}

// checkAllocate compares Allocate with a math/big reference of the largest remainder method.
func checkAllocate(t *testing.T, a int128, ratios []int64) {
	t.Helper()
	parts, err := Money[currency.USD]{amount: a}.Allocate(ratios...)
	if err != nil {
		t.Fatalf("Allocate(%s, %v) error = %v", a.big(), ratios, err)
	}

	total := new(big.Int)
	for _, r := range ratios {
		total.Add(total, big.NewInt(r))
	}
	ab := a.big()
	want := make([]*big.Int, len(ratios))
	rems := make([]*big.Int, len(ratios))
	left := new(big.Int).Set(ab)
	for i, r := range ratios {
		num := new(big.Int).Mul(ab, big.NewInt(r))
		want[i], rems[i] = new(big.Int).QuoRem(num, total, new(big.Int))
		rems[i].Abs(rems[i])
		left.Sub(left, want[i])
	}
	step := big.NewInt(int64(left.Sign()))
	for left.Sign() != 0 {
		best := -1
		for i := range rems {
			if rems[i] != nil && (best < 0 || rems[i].Cmp(rems[best]) > 0) {
				best = i
			}
		}
		want[best].Add(want[best], step)
		rems[best] = nil
		left.Sub(left, step)
	}

	for i := range parts {
		if parts[i].BigInt().Cmp(want[i]) != 0 {
			t.Errorf("Allocate(%s, %v)[%d] = %s, want %s", ab, ratios, i, parts[i].BigInt(), want[i])
		}
	}
}

func TestDifferentialEdges(t *testing.T) {
	t.Parallel()

	for _, a := range arithEdges128() {
		for _, n := range arithEdges64 {
			for _, d := range arithEdges64 {
				checkMoneyOps(t, a, n, d)
			}
		}
		checkAllocate(t, a, []int64{1, 1, 1})
		checkAllocate(t, a, []int64{1, 2, math.MaxInt64 / 4})
		checkAllocate(t, a, []int64{0, 3, 0, 7})
	}
}

func FuzzDifferential(f *testing.F) {
	f.Add(int64(0), uint64(1050), int64(107203), int64(100000), uint32(1), uint32(2), uint32(3))
	f.Add(int64(math.MinInt64), uint64(0), int64(-1), int64(3), uint32(7), uint32(7), uint32(1))
	f.Add(int64(math.MaxInt64), uint64(math.MaxUint64), int64(3), int64(-4), uint32(1), uint32(1), uint32(1))

	f.Fuzz(func(t *testing.T, hi int64, lo uint64, n, d int64, r1, r2, r3 uint32) {
		a := mkInt128(hi, lo)
		checkMoneyOps(t, a, n, d)
		checkAllocate(t, a, []int64{int64(r1), int64(r2) + 1, int64(r3)})
	})
}
