package model

import "time"

type KyberswapSupportedToken struct {
	Id        int64     `json:"id"`
	Address   string    `json:"address"`
	ChainId   int64     `json:"chain_id"`
	ChainName string    `json:"chain_name"`
	Decimals  int64     `json:"decimals"`
	Symbol    string    `json:"symbol"`
	Name      string    `json:"name"`
	LogoUri   string    `json:"logo_uri"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
