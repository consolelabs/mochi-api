package request

type CreateTradeOfferRequest struct {
	OwnerAddress string           `json:"owner_address" form:"owner_address" binding:"required"`
	HaveItems    []TradeOfferItem `json:"have_items" form:"have_items"`
	WantItems    []TradeOfferItem `json:"want_items" form:"want_items"`
}

type TradeOfferItem struct {
	TokenAddress string   `json:"token_address" form:"token_address" binding:"required"`
	TokenIds     []string `json:"token_ids" form:"token_ids" binding:"required"`
}
