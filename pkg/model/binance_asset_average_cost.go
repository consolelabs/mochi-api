package model

type BinanceAssetAverageCost struct {
	ProfileId   string `json:"profile_id"`
	Symbol      string `json:"symbol"`
	AverageCost string `json:"average_cost"`
}
