package abi

type Service interface {
	GetNameAndSymbol(address string, chainId int64) (name string, symbol string, err error)
}
