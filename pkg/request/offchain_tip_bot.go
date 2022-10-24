package request

type OffchainTransferRequest struct {
	Sender       string   `json:"sender"`
	Recipients   []string `json:"recipients"`
	GuildID      string   `json:"guild_id"`
	ChannelID    string   `json:"channel_id"`
	Amount       float64  `json:"amount"`
	Token        string   `json:"token"`
	Each         bool     `json:"each"`
	All          bool     `json:"all"`
	TransferType string   `json:"transfer_type"`
	FullCommand  string   `json:"full_command"`
	Duration     int      `json:"duration"`
}
