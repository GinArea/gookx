package okxv5

import (
	"fmt"

	"github.com/msw-x/moon/ujson"
)

type Error struct {
	Code ujson.Int64
	Text string
}

func (o *Error) Std() error {
	if o.Empty() {
		return nil
	} else {
		return o
	}
}

func (o *Error) Empty() bool {
	// 1 added, as PlaceOrder response contains info inside
	return o.Code == 0 || o.Code == 1
}

func (o *Error) Error() string {
	return fmt.Sprintf("code[%d]: %s", o.Code.Value(), o.Text)
}
