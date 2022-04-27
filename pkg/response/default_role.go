package response

type DefaultRole struct {
	RoleID  string `json:"role_id"`
	GuildID string `json:"guild_id"`
}

type DefaultRoleCreationResponse struct {
	Data    DefaultRole `json:"role"`
	Success bool        `json:"success"`
}

type DefaultRoleGetAllResponse struct {
	Data    []*DefaultRole `json:"data"`
	Success bool           `json:"success"`
}
