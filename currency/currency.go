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
	FormatInfo(locale locale.Locale) CurrencyFormatInfo
}

// CurrencyFormatInfo contains locale-specific currency formatting information
type CurrencyFormatInfo struct {
	Symbol           string // Currency symbol for the locale
	Format           string // Format pattern
	GroupSeparator   string // Thousands separator
	DecimalSeparator string // Decimal separator
	MinusSign        string // Negative number prefix
}
