package entities

import (
	"github.com/defipod/mochi/pkg/model"
)

func (e *Entity) GetUserTransaction(userDiscordId string) ([]model.OffchainTipBotTransferHistory, error) {
	return e.repo.OffchainTipBotTransferHistories.GetByUserDiscordId(userDiscordId)
}

func (e *Entity) GetTransactionsByQuery(senderId, receiverId, token string) ([]model.OffchainTipBotTransferHistory, error) {
	return e.repo.OffchainTipBotTransferHistories.GetTransactionsByQuery(senderId, receiverId, token)
}
