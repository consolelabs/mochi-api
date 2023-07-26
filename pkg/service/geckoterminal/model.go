package geckoterminal

type Search struct {
	Data SearchData `json:"data"`
}

type SearchData struct {
	ID         string           `json:"id"`
	Type       string           `json:"type"`
	Attributes SearchAttributes `json:"attributes"`
}

type SearchAttributes struct {
	// Networks []interface{} `json:"networks"`
	// Dexes    []interface{} `json:"dexes"`
	Pools []SearchPoolElement `json:"pools"`
	// Pairs    []interface{} `json:"pairs"`
}

type SearchPoolElement struct {
	Type               SearchType    `json:"type"`
	Address            string        `json:"address"`
	APIAddress         string        `json:"api_address"`
	PriceInUsd         string        `json:"price_in_usd"`
	ReserveInUsd       string        `json:"reserve_in_usd"`
	FromVolumeInUsd    string        `json:"from_volume_in_usd"`
	ToVolumeInUsd      string        `json:"to_volume_in_usd"`
	PricePercentChange string        `json:"price_percent_change"`
	Network            SearchDex     `json:"network"`
	Dex                SearchDex     `json:"dex"`
	Tokens             []SearchToken `json:"tokens"`
}

type SearchDex struct {
	Name       string `json:"name"`
	Identifier string `json:"identifier"`
	ImageURL   string `json:"image_url"`
}

type SearchToken struct {
	Name        string `json:"name"`
	Symbol      string `json:"symbol"`
	ImageURL    string `json:"image_url"`
	IsBaseToken bool   `json:"is_base_token"`
}

type SearchType string

const (
	SearchTypePool SearchType = "pool"
)

// Pool
type Pool struct {
	Data     PoolData       `json:"data"`
	Included []PoolIncluded `json:"included"`
}

type PoolData struct {
	ID            string             `json:"id"`
	Type          string             `json:"type"`
	Attributes    PoolDataAttributes `json:"attributes"`
	Relationships PoolRelationships  `json:"relationships"`
}

type PoolDataAttributes struct {
	BaseTokenPriceUsd             string                    `json:"base_token_price_usd"`
	BaseTokenPriceNativeCurrency  string                    `json:"base_token_price_native_currency"`
	QuoteTokenPriceUsd            string                    `json:"quote_token_price_usd"`
	QuoteTokenPriceNativeCurrency string                    `json:"quote_token_price_native_currency"`
	Address                       string                    `json:"address"`
	Name                          string                    `json:"name"`
	ReserveInUsd                  string                    `json:"reserve_in_usd"`
	PoolCreatedAt                 string                    `json:"pool_created_at"`
	FdvUsd                        string                    `json:"fdv_usd"`
	MarketCapUsd                  string                    `json:"market_cap_usd"`
	PriceChangePercentage         PoolPriceChangePercentage `json:"price_change_percentage"`
	Transactions                  PoolTransactions          `json:"transactions"`
	VolumeUsd                     PoolVolumeUsd             `json:"volume_usd"`
}

type PoolPriceChangePercentage struct {
	H1  string `json:"h1"`
	H24 string `json:"h24"`
}

type PoolTransactions struct {
	H1  PoolTx `json:"h1"`
	H24 PoolTx `json:"h24"`
}

type PoolTx struct {
	Buys  int64 `json:"buys"`
	Sells int64 `json:"sells"`
}

type PoolVolumeUsd struct {
	H24 string `json:"h24"`
}

type PoolRelationships struct {
	BaseToken  PoolBaseToken `json:"base_token"`
	QuoteToken PoolBaseToken `json:"quote_token"`
	Dex        PoolBaseToken `json:"dex"`
}

type PoolBaseToken struct {
	Data PoolBaseTokenData `json:"data"`
}

type PoolBaseTokenData struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

type PoolIncluded struct {
	ID         string                 `json:"id"`
	Type       string                 `json:"type"`
	Attributes PoolIncludedAttributes `json:"attributes"`
}

type PoolIncludedAttributes struct {
	Address         *string `json:"address,omitempty"`
	Name            string  `json:"name"`
	Symbol          *string `json:"symbol,omitempty"`
	CoingeckoCoinID *string `json:"coingecko_coin_id,omitempty"`
}

// PoolP1
type PoolP1 struct {
	Data     PoolP1Data       `json:"data"`
}

type PoolP1Data struct {
	ID            string             `json:"id"`
	Type          string             `json:"type"`
	Relationships PoolP1Relationships      `json:"relationships"`
}

type PoolP1Relationships struct {
	Pairs PoolP1Pairs `json:"pairs"`
}

type PoolP1Pairs struct {
	Data []PoolP1Datum `json:"data"`
}

type PoolP1Datum struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

// Candlesticks
type Candlesticks struct {
	Data []Candlestick `json:"data"`
}

type Candlestick struct {
	Dt string  `json:"dt"`
	O  float64 `json:"o"`
	H  float64 `json:"h"`
	L  float64 `json:"l"`
	C  float64 `json:"c"`
	V  float64 `json:"v"`
}
