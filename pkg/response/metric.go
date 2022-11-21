package response

type Metric struct {
	NftCollections    int64 `json:"nft_collections"`
	ServerActiveUsers int64 `json:"server_active_users"`
	TotalActiveUsers  int64 `json:"total_active_users"`
	TotalServers      int64 `json:"total_servers"`
}

type DataMetric struct {
	Data Metric `json:"data"`
}

type MetricActiveUser struct {
	ActiveUsers int64  `json:"active_users"`
	GuildId     string `json:"guild_id"`
}
