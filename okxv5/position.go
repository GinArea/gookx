package okxv5

import (
	"github.com/msw-x/moon/ujson"
)

// Get positions
// https://www.okx.com/docs-v5/en/#trading-account-rest-api-get-positions
type PositionsQuery struct {
	InstType Category `url:",omitempty"`
	InstId   string   `url:",omitempty"`
	PosId    string   `url:",omitempty"`
}

type Positions struct {
	InstType       Category
	MgnMode        MarginMode
	PosId          ujson.Int64
	PosSide        string
	Pos            ujson.Float64
	BaseBorrowed   string
	BaseInterest   string
	QuoteBorrowed  string
	QuoteInterest  string
	PosCcy         string
	AvailPos       ujson.Float64
	AvgPx          ujson.Float64
	MarkPx         ujson.Float64
	Upl            ujson.Float64
	UplRatio       ujson.Float64
	UplLastPx      ujson.Float64
	UplRatioLastPx ujson.Float64
	InstId         string
	Lever          ujson.Float64
	LiqPx          ujson.Float64
	Imr            string
	Margin         ujson.Float64
	MgnRatio       ujson.Float64
	Mmr            ujson.Float64
	Liab           string
	LiabCcy        string
	Interest       ujson.Float64
	TradeId        ujson.Int64
	OptVal         string
	NotionalUsd    ujson.Float64
	Adl            ujson.Int64
	Ccy            string
	Last           ujson.Float64
	IdxPx          ujson.Float64
	UsdPx          ujson.Float64
	RealizedPnl    ujson.Float64
	Pnl            ujson.Float64
	Fee            ujson.Float64
	FundingFee     ujson.Float64
	LiqPenalty     ujson.Float64
}

func (o *Client) GetPositions(pq PositionsQuery) Response[[]Positions] {
	return pq.Do(o)
}

func (o PositionsQuery) Do(c *Client) Response[[]Positions] {
	return Get(c, "account/positions", o, forward[[]Positions])
}
