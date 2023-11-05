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

type TradeMode string

const (
	CrossMode    TradeMode = "cross"
	IsolatedMode TradeMode = "isolated"
	CashMode     TradeMode = "cash"
)

type Side string

const (
	Buy  Side = "buy"
	Sell Side = "sell"
)

type OrderType string

const (
	LimitType       OrderType = "limit"
	MarketType      OrderType = "market"
	PostOnly        OrderType = "post_only"
	Fok             OrderType = "fok"
	Ioc             OrderType = "ioc"
	OptimalLimitIoc OrderType = "optimal_limit_ioc"
	Mmp             OrderType = "mmp"
	MmpAndPostOnly  OrderType = "mmp_and_post_only"
)
