package mochipay

import "github.com/consolelabs/mochi-typeset/mochi-pay/typeset"

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

type CreateBatchTokenRequest struct {
	Tokens []CreateTokenRequest `json:"tokens"`
}

type TokenProperties struct {
	// can add more if want
	ChainId string `json:"chain_id"`
	Address string `json:"address"`
}

type Wallet struct {
	ProfileGlobalId string `json:"profile_global_id"`
}

type TransferV2Request struct {
	From     *Wallet                   `json:"from"`
	Tos      []*Wallet                 `json:"tos"`
	Amount   []string                  `json:"amount"`
	TokenId  string                    `json:"token_id"`
	Platform string                    `json:"platform"`
	Action   typeset.TransactionAction `json:"action" binding:"required"`
	Metadata map[string]interface{}    `json:"metadata"`
}
