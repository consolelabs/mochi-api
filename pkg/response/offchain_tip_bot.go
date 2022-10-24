package response

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
