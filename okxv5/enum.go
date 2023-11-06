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

type OrderState string

const (
	Canceled        OrderState = "canceled"
	Live            OrderState = "live"
	PartiallyFilled OrderState = "partially_filled"
	Filled          OrderState = "filled"
	MmpCanceled     OrderState = "mmp_canceled"
)

type TriggerPriceType string

const (
	Last  TriggerPriceType = "last"
	Index TriggerPriceType = "index"
	Mark  TriggerPriceType = "mark"
)

type OrderCategory string

const (
	Normal             OrderCategory = "normal"
	Twap               OrderCategory = "twap"
	Adl                OrderCategory = "adl"
	FullLiquidation    OrderCategory = "full_liquidation"
	PartialLiquidation OrderCategory = "partial_liquidation"
	Delivery           OrderCategory = "delivery"
	Ddh                OrderCategory = "ddh" //Delta dynamic hedge
)

type MarginType string

const (
	Manual     MarginType = "manual"
	AutoBorrow MarginType = "auto_borrow"
	AutoRepay  MarginType = "auto_repay"
)
