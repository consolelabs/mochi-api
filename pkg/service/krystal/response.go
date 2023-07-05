package krystal

type BalanceTokenResponse struct {
	Data []BalanceToken `json:"data"`
}

type BalanceToken struct {
	ChainName string    `json:"chainName"`
	ChainId   int       `json:"chainId"`
	Balances  []Balance `json:"balances"`
	// ChainLogo string    `json:"chainLogo"`
}

type Balance struct {
	Token     Token  `json:"token"`
	TokenType string `json:"tokenType"`
	Balance   string `json:"balance"`
	Quotes    Quote  `json:"quotes"`
	// UserAddress string `json:"userAddress"`
}

type Token struct {
	Address  string `json:"address"`
	Symbol   string `json:"symbol"`
	Name     string `json:"name"`
	Decimals int    `json:"decimals"`
	// Logo     string `json:"logo"`
	// Tag      string `json:"tag"`
}

type Quote struct {
	Usd Usd `json:"usd,omitempty"`
}

type Usd struct {
	Price float64 `json:"price"`
	Value float64 `json:"value"`
	// Symbol                   string  `json:"symbol"`
	// PriceChange24HPercentage float64 `json:"priceChange24hPercentage"`
	// Timestamp                int     `json:"timestamp"`
}
