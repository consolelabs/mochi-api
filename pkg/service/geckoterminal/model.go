package geckoterminal

type Search struct {
	Data SearchData `json:"data,omitempty"`
}

type SearchData struct {
	ID         string           `json:"id,omitempty"`
	Type       string           `json:"type,omitempty"`
	Attributes SearchAttributes `json:"attributes,omitempty"`
}

type SearchAttributes struct {
	Networks []interface{}       `json:"networks,omitempty"`
	Dexes    []interface{}       `json:"dexes,omitempty"`
	Pools    []SearchPoolElement `json:"pools,omitempty"`
	Pairs    []interface{}       `json:"pairs,omitempty"`
}

type SearchPoolElement struct {
	Type               SearchType    `json:"type,omitempty"`
	Address            string        `json:"address,omitempty"`
	APIAddress         string        `json:"api_address,omitempty"`
	PriceInUsd         string        `json:"price_in_usd,omitempty"`
	ReserveInUsd       string        `json:"reserve_in_usd,omitempty"`
	FromVolumeInUsd    string        `json:"from_volume_in_usd,omitempty"`
	ToVolumeInUsd      string        `json:"to_volume_in_usd,omitempty"`
	PricePercentChange string        `json:"price_percent_change,omitempty"`
	Network            *SearchDex    `json:"network,omitempty"`
	Dex                *SearchDex    `json:"dex,omitempty"`
	Tokens             []SearchToken `json:"tokens,omitempty"`
}

type SearchDex struct {
	Name       *string `json:"name,omitempty"`
	Identifier *string `json:"identifier,omitempty"`
	ImageURL   *string `json:"image_url,omitempty"`
}

type SearchToken struct {
	Name        *string `json:"name,omitempty"`
	Symbol      *string `json:"symbol,omitempty"`
	ImageURL    *string `json:"image_url,omitempty"`
	IsBaseToken *bool   `json:"is_base_token,omitempty"`
}

type SearchType string

const (
	SearchTypePool SearchType = "pool"
)

// Pool
type Pool struct {
	Data PoolData `json:"data"`
}

type PoolData struct {
	ID            string        `json:"id"`
	Type          string        `json:"type"`
	Attributes    Attributes    `json:"attributes"`
	Relationships Relationships `json:"relationships"`
}

type Attributes struct {
	BaseTokenPriceUsd             string                `json:"base_token_price_usd"`
	BaseTokenPriceNativeCurrency  string                `json:"base_token_price_native_currency"`
	QuoteTokenPriceUsd            string                `json:"quote_token_price_usd"`
	QuoteTokenPriceNativeCurrency string                `json:"quote_token_price_native_currency"`
	Address                       string                `json:"address"`
	Name                          string                `json:"name"`
	ReserveInUsd                  string                `json:"reserve_in_usd"`
	PoolCreatedAt                 string                `json:"pool_created_at"`
	FdvUsd                        string                `json:"fdv_usd"`
	MarketCapUsd                  *string               `json:"market_cap_usd"`
	PriceChangePercentage         PriceChangePercentage `json:"price_change_percentage"`
	Transactions                  Transactions          `json:"transactions"`
	VolumeUsd                     VolumeUsd             `json:"volume_usd"`
}

type PriceChangePercentage struct {
	H1  string `json:"h1"`
	H24 string `json:"h24"`
}

type Transactions struct {
	H1  H1 `json:"h1"`
	H24 H1 `json:"h24"`
}

type H1 struct {
	Buys  int64 `json:"buys"`
	Sells int64 `json:"sells"`
}

type VolumeUsd struct {
	H24 string `json:"h24"`
}

type Relationships struct {
	BaseToken  BaseToken `json:"base_token"`
	QuoteToken BaseToken `json:"quote_token"`
	Dex        BaseToken `json:"dex"`
}

type BaseToken struct {
	Data BaseTokenData `json:"data"`
}

type BaseTokenData struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}
