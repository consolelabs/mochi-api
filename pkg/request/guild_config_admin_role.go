package request

type CreateGuildAdminRoleRequest struct {
	GuildID string   `json:"guild_id" binding:"required"`
	RoleIds []string `json:"role_ids" binding:"required"`
}
