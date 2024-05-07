package request

type IntegrationBinanceData struct {
	ApiKey        string `json:"api_key"`
	ApiSecret     string `json:"api_secret"`
	DiscordUserId string `json:"discord_user_id"`
}

type UnlinkBinance struct {
	DiscordUserId string `json:"discord_user_id"`
}

type BinanceRequest struct {
	Id        string `json:"id"`
	ApiKey    string `json:"api_key" form:"api_key" binding:"required"`
	ApiSecret string `json:"api_secret" form:"api_secret" binding:"required"`
}

type GetBinanceAssetsRequest struct {
	Id       string `json:"id"`
	Platform string `json:"platform"`
}

type GetBinanceFutureRequest struct {
	Id string `uri:"id"`
}

type GetBinanceSpotTxnsRequest struct {
	Id     string `uri:"id"`
	Status string `json:"status" form:"status"`
	PaginationRequest
}
