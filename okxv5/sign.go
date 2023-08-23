package okxv5

import (
	"net/http"
	"net/url"
)

const (
	tag = "557f6da2e9c0BCDE"
)

type Sign struct {
	Key      string
	Secret   string
	Password string
}

func NewSign(key, secret, password string) *Sign {
	o := new(Sign)
	o.Key = key
	o.Secret = secret
	o.Password = password
	return o
}

func (o *Sign) HeaderGet(h http.Header, v url.Values) {
	//TODO
	//	o.header(h, )
}

func (o *Sign) header(h http.Header, s string) {

	//TODO
}

func (o *Sign) HeaderPost(h http.Header, body []byte) {
	//TODO

	//o.header(h, )
}
