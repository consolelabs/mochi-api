package request

type UpsertGmConfigRequest struct {
	GuildID   string `json:"guild_id"`
	ChannelID string `json:"channel_id"`
}

type UpsertGuildTokenConfigRequest struct {
	GuildID string `json:"guild_id"`
	Symbol  string `json:"symbol"`
	Active  bool   `json:"active"`
}
