package okxv5

type Executor[T any] struct {
	Args          SubscriptionArgs
	subscriptions *Subscriptions
}

func NewExecutor[T any](instId string, channel string, subscriptions *Subscriptions) *Executor[T] {
	o := new(Executor[T])
	o.Args = SubscriptionArgs{
		Channel: channel,
		InstId:  instId,
	}
	o.subscriptions = subscriptions
	return o
}

func (o *Executor[T]) Subscribe(onShot func(Topic[T])) {
	o.subscriptions.subscribe(o.Args, func(raw RawTopic) error {
		topic, err := UnmarshalRawTopic[T](raw)
		if err == nil {
			onShot(topic)
		}
		return err
	})
}

func (o *Executor[T]) Unsubscribe() {
	o.subscriptions.unsubscribe(o.Args)
}
