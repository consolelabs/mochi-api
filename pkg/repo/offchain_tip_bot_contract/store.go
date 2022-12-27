package offchain_tip_bot_contract

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	List(ListQuery) ([]model.OffchainTipBotContract, error)
	GetByID(id string) (model.OffchainTipBotContract, error)
	GetByAddress(addr string) (model.OffchainTipBotContract, error)
	CreateAssignContract(ac *model.OffchainTipBotAssignContract) (*model.OffchainTipBotAssignContract, error)
	GetAssignContract(address, tokenSymbol string) (*model.OffchainTipBotAssignContract, error)
	DeleteExpiredAssignContract() error
}
