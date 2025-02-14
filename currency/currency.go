package currency

// Currency represents an ISO 4217 currency with its minor unit information.
// It contains all standard currency attributes as defined by ISO 4217
// along with commonly used symbols and minor unit representations.
type Currency interface {
	// Code is the three-letter ISO 4217 currency code (e.g., "USD", "EUR")
	Code() string

	// Number is the three-digit ISO 4217 numeric code
	Number() string

	// Name is the official ISO 4217 currency name
	Name() string

	// MinorUnits specifies the number of digits after the decimal separator
	MinorUnits() int

	// Symbol is the currency's commonly used symbol (e.g., "$", "€")
	Symbol() string

	// MinorUnitName is the name of the fractional monetary unit
	MinorUnitName() string

	// MinorUnitSymbol represents the symbol used for the minor unit
	MinorUnitSymbol() string
}
