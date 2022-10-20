package trade_offer

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	GetOne(id string) (*model.TradeItem, error)
	UpsertOne(item *model.TradeItem) error
}
