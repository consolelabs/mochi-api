package onchainassetavgcost

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	Upsert(asset *model.OnchainAssetAvgCost) error
	UpsertMany(assets []model.OnchainAssetAvgCost) error
	GetByWalletAddr(walletAddr string) ([]model.OnchainAssetAvgCost, error)
}
