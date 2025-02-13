package currency

var _ Currency = GBP{}

type GBP struct{}

func (GBP) Code() string            { return "GBP" }
func (GBP) Number() int             { return 826 }
func (GBP) Name() string            { return "Pound Sterling" }
func (GBP) MinorUnits() int         { return 2 }
func (GBP) Symbol() string          { return "£" }
func (GBP) MinorUnitName() string   { return "penny" }
func (GBP) MinorUnitSymbol() string { return "p" }
