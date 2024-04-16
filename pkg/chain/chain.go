package chain

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"math"
	"math/big"
	"strconv"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
	"github.com/nanmu42/etherscan-api"
	"golang.org/x/crypto/sha3"

	"github.com/defipod/mochi/pkg/config"
	"github.com/defipod/mochi/pkg/contracts/erc20"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/model"
)

type Chain struct {
	client   *ethclient.Client
	scan     *etherscan.Client
	hdwallet *hdwallet.Wallet
	log      logger.Logger
	config   *config.Config
}

func NewClient(config *config.Config, hdwallet *hdwallet.Wallet, log logger.Logger, rpcURL, apiKey, baseURL string) (*Chain, error) {
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		return nil, err
	}

	scan := etherscan.NewCustomized(etherscan.Customization{
		Timeout: 15 * time.Second,
		Key:     apiKey,
		BaseURL: baseURL,
		Verbose: false,
	})

	return &Chain{
		client:   client,
		hdwallet: hdwallet,
		scan:     scan,
		log:      log,
		config:   config,
	}, nil
}

func (ch *Chain) nativeBalance(address string, token model.Token) (float64, error) {
	account := common.HexToAddress(address)
	balanceAt, err := ch.client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		return float64(0), err
	}

	balance := new(big.Float)
	balance.SetString(balanceAt.String())
	value := new(big.Float).Quo(balance, big.NewFloat(math.Pow10(token.Decimals)))
	v, _ := value.Float64()

	return v, nil
}

func (ch *Chain) TransferOnchain(toAcc accounts.Account, amount float64, token model.Token, nonce int, all bool) (*types.Transaction, float64, error) {
	switch token.IsNative {
	case true:
		return ch.transferNativeOnchain(toAcc, amount, token, nonce, all)
	default:
		return ch.transferErc20TokenOnchain(toAcc, amount, token, nonce, all)
	}
}

func (ch *Chain) transferNativeOnchain(toAcc accounts.Account, amount float64, token model.Token, prevTxNonce int, all bool) (*types.Transaction, float64, error) {
	privateKey, err := crypto.HexToECDSA(ch.config.CentralizedWalletPrivateKey)
	if err != nil {
		ch.log.Error(err, "[transferNativeOnchain] crypto.HexToECDSA() failed")
		return nil, 0, err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		err = errors.New("error casting public key to ECDSA")
		ch.log.Error(err, "[transferNativeOnchain] "+err.Error())
		return nil, 0, err
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce := uint64(prevTxNonce)
	if prevTxNonce < 0 {
		nonce, err = ch.client.PendingNonceAt(context.Background(), fromAddress)
		if err != nil {
			ch.log.Error(err, "[transferNativeOnchain] client.PendingNonceAt() failed")
			return nil, 0, err
		}
	}

	gasLimit := uint64(21000)
	gasPrice, err := ch.client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, 0, err
	}

	balance, err := ch.nativeBalance(fromAddress.Hex(), token)
	if err != nil {
		ch.log.Fields(logger.Fields{"address": fromAddress.Hex(), "token": token}).Error(err, "[chain.transferNativeOnchain] nativeBalance() failed")
		return nil, 0, err
	}
	maxTxFee := float64(gasPrice.Int64()) * float64(gasLimit) / float64(math.Pow10(18))
	if all {
		if balance <= maxTxFee {
			err = errors.New("insufficient funds for gas")
			ch.log.Error(err, "[transferNativeOnchain] "+err.Error())
			return nil, 0, err
		}
		amount = balance - maxTxFee
	}

	value := new(big.Int)
	value.SetString(strconv.FormatFloat(float64(math.Pow10(18))*amount, 'f', 6, 64), 10)

	toAddress := common.HexToAddress(toAcc.Address.Hex())
	var data []byte
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)

	chainID, err := ch.client.NetworkID(context.Background())
	if err != nil {
		ch.log.Error(err, "[transferNativeOnchain] client.NetworkID() failed")
		return nil, 0, err
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		ch.log.Error(err, "[transferNativeOnchain] types.SignTx() failed")
		return nil, 0, err
	}

	logWithHash := ch.log.Fields(logger.Fields{"hash": signedTx.Hash().Hex()})
	err = ch.client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		logWithHash.Error(err, "[transferNativeOnchain] client.SendTransaction() failed")
		return nil, 0, err
	}

	receipt, err := bind.WaitMined(context.Background(), ch.client, signedTx)
	if err != nil {
		logWithHash.Error(err, "[transferNativeOnchain] bind.WaitMined() failed")
		return nil, 0, err
	}
	if receipt.Status == types.ReceiptStatusFailed {
		logWithHash.Error(err, "[transferNativeOnchain] receipt status === failed")
		return nil, 0, errors.New("unknown reason")
	}

	return signedTx, amount, nil
}

