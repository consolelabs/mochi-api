package entities

import (
	"github.com/defipod/mochi/pkg/model"
)

func (e *Entity) UpsertManyOnchainAssetAvgCost(assets []model.OnchainAssetAvgCost) error {
	err := e.repo.OnchainAssetAverageCost.UpsertMany(assets)
	if err != nil {
		e.log.Error(err, "[entities.UpsertManyOnchainAssetAvgCost] - repo.OnchainAssetAverageCost.UpsertMany failed")
		return err
	}
	return nil
}

func (e *Entity) GetOnchainAssetAvgCostByWalletAddress(walletAddress string) ([]model.OnchainAssetAvgCost, error) {
	assets, err := e.repo.OnchainAssetAverageCost.GetByWalletAddr(walletAddress)
	if err != nil {
		e.log.Error(err, "[entities.GetOnchainAssetAvgCostByWalletAddress] - repo.OnchainAssetAverageCost.GetByWalletAddress failed")
		return nil, err
	}
	return assets, nil
}
