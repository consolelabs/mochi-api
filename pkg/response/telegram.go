package response

import "github.com/defipod/mochi/pkg/model"

type GetLinkedTelegramResponse struct {
	Data *model.UserTelegramDiscordAssociation `json:"data"`
}

type LinkUserTelegramWithDiscordResponseData struct {
	model.UserTelegramDiscordAssociation
	DiscordUsername string `json:"discord_username"`
}

type LinkUserTelegramWithDiscordResponse struct {
	Data *LinkUserTelegramWithDiscordResponseData `json:"data"`
}
