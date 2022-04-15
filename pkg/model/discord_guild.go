package model

type DiscordGuild struct {
	ID         int64                 `json:"id"`
	Name       string                `json:"name"`
	BotScopes  JSONArrayString       `json:"bot_scopes"`
	Alias      string                `json:"alias"`
	Roles      []GuildRole           `json:"roles" gorm:"foreignkey:GuildID"`
	LogChannel GuildConfigLogChannel `json:"log_channel" gorm:"rel:has_one;foreignkey:GuildID;table_name:guild_config_log_channels"`
}
