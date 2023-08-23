package response

type SearchDexPairResponse struct {
	Pairs []DexPair `json:"pairs"`
}

type GetDexPairResponse struct {
	Pair DexPair `json:"pair"`
}

type DexPair struct {
	Id                    string            `json:"id"`
	Name                  string            `json:"name"`
	Address               string            `json:"address"`
	ChainId               string            `json:"chain_id"`
	DexId                 string            `json:"dex_id"`
	Url                   map[string]string `json:"url"`
	Price                 float64           `json:"price"`
	PriceUsd              float64           `json:"price_usd"`
	PricePercentChange24H float64           `json:"price_percent_change_24h"`
	BaseToken             DexToken          `json:"base_token"`
	QuoteToken            DexToken          `json:"quote_token"`
	CreatedAt             int64             `json:"created_at"`
	MarketCapUsd          float64           `json:"market_cap_usd"`
	VolumeUsd24h          float64           `json:"volume_usd_24h"`
	LiquidityUsd          float64           `json:"liquidity_usd"`
	Owner                 string            `json:"owner"`
	Fdv                   float64           `json:"fdv"`
	Txn24hBuy             int               `json:"txn_24h_buy"`
	Txn24hSell            int               `json:"txn_24h_sell"`
	Holders               []DexTokenHolder  `json:"holders"`
}

type DexTokenHolder struct {
	Address string  `json:"address"`
	Alias   string  `json:"alias"`
	Balance float64 `json:"balance"`
	Percent float64 `json:"percent"`
}

type DexToken struct {
	Address string `json:"address"`
	Name    string `json:"name"`
	Symbol  string `json:"symbol"`
}
