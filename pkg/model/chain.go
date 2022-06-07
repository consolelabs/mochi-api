package model

type Chain struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	RPC        string `json:"-"`
	APIBaseURL string `json:"-"`
	APIKey     string `json:"-"`
	TxBaseURL  string `json:"-"`
	Currency   string `json:"currency"`
}
