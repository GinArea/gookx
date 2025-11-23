package okxv5

import "github.com/msw-x/moon/ujson"

// Place order
// https://www.okx.com/docs-v5/en/#order-book-trading-trade-post-place-order
type PlaceOrder struct {
	InstId     string
	TdMode     TradeMode
	Ccy        string `json:",omitempty"`
	Side       Side
	OrdType    OrderType
	Sz         float64
	PosSide    *PosSide `json:",omitempty"`
	ClOrdId    string   `json:",omitempty"`
	ReduceOnly *bool    `json:",omitempty"`
	Tag        string   `json:",omitempty"`
}

type OrderDetail struct {
	ClOrdId string
	OrdId   string
	Tag     string
	SCode   ujson.Int64
	SMsg    string
}

func (o *Client) PlaceOrder(v PlaceOrder) Response[[]OrderDetail] {
	return v.Do(o)
}

func (o PlaceOrder) Do(c *Client) Response[[]OrderDetail] {

	o.Tag = GinAreaTag
	return Post(c, "trade/order", o, forward[[]OrderDetail])
}

// Get order details
// https://www.okx.com/docs-v5/en/#order-book-trading-trade-get-order-details

type GetOrderDetails struct {
	InstId  string
	OrdId   string `json:",omitempty"`
	ClOrdId string `json:",omitempty"`
}

type RetrievedOrderDetails struct {
	InstType           Category
	InstId             string
	TgtCcy             string
	Ccy                string
	OrdId              string
	ClOrdId            string
	Tag                string
	Px                 ujson.Float64
	PxUsd              ujson.Float64
	PxVol              ujson.Float64
	PxType             string
	Sz                 ujson.Float64
	Pnl                ujson.Float64
	OrdType            OrderType
	Side               Side
	PosSide            PosSide
	TdMode             TradeMode
	AccFillSz          ujson.Float64
	FillPx             ujson.Float64
	TradeId            string
	FillSz             ujson.Float64
	FillTime           string
	AvgPx              ujson.Float64
	State              OrderState
	StpId              string
	StpMode            string
	Lever              ujson.Float64
	AttachAlgoClOrdId  string
	LastPx             ujson.Float64
	TpTriggerPx        ujson.Float64
	TpTriggerPxType    TriggerPriceType
	TpOrdPx            ujson.Float64
	SlTriggerPx        ujson.Float64
	SlTriggerPxType    TriggerPriceType
	SlOrdPx            ujson.Float64
	FeeCcy             string
	Fee                ujson.Float64
	RebateCcy          string
	Source             string
	Rebate             string
	Category           OrderCategory
	ReduceOnly         ujson.Bool
	CancelSource       string
	CancelSourceReason string
	QuickMgnType       MarginType
	AlgoClOrdId        string
	AlgoId             string
	UTime              ujson.TimeMs
	CTime              ujson.TimeMs
}

func (o *Client) GetOrderDetails(v GetOrderDetails) Response[[]RetrievedOrderDetails] {
	return v.Do(o)
}

func (o GetOrderDetails) Do(c *Client) Response[[]RetrievedOrderDetails] {
	return Get(c, "trade/order", o, forward[[]RetrievedOrderDetails])
}
