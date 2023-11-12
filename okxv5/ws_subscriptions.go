package okxv5

import (
	"encoding/json"
	"fmt"
	"strings"
	"sync"
)

type Subscriptions struct {
	c     SubscriptionClient
	mutex sync.Mutex
	funcs SubscriptionFuncs
}

func NewSubscriptions(c SubscriptionClient) *Subscriptions {
	o := new(Subscriptions)
	o.c = c
	o.funcs = make(SubscriptionFuncs)
	return o
}

func (o *Subscriptions) subscribe(args SubscriptionArgs, f SubscriptionFunc) {
	if o.c.Ready() {
		o.c.subscribe(args)
	}
	o.mutex.Lock()
	defer o.mutex.Unlock()
	o.funcs[args] = f
}

func (o *Subscriptions) unsubscribe(args SubscriptionArgs) {
	if o.c.Ready() {
		o.c.unsubscribe(args)
	}
	o.mutex.Lock()
	defer o.mutex.Unlock()
	delete(o.funcs, args)
}

func (o *Subscriptions) subscribeAll() {
	o.mutex.Lock()
	defer o.mutex.Unlock()
	for topic, _ := range o.funcs {
		o.c.subscribe(topic)
	}
}

func (o *Subscriptions) processTopic(data []byte) (err error) {
	var topic RawTopic
	err = json.Unmarshal(data, &topic)
	if err == nil {
		f := o.getFunc(topic.Arg)
		if f == nil {
			err = fmt.Errorf("subscription of topic[%s] not found", topic.Arg.Channel+topic.Arg.InstId)
		} else {
			err = f(topic)
		}
	}
	return
}

func (o *Subscriptions) getFunc(passedArgs SubscriptionArgs) (f SubscriptionFunc) {
	o.mutex.Lock()
	defer o.mutex.Unlock()
	for args, fn := range o.funcs {
		if strings.EqualFold(args.Channel, passedArgs.Channel) &&
			strings.EqualFold(args.InstId, passedArgs.InstId) &&
			strings.EqualFold(args.InstType, passedArgs.InstType) {
			f = fn
			break
		}
	}
	return
}

type SubscriptionClient interface {
	Ready() bool
	subscribe(SubscriptionArgs)
	unsubscribe(SubscriptionArgs)
}

type SubscriptionFunc func(RawTopic) error

type SubscriptionFuncs map[SubscriptionArgs]SubscriptionFunc
