package currency

var _ Currency = GBP{}

type GBP struct{}

func (g GBP) Code() string {
	return "GBP"
}

func (g GBP) Symbol() string {
	return "£"
}

func (g GBP) MinorUnit() uint8 {
	return 2
}

func (g GBP) SubUnitSymbol() string {
	return "p"
}
