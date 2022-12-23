package offchain_tip_bot_chain

import (
	"github.com/defipod/mochi/pkg/model"
	"github.com/google/uuid"
)

type Store interface {
	GetAll(f Filter) ([]model.OffchainTipBotChain, error)
	GetByID(id uuid.UUID) (model.OffchainTipBotChain, error)
	GetByChainID(chainID int) (model.OffchainTipBotChain, error)
}

type Filter struct {
	TokenID             string
	TokenSymbol         string
	IsContractAvailable bool
}
