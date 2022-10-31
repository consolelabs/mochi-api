package chain

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"math"
	"math/big"
	"strconv"
	"strings"
	"time"

	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/model"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/k0kubun/pp"
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
		Verbose: false,
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
		t, amount, err = ch.transfer(fromAcc, toAcc, amount, token, nonce, all)
	default:
		t, amount, err = ch.transferToken(fromAcc, toAcc, amount, token, nonce, all)
	}

	return t, amount, err
}

func (ch *Chain) transfer(fromAcc accounts.Account, toAcc accounts.Account, amount float64, token model.Token, prevTxNonce int, all bool) (*types.Transaction, float64, error) {
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

	status := false
	gl := 10000
	priceStep := 3000
	var signedTx *types.Transaction
	for !status {
		ch.log.Infof("gas limit", gl)
		gasPrice, err := ch.client.SuggestGasPrice(context.Background())
		if err != nil {
			return nil, 0, err
		}

		gasLimit := uint64(gl)
		maxTxFee := float64(gasPrice.Int64()) * float64(gasLimit) / float64(math.Pow10(18))

		if all {
			if balance <= maxTxFee {
				return nil, 0, errors.New("insufficient funds for gas")
			}
			amount = balance - maxTxFee
		}
		pp.Println("here 1")

		value := new(big.Int)
		value.SetString(strconv.FormatFloat(float64(math.Pow10(18))*amount, 'f', 6, 64), 10)
		pp.Println("here 2")
		toAddress := common.HexToAddress(toAcc.Address.Hex())
		var data []byte
		tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)
		pp.Println("here 3")
		chainID, err := ch.client.NetworkID(context.Background())
		if err != nil {
			return nil, 0, err
		}
		pp.Println("here 4")

		signedTx, err = types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
		if err != nil {
			return nil, 0, err
		}

		err = ch.client.SendTransaction(context.Background(), signedTx)
		if err != nil {
			fmt.Println(err)
			gl += priceStep
			continue
			// return nil, 0, err
		}
		pp.Println(signedTx.Hash().Hex())
		pp.Println("here 6")

		txDetails, isPending, err := ch.client.TransactionByHash(context.Background(), signedTx.Hash())
		if err != nil {
			return nil, 0, err
		}
		pp.Println(signedTx.Hash().Hex())
		pp.Println(txDetails.Hash().Hex())
		pp.Println("here 7")

		pp.Println("isPending: ", isPending)
		pp.Println("err: ", err)
		for isPending {
			_, isPending, err = ch.client.TransactionByHash(context.Background(), txDetails.Hash())
			pp.Println(err)
			if err != nil {
				return nil, 0, err
			}
		}

		receipt, err := ch.client.TransactionReceipt(context.Background(), signedTx.Hash())
		if err != nil || receipt.Status == 0 {
			gl += priceStep
			continue
		}

		status = true
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
	if all {
		amount, _ = new(big.Float).Quo(big.NewFloat(0).SetInt(tokenBalance.Int()), big.NewFloat(math.Pow10(18))).Float64()
	}

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

	// --------------------------------------------
	status := false
	gl := 50000
	priceStep := 3000
	i := 1
	var signedTx *types.Transaction
	pp.Println("here 1")
	for !status {
		ch.log.Infof("gas limit", gl)
		gasLimit := uint64(gl)

		pp.Println("here 2")
		tx := types.NewTransaction(nonce, tokenAddress, value, gasLimit, gasPrice, data)

		pp.Println("here 3")
		chainID, err := ch.client.NetworkID(context.Background())
		if err != nil {
			return nil, 0, err
		}
		pp.Println("here 4")

		signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
		if err != nil {
			return nil, 0, err
		}
		pp.Println("here 5")

		err = ch.client.SendTransaction(context.Background(), signedTx)

		if err != nil {
			pp.Println(err)
			gl += priceStep * i
			i++
			time.Sleep(1 * time.Second)
			continue
		}

		pp.Println(signedTx.Hash().Hex())
		pp.Println("here 6")

		txDetails, isPending, err := ch.client.TransactionByHash(context.Background(), signedTx.Hash())
		if err != nil && !strings.Contains(err.Error(), "not found") {
			return nil, 0, err
		}
		pp.Println("isPending: ", isPending)
		pp.Println("err: ", err)
		pp.Println("here 7")

		if txDetails == nil {
			txDetails = signedTx
		}

		pp.Println(txDetails.Hash().Hex())
		for isPending {
			_, isPending, err = ch.client.TransactionByHash(context.Background(), txDetails.Hash())
			if err != nil {
				if strings.Contains(err.Error(), "not found") {
					continue
				}
				pp.Println(err)
				return nil, 0, err
			}
			time.Sleep(1 * time.Second)
		}
		pp.Println("here 8")

		receipt, err := ch.client.TransactionReceipt(context.Background(), signedTx.Hash())
		pp.Println(receipt)
		if err != nil && !strings.Contains(err.Error(), "not found") {
			pp.Println(err)
			gl += priceStep * i
			continue
		}

		pp.Println("here 9")

		status = true
	}

	// -----------------------------------------------
	// tx := types.NewTransaction(nonce, tokenAddress, value, gasLimit, gasPrice, data)

	// chainID, err := ch.client.NetworkID(context.Background())
	// if err != nil {
	// 	return nil, 0, err
	// }

	// signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	// if err != nil {
	// 	return nil, 0, err
	// }

	// err = ch.client.SendTransaction(context.Background(), signedTx)
	// if err != nil {
	// 	return nil, 0, err
	// }

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
