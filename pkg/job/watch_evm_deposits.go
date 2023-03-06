package job

import (
	"errors"
	"fmt"
	"math"
	"math/big"
	"sort"
	"strings"

	"github.com/defipod/mochi/pkg/config"
	"github.com/defipod/mochi/pkg/entities"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/service/covalent"
)

type watchEvmDeposits struct {
	entity *entities.Entity
	log    logger.Logger
}

func NewWatchEvmDeposits(e *entities.Entity, l logger.Logger) Job {
	return &watchEvmDeposits{
		entity: e,
		log:    l,
	}
}

func (job *watchEvmDeposits) Run() error {
	l := job.log
	cfg := config.LoadConfig(config.DefaultConfigLoaders())
	covalentSvc := covalent.NewService(&cfg, l)
	isEvm := true
	supportDeposit := true

	contracts, err := job.entity.GetContracts(request.TipBotGetContractsRequest{
		IsEVM:          &isEvm,
		SupportDeposit: &supportDeposit,
	})
	if err != nil {
		l.Error(err, "[watchEvmDeposits] job.entity.GetContracts() failed")
		return err
	}
	for _, contract := range contracts {
		log := l.Fields(logger.Fields{"contractID": contract.ID.String(), "address": contract.ContractAddress})
		log.Infof("[watchEvmDeposits] start watching contract")
		if contract.Chain == nil || contract.Chain.ChainID == nil {
			log.Info("[watchEvmDeposits] no chainID")
			continue
		}
		chainID := *contract.Chain.ChainID
		contractAddr := contract.ContractAddress
		transactionsRes, err := covalentSvc.GetTransactionsByAddress(chainID, contractAddr, 1000, 3)
		if err != nil {
			log.Error(err, "[watchEvmDeposits] covalentSvc.GetTransactionsByAddress() failed")
			continue
		}

		var newTxs []covalent.TransactionItemData
		for _, item := range transactionsRes.Data.Items {
			if !item.Successful {
				continue
			}
			// if tx already existed (handled) -> skip
			_, err := job.entity.GetOneDepositTx(contract.ChainID.String(), item.TxHash)
			if err == nil {
				continue
			}
			isDeposit, err := job.getTxDetails(&item, contractAddr)
			if err != nil {
				log.Error(err, "[watchEvmDeposits] getTxDetails() failed")
				continue
			}
			if !isDeposit {
				l.Fields(logger.Fields{"txHash": item.TxHash}).Error(err, "[watchEvmDeposits] not deposit transaction")
				continue
			}
			newTxs = append(newTxs, item)
		}

		// handle transactions from old -> new
		sort.Slice(newTxs, func(i, j int) bool {
			return newTxs[i].BlockSignedAt.Unix() < newTxs[j].BlockSignedAt.Unix()
		})

		l.Fields(logger.Fields{"contract": contractAddr}).Infof("[watchEvmDeposits] found %d new deposit transactions on chain %d", len(newTxs), chainID)
		for _, newTx := range newTxs {
			l.Fields(logger.Fields{"tx": newTx, "chain": chainID}).Info("[watchEvmDeposits] detect new deposit tx")
			req := request.TipBotDepositRequest{
				ChainID:       *contract.Chain.ChainID,
				FromAddress:   newTx.FromAddress,
				ToAddress:     contractAddr,
				TokenSymbol:   newTx.TokenSymbol,
				TokenContract: newTx.TokenContract,
				Amount:        newTx.Amount,
				TxHash:        newTx.TxHash,
				BlockNumber:   int64(newTx.BlockHeight),
				SignedAt:      newTx.BlockSignedAt,
			}
			err := job.entity.HandleIncomingDeposit(req)
			if err != nil {
				l.Fields(logger.Fields{"req": req}).Error(err, "[watchEvmDeposits] job.entity.HandleIncomingDeposit() failed")
				continue
			}
		}
	}

	l.Infof("watchEvmDeposits finished")
	return nil
}

func (job *watchEvmDeposits) getTxDetails(tx *covalent.TransactionItemData, contractAddress string) (bool, error) {
	l := job.log.Fields(logger.Fields{"txHash": tx.TxHash})
	// covalent address returned in lower case
	if strings.EqualFold(tx.ToAddress, contractAddress) {
		amount := new(big.Int)
		amount, ok := amount.SetString(tx.Value, 10)
		if !ok {
			err := fmt.Errorf("invalid tx amount %s", tx.Value)
			l.Error(err, "[getTxDetails] invalid native amount")
			return false, err
		}
		if amount.Cmp(big.NewInt(0)) == 0 {
			err := errors.New("zero tx amount")
			l.Error(err, "[getTxDetails] zero transaction amount")
			return false, nil
		}
		tx.Amount, _ = new(big.Float).SetInt(amount).Float64()
		return true, nil
	}

	// case erc20 token
	if tx.LogEvents == nil || len(tx.LogEvents) == 0 {
		l.Info("[getTxDetails] not transfer transaction")
		return false, nil
	}

	var isDeposit bool
	for _, event := range tx.LogEvents {
		if event.Decoded.Params == nil || len(event.Decoded.Params) == 0 {
			l.Info("[getTxDetails] no event params")
			continue
		}
		if !strings.EqualFold(event.Decoded.Signature, "Transfer(indexed address from, indexed address to, uint256 value)") {
			l.Info("[getTxDetails] not erc20 transfer tx")
			continue
		}
		to := event.Decoded.Params[1].Value.(string)
		value := event.Decoded.Params[2].Value.(string)
		if !strings.EqualFold(to, contractAddress) {
			continue
		}
		amount, ok := new(big.Int).SetString(value, 10)
		if !ok {
			err := fmt.Errorf("invalid tx amount %s", value)
			l.Error(err, "[getTxDetails] invalid erc20 transfer amount")
			continue
		}
		// assign deposit details
		tx.TokenSymbol = event.SenderContractTickerSymbol
		tx.ToAddress = to
		tx.TokenContract = event.SenderAddress
		tx.Amount, _ = new(big.Float).SetInt(amount).Float64()
		tx.Amount /= math.Pow10(event.SenderContractDecimals)
		isDeposit = true
		break
	}
	return isDeposit, nil
}
