package request

type UpsertUserDeviceRequest struct {
	DeviceID     string `json:"device_id"`
	IosNotiToken string `json:"ios_noti_token"`
}

type DeleteUserDeviceRequest struct {
	DeviceID string `json:"device_id"`
}

type UpsertDiscordUserAlertRequest struct {
	ID        string  `json:"id"`
	IsEnable  bool    `json:"is_enable"`
	TokenID   string  `json:"token_id"`
	Symbol    string  `json:"symbol"`
	DiscordID string  `json:"discord_id"`
	PriceSet  float64 `json:"price_set"`
	Trend     string  `json:"trend"`
	DeviceID  string  `json:"device_id"`
}

type DeleteDiscordUserAlertRequest struct {
	ID string `json:"id"`
}
