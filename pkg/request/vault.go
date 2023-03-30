package request

type CreateVaultRequest struct {
	GuildId   string `json:"guild_id"`
	Name      string `json:"name"`
	Threshold string `json:"threshold"`
}

type CreateVaultConfigChannelRequest struct {
	GuildId   string `json:"guild_id" binding:"required"`
	ChannelId string `json:"channel_id" binding:"required"`
}

type CreateConfigThresholdRequest struct {
	GuildId   string `json:"guild_id" binding:"required"`
	Name      string `json:"name" binding:"required"`
	Threshold string `json:"threshold" binding:"required"`
}
