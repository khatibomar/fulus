package currency

var _ Currency = EUR{}

type EUR struct{}

func (e EUR) Code() string {
	return "EUR"
}

func (e EUR) Symbol() string {
	return "€"
}

func (e EUR) MinorUnit() uint8 {
	return 2
}

func (e EUR) SubUnitSymbol() string {
	return "c"
}
