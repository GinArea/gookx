package okxv5

import (
	"github.com/msw-x/moon/ulog"
)

type WsResponse struct {
	Operation    string        `json:"op"`
	Event        string        `json:"event"`
	Code         string        `json:"code"`
	Message      string        `json:"msg"`
	ConnectionId string        `json:"connId"`
	Args         []interface{} `json:"args"`
}

func (o WsResponse) Log(log *ulog.Log) {
	switch o.Operation {
	case "ping":
	case "pong":
	case "subscribe":
		if len(o.Args) > 0 {
			log.Info("subscribe: success")
		} else {
			log.Error("subscribe:", o.Message)
		}
	case "unsubscribe":
		log.Info("unsubscribe: success")
	default:
		log.Error("invalid response:", o.Operation)
	}
}
