package response

import (
	"math/big"

	"github.com/defipod/mochi/pkg/model"
)

type GetUserBalances struct {
	ID            string  `json:"id"`
	Name          string  `json:"name"`
	Symbol        string  `json:"symbol"`
	Balances      float64 `json:"balances"`
	BalancesInUSD float64 `json:"balances_in_usd"`
	RateInUSD     float64 `json:"rate_in_usd"`
}

type GetUserBalancesResponse struct {
	Data []GetUserBalances `json:"data"`
}

type OffchainTipBotTransferToken struct {
	SenderID    string  `json:"sender_id"`
	RecipientID string  `json:"recipient_id"`
	Amount      float64 `json:"amount"`
	Symbol      string  `json:"symbol"`
	AmountInUSD float64 `json:"amount_in_usd"`
}

type OffchainTipBotTransferTokenResponse struct {
	Data []OffchainTipBotTransferToken `json:"data"`
}

type OffchainTipBotWithdraw struct {
	UserDiscordID  string     `json:"user_discord_id"`
	ToAddress      string     `json:"to_address"`
	Amount         float64    `json:"amount"`
	Symbol         string     `json:"symbol"`
	TxHash         string     `json:"tx_hash"`
	TxUrl          string     `json:"tx_url"`
	WithdrawAmount *big.Float `json:"withdraw_amount"`
	TransactionFee float64    `json:"transaction_fee"`
}

type OffchainTipBotWithdrawResponse struct {
	Data OffchainTipBotWithdraw `json:"data"`
}

type UserTransactionResponse struct {
	Data []model.OffchainTipBotTransferHistory `json:"data"`
}

type TotalBalances struct {
	Symbol      string  `json:"symbol"`
	Amount      float64 `json:"amount"`
	AmountInUsd float64 `json:"amount_in_usd"`
}

type TotalOffchainBalancesInDB struct {
	Total       float64 `json:"total"`
	TokenId     string  `json:"token_id"`
	TokenSymbol string  `json:"token_symbol"`
	CoinGeckoId string  `json:"coin_gecko_id"`
}
type TotalOffchainBalances struct {
	Symbol      string  `json:"symbol"`
	Amount      float64 `json:"amount"`
	AmountInUsd float64 `json:"amount_in_usd"`
}

type TotalFeeWithdraw struct {
	Symbol   string  `json:"symbol"`
	TotalFee float64 `json:"total_fee"`
}

type AllTipBotTokensResponse struct {
	Data []model.OffchainTipBotToken `json:"data"`
}
