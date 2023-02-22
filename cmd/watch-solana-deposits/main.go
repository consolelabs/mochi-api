package main

import (
	"context"
	"math"
	"sync"
	"time"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/gagliardetto/solana-go/rpc/ws"
	"github.com/portto/solana-go-sdk/client"

	"github.com/defipod/mochi/pkg/config"
	"github.com/defipod/mochi/pkg/entities"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/request"
)

func main() {
	cfg := config.LoadConfig(config.DefaultConfigLoaders())
	l := logger.NewLogrusLogger()
	err := entities.Init(cfg, l)
	if err != nil {
		l.Error(err, "entities.Init() failed")
		return
	}
	err = watchSolanaDeposits(cfg, l)
	if err != nil {
		l.Error(err, "watchSolanaDeposits() failed")
		return
	}
}

func watchSolanaDeposits(cfg config.Config, l logger.Logger) error {
	e := entities.Get()
	wsClient, err := ws.Connect(context.Background(), rpc.MainNetBeta_WS)
	if err != nil {
		l.Error(err, "[watchSolanaDeposits] ws.Connect() failed")
		return err
	}
	// TODO: need refactor this
	// for now we hard set solana chain ID
	supportDeposit := true
	contracts, err := e.GetContracts(request.TipBotGetContractsRequest{
		ChainID:        "f26c41dd-7625-4049-b886-8fa23424a37b",
		SupportDeposit: &supportDeposit,
	})
	if err != nil {
		l.Error(err, "[watchSolanaDeposits] e.GetContracts() failed")
		return err
	}
	wg := sync.WaitGroup{}
	wg.Add(len(contracts))
	for _, contract := range contracts {
		go func(contract model.OffchainTipBotContract) error {
			log := l.Fields(logger.Fields{"address": contract.ContractAddress})
			log.Infof("[watchSolanaDeposits] start watching contract")
			program := solana.MustPublicKeyFromBase58(contract.ContractAddress) // serum
			{
				// Subscribe to log events that mention the provided pubkey:
				sub, err := wsClient.LogsSubscribeMentions(
					program,
					rpc.CommitmentFinalized,
				)
				if err != nil {
					log.Error(err, "[watchSolanaDeposits] wsClient.LogsSubscribeMentions()")
					return err
				}
				defer sub.Unsubscribe()

				for {
					got, err := sub.Recv()
					if err != nil {
						log.Error(err, "[watchSolanaDeposits] sub.Recv() failed")
						return err
					}
					signature := got.Value.Signature.String()
					log.Infof("[watchSolanaDeposits] receive signature: %s", signature)
					req, err := getDepositRequest(contract.ContractAddress, signature, 0)
					if err != nil {
						l.Fields(logger.Fields{"address": contract.ContractAddress, "signature": signature}).Error(err, "[watchSolanaDeposits] getDepositRequest() failed")
						continue
					}
					if req == nil {
						continue
					}
					e.HandleIncomingDeposit(*req)
					if err != nil {
						l.Fields(logger.Fields{"req": req}).Error(err, "[watchSolanaDeposits] e.HandleIncomingDeposit() failed")
						continue
					}
				}
			}
		}(contract)
	}
	wg.Wait()
	return nil
}

func getDepositRequest(hostAddr, signature string, retry int) (*request.TipBotDepositRequest, error) {
	l := logger.NewLogrusLogger()
	client := client.NewClient(rpc.MainNetBeta_RPC)
	tx, err := client.GetTransaction(context.Background(), signature)
	log := l.Fields(logger.Fields{"signature": signature, "retry": retry})
	if err != nil {
		log.Error(err, "[getDepositRequest] client.GetTransaction() failed")
		if retry > 3 {
			return nil, err
		}
		time.Sleep(1000)
		return getDepositRequest(hostAddr, signature, retry+1)
	}
	if tx == nil || tx.Meta == nil || tx.Meta.PreBalances == nil || tx.Meta.PostBalances == nil {
		log.Info("[getDepositRequest] no transaction balances data")
		return nil, nil
	}
	if tx.Transaction.Message.Accounts == nil || len(tx.Transaction.Message.Accounts) == 0 {
		log.Info("[getDepositRequest] no transaction accounts data")
		return nil, nil
	}
	accountBalances := make(map[string][2]int64) // value is a slice of 2 int64 [pre-balance, post-balance]
	for i, acc := range tx.Transaction.Message.Accounts {
		accountBalances[acc.String()] = [2]int64{tx.Meta.PreBalances[i], tx.Meta.PostBalances[i]}
	}
	// if pre-balance < post-balance => deposit transaction
	if accountBalances[hostAddr][0] < accountBalances[hostAddr][1] {
		return &request.TipBotDepositRequest{
			ChainID:     1399811149,
			FromAddress: tx.Transaction.Message.Accounts[0].String(),
			ToAddress:   hostAddr,
			Amount:      float64(accountBalances[hostAddr][1]-accountBalances[hostAddr][0]) / math.Pow10(9),
			TokenSymbol: "SOL",
			TxHash:      signature,
			BlockNumber: int64(tx.Slot),
			SignedAt:    time.Unix(*tx.BlockTime, 0),
		}, nil
	}
	log.Info("[getDepositRequest] not a deposit transaction")
	return nil, nil
}
