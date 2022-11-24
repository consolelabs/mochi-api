package request

type UpsertMonikerConfigRequest struct {
	Moniker string  `json:"moniker" binding:"required"`
	Plural  string  `json:"plural"`
	Amount  float64 `json:"amount" binding:"required"`
	Token   string  `json:"token" binding:"required"`
	GuildID string  `json:"guild_id" binding:"required"`
}

type DeleteMonikerConfigRequest struct {
	Moniker string `json:"moniker" binding:"required"`
	GuildID string `json:"guild_id" binding:"required"`
}
