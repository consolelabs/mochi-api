package request

type MessageReactionRequest struct {
	GuildID         string `json:"guild_id"`
	ChannelID       string `json:"channel_id"`
	MessageID       string `json:"message_id"`
	Reaction        string `json:"reaction"`
	ReactionCount   int    `json:"reaction_count"`
	RepostChannelID string `json:"repost_channel_id"`
	UserID          string `json:"user_id"`
	IsStart         bool   `json:"is_start"`
	IsStop          bool   `json:"is_stop"`
}
