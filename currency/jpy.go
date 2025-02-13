package currency

var _ Currency = JPY{}

type JPY struct{}

func (j JPY) Code() string {
	return "JPY"
}

func (j JPY) Symbol() string {
	return "¥"
}

func (j JPY) MinorUnit() uint8 {
	return 0
}

func (j JPY) SubUnitSymbol() string {
	return ""
}
