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

	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/model"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
	"github.com/nanmu42/etherscan-api"

	"golang.org/x/crypto/sha3"
)

type Chain struct {
	client   *ethclient.Client
	scan     *etherscan.Client
	hdwallet *hdwallet.Wallet
	log      logger.Logger
}

func NewClient(hdwallet *hdwallet.Wallet, log logger.Logger, rpcURL, apiKey, baseURL string) (*Chain, error) {
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		return nil, err
	}

	scan := etherscan.NewCustomized(etherscan.Customization{
		Timeout: 15 * time.Second,
		Key:     apiKey,
		BaseURL: baseURL,
		Verbose: true,
	})

	return &Chain{
		client:   client,
		hdwallet: hdwallet,
		scan:     scan,
		log:      log,
	}, nil
}

func (ch *Chain) Balance(address string) (float64, error) {
	account := common.HexToAddress(address)
	balanceAt, err := ch.client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		return float64(0), err
	}

	balance := new(big.Float)
	balance.SetString(balanceAt.String())
	value := new(big.Float).Quo(balance, big.NewFloat(math.Pow10(18)))
	v, _ := value.Float64()

	return v, nil
}

func (ch *Chain) Transfer(fromAcc accounts.Account, toAcc accounts.Account, amount float64, token model.Token, nonce int, all bool) (*types.Transaction, float64, error) {
	var (
		t   *types.Transaction
		err error
	)

	switch token.IsNative {
	case true:
		t, amount, err = ch.transfer(fromAcc, toAcc, amount, nonce, all)
	default:
		t, amount, err = ch.transferToken(fromAcc, toAcc, amount, token, nonce, all)
	}

	return t, amount, err
}

