package entities

import (
	"github.com/defipod/mochi/pkg/model"
)

func (e *Entity) GetUserTransaction(userDiscordId string) ([]model.OffchainTipBotTransferHistory, error) {
	return e.repo.OffchainTipBotTransferHistories.GetByUserDiscordId(userDiscordId)
}
