package okxv5

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
