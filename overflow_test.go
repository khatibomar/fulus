package fulus

import (
	"errors"
	"math/big"
	"testing"

	"github.com/khatibomar/fulus/currency"
)

func TestOverflowAt128Bits(t *testing.T) {
	t.Parallel()

	maxUSD, minUSD := maxMoney[currency.USD](), minMoney[currency.USD]()
	one := NewMoney[currency.USD](1)
	tests := []struct {
		name string
		op   func() error
	}{
		{"add", func() error { _, err := maxUSD.Add(one); return err }},
		{"sub", func() error { _, err := minUSD.Sub(one); return err }},
		{"mul", func() error { _, err := maxUSD.Mul(2); return err }},
		{"mul smallest by -1", func() error { _, err := minUSD.Mul(-1); return err }},
		{"div smallest by -1", func() error { _, err := minUSD.Div(-1, RoundHalfUp); return err }},
		{"abs", func() error { _, err := minUSD.Abs(); return err }},
		{"neg", func() error { _, err := minUSD.Neg(); return err }},
		{"mul factor", func() error { _, err := maxUSD.MulFactor(MustParseFactor("1.5"), RoundTruncate); return err }},
		{"convert", func() error {
			_, err := Convert(maxMoney[currency.EUR](), MustParseRate[currency.EUR, currency.USD]("2"), RoundTruncate)
			return err
		}},
		{"convert with scale", func() error {
			_, err := Convert(maxMoney[currency.JPY](), MustParseRate[currency.JPY, currency.USD]("1"), RoundTruncate)
			return err
		}},
		{"sum", func() error { _, err := Sum(maxUSD, one); return err }},
		{"parse", func() error {
			_, err := ParseMoney[currency.USD]("1701411834604692317316873037158841057.28")
			return err
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if err := tt.op(); !errors.Is(err, ErrOverflow) {
				t.Errorf("error = %v, want ErrOverflow", err)
			}
		})
	}
}

func TestRange(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		got  string
		want string
	}{
		{"largest USD", maxMoney[currency.USD]().Decimal(), "1701411834604692317316873037158841057.27"},
		{"smallest USD", minMoney[currency.USD]().Decimal(), "-1701411834604692317316873037158841057.28"},
		{"beyond int64", NewMoney[currency.USD](1 << 62).MustMul(4).Decimal(), "184467440737095516.16"},
	}
	for _, tt := range tests {
		if tt.got != tt.want {
			t.Errorf("%s = %s, want %s", tt.name, tt.got, tt.want)
		}
	}

	for _, s := range []string{"1701411834604692317316873037158841057.27", "-1701411834604692317316873037158841057.28"} {
		m, err := ParseMoney[currency.USD](s)
		if err != nil || m.Decimal() != s {
			t.Errorf("ParseMoney(%q) = %s, %v", s, m.Decimal(), err)
		}
	}

	large := NewMoney[currency.USD](1 << 62).MustMul(4)
	if _, ok := large.Int64(); ok {
		t.Error("Int64() of 2^64 reports that it fits")
	}
	if _, err := (BigintMoney[currency.USD]{Money: large}).Value(); !errors.Is(err, ErrOverflow) {
		t.Errorf("BigintMoney Value() of 2^64 error = %v", err)
	}
	if v, err := large.Value(); err != nil || v != "184467440737095516.16" {
		t.Errorf("Value() of 2^64 = %v, %v", v, err)
	}
	back, err := NewMoneyFromBigInt[currency.USD](large.BigInt())
	if err != nil || back != large {
		t.Errorf("NewMoneyFromBigInt() = %v, %v", back, err)
	}
	if _, err := NewMoneyFromBigInt[currency.USD](new(big.Int).Lsh(big.NewInt(1), 127)); !errors.Is(err, ErrOverflow) {
		t.Errorf("NewMoneyFromBigInt(2^127) error = %v", err)
	}
}
