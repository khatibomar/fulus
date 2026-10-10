package main

import (
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/khatibomar/fulus-examples/codegen/money"
)

//go:generate go run money/money_generator.go

type Config struct {
	Money struct {
		Currency string `toml:"currency"`
		Min      int64  `toml:"min"`
		Max      int64  `toml:"max"`
	} `toml:"money"`
}

func main() {
	var config Config
	if _, err := toml.DecodeFile("config.toml", &config); err != nil {
		panic(err)
	}

	low, high := money.New(config.Money.Min), money.New(config.Money.Max)
	m1 := money.New(500)
	m2 := money.New(500)

	fmt.Println(m1)

	m1, err := m1.Add(m2)
	printErr(err)

	m1, err = do(m1)
	printErr(err)
	fmt.Println(m1)

	m1, err = m1.Mul(10)
	printErr(err)

	printRange(m1, low, high)

	m1, err = m1.Mul(50)
	printErr(err)

	printRange(m1, low, high)
}

func printRange(m, low, high money.Money) {
	if m.InRange(low, high) {
		fmt.Printf("%s is in range\n", m)
		return
	}
	fmt.Printf("%s is not in [%s, %s]\n", m, low, high)
}

func printErr(err error) {
	if err != nil {
		fmt.Printf("err: %s\n", err.Error())
	}
}

func do(m money.Money) (money.Money, error) {
	other := money.New(500)
	return m.Add(other)
}
