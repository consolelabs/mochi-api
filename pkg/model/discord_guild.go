package model

import "time"

type DiscordGuild struct {
	ID                       string                   `json:"id"`
	Name                     string                   `json:"name"`
	BotScopes                JSONArrayString          `json:"bot_scopes"`
	Alias                    string                   `json:"alias"`
	Roles                    []GuildRole              `json:"roles" gorm:"foreignkey:GuildID"`
	CreatedAt                time.Time                `json:"created_at"`
	GuildConfigInviteTracker GuildConfigInviteTracker `json:"-" gorm:"foreignkey:GuildID"`
}
