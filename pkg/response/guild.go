package response

import "github.com/defipod/mochi/pkg/model"

type GetGuildsResponse struct {
	Data []*GetGuildResponse `json:"data"`
}

type GetGuildResponse struct {
	ID           string   `json:"id"`
	Name         string   `json:"name"`
	BotScopes    []string `json:"bot_scopes"`
	Alias        string   `json:"alias"`
	LogChannel   string   `json:"log_channel"`
	LogChannelID string   `json:"log_channel_id"`
	GlobalXP     bool     `json:"global_xp"`
}

type ListAllCustomTokenResponse struct {
	Data []model.Token `json:"data"`
}
