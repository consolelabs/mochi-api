package request

type SendCollectionIntegrationLogsRequest struct {
	GuildID           string `json:"guild_id"`
	ChannelID         string `json:"channel_id"`
	MessageID         string `json:"message_id"`
	CollectionAddress string `json:"collection_address"`
}
