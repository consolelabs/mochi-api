package response

type OffchainTipBotWithdrawResponse struct {
	FromDiscordID  string  `json:"from_discord_id"`
	ToAddress      string  `json:"to_address"`
	Amount         float32 `json:"amount"`
	Cryptocurrency string  `json:"cryptocurrency"`
	TxHash         string  `json:"tx_hash"`
	TxUrl          string  `json:"tx_url"`
	WithdrawAmount float32 `json:"withdraw_amount"`
	TransactionFee float32 `json:"transaction_fee"`
}
