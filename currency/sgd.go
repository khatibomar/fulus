package currency

var _ Currency = SGD{}

type SGD struct{}

func (s SGD) Code() string {
	return "SGD"
}

func (s SGD) Symbol() string {
	return "S$"
}

func (s SGD) MinorUnit() uint8 {
	return 2
}

func (s SGD) SubUnitSymbol() string {
	return "¢"
}
