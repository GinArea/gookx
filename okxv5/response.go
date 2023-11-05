package okxv5

import "github.com/msw-x/moon/ujson"

type Response[T any] struct {
	Data       T
	Limit      RateLimit
	Error      error
	StatusCode int
	NetError   bool
}

type response[T any] struct {
	Code ujson.Int64
	Data T
	// InTime + OutTime are contained in PlaceOrder, but not in GetPositions
	// InTime  ujson.Int64
	// OutTime ujson.Int64
	Msg string
}

func (o *Response[T]) Ok() bool {
	return o.Error == nil
}

func (o *Response[T]) SetErrorIfNil(err error) {
	if o.Error == nil {
		o.Error = err
	}
}

func (o *response[T]) Error() error {
	e := Error{
		Code: o.Code,
		Text: o.Msg,
	}
	return e.Std()
}
