package okxv5

import (
	"strings"
	"time"

	"github.com/msw-x/moon/ufmt"
	"github.com/msw-x/moon/ulog"
	"github.com/msw-x/moon/uws"
)

type WsPrivate struct {
	c              *WsClient
	s              *Sign
	ready          bool
	onReady        func()
	onConnected    func()
	onDisconnected func()
	subscriptions  *Subscriptions
}

func NewWsPrivate(key, secret, password string) *WsPrivate {
	o := new(WsPrivate)
	o.c = NewWsClient()
	o.s = NewSign(key, secret, password)
	o.subscriptions = NewSubscriptions(o)
	o.c.WithPath("v5/private")
	return o
}

func (o *WsPrivate) Close() {
	o.c.Close()
}

func (o *WsPrivate) Transport() *uws.Options {
	return o.c.Transport()
}

func (o *WsPrivate) WithLog(log *ulog.Log) *WsPrivate {
	o.c.WithLog(log)
	return o
}

func (o *WsPrivate) WithProxy(proxy string) *WsPrivate {
	o.c.WithProxy(proxy)
	return o
}

func (o *WsPrivate) WithLogRequest(enable bool) *WsPrivate {
	o.c.WithLogRequest(enable)
	return o
}

func (o *WsPrivate) WithLogResponse(enable bool) *WsPrivate {
	o.c.WithLogResponse(enable)
	return o
}

func (o *WsPrivate) WithOnDialDelay(f func() time.Duration) *WsPrivate {
	o.c.WithOnDialDelay(f)
	return o
}

func (o *WsPrivate) WithOnDialError(f func(error) bool) *WsPrivate {
	o.c.WithOnDialError(f)
	return o
}

func (o *WsPrivate) WithOnReady(f func()) *WsPrivate {
	o.onReady = f
	return o
}

func (o *WsPrivate) WithOnConnected(f func()) *WsPrivate {
	o.onConnected = f
	return o
}

func (o *WsPrivate) WithOnDisconnected(f func()) *WsPrivate {
	o.onDisconnected = f
	return o
}

func (o *WsPrivate) Run() {
	o.c.WithOnConnected(func() {
		if o.onConnected != nil {
			o.onConnected()
		}
		o.auth()
	})
	o.c.WithOnDisconnected(func() {
		o.ready = false
		if o.onDisconnected != nil {
			o.onDisconnected()
		}
	})
	o.c.WithOnResponse(o.onResponse)
	o.c.WithOnTopic(o.onTopic)
	o.c.Run()
}

func (o *WsPrivate) Running() bool {
	return o.c.Running()
}

func (o *WsPrivate) Connected() bool {
	return o.c.Connected()
}

func (o *WsPrivate) Ready() bool {
	return o.ready
}

func (o *WsPrivate) auth() {
	o.c.Send(WsRequestAuth{
		Operation: "login",
		Args:      o.s.WebSocket(),
	})
}

func (o *WsPrivate) subscribe(topicArgs SubscriptionArgs) {
	o.c.Subscribe(topicArgs)
}

func (o *WsPrivate) unsubscribe(topicArgs SubscriptionArgs) {
	o.c.Unsubscribe(topicArgs)
}

func (o *WsPrivate) onResponse(r WsResponse) error {
	log := o.c.Log()
	if r.Event == "login" {
		success := strings.EqualFold(r.Code, "0")
		log.Info("auth:", ufmt.SuccessFailure(success))
		if success {
			o.ready = true
			if o.onReady != nil {
				o.onReady()
			}
			o.subscriptions.subscribeAll()
		}
	} else {
		r.Log(log)
	}
	return nil
}

func (o *WsPrivate) onTopic(data []byte) error {
	return o.subscriptions.processTopic(data)
}

func (o *WsPrivate) Position(c Category) *Executor[[]Positions] {
	args := SubscriptionArgs{
		Channel:     "positions",
		InstType:    string(c),
		ExtraParams: "{\"updateInterval\": \"0\"}",
	}
	return NewExecutor[[]Positions](args, o.subscriptions)
}

func (o *WsPrivate) Orders(c Category) *Executor[[]RetrievedOrderDetails] {
	args := SubscriptionArgs{
		Channel:  "orders",
		InstType: string(c),
	}
	return NewExecutor[[]RetrievedOrderDetails](args, o.subscriptions)
}

// https://www.okx.com/docs-v5/en/#trading-account-websocket-balance-and-position-channel
func (o *WsPrivate) Wallet() *Executor[[]WalletShot] {
	args := SubscriptionArgs{
		Channel: "balance_and_position",
	}
	return NewExecutor[[]WalletShot](args, o.subscriptions)
}
