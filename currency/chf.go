package currency

var _ Currency = CHF{}

type CHF struct{}

func (CHF) Code() string            { return "CHF" }
func (CHF) Number() int             { return 756 }
func (CHF) Name() string            { return "Swiss Franc" }
func (CHF) MinorUnits() int         { return 2 }
func (CHF) Symbol() string          { return "Fr." }
func (CHF) MinorUnitName() string   { return "rappen" }
func (CHF) MinorUnitSymbol() string { return "Rp." }
