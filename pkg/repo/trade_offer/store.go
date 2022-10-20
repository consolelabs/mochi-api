package trade_offer

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	GetOne(id string) (*model.TradeOffer, error)
	Create(offer *model.TradeOffer) (*model.TradeOffer, error)
}
