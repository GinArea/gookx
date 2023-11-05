package okxv5

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"io"
	"net/http"
	"net/url"
	"sort"
	"time"
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

func (o *Sign) HeaderPost(h http.Header, body []byte, path string) {
	o.header(h, string(body[:]), path, "POST")
}

func (o *Sign) HeaderGet(h http.Header, v url.Values, path string) {
	encodedParams := encodeSortParams(v)
	o.header(h, encodedParams, path, "GET")
}

func encodeSortParams(src url.Values) (s string) {
	if len(src) == 0 {
		return
	}
	keys := make([]string, len(src))
	i := 0
	for k := range src {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	for _, k := range keys {
		s += encodeParam(k, src.Get(k)) + "&"
	}
	s = s[0 : len(s)-1]
	return
}

func encodeParam(name, value string) string {
	params := url.Values{}
	params.Add(name, value)
	return params.Encode()
}

func (o *Sign) timestamp() string {

	// Get the current time in UTC
	currentTime := time.Now().UTC()
	// Format the time in ISO 8601 format
	isoTime := currentTime.Format("2006-01-02T15:04:05.999Z")
	return isoTime
}

func (o *Sign) header(h http.Header, s string, path string, method string) {

	//preSignedString: timestamp + method + requestPath + body

	ts := o.timestamp()
	preSignedString := ts + method + "/" + ApiVersion + "/" + path
	if s != "" {
		var delimeter string
		if method == "GET" {
			delimeter = "?"
		}
		preSignedString = preSignedString + delimeter + s
	}
	kcApiSign := signHmac(preSignedString, o.Secret)
	h.Set("OK-ACCESS-KEY", o.Key)
	h.Set("OK-ACCESS-PASSPHRASE", o.Password)
	h.Set("OK-ACCESS-TIMESTAMP", ts)
	h.Set("OK-ACCESS-SIGN", kcApiSign)
}

func signHmac(preSignedString, secret string) string {

	h := hmac.New(sha256.New, []byte(secret))
	io.WriteString(h, preSignedString)
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}
