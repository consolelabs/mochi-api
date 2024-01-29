package model

type TokenPriceSnapshot struct {
	ID     int     `json:"id"`
	Symbol string  `json:"symbol"`
	Chain  string  `json:"chain"`
	Price  float64 `json:"price"`
}
