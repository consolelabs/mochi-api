package request

type GetSwapRouteRequest struct {
	From      string `form:"from" binding:"required"`
	To        string `form:"to" binding:"required"`
	Amount    string `form:"amount" binding:"required"`
	ProfileId string `form:"profileId" binding:"required"`
	Address   string `form:"address"`
}

type SwapRequest struct {
	UserDiscordId string      `json:"userDiscordId" binding:"required"`
	ChainName     string      `json:"chainName" binding:"required"`
	RouteSummary  interface{} `json:"routeSummary"`
	Aggregator    string      `json:"aggregator"`
	SwapData      interface{} `json:"swapData"`
}

type GetOnchainAssetAvgCost struct {
	WalletAddress string `form:"walletAddress" binding:"required"`
}
