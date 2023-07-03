package chain

import (
	"context"
	"encoding/json"
	"math/big"

	"github.com/portto/solana-go-sdk/client"
	"github.com/portto/solana-go-sdk/common"
	"github.com/portto/solana-go-sdk/program/sysprog"
	"github.com/portto/solana-go-sdk/rpc"
	"github.com/portto/solana-go-sdk/types"

	"github.com/defipod/mochi/pkg/cache"
	"github.com/defipod/mochi/pkg/config"
	"github.com/defipod/mochi/pkg/logger"
)

type Solana struct {
	config *config.Config
	client *client.Client
	logger logger.Logger
	cache  cache.Cache
}

func NewSolanaClient(cfg *config.Config, l logger.Logger, cache cache.Cache) *Solana {
	return &Solana{
		config: cfg,
		logger: l,
		client: client.NewClient(rpc.MainnetRPCEndpoint),
		cache:  cache,
	}
}

var (
	solanaBalanceKey = "solana-balance"
)

func (s *Solana) Balance(address string) (float64, error) {
	s.logger.Debug("start Solana.Balance()")
	defer s.logger.Debug("end Solana.Balance()")

	var res float64
	// check if data cached

	cached, err := s.doCacheBalance(address)
	if err == nil && cached != "" {
		s.logger.Infof("hit cache data solana, address: %s", address)
		go s.doNetworkBalance(address)
		return res, json.Unmarshal([]byte(cached), &res)
	}

	// call network
	return s.doNetworkBalance(address)
}

func (s *Solana) GetTokenBalance(walletAddress, tokenAddress string) (*big.Int, error) {
	balances, err := s.client.GetTokenAccountsByOwner(context.Background(), walletAddress)
	if err != nil {
		s.logger.Fields(logger.Fields{"walletAddress": walletAddress}).Error(err, "[solana.Balance] client.GetBalance() failed")
	}
	var bal uint64
	for _, v := range balances {
		if v.Mint == common.PublicKeyFromString(tokenAddress) {
			bal = v.Amount
			break
		}
	}
	return new(big.Int).SetUint64(bal), err
}

func (s *Solana) Transfer(senderPK, recipientAddr string, amount float64, all bool) (string, float64, error) {
	var txHash string
	fromAcc, _ := types.AccountFromBase58(senderPK)

	// to fetch recent blockhash
	res, err := s.client.GetLatestBlockhash(context.Background())
	if err != nil {
		s.logger.Error(err, "[solana.Transfer] client.GetLatestBlockhash() failed")
		return txHash, 0, err
	}

	// create a message
	message := types.NewMessage(types.NewMessageParam{
		FeePayer:        fromAcc.PublicKey,
		RecentBlockhash: res.Blockhash,
		Instructions: []types.Instruction{
			sysprog.Transfer(sysprog.TransferParam{
				From:   fromAcc.PublicKey,
				To:     common.PublicKeyFromString(recipientAddr),
				Amount: uint64(amount * 1e9),
			}),
		},
	})

	// handle case all
	if all {
		amount, err = s.getTransferAllAmount(fromAcc, message)
		if err != nil {
			s.logger.Error(err, "[solana.Transfer] s.getTransferAllAmount() failed")
			return txHash, 0, err
		}
		message = types.NewMessage(types.NewMessageParam{
			FeePayer:        fromAcc.PublicKey,
			RecentBlockhash: res.Blockhash,
			Instructions: []types.Instruction{
				sysprog.Transfer(sysprog.TransferParam{
					From:   fromAcc.PublicKey,
					To:     common.PublicKeyFromString(recipientAddr),
					Amount: uint64(amount),
				}),
			},
		})
	}

	// create tx by message + signer
	tx, err := types.NewTransaction(types.NewTransactionParam{
		Message: message,
		Signers: []types.Account{fromAcc},
	})
	if err != nil {
		s.logger.Error(err, "[solana.Transfer] types.NewTransaction() failed")
		return txHash, 0, err
	}

	// send tx
	txHash, err = s.client.SendTransaction(context.Background(), tx)
	if err != nil {
		s.logger.Error(err, "[solana.Transfer] s.client.SendTransaction() failed")
		return txHash, 0, err
	}
	return txHash, amount, nil
}

func (s *Solana) getTransferAllAmount(fromAcc types.Account, msg types.Message) (float64, error) {
	senderBal, err := s.Balance(fromAcc.PublicKey.String())
	if err != nil {
		s.logger.Error(err, "[solana.Transfer] s.Balance() failed")
		return 0, err
	}
	amount := senderBal * 1e9
	// estimate fee
	estimatedFee, err := s.client.GetFeeForMessage(context.Background(), msg)
	if err != nil {
		s.logger.Error(err, "[solana.Transfer] s.client.GetFeeForMessage() failed")
		return 0, err
	}
	return amount - float64(*estimatedFee), nil
}

func (s *Solana) GetCentralizedWalletAddress() string {
	centralizedAccount, _ := types.AccountFromBase58(s.config.SolanaCentralizedWalletPrivateKey)
	return string(centralizedAccount.PublicKey.String())
}
