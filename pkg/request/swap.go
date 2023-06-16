package request

import "github.com/defipod/mochi/pkg/model"

type GetSwapRouteRequest struct {
	From      string `json:"from" binding:"required"`
	To        string `json:"to" binding:"required"`
	Amount    string `json:"amount" binding:"required"`
	ProfileId string `json:"profileId" binding:"required"`
}

type SwapRequest struct {
	UserDiscordId string             `json:"userDiscordId" binding:"required"`
	ChainName     string             `json:"chainName" binding:"required"`
	RouteSummary  model.RouteSummary `json:"routeSummary"`
}
