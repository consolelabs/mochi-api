package entities

import (
	"github.com/defipod/mochi/pkg/model"
)

func (e *Entity) OffchainTipBotCreateAssignContract(ac *model.OffchainTipBotAssignContract) (err error) {
	return e.repo.OffchainTipBotContract.CreateAssignContract(ac)
}

func (e *Entity) OffchainTipBotDeleteExpiredAssignContract() (err error) {
	return e.repo.OffchainTipBotContract.DeleteExpiredAssignContract()
}
