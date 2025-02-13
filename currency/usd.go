package currency

var _ Currency = USD{}

type USD struct{}

func (USD) Code() string { return "USD" }

func (USD) Number() int { return 840 }

func (USD) Name() string { return "United States Dollar" }

func (USD) MinorUnits() int { return 2 }

func (USD) Symbol() string { return "$" }

func (USD) MinorUnitName() string { return "cent" }

func (USD) MinorUnitSymbol() string { return "¢" }
