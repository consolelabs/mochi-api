package mochipay

import "time"

type Platform string

const (
	PlatformDiscord Platform = "discord"
	PlatformSol     Platform = "solana-chain"
	PlatformEVM     Platform = "evm-chain"
)

type GetProfileByDiscordResponse struct {
	ID                 string              `json:"id"`
	AssociatedAccounts []AssociatedAccount `json:"associated_accounts"`
	CreatedAt          time.Time           `json:"created_at"`
	UpdatedAt          time.Time           `json:"updated_at"`
}

type AssociatedAccount struct {
	ID                 string    `json:"id"`
	ProfileID          string    `json:"profile_id"`
	Platform           Platform  `json:"platform"`
	PlatformIdentifier string    `json:"platform_identifier"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}

type ErrorResponse struct {
	Msg        string `json:"msg"`
	StatusCode int    `json:"status_code"`
	ErrorCode  string `json:"error_code"`
}

type GetBalanceDataResponse struct {
	Data []GetBalanceResponse `json:"data"`
}

type GetBalanceResponse struct {
	Id        string          `json:"id"`
	ProfileId string          `json:"profile_id"`
	TokenId   string          `json:"token_id"`
	Amount    string          `json:"amount"`
	CreatedAt string          `json:"created_at"`
	UpdatedAt string          `json:"updated_at"`
	Token     GetBalanceToken `json:"token"`
}

type GetBalanceToken struct {
	Id      string      `json:"id"`
	Name    string      `json:"name"`
	Symbol  string      `json:"symbol"`
	Decimal int64       `json:"decimal"`
	ChainId int64       `json:"chain_id"`
	Native  bool        `json:"native"`
	Address string      `json:"address"`
	Icon    string      `json:"icon"`
	Price   float64     `json:"price"`
	Chain   interface{} `json:"chain"`
}
