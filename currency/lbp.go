package currency

var _ Currency = LBP{}

type LBP struct{}

func (LBP) Code() string            { return "LBP" }
func (LBP) Number() int             { return 422 }
func (LBP) Name() string            { return "Lebanese Pound" }
func (LBP) MinorUnits() int         { return 2 }
func (LBP) Symbol() string          { return "ل.ل" }
func (LBP) MinorUnitName() string   { return "piastre" }
func (LBP) MinorUnitSymbol() string { return "p" }
