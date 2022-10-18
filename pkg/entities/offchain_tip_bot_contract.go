package entities

import (
	"github.com/defipod/mochi/pkg/model"
)

func (e *Entity) OffchainTipBotCreateAssignContract(ac *model.OffchainTipBotAssignContract) (err error) {
	return e.repo.OffchainTipBotContract.CreateAssignContract(ac)
}
