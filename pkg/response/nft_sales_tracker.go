package response

type NFTSalesTrackerResponse struct {
	ContractAddress string `json:"contract_address"`
	Platform        string `json:"platform"`
	GuildID         string `json:"guild_id"`
	ChannelID       string `json:"channel_id"`
}

type NFTSalesTrackerGuildResponse struct {
	ID        string            `json:"id"`
	GuildID   string            `json:"guild_id"`
	ChannelID string            `json:"channel_id"`
	Data      []NFTSalesTracker `json:"data"`
}

type NFTSalesTracker struct {
	ID              string `json:"id"`
	ContractAddress string `json:"contract_address"`
	Platform        string `json:"platform"`
}
