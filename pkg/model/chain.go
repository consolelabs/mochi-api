package model

type Chain struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	ShortName   string `json:"short_name"`
	CoinGeckoID string `json:"coin_gecko_id"`
	RPC         string `json:"-"`
	APIBaseURL  string `json:"-"`
	APIKey      string `json:"-"`
	TxBaseURL   string `json:"-"`
	Currency    string `json:"currency"`
}
