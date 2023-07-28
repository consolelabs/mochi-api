package request

type GetSwapRouteRequest struct {
	From      string `form:"from"`
	To        string `form:"to"`
	Amount    string `form:"amount"`
	ProfileId string `form:"profileId"`
	Address   string `form:"address"`
}

type SwapRequest struct {
	UserDiscordId string      `json:"userDiscordId" binding:"required"`
	ChainName     string      `json:"chainName" binding:"required"`
	RouteSummary  interface{} `json:"routeSummary"`
	Aggregator    string      `json:"aggregator"`
	SwapData      interface{} `json:"swapData"`
}
