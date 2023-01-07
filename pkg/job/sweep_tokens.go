package job

import (
	"github.com/defipod/mochi/pkg/config"
	"github.com/defipod/mochi/pkg/entities"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/service/abi"
)

type sweepTokens struct {
	entity *entities.Entity
	log    logger.Logger
}

func NewSweepTokens(e *entities.Entity, l logger.Logger) Job {
	return &sweepTokens{
		entity: e,
		log:    l,
	}
}

func (job *sweepTokens) Run() error {
	l := job.log
	cfg := config.LoadConfig(config.DefaultConfigLoaders())
	abi := abi.NewAbi(&cfg)
	isEVM := true
	supportDeposit := true
	contracts, err := job.entity.GetContracts(request.TipBotGetContractsRequest{SupportDeposit: &supportDeposit, IsEVM: &isEVM})
	if err != nil {
		l.Error(err, "[sweepTokens] job.entity.GetContracts() failed")
		return err
	}
	for _, contract := range contracts {
		log := l.Fields(logger.Fields{"contractID": contract.ID.String(), "address": contract.ContractAddress})
		log.Infof("[sweepTokens] start to sweep tokens")
		if contract.Chain == nil || contract.Chain.ChainID == nil {
			log.Info("[sweepTokens] undefined chain/chainID")
			continue
		}
		tokens, err := job.entity.GetTokensByChainID(*contract.Chain.ChainID)
		if err != nil {
			log.Errorf(err, "[sweepTokens] job.entity.GetTokensByChainID() failed: %d", *contract.Chain.ChainID)
			continue
		}
		if len(tokens) == 0 {
			log.Infof("[sweepTokens] chain %d has no tokens", *contract.Chain.ChainID)
			continue
		}
		for _, token := range tokens {
			tx, err := abi.SweepTokens(contract.ContractAddress, int64(*contract.Chain.ChainID), token)
			if err != nil {
				log.Error(err, "[sweepTokens] abi.SweepTokens() failed")
				continue
			}
			log.Infof("[sweepTokens] sweep tokens tx: %s", tx.Hash().Hex())
		}
	}

	l.Infof("sweepTokens finished")
	return nil
}
