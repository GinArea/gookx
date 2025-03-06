package okxv5

import (
	"bytes"
	"encoding/json"
	"time"

	"github.com/gorilla/websocket"
	"github.com/msw-x/moon/ulog"
	"github.com/msw-x/moon/uws"
	"golang.org/x/exp/slices"
)

type WsClient struct {
	c          *uws.Client
	onResponce func(WsResponse) error
	onTopic    func([]byte) error
}

func NewWsClient() *WsClient {
	o := new(WsClient)
	o.c = uws.NewClient(WebsocketUrl)
	return o
}

func (o *WsClient) Close() {
	o.c.Close()
}

func (o *WsClient) Log() *ulog.Log {
	return o.c.Log()
}

func (o *WsClient) Transport() *uws.Options {
	return &o.c.Options
}

func (o *WsClient) WithLog(log *ulog.Log) *WsClient {
	o.c.WithLog(log)
	return o
}

func (o *WsClient) WithPath(path string) *WsClient {
	o.c.WithPath(path)
	return o
}

func (o *WsClient) WithProxy(proxy string) *WsClient {
	o.c.WithProxy(proxy)
	return o
}

func (o *WsClient) WithLogRequest(enable bool) *WsClient {
	o.Transport().LogSent.Size = enable
	o.Transport().LogSent.Data = enable
	return o
}

func (o *WsClient) WithLogResponse(enable bool) *WsClient {
	o.Transport().LogRecv.Size = enable
	o.Transport().LogRecv.Data = enable
	return o
}

func (o *WsClient) WithOnDialDelay(f func() time.Duration) *WsClient {
	o.c.WithOnDialDelay(f)
	return o
}

func (o *WsClient) WithOnDialError(f func(error) bool) *WsClient {
	o.c.WithOnDialError(f)
	return o
}

func (o *WsClient) WithOnConnected(f func()) *WsClient {
	o.c.WithOnConnected(f)
	return o
}

func (o *WsClient) WithOnDisconnected(f func()) *WsClient {
	o.c.WithOnDisconnected(f)
	return o
}

func (o *WsClient) WithOnResponse(f func(WsResponse) error) *WsClient {
	o.onResponce = f
	return o
}

func (o *WsClient) WithOnTopic(f func([]byte) error) *WsClient {
	o.onTopic = f
	return o
}

func (o *WsClient) Run() {
	o.c.WithOnPing(o.ping)
	o.c.WithOnMessage(o.onMessage)
	o.c.Run()
}

func (o *WsClient) Running() bool {
	return o.c.Running()
}

func (o *WsClient) Connected() bool {
	return o.c.Connected()
}

func (o *WsClient) Send(r any) {
	o.c.SendJson(r)
}

func (o *WsClient) Subscribe(args SubscriptionArgs) {

	req := WsRequest{
		Operation: "subscribe",
		Args:      []SubscriptionArgs{args},
	}
	o.Send(req)
}

func (o *WsClient) Unsubscribe(args SubscriptionArgs) {
	o.Send(WsRequest{
		Operation: "unsubscribe",
		Args:      []SubscriptionArgs{args},
	})
}

func (o *WsClient) ping() {
	o.c.SendText("ping")
}

func (o *WsClient) onMessage(messageType int, data []byte) {
	log := o.c.Log()
	if messageType != websocket.TextMessage {
		log.Warning("invalid message type:", uws.MessageTypeString(messageType))
		return
	}

	if bytes.Equal(data, []byte("pong")) {
		return
	}

	var r WsResponse
	err := json.Unmarshal(data, &r)

	if err == nil {
		if r.isError() {
			out, e := json.Marshal(r)
			if e == nil {
				log.Error(string(out))
			}
			if r.Code == "50100" {
				// {"op":"","event":"error","code":"50100","msg":"API frozen, please contact customer service.","connId":"9dc875db","args":null}
				// TODO: recv: websocket: close 4001: Login failed.
				o.c.Cancel()
			}
		} else {
			if r.Valid() {
				if o.onResponce != nil {
					err = o.onResponce(r)
				}
			} else {
				skipTypes := []string{"subscribe", "unsubscribe"}
				if o.onTopic != nil && !slices.Contains(skipTypes, r.Event) {
					_ = o.onTopic(data)
				}
			}
		}
	}
	if err != nil {
		log.Error(err)
	}

}
