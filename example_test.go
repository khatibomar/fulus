package fulus_test

import (
	"encoding/json"
	"fmt"

	"github.com/khatibomar/fulus"
	"github.com/khatibomar/fulus/currency"
	"github.com/khatibomar/fulus/locale"
)

func ExampleMoney_MulFactor() {
	price := fulus.NewMoney[currency.USD](1000)
	tax, err := price.MulFactor(fulus.MustParseFactor("8.25%"), fulus.RoundHalfUp)
	if err != nil {
		panic(err)
	}
	fmt.Println(tax)
	// Output: USD 0.83
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
	// EUR 12.50 <nil>
	// currency mismatch: expected USD, got EUR
}

func ExampleParseFormatted() {
	loc, _ := locale.Match("en_IN.UTF-8")
	inr := fulus.NewMoney[currency.INR](-1234567890)
	formatted := inr.Format(loc)

	parsed, err := fulus.ParseFormatted[currency.INR](formatted, loc)
	fmt.Println(formatted, parsed.Decimal(), err)
	// Output: -₹1,23,45,678.90 -12345678.90 <nil>
}

func ExampleSum() {
	total, err := fulus.Sum(
		fulus.NewMoney[currency.USD](1000),
		fulus.NewMoney[currency.USD](250),
		fulus.NewMoney[currency.USD](-50),
	)
	fmt.Println(total, err)
	// Output: USD 12.00 <nil>
}

func ExampleMoney_Format() {
	price := fulus.NewMoney[currency.EUR](-123456789)
	fmt.Println(price.Format(locale.EN))
	fmt.Println(price.Format(locale.DE_CH))

	// Some locales use a no-break space U+00A0 or a narrow no-break space U+202F.
	fmt.Printf("%+q\n", price.Format(locale.FR))
	// Output:
	// -€1,234,567.89
	// EUR-1'234'567.89
	// "-1\u202f234\u202f567,89\u00a0\u20ac"
}

func ExampleParseMoney() {
	price, err := fulus.ParseMoney[currency.USD]("19.99")
	fmt.Println(price, price.BigInt(), err)

	_, err = fulus.ParseMoney[currency.USD]("19.999")
	fmt.Println(err != nil)
	// Output:
	// USD 19.99 1999 <nil>
	// true
}

func ExampleNullMoney() {
	var discount fulus.NullMoney[currency.USD]
	if err := discount.Scan(nil); err != nil {
		panic(err)
	}
	data, _ := json.Marshal(discount)
	fmt.Println(discount.Valid, string(data))

	// A driver returns a NUMERIC value as text.
	if err := discount.Scan([]byte("2.50")); err != nil {
		panic(err)
	}
	value, _ := discount.Value()
	fmt.Println(discount.Valid, discount.Money, value)
	// Output:
	// false null
	// true USD 2.50 2.50
}
