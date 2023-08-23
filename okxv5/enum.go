package okxv5

type Category string

const (
	Spot    Category = "SPOT"
	Margin  Category = "MARGIN"
	Swap    Category = "SWAP"
	Futures Category = "Futures"
	Option  Category = "Option"
)
