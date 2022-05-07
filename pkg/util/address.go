package util

import (
	"strings"
)

type ChainType string

const (
	EVM   ChainType = "evm"
	TERRA ChainType = "terra"
)

func GetChainTypeFromAddress(address string) ChainType {
	if strings.HasPrefix(address, string(TERRA)) {
		return TERRA
	}

	return EVM
}
