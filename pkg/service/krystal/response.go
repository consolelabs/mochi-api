package krystal

type BalanceTokenResponse struct {
	Data []BalanceToken `json:"data"`
}

type BalanceToken struct {
	ChainName string    `json:"chainName"`
	ChainId   int       `json:"chainId"`
	ChainLogo string    `json:"chainLogo"`
	Balances  []Balance `json:"balances"`
}

type Balance struct {
	Token       Token  `json:"token"`
	TokenType   string `json:"tokenType"`
	Balance     string `json:"balance"`
	UserAddress string `json:"userAddress"`
	Quotes      Quote  `json:"quotes"`
}

type Token struct {
	Address  string `json:"address"`
	Symbol   string `json:"symbol"`
	Name     string `json:"name"`
	Decimals int    `json:"decimals"`
	Logo     string `json:"logo"`
	Tag      string `json:"tag"`
}

type Quote struct {
	Usd Usd `json:"usd,omitempty"`
}

type Usd struct {
	Symbol                   string  `json:"symbol"`
	Price                    float64 `json:"price"`
	PriceChange24HPercentage float64 `json:"priceChange24hPercentage"`
	Value                    float64 `json:"value"`
	Timestamp                int     `json:"timestamp"`
}
