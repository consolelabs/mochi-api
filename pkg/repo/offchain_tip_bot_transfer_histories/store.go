package offchain_tip_bot_transfer_histories

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	CreateTransferHistories(transferHistories []model.OffchainTipBotTransferHistory) ([]model.OffchainTipBotTransferHistory, error)
}
