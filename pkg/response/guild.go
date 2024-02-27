package response

import (
	"time"

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

type GuildReportRoles struct {
	Id          string                  `json:"id"`
	Name        string                  `json:"name"`
	Roles       []GuildReportRoleDetail `json:"roles"`
	LastUpdated time.Time               `json:"last_updated"`
}

type GuildReportRoleDetail struct {
	Id               string  `json:"id"`
	Name             string  `json:"name"`
	NrOfMember       int64   `json:"nr_of_member"`
	ChangePercentage float64 `json:"change_percentage"`
}

type GuildReportMembers struct {
	Id               string    `json:"id"`
	Name             string    `json:"name"`
	NrOfMember       int64     `json:"nr_of_member"`
	ChangePercentage float64   `json:"change_percentage"`
	LastUpdated      time.Time `json:"last_updated"`
}

type GuildReportAdditionalRole struct {
	Id          string                  `json:"id"`
	Name        string                  `json:"name"`
	LastUpdated time.Time               `json:"last_updated"`
	Roles       []GuildReportRoleDetail `json:"roles"`
	ListMember  []discordgo.User        `json:"list_member"`
}
