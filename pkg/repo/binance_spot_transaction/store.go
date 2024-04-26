package binancespottransaction

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	Create(tx *model.BinanceSpotTransaction) error
	List(q ListQuery) ([]model.BinanceSpotTransaction, error)
}
