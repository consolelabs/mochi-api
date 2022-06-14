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

	for i, chain := range listChain {
		if strings.ToLower(symbol) == strings.ToLower(listChain[i].Currency) {
			return chain, true, nil
		}
	}

	return listChain[0], false, nil
}
