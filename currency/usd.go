package currency

var _ Currency = USD{}

type USD struct{}

func (u USD) Code() string {
	return "USD"
}

func (u USD) Symbol() string {
	return "$"
}

func (u USD) MinorUnit() uint8 {
	return 2
}

func (u USD) SubUnitSymbol() string {
	return "¢"
}
