package response

type Metric struct {
	NftCollections        int64    `json:"nft_collections"`
	ServerActiveUsers     int64    `json:"server_active_users"`
	TotalActiveUsers      int64    `json:"total_active_users"`
	TotalServers          int64    `json:"total_servers"`
	TotalVerifiedWallets  int64    `json:"total_verified_wallets"`
	ServerVerifiedWallets int64    `json:"server_verified_wallets"`
	ServerTokenSupported  int64    `json:"server_token_supported"`
	TotalTokenSupported   int64    `json:"total_token_supported"`
	ServerToken           []string `json:"server_token"`
	TotalToken            []string `json:"total_token"`
	TotalCommandUsage     int64    `json:"total_command_usage"`
	ServerCommandUsage    int64    `json:"server_command_usage"`
}

type DataMetric struct {
	Data Metric `json:"data"`
}

type MetricActiveUser struct {
	ActiveUsers int64  `json:"active_users"`
	GuildId     string `json:"guild_id"`
}
