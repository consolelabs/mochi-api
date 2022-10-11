package offchain_tip_bot_chain

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	GetAll(f Filter) ([]model.OffchainTipBotChain, error)
	GetByID(id string) (model.OffchainTipBotChain, error)
}

type Filter struct {
	TokenID             string
	TokenSymbol         string
	IsContractAvailable bool
}
