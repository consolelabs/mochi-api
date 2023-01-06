package response

import "github.com/defipod/mochi/pkg/model"

type SubmitOnchainTransfer struct {
	SenderID    string  `json:"sender_id"`
	RecipientID string  `json:"recipient_id"`
	Amount      float64 `json:"amount"`
	Symbol      string  `json:"symbol"`
	AmountInUSD float64 `json:"amount_in_usd"`
}

type SubmitOnchainTransferResponse struct {
	Data []SubmitOnchainTransfer `json:"data"`
}

type ClaimOnchainTransfer struct {
	SubmitOnchainTransfer
	RecipientAddress string `json:"recipient_address"`
	TxHash           string `json:"tx_hash"`
	TxUrl            string `json:"tx_url"`
}

type ClaimOnchainTransferResponse struct {
	Data *ClaimOnchainTransfer `json:"data"`
}

type GetOnchainTransfersResponse struct {
	Data []model.OnchainTipBotTransaction `json:"data"`
}
