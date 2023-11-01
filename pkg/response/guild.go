package response

import (
	"github.com/bwmarrin/discordgo"

	"github.com/defipod/mochi/pkg/model"
)

type GetGuildsResponse struct {
	Data []*GetGuildResponse `json:"data"`
}

type GetGuildResponse struct {
	ID            string              `json:"id"`
	Name          string              `json:"name"`
	BotScopes     []string            `json:"bot_scopes"`
	Alias         string              `json:"alias"`
	LogChannel    string              `json:"log_channel"`
	LogChannelID  string              `json:"log_channel_id"`
	GlobalXP      bool                `json:"global_xp"`
	Active        bool                `json:"active"`
	Icon          string              `json:"icon"`
	AvailableCMDs *[]model.DiscordCMD `json:"available_cmds"`
}

type ListAllCustomTokenResponse struct {
	Data []model.Token `json:"data"`
}

type DiscordGuildResponse struct {
	discordgo.UserGuild
	BotAddable bool `json:"bot_addable"`
	BotArrived bool `json:"bot_arrived"`
}

type ListMyGuildsResponse struct {
	Data []DiscordGuildResponse `json:"data"`
}
