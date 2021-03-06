package request

type RoleReactionRequest struct {
	GuildID   string `json:"guild_id"`
	MessageID string `json:"message_id"`
	Reaction  string `json:"reaction"`
}

type RoleReactionUpdateRequest struct {
	GuildID   string `json:"guild_id"`
	MessageID string `json:"message_id"`
	Reaction  string `json:"reaction"`
	RoleID    string `json:"role_id"`
}
