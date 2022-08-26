package response

type DefaultRole struct {
	RoleID  string `json:"role_id"`
	GuildID string `json:"guild_id"`
}

type DefaultRoleResponse struct {
	Data DefaultRole `json:"data"`
	Ok   bool        `json:"ok"`
}
