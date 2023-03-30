package request

type CreateVaultRequest struct {
	GuildId   string `json:"guild_id"`
	Name      string `json:"name"`
	Threshold string `json:"threshold"`
}

type CreateVaultConfigChannelRequest struct {
	GuildId   string `json:"guild_id"`
	ChannelId string `json:"channel_id"`
}

type CreateConfigThresholdRequest struct {
	GuildId   string `json:"guild_id"`
	Name      string `json:"name"`
	Threshold string `json:"threshold"`
}
