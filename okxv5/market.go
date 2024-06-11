package okxv5

import (
	"github.com/msw-x/moon/ujson"
)

// Get instruments
// Retrieve a list of instruments with open contracts.
// https://www.okx.com/docs-v5/en/#public-data-rest-api-get-instruments

type InstrumentsQuery struct {
	InstType   Category
	Uly        string `url:",omitempty"`
	InstFamily string `url:",omitempty"`
	InstId     string `url:",omitempty"`
}

type Instruments struct {
	Alias        string
	BaseCcy      string
	Category     ujson.Int64
	CtMult       string
	CtType       string
	CtVal        ujson.Float64
	CtValCcy     string
	ExpTime      string
	InstFamily   string
	InstId       string
	InstType     string
	Lever        ujson.Float64
	ListTime     ujson.Int64
	LotSz        ujson.Float64
	MaxIcebergSz ujson.Float64
	MaxLmtSz     ujson.Float64
	MaxMktSz     ujson.Float64
	MaxStopSz    ujson.Float64
	MaxTriggerSz ujson.Float64
	MaxTwapSz    ujson.Float64
	MinSz        ujson.Float64
	OptType      string
	QuoteCcy     string
	SettleCcy    string
	State        string
	Stk          string
	TickSz       ujson.Float64
	Uly          string
}

func (o *Client) GetInstruments(i InstrumentsQuery) Response[[]Instruments] {
	return i.Do(o)
}

func (o InstrumentsQuery) Do(c *Client) Response[[]Instruments] {
	return GetPub(c.public(), "instruments", o, forward[[]Instruments])
}

// GET / Tickers
// Retrieve the latest price snapshot, best bid/ask price, and trading volume in the last 24 hours.
// https://www.okx.com/docs-v5/en/#order-book-trading-market-data-get-tickers

type MarketsQuery struct {
	InstType   Category //exclude Margin
	Uly        string   `url:",omitempty"`
	InstFamily string   `url:",omitempty"`
}

// GET / Ticker
// Retrieve the latest price snapshot, best bid/ask price, and trading volume in the last 24 hours.
// https://www.okx.com/docs-v5/en/#order-book-trading-market-data-get-ticker

type MarketQuery struct {
	InstId string
}

type Market struct {
	InstType  string
	InstId    string
	Last      ujson.Float64
	LastSz    ujson.Float64
	AskPx     ujson.Float64
	AskSz     ujson.Float64
	BidPx     ujson.Float64
	BidSz     ujson.Float64
	Open24h   ujson.Float64
	High24h   ujson.Float64
	Low24h    ujson.Float64
	VolCcy24h ujson.Float64
	Vol24h    ujson.Float64
	Ts        ujson.Int64
	SodUtc0   ujson.Float64
	SodUtc8   ujson.Float64
}

func (o *Client) GetMarkets(mq MarketsQuery) Response[[]Market] {
	return mq.Do(o)
}

func (mq MarketsQuery) Do(c *Client) Response[[]Market] {
	return GetPub(c.market(), "tickers", mq, forward[[]Market])
}

func (o *Client) GetMarket(mq MarketQuery) Response[[]Market] {
	return mq.Do(o)
}

func (mq MarketQuery) Do(c *Client) Response[[]Market] {
	return GetPub(c.market(), "ticker", mq, forward[[]Market])
}

// GET / Order book
// Retrieve order book of the instrument.
// https://www.okx.com/docs-v5/en/#order-book-trading-market-data-get-order-book
type GetOrderbook struct {
	InstId string
	Size   int `url:"sz,omitempty"`
}

type Orderbook struct {
	Asks  [][]ujson.Float64 `json:"asks"`
	Bids  [][]ujson.Float64 `json:"bids"`
	Ts    ujson.Int64       `json:"ts"`
	SeqId ujson.Int64       `json:"seqId"`
}

func (o GetOrderbook) Do(c *Client) Response[[]Orderbook] {
	return GetPub(c.market(), "books", o, forward[[]Orderbook])
}

func (o *Client) GetOrderbook(v GetOrderbook) Response[[]Orderbook] {
	return v.Do(o)
}

// GET / Trades history
// Retrieve the recent transactions of an instrument from the last 3 months with pagination.
// https://www.okx.com/docs-v5/en/#order-book-trading-market-data-get-trades-history
type GetTradesHistory struct {
	InstId string
	Type   int    `url:",omitempty"`
	After  string `url:",omitempty"`
	Before string `url:",omitempty"`
	Limit  int    `url:",omitempty"`
}

type TradesHistory struct {
	InstId  string
	Side    Side
	Size    ujson.Float64 `json:"sz"`
	Price   ujson.Float64 `json:"px"`
	TradeId string
	Ts      ujson.TimeMs
}

func (o GetTradesHistory) Do(c *Client) Response[[]TradesHistory] {
	return GetPub(c.market(), "history-trades", o, forward[[]TradesHistory])
}

func (o *Client) GetTradesHistory(v GetTradesHistory) Response[[]TradesHistory] {
	return v.Do(o)
}
