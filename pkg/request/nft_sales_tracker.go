package request

type NFTSalesTrackerRequest struct {
	ContractAddress string `json:"contract_address"`
	Platform        string `json:"platform"`
	GuildID         string `json:"guild_id"`
	ChannelID       string `json:"channel_id"`
}
