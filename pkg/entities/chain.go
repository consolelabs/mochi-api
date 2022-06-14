package entities

import (
	"strings"

	"github.com/defipod/mochi/pkg/model"
)

func (e *Entity) GetChainIdBySymbol(symbol string) ([]model.Chain, error) {
	listChain, err := e.repo.Chain.GetAll()
	var returnChain []model.Chain
	if err != nil {
		return listChain, err
	}

	for i := 0; i < len(listChain); i++ {
		if strings.ToUpper(symbol) == listChain[i].Currency {
			returnChain = append(returnChain, listChain[i])
			return returnChain, nil
		}
	}

	return listChain, nil
}
