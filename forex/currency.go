package forex

var _ Currency = LBP{}
var _ Currency = USD{}

type LBP struct{}

type USD struct{}

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

func (u USD) Code() string {
	return "USD"
}

func (u USD) Symbol() string {
	return "$"
}

func (u USD) MinorUnit() uint8 {
	return 2
}

func (u USD) SubUnitSymbol() string {
	return "¢"
}
