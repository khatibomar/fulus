package locale

import (
	"slices"
	"strings"
)

// Locale represents a supported CLDR locale.
// The zero Locale is not a CLDR locale. It uses generic number data and English currency symbols.
type Locale struct {
	id uint16
}

// Numbers holds the CLDR number data that formats currency amounts in a locale.
type Numbers struct {
	// CurrencyFormat is the CLDR currency pattern, for example "#,##0.00 ¤" or "¤#,##0.00;¤-#,##0.00".
	CurrencyFormat string
	// GroupSeparator separates digit groups, for example "," in en.
	GroupSeparator string
	// DecimalSeparator separates the fraction, for example "." in en.
	DecimalSeparator string
	// MinusSign replaces "-" in the pattern.
	MinusSign string
	// MinimumGroupingDigits is the smallest number of digits before the first group separator, for example 2 in es.
	MinimumGroupingDigits int
}

type data struct {
	code    string
	numbers Numbers
	symbols uint16
	// currencyNumbers is 1 + the index of the first currencyNumbers entry of the locale, or 0 if it has none.
	currencyNumbers uint16
}

type currencyData struct {
	locale  uint16
	code    string
	numbers Numbers
}

type symbol struct {
	code   string
	symbol string
}

// ParseLocale returns the locale with the exact CLDR code, for example "en-IN".
// It returns false if the locale is not supported. Use Match for a tag in another form.
func ParseLocale(code string) (Locale, bool) {
	supported := localeData[1:]
	i, found := slices.BinarySearchFunc(supported, code, func(d data, code string) int {
		return strings.Compare(d.code, code)
	})
	if !found {
		return Locale{}, false
	}
	return Locale{id: uint16(i + 1)}, true
}

// String returns the CLDR code of the locale, or "" for the zero Locale.
func (l Locale) String() string {
	return localeData[l.id].code
}

// Numbers returns the CLDR number data of the locale.
func (l Locale) Numbers() Numbers {
	return localeData[l.id].numbers
}

// CurrencyNumbers returns the CLDR number data of a currency in the locale, for example "¤#,##0.00" for EUR in en-SK.
func (l Locale) CurrencyNumbers(code string) Numbers {
	if first := localeData[l.id].currencyNumbers; first > 0 {
		for _, d := range currencyNumbers[first-1:] {
			if d.locale != l.id {
				break
			}
			if d.code == code {
				return d.numbers
			}
		}
	}
	return l.Numbers()
}

// CurrencySymbol returns the CLDR symbol of a currency in the locale, for example "US$" for "USD" in en-CA.
// It returns the code if the locale has no other symbol for the currency.
func (l Locale) CurrencySymbol(code string) string {
	set := symbolSets[localeData[l.id].symbols]
	i, found := slices.BinarySearchFunc(set, code, func(s symbol, code string) int {
		return strings.Compare(s.code, code)
	})
	if !found {
		return code
	}
	return set[i].symbol
}
