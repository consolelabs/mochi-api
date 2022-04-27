package response

type Role struct {
	ID       string `json:"id"`
	Reaction string `json:"reaction"`
}

type RoleReactionResponse struct {
	MessageID string `json:"message_id"`
	GuildID   string `json:"guild_id"`
	Role      Role   `json:"role"`
}

type RoleReactionConfigResponse struct {
	MessageID string `json:"message_id"`
	GuildID   string `json:"guild_id"`
	Roles     []Role `json:"roles"`
	Success   bool   `json:"success"`
}
