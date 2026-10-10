package currency_test

import (
	"fmt"

	"github.com/khatibomar/fulus/currency"
)

func ExampleByCode() {
	c, ok := currency.ByCode("jpy")
	fmt.Println(c.Code(), c.Number(), c.MinorUnits(), ok)

	_, ok = currency.ByCode("XYZ")
	fmt.Println(ok)
	// Output:
	// JPY 392 0 true
	// false
}

func ExampleByNumber() {
	c, ok := currency.ByNumber("978")
	fmt.Println(c.Code(), c.Name(), ok)
	// Output: EUR Euro true
}
