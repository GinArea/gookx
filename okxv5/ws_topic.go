package okxv5

import (
	"encoding/json"

	"github.com/msw-x/moon/ujson"
)

type Topic[T any] struct {
	Arg    SubscriptionArgs
	Action string
	Data   T
}

func UnmarshalRawTopic[T any](raw RawTopic) (ret Topic[T], err error) {
	ret.Action = raw.Action
	ret.Arg = raw.Arg
	err = json.Unmarshal(raw.Data, &ret.Data)
	return
}

type RawTopic Topic[json.RawMessage]

type WalletShot struct {
	PTime     ujson.Int64
	EventType string
	BalData   []BalanceData
	PosData   []PositionData
	Trades    []Trades
}

type BalanceData struct {
	CashBal ujson.Float64
	Ccy     string
	UTime   ujson.Int64
}

type PositionData struct {
	AvgPx    ujson.Float64
	BaseDeal string
	Ccy      string
	InstId   string
	InstType string
	MgnMode  string
	Pos      ujson.Float64
	PosCcy   string
	PosId    ujson.Int64
	PosSide  string
	QuoteBal ujson.Float64
	TradeId  ujson.Int64
	UTime    ujson.Int64
}

type Trades struct {
	InstId  string
	TradeId ujson.Int64
}

// struct for websocket orderbook
// https://www.okx.com/docs-v5/en/#order-book-trading-market-data-ws-order-book-channel
type WsOrderbook struct {
	Asks  [][]ujson.Float64 `json:"asks"`
	Bids  [][]ujson.Float64 `json:"bids"`
	Ts    ujson.Int64       `json:"ts"`
	SeqId ujson.Int64       `json:"seqId"`
}

/*
{
   "Arg":{
      "channel":"balance_and_position"
   },
   "Action":"",
   "Data":[
      {
         "balData":[
            {
               "cashBal":"44.6662370515664886",
               "ccy":"USDT",
               "uTime":"1699789825443"
            }
         ],
         "eventType":"filled",
         "pTime":"1699789825442",
         "posData":[
            {
               "avgPx":"88.42",
               "baseBal":"",
               "ccy":"USDT",
               "instId":"TRB-USDT-SWAP",
               "instType":"SWAP",
               "mgnMode":"cross",
               "pos":"1",
               "posCcy":"",
               "posId":"624362342386012199",
               "posSide":"net",
               "quoteBal":"",
               "tradeId":"143360323",
               "uTime":"1699789825443"
            }
         ],
         "trades":[
            {
               "instId":"TRB-USDT-SWAP",
               "tradeId":"143360323"
            }
         ]
      }
   ]
}

*/
