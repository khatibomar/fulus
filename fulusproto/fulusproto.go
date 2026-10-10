// Package fulusproto converts fulus values to and from google.type.Money, the protocol buffers type for money.
//
// google.type.Money holds whole units and nanos, so it has 9 fraction digits and an int64 range of units.
// A conversion never rounds. It returns an error when a value does not fit.
package fulusproto

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/khatibomar/fulus"
	"github.com/khatibomar/fulus/currency"
	"google.golang.org/genproto/googleapis/type/money"
)

// nanoDigits is the number of fraction digits of google.type.Money.
const nanoDigits = 9

// ToMoney returns m as a google.type.Money.
// Returns fulus.ErrOverflow if the whole units do not fit in int64,
// and fulus.ErrInexact if m has more than 9 fraction digits that are not zero.
func ToMoney[T currency.Unit](m fulus.Money[T]) (*money.Money, error) {
	return toMoney(m.Decimal(), m.Currency().Code())
}

// ToMoneyAny returns m as a google.type.Money, like ToMoney.
func ToMoneyAny(m fulus.AnyMoney) (*money.Money, error) {
	if m.Currency() == nil {
		return nil, fmt.Errorf("%w: AnyMoney has no currency", fulus.ErrUnknownCurrency)
	}
	return toMoney(m.Decimal(), m.Currency().Code())
}

func toMoney(decimal, code string) (*money.Money, error) {
	digits, negative := strings.CutPrefix(decimal, "-")
	whole, fraction, _ := strings.Cut(digits, ".")

	if len(fraction) > nanoDigits {
		if strings.Trim(fraction[nanoDigits:], "0") != "" {
			return nil, fmt.Errorf("%w: %s has more than %d fraction digits", fulus.ErrInexact, decimal, nanoDigits)
		}
		fraction = fraction[:nanoDigits]
	}
	sign := ""
	if negative {
		sign = "-"
	}
	// The sign is parsed with the units, so that math.MinInt64 fits.
	units, err := strconv.ParseInt(sign+whole, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("%w: units of %s do not fit in int64", fulus.ErrOverflow, decimal)
	}
	nanos := 0
	if fraction != "" {
		// The fraction has at most 9 digits, so it fits in int32.
		nanos, _ = strconv.Atoi(fraction + strings.Repeat("0", nanoDigits-len(fraction)))
	}
	if negative {
		nanos = -nanos
	}
	return &money.Money{CurrencyCode: code, Units: units, Nanos: int32(nanos)}, nil
}

// FromMoney returns p as a Money[T].
// Returns fulus.ErrCurrencyMismatch if the currency code of p is not the code of T. The code is not case-sensitive.
// fulus.ErrInvalidAmountFormat if p is nil or its nanos are not valid,
// and fulus.ErrScaleMismatch if p has a fraction digit that is not zero after the minor units of T.
func FromMoney[T currency.Unit](p *money.Money) (fulus.Money[T], error) {
	var c T
	if p != nil && !strings.EqualFold(p.GetCurrencyCode(), c.Code()) {
		return fulus.Money[T]{}, fmt.Errorf("%w: expected %s, got %s", fulus.ErrCurrencyMismatch, c.Code(), p.GetCurrencyCode())
	}
	decimal, err := fromMoney(p)
	if err != nil {
		return fulus.Money[T]{}, err
	}
	return fulus.ParseMoney[T](decimal)
}

// FromMoneyAny returns p as an AnyMoney. The currency code must be in currency.Default().
func FromMoneyAny(p *money.Money) (fulus.AnyMoney, error) {
	decimal, err := fromMoney(p)
	if err != nil {
		return fulus.AnyMoney{}, err
	}
	return fulus.ParseAnyMoney(decimal, p.GetCurrencyCode())
}

// fromMoney returns the canonical decimal of p, without trailing fraction zeros.
func fromMoney(p *money.Money) (string, error) {
	if p == nil {
		return "", fmt.Errorf("%w: nil google.type.Money", fulus.ErrInvalidAmountFormat)
	}
	units, nanos := p.GetUnits(), p.GetNanos()
	if nanos <= -1e9 || nanos >= 1e9 || (units > 0 && nanos < 0) || (units < 0 && nanos > 0) {
		return "", fmt.Errorf("%w: units %d and nanos %d are not valid", fulus.ErrInvalidAmountFormat, units, nanos)
	}

	var b strings.Builder
	if units < 0 || nanos < 0 {
		b.WriteByte('-')
	}
	b.WriteString(strconv.FormatUint(absInt64(units), 10))
	if nanos != 0 {
		fraction := strconv.Itoa(int(max(nanos, -nanos)))
		fraction = strings.Repeat("0", nanoDigits-len(fraction)) + fraction
		b.WriteByte('.')
		b.WriteString(strings.TrimRight(fraction, "0"))
	}
	return b.String(), nil
}

func absInt64(x int64) uint64 {
	if x < 0 {
		return -uint64(x)
	}
	return uint64(x)
}
