package model

type CoingeckoSupportedTokens struct {
	ID           string  `json:"id"`
	Symbol       string  `json:"symbol"`
	Name         string  `json:"name"`
	CurrentPrice float64 `json:"current_price" gorm:"-"`
}
