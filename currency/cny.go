package currency

var _ Currency = CNY{}

type CNY struct{}

func (c CNY) Code() string {
	return "CNY"
}

func (c CNY) Symbol() string {
	return "¥"
}

func (c CNY) MinorUnit() uint8 {
	return 2
}

func (c CNY) SubUnitSymbol() string {
	return "分"
}
