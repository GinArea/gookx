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

type PosSide string

const (
	LongSide  PosSide = "long"
	ShortSide PosSide = "short"
	NetSide   PosSide = "net"
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

// https://www.okx.com/docs-v5/en/#order-book-trading-market-data-ws-order-book-channel

type OrderbookType string

const (
	Books  OrderbookType = "books"
	Books5 OrderbookType = "books5"
	Bbotbt OrderbookType = "bbo-tbt"
	//Log in required:
	//Booksl2tbt   OrderbookType = "books-l2-tbt"
	//Books50l2tbt OrderbookType = "books50-l2-tbt"
)

type ContractIsolatedMarginTradingSettings string

const (
	Auto    ContractIsolatedMarginTradingSettings = "automatic"
	Manualy ContractIsolatedMarginTradingSettings = "autonomy"
)

type GreekType string

const (
	Pa GreekType = "PA"
	Bs GreekType = "BS"
)

type MgnIsoMode string

const (
	Automatic   MgnIsoMode = "automatic"
	QuickMargin MgnIsoMode = "quick_margin"
)

type PositionMode string

const (
	LongShortMode PositionMode = "long_short_mode"
	NetMode       PositionMode = "net_mode"
)

// bar size (https://www.okx.com/docs-v5/en/#public-data-rest-api-get-mark-price-candlesticks-history)
// 1m/3m/5m/15m/30m/1H/2H/4H
// Hong Kong time opening price k-line: 6H/12H/1D/1W/1M
// UTC time opening price k-line: 6Hutc/12Hutc/1Dutc/1Wutc/1Mutc
type Bar string

const (
	Bar1m     Bar = "1m"
	Bar3m     Bar = "3m"
	Bar5m     Bar = "5m"
	Bar15m    Bar = "15m"
	Bar30m    Bar = "30m"
	Bar1H     Bar = "1H"
	Bar2H     Bar = "2H"
	Bar4H     Bar = "4H"
	Bar6H     Bar = "6H"
	Bar12H    Bar = "12H"
	Bar1D     Bar = "1D"
	Bar2D     Bar = "2D"
	Bar3D     Bar = "3D"
	Bar5D     Bar = "5D"
	Bar1W     Bar = "1W"
	Bar1M     Bar = "1M"
	Bar3M     Bar = "3M"
	Bar6Hutc  Bar = "6Hutc"
	Bar12Hutc Bar = "12Hutc"
	Bar1Dutc  Bar = "1Dutc"
	Bar2Dutc  Bar = "2Dutc"
	Bar3Dutc  Bar = "3Dutc"
	Bar5Dutc  Bar = "5Dutc"
	Bar1wutc  Bar = "1Wutc"
	Bar1Mutc  Bar = "1Mutc"
	Bar3Mutc  Bar = "3Mutc"
)

// mark-price-candle (https://www.okx.com/docs-v5/en/#public-data-websocket-mark-price-candlesticks-channel)

type MarkPriceCandle string

func (o Bar) MarkPriceCandle() MarkPriceCandle {
	return MarkPriceCandle("mark-price-candle" + o)
}
