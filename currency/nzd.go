package currency

var _ Currency = NZD{}

type NZD struct{}

func (NZD) Code() string            { return "NZD" }
func (NZD) Number() int             { return 554 }
func (NZD) Name() string            { return "New Zealand Dollar" }
func (NZD) MinorUnits() int         { return 2 }
func (NZD) Symbol() string          { return "$" }
func (NZD) MinorUnitName() string   { return "cent" }
func (NZD) MinorUnitSymbol() string { return "c" }
