package response

import (
	"github.com/defipod/mochi/pkg/model"
)

type TransactionsResponse struct {
	Data []model.OffchainTipBotTransferHistory `json:"data"`
}
