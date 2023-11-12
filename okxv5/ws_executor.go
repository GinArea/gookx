package okxv5

type Executor[T any] struct {
	Args          SubscriptionArgs
	subscriptions *Subscriptions
}

func NewExecutor[T any](args SubscriptionArgs, subscriptions *Subscriptions) *Executor[T] {
	o := new(Executor[T])
	o.Args = args
	o.subscriptions = subscriptions
	return o
}

func (o *Executor[T]) Subscribe(onShot func(Topic[T])) {
	o.subscriptions.subscribe(o.Args, func(raw RawTopic) error {

		// out, _ := json.Marshal(raw)
		// fmt.Print(string(out) + "\n\n")

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
