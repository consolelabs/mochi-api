package request

type CreateTipConfigNotify struct {
	GuildId   string `json:"guild_id"`
	ChannelId string `json:"channel_id"`
	Token     string `json:"token"`
}
