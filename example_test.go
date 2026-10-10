package fulus_test

import (
	"encoding/json"
	"fmt"

	"github.com/khatibomar/fulus"
	"github.com/khatibomar/fulus/currency"
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

func ExampleSum() {
	total, err := fulus.Sum(
		fulus.NewMoney[currency.USD](1000),
		fulus.NewMoney[currency.USD](250),
		fulus.NewMoney[currency.USD](-50),
	)
	fmt.Println(total, err)
	// Output: USD 12.00 <nil>
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

func ExampleMoney_InRange() {
	low, high := fulus.NewMoney[currency.USD](100), fulus.NewMoney[currency.USD](10000)
	fmt.Println(fulus.NewMoney[currency.USD](2500).InRange(low, high))
	fmt.Println(fulus.NewMoney[currency.USD](50).InRange(low, high))
	// Output:
	// true
	// false
}

func ExampleAnyMoney_MulFactor() {
	price, err := fulus.ParseAnyMoney("12.50", "EUR")
	if err != nil {
		panic(err)
	}
	vat, err := price.MulFactor(fulus.Percent(19), fulus.RoundHalfUp)
	fmt.Println(vat, err)
	// Output: EUR 2.38 <nil>
}

func ExampleAnyMoney_ScanColumns() {
	var price fulus.AnyMoney
	amount, code := price.ScanColumns()

	// rows.Scan(amount, code) calls Scan for each column, with the values from the driver.
	if err := amount.Scan([]byte("12.50")); err != nil {
		panic(err)
	}
	if err := code.Scan([]byte("EUR")); err != nil {
		panic(err)
	}
	fmt.Println(price)
	// Output: EUR 12.50
}
