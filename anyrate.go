package fulus

import (
	"fmt"

	"github.com/khatibomar/fulus/currency"
)

// AnyRate is an exchange rate between two currencies that are known only at run time,
// for example from a rate feed or a database row. It is the run-time form of Rate[Base, Quote].
// It is the price of one major unit of Base in major units of Quote, as an exact positive fraction in lowest terms.
// The zero value is not a valid rate. AnyMoney.Convert returns ErrInvalidExchangeRate for it.
type AnyRate struct {
	base, quote currency.Currency
	num, den    int64
}

// NewAnyRate returns the rate numerator/denominator from base to quote.
// Returns ErrUnknownCurrency if a currency is nil, and ErrInvalidExchangeRate if a term is not positive.
func NewAnyRate(numerator, denominator int64, base, quote currency.Currency) (AnyRate, error) {
	if err := checkRateCurrencies(base, quote); err != nil {
		return AnyRate{}, err
	}
	num, den, err := reduceRate(numerator, denominator)
	if err != nil {
		return AnyRate{}, err
	}
	return AnyRate{base: base, quote: quote, num: num, den: den}, nil
}

// ParseAnyRate parses a rate from base to quote with the same rules as ParseRate.
// Returns ErrUnknownCurrency if a currency is nil, ErrInvalidExchangeRate if the rate cannot be parsed
// or is not positive, and ErrOverflow if a term of the fraction in lowest terms does not fit in int64.
func ParseAnyRate(s string, base, quote currency.Currency) (AnyRate, error) {
	if err := checkRateCurrencies(base, quote); err != nil {
		return AnyRate{}, err
	}
	num, den, err := parseRate(s)
	if err != nil {
		return AnyRate{}, err
	}
	return AnyRate{base: base, quote: quote, num: num, den: den}, nil
}

func checkRateCurrencies(base, quote currency.Currency) error {
	if base == nil || quote == nil {
		return fmt.Errorf("%w: nil currency in rate", ErrUnknownCurrency)
	}
	return nil
}

// Any returns the rate as an AnyRate. The zero Rate gives an AnyRate that is not valid.
func (r Rate[Base, Quote]) Any() AnyRate {
	var base Base
	var quote Quote
	return AnyRate{base: base, quote: quote, num: r.num, den: r.den}
}

// AsRate returns r as a Rate[Base, Quote].
// Returns ErrInvalidExchangeRate if r is not valid,
// and ErrCurrencyMismatch if the currencies of r do not have the codes and the minor units of Base and Quote.
func AsRate[Base, Quote currency.Unit](r AnyRate) (Rate[Base, Quote], error) {
	if !r.IsValid() {
		return Rate[Base, Quote]{}, ErrInvalidExchangeRate
	}
	var base Base
	var quote Quote
	if !sameCurrency(r.base, base) || !sameCurrency(r.quote, quote) {
		return Rate[Base, Quote]{}, fmt.Errorf("%w: expected %s/%s, got %s/%s", ErrCurrencyMismatch,
			describe(base, r.base), describe(quote, r.quote), describe(r.base, base), describe(r.quote, quote))
	}
	return Rate[Base, Quote]{num: r.num, den: r.den}, nil
}

// Base returns the currency that the rate converts from, or nil for the zero value.
func (r AnyRate) Base() currency.Currency {
	return r.base
}

// Quote returns the currency that the rate converts to, or nil for the zero value.
func (r AnyRate) Quote() currency.Currency {
	return r.quote
}

// Fraction returns the rate as numerator/denominator in lowest terms.
// Both terms are zero for the zero AnyRate.
func (r AnyRate) Fraction() (numerator, denominator int64) {
	return r.num, r.den
}

// IsValid reports whether r is a valid rate. Only the zero AnyRate is not valid.
func (r AnyRate) IsValid() bool {
	return r.den > 0 && r.base != nil && r.quote != nil
}

// Invert returns the rate in the other direction, like Rate.Invert.
func (r AnyRate) Invert() AnyRate {
	return AnyRate{base: r.quote, quote: r.base, num: r.den, den: r.num}
}

// String returns the codes and the rate in the form of Rate.String, for example "EUR/USD 1.07203".
// The zero value gives "0".
func (r AnyRate) String() string {
	if !r.IsValid() {
		return "0"
	}
	return r.base.Code() + "/" + r.quote.Code() + " " + rateString(r.num, r.den)
}

// Convert changes m to the quote currency of r and rounds the result with mode, like Convert.
// Returns ErrInvalidExchangeRate if r is not valid, ErrCurrencyMismatch if the currency of m is not
// the base currency of r, and ErrOverflow if the result does not fit.
func (m AnyMoney) Convert(r AnyRate, mode RoundingMode) (AnyMoney, error) {
	if !r.IsValid() {
		return AnyMoney{}, ErrInvalidExchangeRate
	}
	if !mode.valid() {
		return AnyMoney{}, ErrInvalidRoundingMode
	}
	if !m.sameCurrency(r.base) {
		return AnyMoney{}, fmt.Errorf("%w: rate from %s cannot convert %s", ErrCurrencyMismatch, describe(r.base, m.currency), describe(m.currency, r.base))
	}
	amount, err := convertAmount(m.amount, r.num, r.den, r.base, r.quote, mode)
	if err != nil {
		return AnyMoney{}, err
	}
	return AnyMoney{amount: amount, currency: r.quote}, nil
}
