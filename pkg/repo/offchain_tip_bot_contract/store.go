package offchain_tip_bot_contract

import (
	"time"

	"github.com/defipod/mochi/pkg/model"
)

type Store interface {
	List(ListQuery) ([]model.OffchainTipBotContract, error)
	GetByID(id string) (model.OffchainTipBotContract, error)
	GetByAddress(addr string) (model.OffchainTipBotContract, error)
	CreateAssignContract(ac *model.OffchainTipBotAssignContract) (*model.OffchainTipBotAssignContract, error)
	GetAssignContract(GetAssignContractQuery) (*model.OffchainTipBotAssignContract, error)
	DeleteExpiredAssignContract() error
	UpdateSweepTime(contractID string, sweepTime time.Time) error
}
