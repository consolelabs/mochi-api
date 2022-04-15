package model

type GuildConfigLogChannel struct {
	GuildID   int64 `json:"guild_id" gorm:"primary_key"`
	ChannelID int64 `json:"channel_id" gorm:"primary_key"`
}
