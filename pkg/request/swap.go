package request

type GetSwapRouteRequest struct {
	From   string `json:"from" binding:"required"`
	To     string `json:"to" binding:"required"`
	Amount string `json:"amount" binding:"required"`
}
