package request

type UpsertUserDeviceRequest struct {
	DeviceID     string `json:"device_id"`
	IosNotiToken string `json:"ios_noti_token"`
}

type DeleteUserDeviceRequest struct {
	DeviceID string `json:"device_id"`
}

type UpsertDiscordUserAlertRequest struct {
	TokenID   string  `json:"token_id"`
	DiscordID string  `json:"discord_id"`
	PriceSet  float64 `json:"price_set"`
	Trend     string  `json:"trend"`
	DeviceID  string  `json:"device_id"`
}

type DeleteDiscordUserAlertRequest struct {
	TokenID   string `json:"token_id"`
	DiscordID string `json:"discord_id"`
	Trend     string `json:"trend"`
}
