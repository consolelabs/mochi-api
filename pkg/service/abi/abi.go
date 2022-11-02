package abi

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"math/big"

	"github.com/defipod/mochi/pkg/config"
	depositabi "github.com/defipod/mochi/pkg/contract/deposit"
	abi "github.com/defipod/mochi/pkg/contract/erc721"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type abiEntity struct {
	config *config.Config
	log    logger.Logger
}

func NewAbi(cfg *config.Config, log logger.Logger) Service {
	return &abiEntity{
		config: cfg,
		log:    log,
	}
}

func (e *abiEntity) GetNameAndSymbol(address string, chainId int64) (name string, symbol string, err error) {
	if chainId == 99999999 {
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

// deposit logic
func (e *abiEntity) SweepNativeToken(rpcUrl string, contractAddress string) (*types.Transaction, error) {
	client, err := ethclient.Dial(rpcUrl)
	if err != nil {
		e.log.Fields(logger.Fields{"rpcUrl": rpcUrl, "contractAddress": contractAddress}).Error(err, "failed to init client")
		return nil, err
	}

	privateKey, err := crypto.HexToECDSA(e.config.CentralizedWalletPrivateKey)
	if err != nil {
		e.log.Fields(logger.Fields{"rpcUrl": rpcUrl, "contractAddress": contractAddress}).Error(err, "failed to encrypt private key")
		return nil, err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		e.log.Fields(logger.Fields{"rpcUrl": rpcUrl, "contractAddress": contractAddress}).Error(err, "failed to get public key")
		return nil, err
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		e.log.Fields(logger.Fields{"rpcUrl": rpcUrl, "contractAddress": contractAddress}).Error(err, "failed to get nonce")
		return nil, err
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		e.log.Fields(logger.Fields{"rpcUrl": rpcUrl, "contractAddress": contractAddress}).Error(err, "failed to get suggest gas price")
		return nil, err
	}

	chainId, err := client.NetworkID(context.Background())
	if err != nil {
		e.log.Fields(logger.Fields{"rpcUrl": rpcUrl, "contractAddress": contractAddress}).Error(err, "failed to get network id")
		return nil, err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	if err != nil {
		e.log.Fields(logger.Fields{"rpcUrl": rpcUrl, "contractAddress": contractAddress}).Error(err, "failed to get get key transactor")
		return nil, err
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(500000) // in units
	auth.GasPrice = gasPrice

	address := common.HexToAddress(contractAddress)
	instance, err := depositabi.NewDeposit(address, client)
	if err != nil {
		e.log.Fields(logger.Fields{"rpcUrl": rpcUrl, "contractAddress": contractAddress}).Error(err, "failed to create instance")
		return nil, err
	}

	tx, err := instance.SweepNativeToken(auth)
	if err != nil {
		e.log.Fields(logger.Fields{"rpcUrl": rpcUrl, "contractAddress": contractAddress}).Error(err, "failed to sweep native token")
		return nil, err
	}

	return tx, nil
}
