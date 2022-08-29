package response

import "github.com/defipod/mochi/pkg/model"

type CreateCustomCommandResponse struct {
	Data model.GuildCustomCommand `json:"data"`
}

type UpdateCustomCommandResponse struct {
	Data model.GuildCustomCommand `json:"data"`
}

type ListCustomCommandsResponse struct {
	Data []model.GuildCustomCommand `json:"data"`
}

type GetCustomCommandResponse struct {
	Data *model.GuildCustomCommand `json:"data"`
}