func (ch *Chain) transferErc20TokenOnchain(toAcc accounts.Account, amount float64, token model.Token, prevTxNonce int, all bool) (*types.Transaction, float64, error) {
	privateKey, err := crypto.HexToECDSA(ch.config.CentralizedWalletPrivateKey)
	if err != nil {
		ch.log.Error(err, "[transferErc20TokenOnchain] crypto.HexToECDSA() failed")
		return nil, 0, err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		err = errors.New("error casting public key to ECDSA")
		ch.log.Error(err, "[transferErc20TokenOnchain] "+err.Error())
		return nil, 0, err
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce := uint64(prevTxNonce)
	if prevTxNonce < 0 {
		nonce, err = ch.client.PendingNonceAt(context.Background(), fromAddress)
		if err != nil {
			ch.log.Error(err, "[transferErc20TokenOnchain] client.PendingNonceAt() failed")
			return nil, 0, err
		}
	}

	value := big.NewInt(0) // in wei (0 eth)
	gasPrice, err := ch.client.SuggestGasPrice(context.Background())
	if err != nil {
		ch.log.Error(err, "[transferErc20TokenOnchain] client.SuggestGasPrice() failed")
		return nil, 0, err
	}

	toAddress := common.HexToAddress(toAcc.Address.Hex())
	tokenAddress := common.HexToAddress(token.Address)

	transferFnSignature := []byte("transfer(address,uint256)")
	hash := sha3.NewLegacyKeccak256()
	hash.Write(transferFnSignature)
	methodID := hash.Sum(nil)[:4]

	paddedAddress := common.LeftPadBytes(toAddress.Bytes(), 32)

	balance, err := ch.erc20TokenBalance(fromAddress.Hex(), token)
	if err != nil {
		ch.log.Fields(logger.Fields{"address": fromAddress.Hex(), "token": token}).Error(err, "[chain.transferErc20TokenOnchain] erc20TokenBalance() failed")
		return nil, 0, err
	}
	if all {
		amount = balance
	}

	amt := new(big.Int)
	amt.SetString(strconv.FormatFloat(math.Pow10(token.Decimals)*amount, 'f', 6, 64), 10)

	if !all && balance < amount {
		err = errors.New("balance is not enough")
		ch.log.Error(err, "[transferErc20TokenOnchain] "+err.Error())
		return nil, 0, err
	}

	paddedAmount := common.LeftPadBytes(amt.Bytes(), 32)

	var data []byte
	data = append(data, methodID...)
	data = append(data, paddedAddress...)
	data = append(data, paddedAmount...)

	gasLimit, err := ch.client.EstimateGas(context.Background(), ethereum.CallMsg{
		To:   &toAddress,
		Data: data,
	})
	if err != nil {
		ch.log.Error(err, "[transferErc20TokenOnchain] client.EstimateGas() failed")
		return nil, 0, err
	}
	gasLimit *= 3
	// TODO: temp workaround -> find other generic solution
	if strings.EqualFold(token.Symbol, "fbomb") {
		gasLimit = uint64(1000000)
	}

	tx := types.NewTransaction(nonce, tokenAddress, value, gasLimit, gasPrice, data)

	chainID, err := ch.client.NetworkID(context.Background())
	if err != nil {
		ch.log.Error(err, "[transferErc20TokenOnchain] client.NetworkID() failed")
		return nil, 0, err
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		ch.log.Error(err, "[transferErc20TokenOnchain] types.SignTx() failed")
		return nil, 0, err
	}

	logWithHash := ch.log.Fields(logger.Fields{"hash": signedTx.Hash().Hex()})
	err = ch.client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		logWithHash.Error(err, "[transferErc20TokenOnchain] client.SendTransaction() failed")
		return nil, 0, err
	}

	receipt, err := bind.WaitMined(context.Background(), ch.client, signedTx)
	if err != nil {
		logWithHash.Error(err, "[transferErc20TokenOnchain] bind.WaitMined() failed")
		return nil, 0, err
	}
	if receipt.Status == types.ReceiptStatusFailed {
		logWithHash.Error(err, "[transferErc20TokenOnchain] receipt status === failed")
		return nil, 0, errors.New("unknown reason")
	}

	return signedTx, amount, nil
}

