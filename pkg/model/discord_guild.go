package model

import "time"

type DiscordCMD struct {
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

type DiscordGuild struct {
	ID            string          `json:"id"`
	Name          string          `json:"name"`
	BotScopes     JSONArrayString `json:"bot_scopes"`
	Alias         string          `json:"alias"`
	Roles         []GuildRole     `json:"roles" gorm:"foreignkey:GuildID"`
	CreatedAt     time.Time       `json:"created_at"`
	GlobalXP      bool            `json:"global_xp"`
	LogChannel    string          `json:"log_channel"`
	Active        bool            `json:"active"`
	JoinedAt      time.Time       `json:"-"`
	LeftAt        *time.Time      `json:"-"`
	AvailableCMDs JSONNullString  `json:"available_cmds"`
}
