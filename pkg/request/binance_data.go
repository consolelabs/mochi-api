package request

type IntegrationBinanceData struct {
	ApiKey        string `json:"apiKey"`
	ApiSecret     string `json:"apiSecret"`
	DiscordUserId string `json:"discordUserId"`
}

type BinanceRequest struct {
	ProfileId string `json:"profile_id" form:"profile_id"`
	ApiKey    string `json:"api_key" form:"api_key" binding:"required"`
	ApiSecret string `json:"api_secret" form:"api_secret" binding:"required"`
}
