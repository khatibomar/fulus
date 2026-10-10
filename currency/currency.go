package currency

import "github.com/khatibomar/fulus/locale"

// Currency represents an ISO 4217 currency with locale support
type Currency interface {
	// Code returns the three-letter ISO 4217 currency code
	Code() string

	// Number returns the three-digit ISO 4217 numeric code
	Number() string

	// Name returns the official ISO 4217 currency name
	Name() string

	// MinorUnits returns the number of digits after the decimal separator
	MinorUnits() int

	// FormatInfo returns the currency formatting information for a given locale
	FormatInfo(locale locale.Locale) FormatInfo
}

// FormatInfo contains locale-specific currency formatting information
type FormatInfo struct {
	Symbol           string // Currency symbol for the locale
	Format           string // Format pattern
	GroupSeparator   string // Thousands separator
	DecimalSeparator string // Decimal separator
	MinusSign        string // Negative number prefix
	// MinimumGroupingDigits is the smallest number of digits before the first group separator. Zero means 1.
	MinimumGroupingDigits int
}

// formatInfo returns the CLDR format information of a currency code in a locale.
func formatInfo(loc locale.Locale, code string) FormatInfo {
	n := loc.CurrencyNumbers(code)
	return FormatInfo{
		Symbol:                loc.CurrencySymbol(code),
		Format:                n.CurrencyFormat,
		GroupSeparator:        n.GroupSeparator,
		DecimalSeparator:      n.DecimalSeparator,
		MinusSign:             n.MinusSign,
		MinimumGroupingDigits: n.MinimumGroupingDigits,
	}
}
