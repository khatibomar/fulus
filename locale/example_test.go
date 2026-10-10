package locale_test

import (
	"fmt"

	"github.com/khatibomar/fulus/locale"
)

func ExampleMatch() {
	for _, tag := range []string{"en_US.UTF-8", "de-ch", "fr-CA-u-nu-latn", "xx"} {
		loc, ok := locale.Match(tag)
		fmt.Println(tag, loc, ok)
	}
	// Output:
	// en_US.UTF-8 en true
	// de-ch de-CH true
	// fr-CA-u-nu-latn fr-CA true
	// xx  false
}

func ExampleLocale_CurrencySymbol() {
	fmt.Println(locale.EN.CurrencySymbol("USD"))
	fmt.Println(locale.EN_CA.CurrencySymbol("USD"))
	fmt.Println(locale.EN.CurrencySymbol("CHF"))
	// Output:
	// $
	// US$
	// CHF
}
