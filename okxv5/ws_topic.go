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
	PTime ujson.Int64
}
