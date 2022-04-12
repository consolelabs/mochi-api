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

	"github.com/defipod/mochi/pkg/config"
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

type FTM struct {
	client   *ethclient.Client
	scan     *etherscan.Client
	hdwallet *hdwallet.Wallet
}

func NewFTMClient(cfg config.Config, hdwallet *hdwallet.Wallet) (Chain, error) {
	client, err := ethclient.Dial(cfg.FantomRPC)
	if err != nil {
		return nil, err
	}

	scan := etherscan.NewCustomized(etherscan.Customization{
		Timeout: 15 * time.Second,
		Key:     cfg.FantomScanAPIKey,
		BaseURL: cfg.FantomScan,
		Verbose: true,
	})

	return &FTM{
		client:   client,
		hdwallet: hdwallet,
		scan:     scan,
	}, nil
}

func (ch *FTM) Balance(address string) (float64, error) {
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

func (ch *FTM) Transfer(fromAcc accounts.Account, toAcc accounts.Account, amount float64, token model.Token, nonce int) (*types.Transaction, error) {
	var (
		t   *types.Transaction
		err error
	)

	switch strings.ToUpper(token.Symbol) {
	case "FTM":
		t, err = ch.transfer(fromAcc, toAcc, amount, nonce)
	default:
		t, err = ch.transferToken(fromAcc, toAcc, amount, token, nonce)
	}

	return t, err
}

func (ch *FTM) transfer(fromAcc accounts.Account, toAcc accounts.Account, amount float64, prevTxNonce int) (*types.Transaction, error) {
	balance, err := ch.Balance(fromAcc.Address.Hex())
	if err != nil {
		return nil, err
	}
	if balance < amount {
		return nil, errors.New("balance is not enough")
	}

	priv, _ := ch.hdwallet.PrivateKeyHex(fromAcc)
	privateKey, err := crypto.HexToECDSA(priv)
	if err != nil {
		return nil, err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, errors.New("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce := uint64(prevTxNonce)
	if prevTxNonce < 0 {
		nonce, err = ch.client.PendingNonceAt(context.Background(), fromAddress)
		if err != nil {
			return nil, err
		}
	}

	value := new(big.Int)
	value.SetString(strconv.FormatFloat(float64(math.Pow10(18))*amount, 'f', 6, 64), 10)

	gasLimit := uint64(21000) // in units
	gasPrice, err := ch.client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}

	toAddress := common.HexToAddress(toAcc.Address.Hex())
	var data []byte
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)

	chainID, err := ch.client.NetworkID(context.Background())
	if err != nil {
		return nil, err
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		return nil, err
	}

	err = ch.client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return nil, err
	}

	return signedTx, nil
}

func (ch *FTM) transferToken(fromAcc accounts.Account, toAcc accounts.Account, amount float64, token model.Token, prevTxNonce int) (*types.Transaction, error) {
	priv, _ := ch.hdwallet.PrivateKeyHex(fromAcc)
	privateKey, err := crypto.HexToECDSA(priv)
	if err != nil {
		return nil, err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, errors.New("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce := uint64(prevTxNonce)
	if prevTxNonce < 0 {
		nonce, err = ch.client.PendingNonceAt(context.Background(), fromAddress)
		if err != nil {
			return nil, err
		}
	}

	value := big.NewInt(0) // in wei (0 eth)
	gasPrice, err := ch.client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}

	toAddress := common.HexToAddress(toAcc.Address.Hex())
	tokenAddress := common.HexToAddress(token.Address)

	transferFnSignature := []byte("transfer(address,uint256)")
	hash := sha3.NewLegacyKeccak256()
	hash.Write(transferFnSignature)
	methodID := hash.Sum(nil)[:4]

	paddedAddress := common.LeftPadBytes(toAddress.Bytes(), 32)

	amt := new(big.Int)
	amt.SetString(strconv.FormatFloat(float64(math.Pow10(token.Decimals))*amount, 'f', 6, 64), 10)

	tokenBalance, err := ch.scan.TokenBalance(token.Address, fromAcc.Address.Hex())
	if err != nil {
		return nil, err
	}
	if tokenBalance.Int().Cmp(amt) == -1 {
		return nil, errors.New("balance is not enough")
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
		return nil, err
	}
	gasLimit *= 2

	tx := types.NewTransaction(nonce, tokenAddress, value, gasLimit, gasPrice, data)

	chainID, err := ch.client.NetworkID(context.Background())
	if err != nil {
		return nil, err
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		return nil, err
	}

	err = ch.client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return nil, err
	}

	return signedTx, nil
}

func (ch *FTM) erc20TokenBalance(address string, token model.Token) (float64, error) {
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

func (ch *FTM) Balances(address string, tokens []model.Token) (map[string]float64, error) {
	balances := make(map[string]float64, 0)
	for _, token := range tokens {
		key := strings.ToUpper(token.Symbol)
		switch key {
		case "FTM":
			ftmBalance, err := ch.Balance(address)
			if err != nil {
				return nil, err
			}
			balances[key] = ftmBalance
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
