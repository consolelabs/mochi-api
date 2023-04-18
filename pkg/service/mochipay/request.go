package mochipay

type CreateTokenRequest struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Symbol      string `json:"symbol"`
	Decimal     int64  `json:"decimal"`
	ChainId     string `json:"chain_id"`
	Address     string `json:"address"`
	Icon        string `json:"icon"`
	CoinGeckoId string `json:"coin_gecko_id"`
}

type GetTokenRequest struct {
	Symbol  string `json:"symbol"`
	ChainId string `json:"chain_id"`
}
