package job

import (
	"fmt"
	"strings"
	"time"

	"github.com/defipod/mochi/pkg/entities"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/service"
	"github.com/defipod/mochi/pkg/service/sentrygo"
)

type updateUserOnchainAssetAvgCostJob struct {
	entity  *entities.Entity
	service *service.Service
	log     logger.Logger
}

func NewUpdateUserOnchainAssetAvgCostJob(e *entities.Entity) Job {
	return &updateUserOnchainAssetAvgCostJob{
		entity:  e,
		service: e.GetSvc(),
		log:     e.GetLogger(),
	}
}

func (j *updateUserOnchainAssetAvgCostJob) Run() error {
	j.log.Info("update_user_asset_average_cost job started")

	// 1.0 Get top 50 active evm addresses
	evmAddrs, err := j.getTop50ActiveEvmAddresses()
	if err != nil {
		j.log.Error(err, "failed to get top 50 active evm addresses")
		j.logSentry(err, map[string]interface{}{"task": "getTop50ActiveEvmAddresses"})
		return err
	}

	// 2.0 Call dune to get the average cost of the assets of each address
	avgCosts, err := j.getEvmAssetAvgCosts(evmAddrs)
	if err != nil {
		j.log.Error(err, "failed to get avg cost of evm assets")
		j.logSentry(err, map[string]interface{}{"task": "getEvmAssetAvgCosts"})
		return err
	}

	// 3.0 Update the average cost of the assets of each address in the database
	if err := j.updateEvmAssetAvgCosts(avgCosts); err != nil {
		j.log.Error(err, "failed to update avg cost of evm assets")
		j.logSentry(err, map[string]interface{}{"task": "updateEvmAssetAvgCosts"})
		return err
	}

	j.log.Info("update_user_asset_average_cost job finished")
	return nil
}

func (j *updateUserOnchainAssetAvgCostJob) getTop50ActiveEvmAddresses() ([]string, error) {
	profiles, err := j.service.MochiProfile.GetTopActiveUsers(50)
	if err != nil {
		return nil, err
	}
	evmAddrs := make([]string, 0)
	for _, p := range profiles {
		for _, a := range p.AssociatedAccounts {
			if a.Platform == "evm-chain" {
				evmAddrs = append(evmAddrs, a.PlatformIdentifier)
			}
		}
	}
	return evmAddrs, nil
}

func (j *updateUserOnchainAssetAvgCostJob) getEvmAssetAvgCosts(evmAddrs []string) ([]model.OnchainAssetAvgCost, error) {
	avgCosts := make([]model.OnchainAssetAvgCost, 0)

	// Call dune to get the average cost of the assets of each address
	j.log.Info("executing query to get average cost of assets of each address")
	var queryId int64 = 3782999 // https://dune.com/queries/3782999
	params := make(map[string]interface{})
	params["wallet_addresses"] = strings.Join(evmAddrs, ",")
	executeQueryResp, err := j.service.Dune.ExecuteQuery(queryId, map[string]interface{}{})
	if err != nil {
		return nil, err
	}
	executionId := executeQueryResp.ExecutionId
	j.log.Infof("finished execute query, executeId: %s", executionId)

	// Execute query, wait for the query execution to finish
	j.log.Info("waiting for query execution to finish")

	for {
		time.Sleep(30 * time.Second)

		execStatus, err := j.service.Dune.GetExecutionStatus(executionId)
		if err != nil {
			return nil, fmt.Errorf("failed to get execution status: %w", err)
		}

		// If the query execution is finished, break the loop
		if execStatus.IsExecutionFinished {
			break
		}
	}

	j.log.Info("execution finished, start getting execution result")

	// Get the execution result

	var (
		limit  int64 = 500
		offset int64 = 0
	)
	for {
		res, err := j.service.Dune.GetExecutionResult(executionId, limit, offset)
		if err != nil {
			return nil, fmt.Errorf("failed to get execution result: %w", err)
		}
		for _, row := range res.Result.Rows {
			// Parse the data
			walletAddr, wOk := row["wallet_address"]
			blockchain, bOk := row["blockchain"]
			tokenAddress, tOk := row["token_address"]
			symbol, sOk := row["symbol"]
			averageCost, cOk := row["average_cost"]
			if !wOk || !bOk || !tOk || !sOk || !cOk {
				continue
			}

			// Convert the data to the correct type
			walletAddrStr, wOk := walletAddr.(string)
			blockchainStr, bOk := blockchain.(string)
			tokenAddressStr, tOk := tokenAddress.(string)
			symbolStr, sOk := symbol.(string)
			averageCostFloat, cOk := averageCost.(float64)
			if !wOk || !bOk || !tOk || !sOk || !cOk {
				continue
			}

			avgCosts = append(avgCosts, model.OnchainAssetAvgCost{
				WalletAddress: walletAddrStr,
				Blockchain:    blockchainStr,
				TokenAddress:  tokenAddressStr,
				Symbol:        symbolStr,
				AverageCost:   averageCostFloat,
			})
		}

		// If there is no more data, break the loop
		if res.NextOffset == nil || *res.NextOffset == 0 {
			break
		}
		// Otherwise update the offset
		offset = *res.NextOffset
	}

	j.log.Info("finished getting execution result")
	return avgCosts, nil
}

func (j *updateUserOnchainAssetAvgCostJob) updateEvmAssetAvgCosts(assets []model.OnchainAssetAvgCost) error {
	if err := j.entity.UpsertManyOnchainAssetAvgCost(assets); err != nil {
		return err
	}

	return nil
}

func (j *updateUserOnchainAssetAvgCostJob) logSentry(err error, extra map[string]interface{}) {
	sentryTags := map[string]string{
		"type": "system",
	}
	j.service.Sentry.CaptureErrorEvent(sentrygo.SentryCapturePayload{
		Message: fmt.Sprintf("[CJ prod mochi] - update_user_asset_average_cost failed - %v", err),
		Tags:    sentryTags,
		Extra:   extra,
	})
}
