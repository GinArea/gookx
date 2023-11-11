package okxv5

type WsRequest struct {
	Operation string             `json:"op"`
	Args      []SubscriptionArgs `json:"args"`
}

type SubscriptionArgs struct {
	Channel string `json:"channel"`
	InstId  string `json:"InstId"`
}
