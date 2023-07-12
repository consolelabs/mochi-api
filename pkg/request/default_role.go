package request

type CreateDefaultRoleRequest struct {
	RoleID  string `json:"role_id"`
	GuildID string `uri:"guild_id" binding:"required"`
}
