package fulus

import (
	"testing"
	"testing/quick"

	"github.com/khatibomar/fulus/currency"
)

func TestPropertyAddIsCommutative(t *testing.T) {
	property := func(a, b int64) bool {
		left, leftErr := NewMoney[currency.USD](a).Add(NewMoney[currency.USD](b))
		right, rightErr := NewMoney[currency.USD](b).Add(NewMoney[currency.USD](a))

		if leftErr != nil || rightErr != nil {
			return leftErr == ErrOverflow && rightErr == ErrOverflow
		}

		return left.Amount() == right.Amount()
	}

	if err := quick.Check(property, nil); err != nil {
		t.Error(err)
	}
}

func TestPropertyAddThenSubReturnsOriginal(t *testing.T) {
	property := func(a, b int64) bool {
		sum, err := NewMoney[currency.USD](a).Add(NewMoney[currency.USD](b))
		if err != nil {
			return true
		}

		recovered, err := sum.Sub(NewMoney[currency.USD](b))
		if err != nil {
			return false
		}

		return recovered.Amount() == a
	}

	if err := quick.Check(property, nil); err != nil {
		t.Error(err)
	}
}

func TestPropertyDistributePreservesAmount(t *testing.T) {
	property := func(amount int64, rawChunks uint8) bool {
		chunks := int64(rawChunks) + 1

		dist, err := NewMoney[currency.USD](amount).Distribute(chunks)
		if err != nil {
			return false
		}

		total := (dist.SmallerChunkSize * dist.SmallerCount) +
			(dist.LargerChunkSize * dist.LargerCount)

		return total == amount &&
			dist.SmallerCount+dist.LargerCount == chunks
	}

	if err := quick.Check(property, nil); err != nil {
		t.Error(err)
	}
}

func TestPropertyAllocatePreservesAmount(t *testing.T) {
	property := func(amount int64, r1, r2, r3 uint8) bool {
		ratios := []int64{int64(r1) + 1, int64(r2) + 1, int64(r3) + 1}

		allocation, err := NewMoney[currency.USD](amount).Allocate(ratios)
		if err != nil {
			return false
		}

		sum := int64(0)
		for _, part := range allocation.Parts {
			sum += part.Amount()
		}

		return sum == amount && allocation.Total.Amount() == amount
	}

	if err := quick.Check(property, nil); err != nil {
		t.Error(err)
	}
}
