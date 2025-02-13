package currency

var _ Currency = EUR{}

type EUR struct{}

func (EUR) Code() string            { return "EUR" }
func (EUR) Number() int             { return 978 }
func (EUR) Name() string            { return "Euro" }
func (EUR) MinorUnits() int         { return 2 }
func (EUR) Symbol() string          { return "€" }
func (EUR) MinorUnitName() string   { return "cent" }
func (EUR) MinorUnitSymbol() string { return "c" }
