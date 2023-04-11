package abi

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"errors"
	"log"
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rlp"

	"github.com/defipod/mochi/pkg/config"
	abi "github.com/defipod/mochi/pkg/contract/erc721"
	"github.com/defipod/mochi/pkg/contracts/deposit"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/util"
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
	if chainId == 9999 {
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
	case 137:
		return e.config.RpcUrl.Polygon
	case 42161:
		return e.config.RpcUrl.Arbitrum
	case 66:
		return e.config.RpcUrl.Okc
	case 1975:
		return e.config.RpcUrl.Onus
	default:
		return e.config.RpcUrl.Eth
	}
}

func (e *abiEntity) SweepTokens(contractAddr string, chainID int64, token model.Token) (*types.Transaction, error) {
	l := logger.NewLogrusLogger()
	rpcUrl := e.selectRpcUrl(chainID)
	client, err := ethclient.Dial(rpcUrl)
	if err != nil {
		return nil, err
	}

	privateKey, err := crypto.HexToECDSA(e.config.CentralizedWalletPrivateKey)
	if err != nil {
		return nil, err
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, errors.New("error casting public key to ECDSA")
	}

	centralizedAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), centralizedAddress)
	if err != nil {
		return nil, err
	}
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}

	networkID, err := client.NetworkID(context.Background())
	if err != nil {
		return nil, err
	}
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, networkID)
	if err != nil {
		return nil, err
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(300000)
	auth.GasPrice = gasPrice

	depositContract, err := deposit.NewDeposit(common.HexToAddress(contractAddr), client)
	if err != nil {
		return nil, err
	}

	var tx *types.Transaction
	if token.IsNative {
		tx, err = depositContract.SweepNativeToken(auth)
	} else {
		tx, err = depositContract.SweepToken(auth, common.HexToAddress(token.Address))
	}
	if err != nil {
		return nil, err
	}
	receipt, err := bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		return nil, err
	}
	log := l.Fields(logger.Fields{"txHash": receipt.TxHash.Hex(), "chainID": chainID})
	if receipt.Status == 0 {
		log.Info("sweep tokens tx failed")
	} else if receipt.Status == 1 {
		log.Info("sweep tokens tx succeeded")
	}
	return tx, nil
}

func (e *abiEntity) PrepareTxOpts(client *ethclient.Client) (*bind.TransactOpts, error) {
	privateKey, err := crypto.HexToECDSA(e.config.CentralizedWalletPrivateKey)
	if err != nil {
		return nil, err
	}

	publicKey := privateKey.Public()

	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, errors.New("error casting public key to ECDSA")
	}

	centralizedAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	nonce, err := client.PendingNonceAt(context.Background(), centralizedAddress)
	if err != nil {
		return nil, err
	}
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}

	networkID, err := client.NetworkID(context.Background())
	if err != nil {
		return nil, err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, networkID)
	if err != nil {
		return nil, err
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(300000)
	auth.GasPrice = gasPrice

	return auth, nil
}

func (e *abiEntity) SwapTokenOnKyber(req request.KyberSwapRequest) (*types.Transaction, error) {
	// l := logger.NewLogrusLogger()
	chainID := util.ConvertChainNameToChainId(req.ChainName)
	rpcUrl := e.selectRpcUrl(chainID)

	// rpcUrl = https://rpc.ankr.com/fantom
	client, err := ethclient.Dial(rpcUrl)
	if err != nil {
		return nil, err
	}

	privateKey, err := crypto.HexToECDSA(e.config.CentralizedWalletPrivateKey)
	if err != nil {
		return nil, err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, errors.New("error casting public key to ECDSA")
	}

	centralizedAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), centralizedAddress)
	if err != nil {
		return nil, err
	}

	gasInt, _ := strconv.Atoi(req.Gas)
	gasLimit := uint64(gasInt) * 5
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}

	networkID, err := client.NetworkID(context.Background())
	if err != nil {
		return nil, err
	}

	data := common.FromHex(req.EncodedData)

	//
	tx := types.NewTransaction(nonce, common.HexToAddress(req.RouterAddress), big.NewInt(req.Amount.Int64()), gasLimit, gasPrice, data)
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(networkID), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	ts := types.Transactions{signedTx}
	rawTxBytes, _ := rlp.EncodeToBytes(ts[0])
	rawTxHex := hex.EncodeToString(rawTxBytes)

	// send tx
	rawTxSendBytes, _ := hex.DecodeString(rawTxHex)

	txSend := new(types.Transaction)

	rlp.DecodeBytes(rawTxSendBytes, &txSend)

	err = client.SendTransaction(context.Background(), txSend)
	if err != nil {
		log.Fatal(err)
	}

	return nil, nil
}
