package coingecko

type CoinGeckoSearchResponse struct {
	Coins      []Coin        `json:"coins,omitempty"`
	Exchanges  []interface{} `json:"exchanges,omitempty"`
	Icos       []interface{} `json:"icos,omitempty"`
	Categories []Category    `json:"categories,omitempty"`
	Nfts       []Nft         `json:"nfts,omitempty"`
}

type Category struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type Coin struct {
	ID            *string `json:"id,omitempty"`
	Name          *string `json:"name,omitempty"`
	APISymbol     *string `json:"api_symbol,omitempty"`
	Symbol        *string `json:"symbol,omitempty"`
	MarketCapRank *int64  `json:"market_cap_rank"`
	Thumb         *string `json:"thumb,omitempty"`
	Large         *string `json:"large,omitempty"`
}

type Nft struct {
	ID     *string `json:"id,omitempty"`
	Name   *string `json:"name,omitempty"`
	Symbol *string `json:"symbol,omitempty"`
	Thumb  *string `json:"thumb,omitempty"`
}
