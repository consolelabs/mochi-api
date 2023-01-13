package request

type CreateSalesTrackerConfigRequest struct {
	GuildID         string `json:"guild_id"`
	ChannelID       string `json:"channel_id"`
	ContractAddress string `json:"contract_address"`
	Chain           string `json:"chain"`
}
