package request

import "github.com/defipod/mochi/pkg/model"

type GetSwapRouteRequest struct {
	From      string `json:"from" binding:"required"`
	To        string `json:"to" binding:"required"`
	Amount    string `json:"amount" binding:"required"`
	ChainId   int64  `json:"chain_id" binding:"required"`
	ChainName string `json:"chain_name" binding:"required"`
}

type SwapRequest struct {
	ChainName    string             `json:"chainName" binding:"required"`
	Recipient    string             `json:"recipient" binding:"required"`
	Sender       string             `json:"sender" binding:"required"`
	RouteSummary model.RouteSummary `json:"routeSummary"`
}
