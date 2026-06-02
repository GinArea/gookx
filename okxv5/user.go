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

// Set position mode
// posMode: "long_short_mode" (long/short, FUTURES/SWAP only) or "net_mode" (net)
// https://www.okx.com/docs-v5/en/#trading-account-rest-api-set-position-mode

type SetPositionMode struct {
	// PosMode - Position mode: "long_short_mode" or "net_mode"
	PosMode PositionMode
}

type PositionModeResult struct {
	// PosMode - The position mode now in effect for the account
	PosMode PositionMode
}

func (o *Client) SetPositionMode(posMode PositionMode) Response[[]PositionModeResult] {
	return SetPositionMode{PosMode: posMode}.Do(o)
}

func (o SetPositionMode) Do(c *Client) Response[[]PositionModeResult] {
	return Post(c, "account/set-position-mode", o, forward[[]PositionModeResult])
}

// Set leverage
// instId/ccy: one is required. mgnMode: "isolated" or "cross".
// posSide ("long"/"short") only required in long/short mode under isolated margin.
// https://www.okx.com/docs-v5/en/#trading-account-rest-api-set-leverage

type SetLeverage struct {
	// InstId - Instrument ID (e.g. "BTC-USDT-SWAP"); either InstId or Ccy is required
	InstId string `json:",omitempty"`
	// Ccy - Margin currency, only for cross-margin MARGIN at currency level
	Ccy string `json:",omitempty"`
	// Lever - Leverage value
	Lever int
	// MgnMode - Margin mode: "isolated" or "cross"
	MgnMode MarginMode
	// PosSide - Position side "long"/"short"; required only in long/short isolated margin mode
	PosSide PosSide `json:",omitempty"`
}

type LeverageResult struct {
	// Lever - Leverage now in effect
	Lever ujson.Float64
	// MgnMode - Margin mode: "isolated" or "cross"
	MgnMode MarginMode
	// InstId - Instrument ID
	InstId string
	// PosSide - Position side
	PosSide PosSide
}

func (o *Client) SetLeverage(v SetLeverage) Response[[]LeverageResult] {
	return v.Do(o)
}

func (o SetLeverage) Do(c *Client) Response[[]LeverageResult] {
	return Post(c, "account/set-leverage", o, forward[[]LeverageResult])
}
