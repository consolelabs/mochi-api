package response

type SetUpvoteMessageCacheResponse struct {
	UserID    string `json:"user_id"`
	GuildID   string `json:"guild_id"`
	ChannelID string `json:"channel_id"`
	MessageID string `json:"message_id"`
}

type ZSetWithScoreData struct {
	Score  float64 `json:"score"`
	Member string  `json:"member"`
}
