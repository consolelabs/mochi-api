package request

type CreateTradeOfferRequest struct {
	FromAddress string           `json:"from_address" form:"from_address" binding:"required"`
	ToAddress   string           `json:"to_address" form:"to_address" binding:"required"`
	FromItems   []TradeOfferItem `json:"from_items" form:"form_items"`
	ToItems     []TradeOfferItem `json:"to_items" form:"to_items"`
}

type TradeOfferItem struct {
	TokenAddress string   `json:"token_address" form:"token_address" binding:"required"`
	TokenIds     []string `json:"token_ids" form:"token_address" binding:"required"`
}