func (ch *Chain) transfer(fromAcc accounts.Account, toAcc accounts.Account, amount float64, prevTxNonce int, all bool) (*types.Transaction, float64, error) {
	balance, err := ch.Balance(fromAcc.Address.Hex())
	if err != nil {
		return nil, 0, err
	}
	if !all && balance < amount {
		return nil, 0, errors.New("balance is not enough")
	}

	priv, _ := ch.hdwallet.PrivateKeyHex(fromAcc)
	privateKey, err := crypto.HexToECDSA(priv)
	if err != nil {
		return nil, 0, err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, 0, errors.New("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce := uint64(prevTxNonce)
	if prevTxNonce < 0 {
		nonce, err = ch.client.PendingNonceAt(context.Background(), fromAddress)
		if err != nil {
			return nil, 0, err
		}
	}

	gasLimit := uint64(21000)
	gasPrice, err := ch.client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, 0, err
	}

	maxTxFee := float64(gasPrice.Int64()) * float64(gasLimit) / float64(math.Pow10(18))
	if all {
		if balance <= maxTxFee {
			return nil, 0, errors.New("insufficient funds for gas")
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
		return nil, 0, err
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		return nil, 0, err
	}

	err = ch.client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return nil, 0, err
	}

	return signedTx, amount, nil
}

func (ch *Chain) transferToken(fromAcc accounts.Account, toAcc accounts.Account, amount float64, token model.Token, prevTxNonce int, all bool) (*types.Transaction, float64, error) {
	priv, _ := ch.hdwallet.PrivateKeyHex(fromAcc)
	privateKey, err := crypto.HexToECDSA(priv)
	if err != nil {
		return nil, 0, err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, 0, errors.New("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce := uint64(prevTxNonce)
	if prevTxNonce < 0 {
		nonce, err = ch.client.PendingNonceAt(context.Background(), fromAddress)
		if err != nil {
			return nil, 0, err
		}
	}

	value := big.NewInt(0) // in wei (0 eth)
	gasPrice, err := ch.client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, 0, err
	}

	toAddress := common.HexToAddress(toAcc.Address.Hex())
	tokenAddress := common.HexToAddress(token.Address)

	transferFnSignature := []byte("transfer(address,uint256)")
	hash := sha3.NewLegacyKeccak256()
	hash.Write(transferFnSignature)
	methodID := hash.Sum(nil)[:4]

	paddedAddress := common.LeftPadBytes(toAddress.Bytes(), 32)

	tokenBalance, err := ch.scan.TokenBalance(token.Address, fromAcc.Address.Hex())
	if err != nil {
		return nil, 0, err
	}
	// if all {
	// 	amount, _ = new(big.Float).Quo(big.NewFloat(0).SetInt(tokenBalance.Int()), big.NewFloat(math.Pow10(18))).Float64()
	// }

	amt := new(big.Int)
	amt.SetString(strconv.FormatFloat(math.Pow10(token.Decimals)*amount, 'f', 6, 64), 10)

	if !all && tokenBalance.Int().Cmp(amt) == -1 {
		return nil, 0, errors.New("balance is not enough")
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
		return nil, 0, err
	}
	gasLimit *= 3

	tx := types.NewTransaction(nonce, tokenAddress, value, gasLimit, gasPrice, data)

	chainID, err := ch.client.NetworkID(context.Background())
	if err != nil {
		return nil, 0, err
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		return nil, 0, err
	}

	err = ch.client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return nil, 0, err
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
	value := new(big.Float).Quo(balance, big.NewFloat(math.Pow10(18)))
	v, _ := value.Float64()
	return v, nil
}

func (ch *Chain) Balances(address string, tokens []model.Token) (map[string]float64, error) {
	balances := make(map[string]float64, 0)
	for _, token := range tokens {
		key := strings.ToUpper(token.Symbol)
		switch token.IsNative {
		case true:
			nativeCryptoBal, err := ch.Balance(address)
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

// Ver2
func (ch *Chain) TransferVer2(fromAcc accounts.Account, toAcc accounts.Account, amount float64, token model.Token, all bool) (*types.Transaction, float64, error) {
	var (
		t   *types.Transaction
		err error
	)

	switch token.IsNative {
	case true:
		t, amount, err = ch.transferVer2(fromAcc, toAcc, amount, all)
	default:
		t, amount, err = ch.transferTokenVer2(fromAcc, toAcc, amount, token, all)
	}

	return t, amount, err
}

func (ch *Chain) transferVer2(fromAcc accounts.Account, toAcc accounts.Account, amount float64, all bool) (*types.Transaction, float64, error) {
	balance, err := ch.Balance(fromAcc.Address.Hex())
	if err != nil {
		return nil, 0, err
	}
	if balance <= amount {
		return nil, 0, errors.New("Mochi fund's balance is not enough")
	}

	priv, _ := ch.hdwallet.PrivateKeyHex(fromAcc)
	privateKey, err := crypto.HexToECDSA(priv)
	if err != nil {
		return nil, 0, err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, 0, errors.New("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := ch.client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return nil, 0, err
	}

	gasLimit := uint64(21000)
	gasPrice, err := ch.client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, 0, err
	}

	maxTxFee := float64(gasPrice.Int64()) * float64(gasLimit) / float64(math.Pow10(18))
	// if all {
	// if balance <= maxTxFee {
	// 	return nil, 0, errors.New("insufficient funds for gas")
	// }
	amount -= maxTxFee
	// }

	value := new(big.Int)
	value.SetString(strconv.FormatFloat(float64(math.Pow10(18))*amount, 'f', 6, 64), 10)

	toAddress := common.HexToAddress(toAcc.Address.Hex())
	var data []byte
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)

	chainID, err := ch.client.NetworkID(context.Background())
	if err != nil {
		return nil, 0, err
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		return nil, 0, err
	}

	err = ch.client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return nil, 0, err
	}

	return signedTx, amount, nil
}

func (ch *Chain) transferTokenVer2(fromAcc accounts.Account, toAcc accounts.Account, amount float64, token model.Token, all bool) (*types.Transaction, float64, error) {
	priv, _ := ch.hdwallet.PrivateKeyHex(fromAcc)
	privateKey, err := crypto.HexToECDSA(priv)
	if err != nil {
		return nil, 0, err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, 0, errors.New("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := ch.client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return nil, 0, err
	}

	value := big.NewInt(0) // in wei (0 eth)
	gasPrice, err := ch.client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, 0, err
	}

	toAddress := common.HexToAddress(toAcc.Address.Hex())
	tokenAddress := common.HexToAddress(token.Address)

	transferFnSignature := []byte("transfer(address,uint256)")
	hash := sha3.NewLegacyKeccak256()
	hash.Write(transferFnSignature)
	methodID := hash.Sum(nil)[:4]

	paddedAddress := common.LeftPadBytes(toAddress.Bytes(), 32)

	tokenBalance, err := ch.scan.TokenBalance(token.Address, fromAcc.Address.Hex())
	if err != nil {
		return nil, 0, err
	}

	amt := new(big.Int)
	amt.SetString(strconv.FormatFloat(math.Pow10(token.Decimals)*amount, 'f', 6, 64), 10)

	if tokenBalance.Int().Cmp(amt) == -1 {
		return nil, 0, errors.New("balance is not enough")
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
		return nil, 0, err
	}
	gasLimit *= 3

	tx := types.NewTransaction(nonce, tokenAddress, value, gasLimit, gasPrice, data)

	chainID, err := ch.client.NetworkID(context.Background())
	if err != nil {
		return nil, 0, err
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		return nil, 0, err
	}

	err = ch.client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return nil, 0, err
	}

	return signedTx, amount, nil
}
