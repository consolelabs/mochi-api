package chain

import (
	"github.com/defipod/mochi/pkg/model"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/core/types"
)

type Chain interface {
	Balance(address string) (float64, error)
	Transfer(fromAccount accounts.Account, toAccount accounts.Account, amount float64, token model.Token, nonce int, all bool) (*types.Transaction, error)
	Balances(address string, tokens []model.Token) (map[string]float64, error)
}
