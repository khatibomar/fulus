package fulus_test

import (
	"fmt"

	"github.com/khatibomar/fulus"
	"github.com/khatibomar/fulus/currency"
	"github.com/khatibomar/fulus/locale"
)

func ExampleMoney_MulDecimal() {
	price := fulus.NewMoney[currency.USD](1000)
	tax, err := price.MulDecimal("0.0825", fulus.RoundHalfUp)
	if err != nil {
		panic(err)
	}
	fmt.Println(tax)
	// Output: $0.83
}

func ExampleMoney_RoundCash() {
	chf := fulus.NewMoney[currency.CHF](1003)
	cash, err := chf.RoundCash(fulus.RoundHalfUp)
	if err != nil {
		panic(err)
	}
	fmt.Println(cash.Decimal())
	// Output: 10.05
}

func ExampleAs() {
	price, err := fulus.ParseAnyMoney("12.50", "eur")
	if err != nil {
		panic(err)
	}

	eur, err := fulus.As[currency.EUR](price)
	fmt.Println(eur, err)

	_, err = fulus.As[currency.USD](price)
	fmt.Println(err)
	// Output:
	// €12.50 <nil>
	// currency mismatch: expected USD, got EUR
}

func ExampleParseFormatted() {
	loc, _ := locale.Match("en_IN.UTF-8")
	inr := fulus.NewMoney[currency.INR](-1234567890)
	formatted := inr.Format(loc)

	parsed, err := fulus.ParseFormatted[currency.INR](formatted, loc)
	fmt.Println(formatted, parsed.Amount(), err)
	// Output: -₹1,23,45,678.90 -1234567890 <nil>
}

func ExampleSum() {
	total, err := fulus.Sum(
		fulus.NewMoney[currency.USD](1000),
		fulus.NewMoney[currency.USD](250),
		fulus.NewMoney[currency.USD](-50),
	)
	fmt.Println(total, err)
	// Output: $12.00 <nil>
}
