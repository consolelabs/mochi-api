package entities

import (
	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/request"
)

func (e *Entity) CreateTradeOffer(req request.CreateTradeOfferRequest) (*model.TradeOffer, error) {
	haveItems := make([]model.TradeItem, 0)
	for _, item := range req.HaveItems {
		haveItems = append(haveItems, model.TradeItem{
			TokenAddress: item.TokenAddress,
			TokenIds:     item.TokenIds,
			IsFrom:       true,
		})
	}

	wantItems := make([]model.TradeItem, 0)
	for _, item := range req.WantItems {
		wantItems = append(wantItems, model.TradeItem{
			TokenAddress: item.TokenAddress,
			TokenIds:     item.TokenIds,
		})
	}
	offer := &model.TradeOffer{
		OwnerAddress: req.OwnerAddress,
		HaveItems:    haveItems,
		WantItems:    wantItems,
	}
	return e.repo.TradeOffer.Create(offer)
}

func (e *Entity) GetTradeOffer(id string) (*model.TradeOffer, error) {
	return e.repo.TradeOffer.GetOne(id)
}
