package model

type DiscordGuild struct {
	ID        int64           `json:"id"`
	Name      string          `json:"name"`
	BotScopes JSONArrayString `json:"bot_scopes"`
	Alias     string          `json:"alias"`
	Roles     []GuildRole     `json:"roles" gorm:"foreignkey:GuildID"`
}
