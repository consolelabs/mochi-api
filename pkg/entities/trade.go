package entities

import (
	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/request"
)

func (e *Entity) CreateTradeOffer(req request.CreateTradeOfferRequest) (*model.TradeOffer, error) {
	fromItems := make([]model.TradeItem, 0)
	for _, item := range req.FromItems {
		fromItems = append(fromItems, model.TradeItem{
			TokenAddress: item.TokenAddress,
			TokenIds:     item.TokenIds,
			IsFrom:       true,
		})
	}

	toItems := make([]model.TradeItem, 0)
	for _, item := range req.ToItems {
		toItems = append(toItems, model.TradeItem{
			TokenAddress: item.TokenAddress,
			TokenIds:     item.TokenIds,
		})
	}
	offer := &model.TradeOffer{
		FromAddress: req.FromAddress,
		ToAddress:   req.ToAddress,
		FromItems:   fromItems,
		ToItems:     toItems,
	}
	return e.repo.TradeOffer.Create(offer)
}

func (e *Entity) GetTradeOffer(id string) (*model.TradeOffer, error) {
	return e.repo.TradeOffer.GetOne(id)
}
