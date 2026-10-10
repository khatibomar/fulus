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

// Registry is a set of currencies that can be found by code or by ISO 4217 number.
// It is safe for concurrent use. The zero Registry is empty and ready to use.
//
// The package functions ByCode, ByNumber, Register and All use the Default registry.
// Use a separate Registry to keep custom currencies out of the global state, for example in a test or a tenant.
type Registry struct {
	mu       sync.RWMutex
	byCode   map[string]Currency
	byNumber map[string]Currency
}

// NewRegistry returns a registry with the given currencies.
// Use Builtin() to start from the generated ISO 4217 currencies.
// Returns an error if a currency is not valid or is a duplicate. See Register.
func NewRegistry(currencies ...Currency) (*Registry, error) {
	r := &Registry{}
	for _, c := range currencies {
		if err := r.Register(c); err != nil {
			return nil, err
		}
	}
	return r, nil
}

var defaultRegistry = sync.OnceValue(func() *Registry {
	r, err := NewRegistry(builtin...)
	if err != nil {
		panic(err)
	}
	return r
})

// Default returns the registry that the package functions use. It starts with the generated currencies.
func Default() *Registry {
	return defaultRegistry()
}

// Builtin returns the generated ISO 4217 currencies, sorted by code.
func Builtin() []Currency {
	return slices.Clone(builtin)
}

// ByCode returns the currency with the given code. The lookup is not case-sensitive.
func (r *Registry) ByCode(code string) (Currency, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	c, ok := r.byCode[strings.ToUpper(code)]
	return c, ok
}

// ByNumber returns the currency with the given ISO 4217 numeric code, for example "840".
func (r *Registry) ByNumber(number string) (Currency, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	c, ok := r.byNumber[number]
	return c, ok
}

// Register adds a currency, so that ByCode and ByNumber can find it.
// The code must be upper case. If c does not implement Numbered or its number is empty, ByNumber does not find it.
// Returns ErrInvalidCurrency if the code is not valid or the minor units are negative,
// and ErrDuplicateCurrency if the code or the number is already registered.
func (r *Registry) Register(c Currency) error {
	if c == nil {
		return fmt.Errorf("%w: nil currency", ErrInvalidCurrency)
	}
	code := c.Code()
	if code == "" || code != strings.ToUpper(code) {
		return fmt.Errorf("%w: code %q must be upper case and not empty", ErrInvalidCurrency, code)
	}
	if c.MinorUnits() < 0 {
		return fmt.Errorf("%w: %s has negative minor units", ErrInvalidCurrency, code)
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.byCode[code]; ok {
		return fmt.Errorf("%w: code %s", ErrDuplicateCurrency, code)
	}
	number := Number(c)
	if _, ok := r.byNumber[number]; ok && number != "" {
		return fmt.Errorf("%w: number %s", ErrDuplicateCurrency, number)
	}

	if r.byCode == nil {
		r.byCode = make(map[string]Currency)
		r.byNumber = make(map[string]Currency)
	}
	r.byCode[code] = c
	if number != "" {
		r.byNumber[number] = c
	}
	return nil
}

// All returns all currencies in the registry, sorted by code.
func (r *Registry) All() []Currency {
	r.mu.RLock()
	all := make([]Currency, 0, len(r.byCode))
	for _, c := range r.byCode {
		all = append(all, c)
	}
	r.mu.RUnlock()

	slices.SortFunc(all, func(a, b Currency) int { return strings.Compare(a.Code(), b.Code()) })
	return all
}

// ByCode returns the currency with the given code from the Default registry. The lookup is not case-sensitive.
func ByCode(code string) (Currency, bool) {
	return Default().ByCode(code)
}

// ByNumber returns the currency with the given ISO 4217 numeric code from the Default registry.
func ByNumber(number string) (Currency, bool) {
	return Default().ByNumber(number)
}

// Register adds a custom currency to the Default registry. See Registry.Register.
func Register(c Currency) error {
	return Default().Register(c)
}

// All returns all currencies in the Default registry, sorted by code.
func All() []Currency {
	return Default().All()
}
