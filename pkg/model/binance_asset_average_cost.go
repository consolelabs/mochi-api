package model

type BinanceAssetAverageCost struct {
	ProfileId   string `json:"profile_id"`
	Symbol      string `json:"symbol"`
	TotalAmount string `json:"total_amount"`
	AverageCost string `json:"average_cost"`
}
