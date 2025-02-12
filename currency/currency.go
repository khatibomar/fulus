package currency

// Currency represents a monetary unit that follows ISO 4217 standards.
type Currency interface {
	// Code returns the three-letter ISO 4217 currency code
	// Example: "USD", "EUR", "LBP"
	Code() string

	// Symbol returns the currency symbol
	// Example: "$" for USD, "€" for EUR, "ل.ل" for LBP
	Symbol() string

	// MinorUnit returns the number of decimal places used for the currency
	// as defined in ISO 4217. For example:
	// 2 for USD (cents: 100)
	// 3 for BHD (fils: 1000)
	// 0 for JPY (no minor unit)
	MinorUnit() uint8

	// SubUnitSymbol returns the symbol or abbreviation for the minor unit of the currency
	// Example: "¢" for USD cents, "p" for GBP pence
	SubUnitSymbol() string
}
