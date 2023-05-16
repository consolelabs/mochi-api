package model

type GuildConfigAdminRole struct {
	Id      int64  `json:"id"`
	GuildId string `json:"guild_id"`
	RoleId  string `json:"role_id"`
}
