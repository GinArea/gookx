package okxv5

import (
	"time"

	"github.com/msw-x/moon/uhttp"
)

type Client struct {
	c          *uhttp.Client
	s          *Sign
	onNetError OnNetError
}

func NewClient() *Client {
	o := new(Client)
	o.c = uhttp.NewClient()
	o.WithBaseUrl(MainBaseUrl)
	o.WithPath(ApiVersion)
	return o
}

func (o *Client) WithTimeout(timeout time.Duration) *Client {
	o.c.WithTimeout(timeout)
	return o
}

func (o *Client) WithProxy(proxy string) *Client {
	o.c.WithProxy(proxy)
	return o
}

func (o *Client) Copy() *Client {
	r := new(Client)
	r.c = o.c.Copy()
	r.s = o.s
	r.onNetError = o.onNetError
	return r
}

func (o *Client) WithBaseUrl(url string) *Client {
	o.c.WithBase(url)
	return o
}

func (o *Client) WithPath(path string) *Client {
	o.c.WithPath(path)
	return o
}

func (o *Client) WithAppendPath(path string) *Client {
	o.c.WithAppendPath(path)
	return o
}

func (o *Client) WithAuth(key, secret, password string) *Client {
	o.s = NewSign(key, secret, password)
	return o
}

func (o *Client) WithOnNetError(f OnNetError) *Client {
	o.onNetError = f
	return o
}

func (o *Client) public() *Client {
	return o.Copy().WithAppendPath("public")
}

func (o *Client) market() *Client {
	return o.Copy().WithAppendPath("market")
}

type OnNetError func(err error, statusCode int, attempt int) bool
