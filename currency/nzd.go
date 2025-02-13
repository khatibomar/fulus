package currency

var _ Currency = NZD{}

type NZD struct{}

func (n NZD) Code() string {
	return "NZD"
}

func (n NZD) Symbol() string {
	return "NZ$"
}

func (n NZD) MinorUnit() uint8 {
	return 2
}

func (n NZD) SubUnitSymbol() string {
	return "c"
}
