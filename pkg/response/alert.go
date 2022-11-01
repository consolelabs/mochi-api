package response

import "github.com/defipod/mochi/pkg/model"

type UserDeviceResponse struct {
	DeviceID     string `json:"device_id"`
	IosNotiToken string `json:"ios_noti_token"`
}

type DiscordUserTokenAlertResponse struct {
	Data []model.DiscordUserTokenAlert `json:"data"`
}
