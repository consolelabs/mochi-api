package model

type UserTokenWatchlistItem struct {
	ProfileID   string `json:"profile_id"`
	CoinGeckoID string `json:"coin_gecko_id"`
	Symbol      string `json:"symbol"`
	IsFiat      bool   `json:"is_fiat"`
}
