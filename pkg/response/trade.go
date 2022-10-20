package response

import "github.com/defipod/mochi/pkg/model"

// For swagger
type CreateTradeOfferResponse struct {
	Data *model.TradeOffer `json:"data"`
}

type GetTradeOfferResponse struct {
	Data *model.TradeOffer `json:"data"`
}
