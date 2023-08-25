package ethplorer

type TokenHoldersResponse struct {
	Holders []TokenHolder `json:"holders"`
}

type TokenHolder struct {
	Address string  `json:"address"`
	Balance float64 `json:"balance"`
	Share   float64 `json:"share"`
}

type TokenInfoResponse struct {
	Address           string `json:"address"`
	Decimals          string `json:"decimals"`
	Name              string `json:"name"`
	Owner             string `json:"owner"`
	Price             Price  `json:"price"`
	Symbol            string `json:"symbol"`
	TotalSupply       string `json:"totalSupply"`
	TransfersCount    int64  `json:"transfersCount"`
	TxsCount          int64  `json:"txsCount"`
	IssuancesCount    int64  `json:"issuancesCount"`
	LastUpdated       int64  `json:"lastUpdated"`
	HoldersCount      int64  `json:"holdersCount"`
	Website           string `json:"website"`
	Image             string `json:"image"`
	EthTransfersCount int64  `json:"ethTransfersCount"`
	CountOps          int64  `json:"countOps"`
}

type Price struct {
	Rate            float64 `json:"rate"`
	Diff            float64 `json:"diff"`
	Diff7D          float64 `json:"diff7d"`
	Ts              int64   `json:"ts"`
	MarketCapUsd    float64 `json:"marketCapUsd"`
	AvailableSupply int64   `json:"availableSupply"`
	Volume24H       float64 `json:"volume24h"`
	VolDiff1        float64 `json:"volDiff1"`
	VolDiff7        float64 `json:"volDiff7"`
	VolDiff30       float64 `json:"volDiff30"`
	Currency        string  `json:"currency"`
}
