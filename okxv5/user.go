package okxv5

import "github.com/msw-x/moon/ujson"

// Get account configuration
// https://www.okx.com/docs-v5/en/#trading-account-rest-api-get-account-configuration

type GetAccountConf struct{}

type AccountConf struct {
	Uid            string
	MainUid        string
	AcctLv         ujson.Int64
	AcctStpMode    string
	PosMode        PositionMode
	AutoLoan       ujson.Bool
	GreeksType     GreekType
	Level          string
	LevelTmp       string
	CtIsoMode      ContractIsolatedMarginTradingSettings
	MgnIsoMode     MgnIsoMode
	SpotOffsetType ujson.Int64
	RoleType       ujson.Int64
	SpotRoleType   ujson.Int64
	OpAuth         ujson.Int64
	KycLv          ujson.Int64
	Label          string
	Ip             string
	Perm           string
}

func (o *Client) GetAccountConfiguration() Response[[]AccountConf] {
	return GetAccountConf{}.Do(o)
}

func (o GetAccountConf) Do(c *Client) Response[[]AccountConf] {
	return Get(c, "account/config", o, forward[[]AccountConf])
}
