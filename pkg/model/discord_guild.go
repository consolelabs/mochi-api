package model

import "time"

type DiscordGuild struct {
	ID         string          `json:"id"`
	Name       string          `json:"name"`
	BotScopes  JSONArrayString `json:"bot_scopes"`
	Alias      string          `json:"alias"`
	Roles      []GuildRole     `json:"roles" gorm:"foreignkey:GuildID"`
	CreatedAt  time.Time       `json:"created_at"`
	GlobalXP   bool            `json:"global_xp"`
	LogChannel string          `json:"log_channel"`
	Active     bool            `json:"active"`
	JoinedAt   time.Time       `json:"-"`
	LeftAt     *time.Time      `json:"-"`
}
