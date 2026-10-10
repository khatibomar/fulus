package fulus

import "github.com/khatibomar/fulus/currency"

// amount64 returns the amount in minor units. It panics if the amount does not fit in int64.
func (m Money[T]) amount64() int64 {
	v, ok := m.Int64()
	if !ok {
		panic("amount does not fit in int64")
	}
	return v
}

// amount64 returns the amount in minor units. It panics if the amount does not fit in int64.
func (m AnyMoney) amount64() int64 {
	v, ok := m.Int64()
	if !ok {
		panic("amount does not fit in int64")
	}
	return v
}

// maxMoney and minMoney return the largest and the smallest Money values.
func maxMoney[T currency.Unit]() Money[T] { return Money[T]{amount: maxInt128} }
func minMoney[T currency.Unit]() Money[T] { return Money[T]{amount: minInt128} }
