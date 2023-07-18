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

type GetEarningOptionsResponse struct {
	Result []EarningOption `json:"result"`
}

type EarningOption struct {
	Apy       float64     `json:"apy"`
	Chain     Chain       `json:"chain"`
	Platforms []Platforms `json:"platforms"`
	Token     Token       `json:"token"`
	Tvl       float64     `json:"tvl"`
}

type Chain struct {
	ID   int    `json:"id"`
	Logo string `json:"logo"`
	Name string `json:"name"`
}

type Status struct {
	Detail string `json:"detail"`
	Value  string `json:"value"`
}

type Platforms struct {
	Apy       float64 `json:"apy"`
	Desc      string  `json:"desc"`
	Logo      string  `json:"logo"`
	Name      string  `json:"name"`
	RewardAPY float64 `json:"rewardAPY"`
	Status    Status  `json:"status"`
	Tvl       float64 `json:"tvl"`
	Type      string  `json:"type"`
}
