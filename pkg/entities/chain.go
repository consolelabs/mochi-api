package entities

import (
	"strings"

	"github.com/defipod/mochi/pkg/model"
)

func (e *Entity) GetChainIdBySymbol(symbol string) (model.Chain, bool, error) {
	var returnChain model.Chain

	listChain, err := e.repo.Chain.GetAll()
	if err != nil {
		return returnChain, false, err
	}

	for i := 0; i < len(listChain); i++ {
		if symbol == strings.ToLower(listChain[i].Currency) {
			return listChain[i], true, nil
		}
	}

	return listChain[0], false, nil
}
