package request

type CreateVaultRequest struct {
	VaultCreator string `json:"vault_creator" binding:"required"`
	GuildId      string `json:"guild_id"`
	Name         string `json:"name"`
	Threshold    string `json:"threshold"`
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
	GuildId       string `json:"guild_id" binding:"required"`
	VaultId       int64  `json:"vault_id" binding:"required"`
	UserDiscordID string `json:"user_discord_id" binding:"required"`
	ChannelId     string `json:"channel_id" binding:"required"`
}

type CreateTreasurerResultRequest struct {
	GuildId       string `json:"guild_id" binding:"required"`
	VaultId       int64  `json:"vault_id" binding:"required"`
	UserDiscordID string `json:"user_discord_id" binding:"required"`
	ChannelId     string `json:"channel_id" binding:"required"`
	Type          string `json:"type" binding:"required"`
	Status        string `json:"status" binding:"required"`
}

type CreateTreasurerRequest struct {
	GuildId       string `json:"guild_id" binding:"required"`
	Requester     string `json:"requester" binding:"required"`
	VaultName     string `json:"vault_name" binding:"required"`
	UserDiscordId string `json:"user_discord_id" binding:"required"`
	Message       string `json:"message" binding:"required"`
	Type          string `json:"type" binding:"required"`
}

type CreateTreasurerSubmission struct {
	VaultId   int64  `json:"vault_id" binding:"required"`
	RequestId int64  `json:"request_id" binding:"required"`
	Sumitter  string `json:"submitter" binding:"required"`
	Choice    string `json:"choice" binding:"required"`
}
