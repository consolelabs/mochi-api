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

type AddTreasurerToVaultRequest struct {
	RequestId int64 `json:"request_id" binding:"required"`
}

type CreateAddTreasurerRequest struct {
	GuildId       string `json:"guild_id" binding:"required"`
	VaultName     string `json:"vault_name" binding:"required"`
	UserDiscordId string `json:"user_discord_id" binding:"required"`
	Message       string `json:"message" binding:"required"`
}
