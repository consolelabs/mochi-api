package model

type OnchainAssetAvgCost struct {
	WalletAddress string  `json:"wallet_address"`
	TokenAddress  string  `json:"token_address"`
	Blockchain    string  `json:"blockchain"`
	Symbol        string  `json:"symbol"`
	AverageCost   float64 `json:"average_cost"`
}
