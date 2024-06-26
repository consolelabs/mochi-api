package binancespottransaction

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	Create(tx *model.BinanceSpotTransaction) error
	List(q ListQuery) ([]model.BinanceSpotTransaction, error)
	GetUserAverageCost(profileId string) ([]model.BinanceAssetAverageCost, error)
	Update(tx *model.BinanceSpotTransaction) error
}
