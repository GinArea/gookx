package okxv5

import "github.com/msw-x/moon/ujson"

// Get instruments
// Retrieve a list of instruments with open contracts.
// https://www.okx.com/docs-v5/en/#public-data-rest-api-get-instruments

type GetInstruments struct {
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
	CtVal        string
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
