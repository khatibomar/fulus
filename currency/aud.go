package currency

var _ Currency = AUD{}

type AUD struct{}

func (AUD) Code() string            { return "AUD" }
func (AUD) Number() int             { return 036 }
func (AUD) Name() string            { return "Australian Dollar" }
func (AUD) MinorUnits() int         { return 2 }
func (AUD) Symbol() string          { return "$" }
func (AUD) MinorUnitName() string   { return "cent" }
func (AUD) MinorUnitSymbol() string { return "c" }
