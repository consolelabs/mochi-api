package request

type AutoTriggerRequest struct {
	GuildId         string   `json:"guild_id"`
	ChannelId       string   `json:"channel_id"`
	MessageId       string   `json:"message_id"`
	Reaction        string   `json:"reaction"`
	ReactionCount   int      `json:"reaction_count"`
	RepostChannelId string   `json:"repost_channel_id"`
	UserID          string   `json:"user_id"`
	IsStart         bool     `json:"is_start"`
	IsStop          bool     `json:"is_stop"`
	UserRoles       []string `json:"user_roles"`
	AuthorRoles     []string `json:"author_roles"`
	Content         string   `json:"content"`
	AuthorId        string   `json:"author_id"`
	Source          string   `json:"source"`
}
