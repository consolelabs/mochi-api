package dexscreener

type PairResponse struct {
	SchemaVersion string `json:"schemaVersion"`
	Pairs         []Pair `json:"pairs"`
}

type Pair struct {
	ChainID       string      `json:"chainId"`
	DexID         string      `json:"dexId"`
	URL           string      `json:"url"`
	PairAddress   string      `json:"pairAddress"`
	Labels        []string    `json:"labels"`
	BaseToken     EToken      `json:"baseToken"`
	QuoteToken    EToken      `json:"quoteToken"`
	PriceNative   string      `json:"priceNative"`
	PriceUsd      string      `json:"priceUsd"`
	Txns          Txns        `json:"txns"`
	Volume        PriceChange `json:"volume"`
	PriceChange   PriceChange `json:"priceChange"`
	Liquidity     Liquidity   `json:"liquidity"`
	Fdv           float64     `json:"fdv"`
	PairCreatedAt int64       `json:"pairCreatedAt"`
}

type EToken struct {
	Address string `json:"address"`
	Name    string `json:"name"`
	Symbol  string `json:"symbol"`
}

type Liquidity struct {
	Usd   float64 `json:"usd"`
	Base  float64 `json:"base"`
	Quote float64 `json:"quote"`
}

type PriceChange struct {
	H1  float64 `json:"h1"`
	H24 float64 `json:"h24"`
	H6  float64 `json:"h6"`
	M5  float64 `json:"m5"`
}

type Txns struct {
	H1  Txn `json:"h1"`
	H24 Txn `json:"h24"`
	H6  Txn `json:"h6"`
	M5  Txn `json:"m5"`
}

type Txn struct {
	Buys  int64 `json:"buys"`
	Sells int64 `json:"sells"`
}
