package main

import (
	"fmt"

	"github.com/khatibomar/fulus"
	"github.com/khatibomar/fulus/currency"
)

// Embed currency.USD to inherit all methods
type localUSD struct {
	currency.USD
}

// Override Symbol() method
func (l localUSD) Symbol() string {
	return "$US"
}

func main() {
	usd := fulus.NewMoney[localUSD](1000)
	fmt.Println(usd) // $US10.00
}
