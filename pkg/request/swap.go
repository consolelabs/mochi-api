package request

import "github.com/defipod/mochi/pkg/model"

type GetSwapRouteRequest struct {
	From        string `json:"from" binding:"required"`
	To          string `json:"to" binding:"required"`
	Amount      string `json:"amount" binding:"required"`
	ChainId     int64  `json:"chain_id" binding:"required"`
	ChainName   string `json:"chain_name" binding:"required"`
	FromTokenId string `json:"from_token_id"`
	ToTokenId   string `json:"to_token_id"`
}

type SwapRequest struct {
	UserDiscordId string             `json:"userDiscordId" binding:"required"`
	ChainName     string             `json:"chainName" binding:"required"`
	RouteSummary  model.RouteSummary `json:"routeSummary"`
}
