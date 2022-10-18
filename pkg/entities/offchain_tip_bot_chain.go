package entities

import (
	"github.com/defipod/mochi/pkg/model"
	offchaintipbotchain "github.com/defipod/mochi/pkg/repo/offchain_tip_bot_chain"
)

func (e *Entity) OffchainTipBotListAllChains(f offchaintipbotchain.Filter) (returnChain []model.OffchainTipBotChain, err error) {
	returnChain, err = e.repo.OffchainTipBotChain.GetAll(f)
	if err != nil {
		return returnChain, err
	}

	return returnChain, nil
}
