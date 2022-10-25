package abi

import (
	"errors"

	"github.com/defipod/mochi/pkg/config"
	abi "github.com/defipod/mochi/pkg/contract/erc721"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type abiEntity struct {
	config *config.Config
}

func NewAbi(cfg *config.Config) Service {
	return &abiEntity{
		config: cfg,
	}
}

func (e *abiEntity) GetNameAndSymbol(address string, chainId int64) (name string, symbol string, err error) {
	if chainId == 2 {
		return "", "", nil
	}
	rpcUrl := e.selectRpcUrl(chainId)
	client, err := ethclient.Dial(rpcUrl)
	if err != nil {
		return "", "", err
	}

	addressHash := common.HexToAddress(address)
	instance, err := abi.NewERC721(addressHash, client)
	if err != nil {
		return "", "", err
	}
	name, err = instance.Name(&bind.CallOpts{})
	if err != nil {
		if err.Error() == "execution reverted" {
			return "", "", errors.New("This collection does not support collection name")
		} else {
			return "", "", err
		}

	}
	symbol, err = instance.Symbol(&bind.CallOpts{})
	if err != nil {
		if err.Error() == "execution reverted" {
			return "", "", errors.New("This collection does not support collection symbol")
		} else {
			return "", "", err
		}
	}
	return name, symbol, nil
}

func (e *abiEntity) selectRpcUrl(chainId int64) string {
	switch chainId {
	case 1:
		return e.config.RpcUrl.Eth
	case 250:
		return e.config.RpcUrl.Ftm
	case 10:
		return e.config.RpcUrl.Opt
	case 56:
		return e.config.RpcUrl.Bsc
	default:
		return e.config.RpcUrl.Eth
	}
}
