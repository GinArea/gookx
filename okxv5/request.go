package okxv5

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/msw-x/moon/ufmt"
	"github.com/msw-x/moon/uhttp"
)

func GetPub[R, T any](c *Client, path string, req any, transform func(R) (T, error)) Response[T] {
	return request(c, http.MethodGet, path, req, transform, false)
}

func Get[R, T any](c *Client, path string, req any, transform func(R) (T, error)) Response[T] {
	return request(c, http.MethodGet, path, req, transform, true)
}

func Post[R, T any](c *Client, path string, req any, transform func(R) (T, error)) Response[T] {
	return request(c, http.MethodPost, path, req, transform, true)
}

func request[R, T any](c *Client, method string, path string, request any, transform func(R) (T, error), sign bool) (r Response[T]) {
	var attempt int
	for {
		r = req(c, method, path, request, transform, sign)
		if r.StatusCode != http.StatusOK && c.onTransportError != nil {
			if c.onTransportError(r.Error, method, r.StatusCode, attempt) {
				attempt++
				continue
			}
		}
		break
	}
	return
}

func req[R, T any](c *Client, method string, path string, request any, transform func(R) (T, error), sign bool) (r Response[T]) {
	var perf *uhttp.Performer
	switch method {
	case http.MethodGet:
		perf = c.c.Get(path).Params(request)
	case http.MethodPost:
		perf = c.c.Post(path).Json(request)
	default:
		r.Error = fmt.Errorf("forbidden method: %s", method)
		return
	}
	if sign && c.s != nil {
		if perf.Request.Header == nil {
			perf.Request.Header = make(http.Header)
		}
		switch method {
		case http.MethodGet:
			c.s.HeaderGet(perf.Request.Header, perf.Request.Params, path)
		case http.MethodPost:
			c.s.HeaderPost(perf.Request.Header, perf.Request.Body, path)
		}
	}
	h := perf.Do()
	if h.Error == nil {
		r.StatusCode = h.StatusCode
		// if h.BodyExists() {
		// 	fmt.Println(string(h.Body))
		// }
		if h.BodyExists() &&
			r.StatusCode != http.StatusInternalServerError &&
			r.StatusCode != http.StatusBadGateway &&
			r.StatusCode != http.StatusServiceUnavailable &&
			r.StatusCode != http.StatusGatewayTimeout &&
			r.StatusCode != 520 &&
			r.StatusCode != 521 &&
			r.StatusCode != 522 &&
			r.StatusCode != 523 &&
			r.StatusCode != 524 &&
			r.StatusCode != 525 &&
			r.StatusCode != 526 &&
			r.StatusCode != 527 &&
			r.StatusCode != 530 {
			raw := new(response[R])
			// check that the server response is json
			r.Error = h.Json(raw) //json parsing
			if r.Ok() {
				r.Error = raw.Error() // returns error if response code not in (0,1)
				if r.Ok() {
					r.Data, r.Error = transform(raw.Data)
				}
			}
		} else {
			r.Error = errors.New(ufmt.Join(h.Status))
		}
		if sign {
			r.SetErrorIfNil(h.HeaderTo(&r.Limit))
		}
	} else {
		r.Error = h.Error
		r.NetError = true
	}
	return
}