func (ch *Chain) erc20TokenBalance(address string, token model.Token) (float64, error) {
	tokenBalance, err := ch.scan.TokenBalance(token.Address, address)
	if err != nil {
		return 0, err
	}

	balance := new(big.Float)
	balance.SetString(tokenBalance.Int().String())
	value := new(big.Float).Quo(balance, big.NewFloat(math.Pow10(token.Decimals)))
	v, _ := value.Float64()
	return v, nil
}

func (ch *Chain) Balances(address string, tokens []model.Token) (map[string]float64, error) {
	balances := make(map[string]float64, 0)
	for _, token := range tokens {
		key := strings.ToUpper(token.Symbol)
		switch token.IsNative {
		case true:
			nativeCryptoBal, err := ch.nativeBalance(address, token)
			if err != nil {
				return nil, err
			}
			balances[key] = nativeCryptoBal
		default:
			tokenBalance, err := ch.erc20TokenBalance(address, token)
			if err != nil {
				return nil, err
			}
			balances[key] = tokenBalance
		}
	}
	return balances, nil
}

func (ch *Chain) TokenTotalSupply(address string, decimal int) (float64, error) {
	v, err := ch.scan.TokenTotalSupply(address)
	if err != nil {
		return 0, err
	}

	if v == nil {
		return 0, nil
	}

	str := v.Int().String()
	f, ok := new(big.Float).SetString(str)
	f.Quo(f, big.NewFloat(math.Pow10(decimal)))
	if !ok {
		return 0, nil
	}

	supply, _ := f.Float64()
	return supply, nil
}

// func (ch *Chain) RawBalances(address string, tokens []model.Token) (map[string]*big.Int, error) {
// 	balances := make(map[string]*big.Int, 0)
// 	for _, token := range tokens {
// 		key := strings.ToUpper(token.Symbol)
// 		switch token.IsNative {
// 		case true:
// 			nativeCryptoBal, err := ch.rawNativeBalance(address, token)
// 			if err != nil {
// 				return nil, err
// 			}
// 			balances[key] = nativeCryptoBal
// 		default:
// 			tokenBalance, err := ch.rawErc20TokenBalance(address, token)
// 			if err != nil {
// 				return nil, err
// 			}
// 			balances[key] = tokenBalance
// 		}
// 	}
// 	return balances, nil
// }

func (ch *Chain) RawNativeBalance(address string, token model.Token) (*big.Int, error) {
	account := common.HexToAddress(address)
	balanceAt, err := ch.client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		return nil, err
	}
	return balanceAt, nil
}

func (ch *Chain) RawErc20TokenBalance(address string, token model.Token) (*big.Int, error) {
	erc20, err := erc20.NewErc20Caller(common.HexToAddress(token.Address), ch.client)
	tokenBalance, err := erc20.BalanceOf(nil, common.HexToAddress(address))
	if err != nil {
		return nil, err
	}
	return tokenBalance, nil
}
