package okxv5

import "github.com/msw-x/moon/ujson"

// Get Balance
// https://www.okx.com/docs-v5/en/#trading-account-rest-api-get-balance

type GetBalance struct {
	Currency string `url:"ccy,omitempty"`
}

type BalanceOverview struct {
	AdjEq       ujson.Float64
	BorrowFroz  ujson.Float64
	Details     []BalanceDetails
	Imr         ujson.Float64
	IsoEq       ujson.Float64
	MgnRatio    ujson.Float64
	Mmr         ujson.Float64
	NotionalUsd ujson.Float64
	OrdFroz     ujson.Float64
	TotalEq     ujson.Float64
	UTime       ujson.TimeMs
}

type BalanceDetails struct {
	AvailBal      ujson.Float64
	AvailEq       ujson.Float64
	BorrowFroz    ujson.Float64
	CashBal       ujson.Float64
	Ccy           string
	CrossLiab     ujson.Float64
	DisEq         ujson.Float64
	Eq            ujson.Float64
	EqUsd         ujson.Float64
	FixedBal      ujson.Float64
	FrozenBal     ujson.Float64
	Interest      ujson.Float64
	IsoEq         ujson.Float64
	IsoLiab       ujson.Float64
	IsoUpl        ujson.Float64
	Liab          ujson.Float64
	MaxLoan       ujson.Float64
	MgnRatio      ujson.Float64
	NotionalLever ujson.Float64
	OrdFrozen     ujson.Float64
	SpotInUseAmt  ujson.Float64
	SpotIsoBal    ujson.Float64
	StgyEq        ujson.Float64
	Twap          ujson.Float64
	Upl           ujson.Float64
	UplLiab       ujson.Float64
	UTime         ujson.TimeMs
}

func (o *Client) GetBalance(v GetBalance) Response[[]BalanceOverview] {
	return v.Do(o)
}

func (o GetBalance) Do(c *Client) Response[[]BalanceOverview] {
	return Get(c, "account/balance", o, forward[[]BalanceOverview])
}
