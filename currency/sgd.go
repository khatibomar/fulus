package currency

var _ Currency = SGD{}

type SGD struct{}

func (SGD) Code() string            { return "SGD" }
func (SGD) Number() int             { return 702 }
func (SGD) Name() string            { return "Singapore Dollar" }
func (SGD) MinorUnits() int         { return 2 }
func (SGD) Symbol() string          { return "$" }
func (SGD) MinorUnitName() string   { return "cent" }
func (SGD) MinorUnitSymbol() string { return "¢" }
