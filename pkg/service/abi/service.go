package abi

import (
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/request"
)

type Service interface {
	GetNameAndSymbol(address string, chainId int64) (name string, symbol string, err error)
	SweepTokens(contractAddr string, chainID int64, token model.Token) (*types.Transaction, error)
	SwapTokenOnKyber(req request.KyberSwapRequest) (*types.Transaction, error)
}
