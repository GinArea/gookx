package okxv5

type Category string

const (
	Spot    Category = "SPOT"
	Margin  Category = "MARGIN"
	Swap    Category = "SWAP"
	Futures Category = "FUTURES"
	Option  Category = "OPTION"
)

type MarginMode string

const (
	Cross    MarginMode = "cross"
	Isolated MarginMode = "isolated"
)
