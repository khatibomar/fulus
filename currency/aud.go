package currency

var _ Currency = AUD{}

type AUD struct{}

func (a AUD) Code() string {
	return "AUD"
}

func (a AUD) Symbol() string {
	return "A$"
}

func (a AUD) MinorUnit() uint8 {
	return 2
}

func (a AUD) SubUnitSymbol() string {
	return "c"
}
