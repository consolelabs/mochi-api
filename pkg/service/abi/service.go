package abi

import "github.com/ethereum/go-ethereum/core/types"

type Service interface {
	GetNameAndSymbol(address string, chainId int64) (name string, symbol string, err error)
	SweepNativeToken(rpcUrl string, contractAddress string) (*types.Transaction, error)
}
