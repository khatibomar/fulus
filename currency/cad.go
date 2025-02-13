package currency

var _ Currency = CAD{}

type CAD struct{}

func (CAD) Code() string            { return "CAD" }
func (CAD) Number() int             { return 124 }
func (CAD) Name() string            { return "Canadian Dollar" }
func (CAD) MinorUnits() int         { return 2 }
func (CAD) Symbol() string          { return "$" }
func (CAD) MinorUnitName() string   { return "cent" }
func (CAD) MinorUnitSymbol() string { return "¢" }
