package request

type IntegrationBinanceData struct {
	ApiKey        string `json:"apiKey"`
	ApiSecret     string `json:"apiSecret"`
	DiscordUserId string `json:"discordUserId"`
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
