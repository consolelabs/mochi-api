package request

import "time"

type CreateGuildRequest struct {
	ID       string    `json:"id"`
	Name     string    `json:"name"`
	JoinedAt time.Time `json:"-"`
}

type AvailableCMD struct {
	ID                       string `json:"id"`
	ApplicationID            string `json:"application_id"`
	Version                  string `json:"version"`
	DefaultMemberPermissions *int   `json:"default_member_permissions"`
	Type                     int    `json:"type"`
	Name                     string `json:"name"`
	NameLocalizations        string `json:"name_localizations"`
	Description              string `json:"description"`
	DescriptionLocalizations string `json:"description_localizations"`
	GuildID                  string `json:"guild_id"`
	NSFW                     bool   `json:"nsfw"`
}

type UpdateGuildRequest struct {
	GlobalXP      *bool           `json:"global_xp"`
	LogChannel    *string         `json:"log_channel"`
	Active        *bool           `json:"active"`
	LeftAt        *time.Time      `json:"left_at"`
	AvailableCMDs *[]AvailableCMD `json:"available_cmds"`
}

type HandleGuildDeleteRequest struct {
	GuildID   string `json:"guild_id"`
	GuildName string `json:"guild_name"`
	IconURL   string `json:"icon_url"`
}

type ValidateUserRequest struct {
	Ids     string `form:"ids"`
	GuildId string `form:"guild_id"`
}
