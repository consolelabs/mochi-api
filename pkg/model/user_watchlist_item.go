package model

type UserWatchlistItem struct {
	UserID      string `json:"user_id"`
	CoinGeckoID string `json:"coin_gecko_id"`
	Symbol      string `json:"symbol"`
	IsFiat      bool   `json:"is_fiat"`
}
