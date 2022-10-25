package response

import "math/big"

type GetUserBalances struct {
	ID            string  `json:"id"`
	Name          string  `json:"name"`
	Symbol        string  `json:"symbol"`
	Balances      float64 `json:"balances"`
	BalancesInUSD float64 `json:"balances_in_usd"`
}

type GetUserBalancesResponse struct {
	Data []GetUserBalances `json:"data"`
}

type OffchainTipBotTransferToken struct {
	SenderID    string  `json:"sender_id"`
	RecipientID string  `json:"recipient_id"`
	Amount      float64 `json:"amount"`
	Symbol      string  `json:"symbol"`
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
