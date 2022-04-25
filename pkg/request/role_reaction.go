package request

type RoleReactionRequest struct {
	GuildID    string `json:"guild_id"`
	MessageID  string `json:"message_id"`
	Reaction   string `json:"reaction"`
	ActionType string `json:"action_type"`
}
