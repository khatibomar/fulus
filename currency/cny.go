package currency

var _ Currency = CNY{}

type CNY struct{}

func (CNY) Code() string            { return "CNY" }
func (CNY) Number() int             { return 156 }
func (CNY) Name() string            { return "Yuan Renminbi" }
func (CNY) MinorUnits() int         { return 2 }
func (CNY) Symbol() string          { return "¥" }
func (CNY) MinorUnitName() string   { return "fen" }
func (CNY) MinorUnitSymbol() string { return "分" }
