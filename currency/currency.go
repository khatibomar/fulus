package currency

// Currency is a currency, such as an ISO 4217 currency or a custom token.
// A built-in currency also implements Numbered, Named and, if it has a cash rule, CashRounder.
// The format package gives CLDR format data by the code. A custom currency can implement format.Formatter.
type Currency interface {
	// Code returns the currency code, for example "USD". For an ISO 4217 currency it has three letters.
	Code() string

	// MinorUnits returns the number of digits after the decimal separator, for example 2 for USD.
	MinorUnits() int
}

// Unit is the constraint for the currency type parameter of fulus.Money.
// A Unit must be an empty struct type, so that its zero value is the only value
// and the methods cannot depend on state. For example, fulus.Money[currency.Currency] does not compile.
type Unit interface {
	~struct{}
	Currency
}

// Numbered is a Currency with an ISO 4217 numeric code.
type Numbered interface {
	// Number returns the three-digit ISO 4217 numeric code, for example "840" for USD.
	Number() string
}

// Named is a Currency with a name.
type Named interface {
	// Name returns the ISO 4217 currency name, for example "US Dollar".
	Name() string
}

// Number returns the ISO 4217 numeric code of c, or "" if c does not implement Numbered.
func Number(c Currency) string {
	if n, ok := c.(Numbered); ok {
		return n.Number()
	}
	return ""
}
