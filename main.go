package main

import (
	"fmt"

	"github.com/BurntSushi/toml"

	"github.com/khatibomar/money-experiment/forex"
)

//go:generate go run money_generator.go

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

	m1 := forex.NewMoney(500)
	m2 := forex.NewMoney(500)

	err := m1.Add(m2)
	if err != nil {
		fmt.Printf("err: %s\n", err.Error())
	}
	err = m1.Add(m2)
	if err != nil {
		fmt.Printf("err: %s\n", err.Error())
	}
	fmt.Println(m1)

	err = m1.Mul(10)
	if err != nil {
		fmt.Printf("err: %s\n", err.Error())
	}

	err = m1.Validate(config.Money.Min, config.Money.Max)
	if err != nil {
		fmt.Printf("err: %s\n", err.Error())
	} else {
		fmt.Printf("%s is valid\n", m1.String())
	}

	fmt.Println(m1)
}
