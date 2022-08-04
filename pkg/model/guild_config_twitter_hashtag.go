package model

type GuildConfigTwitterHashtag struct {
	GuildID   string `json:"guild_id"`
	ChannelID string `json:"channel_id"`
	Hashtag   string `json:"hashtag"`
}
