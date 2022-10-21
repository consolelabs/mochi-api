package offchain_tip_bot_contract

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	GetByID(id string) (model.OffchainTipBotContract, error)
	GetByAddress(addr string) (model.OffchainTipBotContract, error)
	GetByChainID(id string) ([]model.OffchainTipBotContract, error)
	CreateAssignContract(ac *model.OffchainTipBotAssignContract) (*model.OffchainTipBotAssignContract, error)
	DeleteExpiredAssignContract() error
}
