package format_test

import (
	"fmt"

	"github.com/khatibomar/fulus"
	"github.com/khatibomar/fulus/currency"
	"github.com/khatibomar/fulus/format"
	"github.com/khatibomar/fulus/locale"
)

func ExampleParse() {
	loc, _ := locale.Match("en_IN.UTF-8")
	inr := fulus.NewMoney[currency.INR](-1234567890)
	formatted := format.Money(inr, loc)

	parsed, err := format.Parse[currency.INR](formatted, loc)
	fmt.Println(formatted, parsed.Decimal(), err)
	// Output: -₹1,23,45,678.90 -12345678.90 <nil>
}

func ExampleMoney() {
	price := fulus.NewMoney[currency.EUR](-123456789)
	fmt.Println(format.Money(price, locale.EN))
	fmt.Println(format.Money(price, locale.DE_CH))

	// Some locales use a no-break space U+00A0 or a narrow no-break space U+202F.
	fmt.Printf("%+q\n", format.Money(price, locale.FR))
	// Output:
	// -€1,234,567.89
	// EUR-1'234'567.89
	// "-1\u202f234\u202f567,89\u00a0\u20ac"
}
