package response

import (
	"github.com/defipod/mochi/pkg/model"
)

type OffchainTipBotTransferToken struct {
	Id          string  `json:"id"`
	AmountEach  float64 `json:"amount_each"`
	TotalAmount float64 `json:"total_amount"`
	TxId        int64   `json:"tx_id"`
}

type OffchainTipBotTransferTokenResponse struct {
	Data []OffchainTipBotTransferToken `json:"data"`
}

type TotalBalances struct {
	Symbol      string  `json:"symbol"`
	Amount      float64 `json:"amount"`
	AmountInUsd float64 `json:"amount_in_usd"`
}

type AllTipBotTokensResponse struct {
	Data []model.OffchainTipBotToken `json:"data"`
}

type TransferTokenV2Data struct {
	Id          string  `json:"id"`
	AmountEach  float64 `json:"amount_each"`
	TotalAmount float64 `json:"total_amount"`
	TxId        int64   `json:"tx_id"`
}

type TransferTokenV2Response struct {
	Data *TransferTokenV2Data `json:"data"`
}
