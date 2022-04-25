package response

type Role struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Reaction string `json:"reaction"`
}

type RoleReactionResponse struct {
	MessageID string `json:"message_id"`
	GuildID   string `json:"guild_id"`
	ChannelID string `json:"channel_id"`
	Role      Role   `json:"role"`
}
