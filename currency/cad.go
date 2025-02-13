package currency

var _ Currency = CAD{}

type CAD struct{}

func (c CAD) Code() string {
	return "CAD"
}

func (c CAD) Symbol() string {
	return "C$"
}

func (c CAD) MinorUnit() uint8 {
	return 2
}

func (c CAD) SubUnitSymbol() string {
	return "¢"
}
