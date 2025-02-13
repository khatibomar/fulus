package currency

var _ Currency = CHF{}

type CHF struct{}

func (c CHF) Code() string {
	return "CHF"
}

func (c CHF) Symbol() string {
	return "CHF"
}

func (c CHF) MinorUnit() uint8 {
	return 2
}

func (c CHF) SubUnitSymbol() string {
	return "Rp."
}
