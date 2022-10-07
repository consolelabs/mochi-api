package request

type CreateDefaultRoleRequest struct {
	RoleID  string `json:"role_id" binding:"required"`
	GuildID string `json:"guild_id" binding:"required"`
}
