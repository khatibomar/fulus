package fulus

import (
	"encoding/json"
	"testing"

	"github.com/khatibomar/fulus/currency"
	"github.com/khatibomar/fulus/locale"
)

var (
	benchMoney    Money[currency.USD]
	benchString   string
	benchErr      error
	benchBytes    []byte
	benchAllocate []Money[currency.USD]
)

func BenchmarkArithmetic(b *testing.B) {
	m := NewMoney[currency.USD](123456789)
	other := NewMoney[currency.USD](987654)
	values := make([]Money[currency.USD], 10)
	for i := range values {
		values[i] = NewMoney[currency.USD](int64(i) * 1000)
	}
	ratio := MustParseRate[currency.USD, currency.USD]("1.07203")

	b.Run("Add", func(b *testing.B) {
		for b.Loop() {
			benchMoney, benchErr = m.Add(other)
		}
	})
	b.Run("Sub", func(b *testing.B) {
		for b.Loop() {
			benchMoney, benchErr = m.Sub(other)
		}
	})
	b.Run("Mul", func(b *testing.B) {
		for b.Loop() {
			benchMoney, benchErr = m.Mul(7)
		}
	})
	b.Run("Div", func(b *testing.B) {
		for b.Loop() {
			benchMoney, benchErr = m.Div(7, RoundHalfEven)
		}
	})
	tax := MustParseFactor("0.0825")
	b.Run("MulFactor", func(b *testing.B) {
		for b.Loop() {
			benchMoney, benchErr = m.MulFactor(tax, RoundHalfUp)
		}
	})
	b.Run("ParseFactor", func(b *testing.B) {
		for b.Loop() {
			tax, benchErr = ParseFactor("0.0825")
		}
	})
	b.Run("Convert", func(b *testing.B) {
		for b.Loop() {
			benchMoney, benchErr = Convert(m, ratio, RoundHalfEven)
		}
	})
	b.Run("RoundCash", func(b *testing.B) {
		chf := NewMoney[currency.CHF](123457)
		for b.Loop() {
			_, benchErr = chf.RoundCash(RoundHalfUp)
		}
	})
	b.Run("Allocate", func(b *testing.B) {
		for b.Loop() {
			benchAllocate, benchErr = m.Allocate(1, 2, 3)
		}
	})
	b.Run("Sum10", func(b *testing.B) {
		for b.Loop() {
			benchMoney, benchErr = Sum(values...)
		}
	})
}

func BenchmarkFormatting(b *testing.B) {
	m := NewMoney[currency.USD](-123456789)
	inr := NewMoney[currency.INR](-123456789)

	b.Run("Format/en", func(b *testing.B) {
		for b.Loop() {
			benchString = m.Format(locale.EN)
		}
	})
	b.Run("Format/en-IN", func(b *testing.B) {
		for b.Loop() {
			benchString = inr.Format(locale.EN_IN)
		}
	})
	b.Run("String", func(b *testing.B) {
		for b.Loop() {
			benchString = m.String()
		}
	})
	b.Run("Decimal", func(b *testing.B) {
		for b.Loop() {
			benchString = m.Decimal()
		}
	})
	b.Run("MarshalJSON", func(b *testing.B) {
		for b.Loop() {
			benchBytes, benchErr = json.Marshal(m)
		}
	})
}

func BenchmarkParsing(b *testing.B) {
	formatted := NewMoney[currency.INR](-123456789).Format(locale.EN_IN)
	payload := []byte(`{"amount":"-123456789","currency":"USD"}`)

	b.Run("ParseMoney", func(b *testing.B) {
		for b.Loop() {
			benchMoney, benchErr = ParseMoney[currency.USD]("-1234567.89")
		}
	})
	b.Run("ParseFormatted", func(b *testing.B) {
		for b.Loop() {
			_, benchErr = ParseFormatted[currency.INR](formatted, locale.EN_IN)
		}
	})
	b.Run("UnmarshalJSON", func(b *testing.B) {
		for b.Loop() {
			benchErr = json.Unmarshal(payload, &benchMoney)
		}
	})
	b.Run("Scan", func(b *testing.B) {
		for b.Loop() {
			benchErr = benchMoney.Scan([]byte("-123456789"))
		}
	})
}

func BenchmarkAnyMoneyAdd(b *testing.B) {
	x := NewMoney[currency.USD](1050).Any()
	y := NewMoney[currency.USD](250).Any()
	b.ReportAllocs()
	for b.Loop() {
		if _, err := x.Add(y); err != nil {
			b.Fatal(err)
		}
	}
}
