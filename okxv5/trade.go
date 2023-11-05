package okxv5

import "github.com/msw-x/moon/ujson"

// Place order
// https://www.okx.com/docs-v5/en/#order-book-trading-trade-post-place-order
type PlaceOrder struct {
	InstId     string
	TdMode     TradeMode
	Side       Side
	OrdType    OrderType
	Sz         ujson.Float64
	ClOrdId    string `json:",omitempty"`
	ReduceOnly bool   `json:",omitempty"`
}

type OrderDetail struct {
	ClOrdId string
	OrdId   string
	Tag     string
	SCode   string
	SMsg    string
}

func (o *Client) PlaceOrder(v PlaceOrder) Response[[]OrderDetail] {
	return v.Do(o)
}

func (o PlaceOrder) Do(c *Client) Response[[]OrderDetail] {

	return Post(c, "trade/order", o, forward[[]OrderDetail])
}
