package request

type CreateGuildRequest struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type UpdateGuildRequest struct {
	GlobalXP   *bool   `json:"global_xp"`
	LogChannel *string `json:"log_channel"`
	Active     *bool   `json:"active"`
}

type HandleGuildDeleteRequest struct {
	GuildID   string `json:"guild_id"`
	GuildName string `json:"guild_name"`
	IconURL   string `json:"icon_url"`
}
