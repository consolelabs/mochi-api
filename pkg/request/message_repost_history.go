package request

type CreateMessageRepostHistRequest struct {
	GuildID         string `json:"guild_id"`
	ChannelID       string `json:"channel_id"`
	MessageID       string `json:"message_id"`
	Reaction        string `json:"reaction"`
	ReactionCount   int    `json:"reaction_count"`
	RepostChannelID string `json:"repost_channel_id"`
}
