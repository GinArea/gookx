package okxv5

import "github.com/msw-x/moon/ujson"

// Get account configuration
// https://www.okx.com/docs-v5/en/#trading-account-rest-api-get-account-configuration

type GetAccountConf struct{}

type AccountConf struct {
	AcctLv         ujson.Int64
	AutoLoan       ujson.Bool
	CtIsoMode      ContractIsolatedMarginTradingSettings
	GreeksType     GreekType
	Level          string
	LevelTmp       string
	MgnIsoMode     MgnIsoMode
	PosMode        PositionMode
	SpotOffsetType ujson.Int64
	Uid            string
	Label          string
	RoleType       ujson.Int64
	SpotRoleType   ujson.Int64
	OpAuth         ujson.Int64
	KycLv          ujson.Int64
	Ip             string
	Perm           string
	MainUid        string
}

func (o *Client) GetAccountConfiguration() Response[[]AccountConf] {
	return GetAccountConf{}.Do(o)
}

func (o GetAccountConf) Do(c *Client) Response[[]AccountConf] {
	return Get(c, "account/config", o, forward[[]AccountConf])
}
