package response

type ListConfigNotifyResponse struct {
	Data []ConfigNotifyResponse `json:"data"`
}

type ConfigNotifyResponse struct {
	Id               string `json:"id"`
	GuildId          string `json:"guild_id"`
	ChannelId        string `json:"channel_id"`
	Token            string `json:"token"`
	CreatedAt        string `json:"created_at"`
	UpdatedAt        string `json:"updated_at"`
	TotalTransaction int64  `json:"total_transaction"`
}
