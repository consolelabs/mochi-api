package request

type CreateGmConfigRequest struct {
	GuildID   string `json:"guild_id"`
	ChannelID string `json:"channel_id"`
}
