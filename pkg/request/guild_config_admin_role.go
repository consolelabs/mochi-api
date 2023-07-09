package request

type CreateGuildAdminRoleRequest struct {
	GuildID string   `uri:"guild_id"`
	RoleIds []string `json:"role_ids"`
}
