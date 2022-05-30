package response

type GetGuildsResponse struct {
	Data []*GetGuildResponse `json:"data"`
}

type GetGuildResponse struct {
	ID           string   `json:"id"`
	Name         string   `json:"name"`
	BotScopes    []string `json:"bot_scopes"`
	Alias        string   `json:"alias"`
	LogChannelID string   `json:"log_channel_id"`
	GlobalXP     bool     `json:"global_xp"`
}
