package job

import (
	"math"
	"sort"
	"strconv"
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
		log := l.Fields(logger.Fields{"contractID": contract.ID.String()})
		log.Infof("[watchEvmDeposits] start watching contract")
		if contract.Chain == nil || contract.Chain.ChainID == nil {
			log.Info("[watchEvmDeposits] no chainID")
			continue
		}
		transactionsRes, err := covalentSvc.GetTransactionsByAddress(*contract.Chain.ChainID, contract.ContractAddress)
		if err != nil {
			log.Error(err, "[watchEvmDeposits] covalentSvc.GetTransactionsByAddress() failed")
			continue
		}
		latestDeposit, err := job.entity.GetLatestDepositTx(request.GetLatestDepositRequest{
			ChainID:         contract.ChainID.String(),
			ContractAddress: contract.ContractAddress,
		})
		if err != nil {
			log.Error(err, "[watchEvmDeposits] job.entity.GetLatestDepositTx() failed")
			continue
		}

		var newTxs []covalent.TransactionItemData
		for _, item := range transactionsRes.Data.Items {
			if !item.Successful {
				continue
			}
			if latestDeposit.SignedAt.Unix() >= item.BlockSignedAt.Unix() {
				continue
			}
			if latestDeposit.TxHash == item.TxHash {
				continue
			}
			isDeposit, err := job.getTxDetails(&item, contract.ContractAddress)
			if err != nil {
				log.Error(err, "[watchEvmDeposits] getTxDetails() failed")
				return err
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

		for _, newTx := range newTxs {
			req := request.TipBotDepositRequest{
				ChainID:       *contract.Chain.ChainID,
				FromAddress:   newTx.FromAddress,
				ToAddress:     newTx.ToAddress,
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
				break
			}
		}
	}

	l.Infof("watchEvmDeposits finished")
	return nil
}

func (job *watchEvmDeposits) getTxDetails(tx *covalent.TransactionItemData, contractAddress string) (bool, error) {
	l := job.log.Fields(logger.Fields{"txHash": tx.TxHash})
	if tx.ToAddress == contractAddress {
		amount, err := strconv.Atoi(tx.Value)
		if err != nil {
			l.Error(err, "[getTxDetails] strconv.Atoi() failed: invalid native amount")
			return false, err
		}
		if amount == 0 {
			l.Error(err, "[getTxDetails] zero amount transaction")
			return false, nil
		}
		tx.Amount = float64(amount)
		return true, nil
	}

	// case erc20 token
	if tx.LogEvents == nil || len(tx.LogEvents) == 0 || !strings.EqualFold(tx.LogEvents[0].Decoded.Name, "Transfer") {
		l.Info("[getTxDetails] not transfer transaction")
		return false, nil
	}
	event := tx.LogEvents[0]
	if event.Decoded.Params == nil || len(event.Decoded.Params) == 0 {
		l.Info("[getTxDetails] no event params")
		return false, nil
	}
	decimals := event.SenderContractDecimals
	tx.TokenSymbol = event.SenderContractTickerSymbol
	for _, p := range event.Decoded.Params {
		if strings.EqualFold(p.Name, "to") && p.Value != contractAddress {
			l.Info("[getTxDetails] different recipient address")
			return false, nil
		}
		if strings.EqualFold(p.Name, "value") {
			amount, err := strconv.Atoi(p.Value)
			if err != nil {
				l.Error(err, "[getTxDetails] strconv.Atoi() failed: invalid erc20 amount")
				return false, err
			}
			tx.ToAddress = contractAddress
			tx.TokenContract = event.SenderAddress
			tx.Amount = float64(amount) / math.Pow10(decimals)
		}
	}
	return true, nil
}
