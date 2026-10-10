package currency

import (
	"errors"
	"fmt"
	"slices"
	"strings"
	"sync"
)

var (
	// ErrInvalidCurrency indicates a currency with an empty code.
	ErrInvalidCurrency = errors.New("invalid currency")

	// ErrDuplicateCurrency indicates a currency code or number that is already registered.
	ErrDuplicateCurrency = errors.New("currency already registered")
)

// CashRounder is a Currency that has a cash rounding rule.
// For example, Swiss francs in cash use steps of 0.05, so CHF returns 5.
type CashRounder interface {
	// CashIncrement returns the smallest cash amount in minor units.
	CashIncrement() int64
}

var registry = sync.OnceValue(func() *currencyRegistry {
	r := &currencyRegistry{
		byCode:   make(map[string]Currency, len(builtin)),
		byNumber: make(map[string]Currency, len(builtin)),
	}
	for _, c := range builtin {
		r.byCode[c.Code()] = c
		r.byNumber[Number(c)] = c
	}
	return r
})

type currencyRegistry struct {
	mu       sync.RWMutex
	byCode   map[string]Currency
	byNumber map[string]Currency
}

// ByCode returns the registered currency with the given ISO 4217 code.
// The lookup is not case-sensitive.
func ByCode(code string) (Currency, bool) {
	r := registry()
	r.mu.RLock()
	defer r.mu.RUnlock()
	c, ok := r.byCode[strings.ToUpper(code)]
	return c, ok
}

// ByNumber returns the registered currency with the given ISO 4217 numeric code, for example "840".
func ByNumber(number string) (Currency, bool) {
	r := registry()
	r.mu.RLock()
	defer r.mu.RUnlock()
	c, ok := r.byNumber[number]
	return c, ok
}

// Register adds a custom currency, so that ByCode and ByNumber can find it.
// The code must be upper case. If c does not implement Numbered or its number is empty, ByNumber does not find it.
// Register returns ErrDuplicateCurrency if the code or the number is already registered.
func Register(c Currency) error {
	code := c.Code()
	if code == "" || code != strings.ToUpper(code) {
		return fmt.Errorf("%w: code %q must be upper case and not empty", ErrInvalidCurrency, code)
	}

	r := registry()
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.byCode[code]; ok {
		return fmt.Errorf("%w: code %s", ErrDuplicateCurrency, code)
	}
	number := Number(c)
	if _, ok := r.byNumber[number]; ok && number != "" {
		return fmt.Errorf("%w: number %s", ErrDuplicateCurrency, number)
	}

	r.byCode[code] = c
	if number != "" {
		r.byNumber[number] = c
	}
	return nil
}

// All returns all registered currencies, sorted by code.
func All() []Currency {
	r := registry()
	r.mu.RLock()
	all := make([]Currency, 0, len(r.byCode))
	for _, c := range r.byCode {
		all = append(all, c)
	}
	r.mu.RUnlock()

	slices.SortFunc(all, func(a, b Currency) int { return strings.Compare(a.Code(), b.Code()) })
	return all
}
