package currency

var _ Currency = LBP{}

type LBP struct{}

func (l LBP) Code() string {
	return "LBP"
}

func (l LBP) Symbol() string {
	return "ل.ل"
}

func (l LBP) MinorUnit() uint8 {
	return 0
}

func (l LBP) SubUnitSymbol() string {
	return ""
}
