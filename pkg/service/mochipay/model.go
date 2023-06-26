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
	Id        string `json:"id"`
	ProfileId string `json:"profile_id"`
	TokenId   string `json:"token_id"`
	Amount    string `json:"amount"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Token     Token  `json:"token"`
}
type GetChainDataResponse struct {
	Data []Chain `json:"data"`
}

type Token struct {
	Id          string  `json:"id"`
	Name        string  `json:"name"`
	Symbol      string  `json:"symbol"`
	Decimal     int64   `json:"decimal"`
	ChainId     string  `json:"chain_id"`
	Native      bool    `json:"native"`
	Address     string  `json:"address"`
	Icon        string  `json:"icon"`
	Price       float64 `json:"price"`
	Chain       Chain   `json:"chain"`
	CoinGeckoId string  `json:"coin_gecko_id"`
}

type Chain struct {
	Id       string `json:"id"`
	ChainId  string `json:"chain_id"`
	Name     string `json:"name"`
	Symbol   string `json:"symbol"`
	Rpc      string `json:"rpc"`
	Explorer string `json:"explorer"`
	Icon     string `json:"icon"`
	IsEvm    bool   `json:"is_evm"`
}

type ListTokensResponse struct {
	Data []Token `json:"data"`
}

type GetTokenResponse struct {
	Data *Token `json:"data"`
}

type VaultResponse struct {
	Data VaultTransferToken `json:"data"`
}

type VaultTransferToken struct {
	TxHash string `json:"tx_hash"`
}

type TipResponse struct {
	Data TipDataResponse `json:"data"`
}

type TipDataResponse struct {
	Id             string      `json:"id"`
	TxId           int64       `json:"tx_id"`
	ProfileId      string      `json:"profile_id"`
	OtherProfileId string      `json:"other_profile_id"`
	Type           string      `json:"type"`
	TokenId        string      `json:"token_id"`
	Amount         string      `json:"amount"`
	CreatedAt      string      `json:"created_at"`
	UpdatedAt      string      `json:"updated_at"`
	Token          interface{} `json:"token"`
}
