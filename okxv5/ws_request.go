package okxv5

type WsRequest struct {
	Operation string             `json:"op"`
	Args      []SubscriptionArgs `json:"args"`
}

type SubscriptionArgs struct {
	Channel  string `json:"channel"`
	InstId   string `json:"instId,omitempty"`
	InstType string `json:"instType,omitempty"`
}

type WsRequestAuth struct {
	Operation string                 `json:"op"`
	Args      []SubscriptionArgsAuth `json:"args"`
}

type SubscriptionArgsAuth struct {
	ApiKey    string `json:"apiKey"`
	Passprase string `json:"passphrase"`
	Timestamp string `json:"timestamp"`
	Sign      string `json:"sign"`
}
