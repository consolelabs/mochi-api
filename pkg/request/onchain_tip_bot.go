package request

type SubmitOnchainTransferRequest struct {
	Sender       string   `json:"sender"`
	Recipients   []string `json:"recipients"`
	Platform     string   `json:"platform"`
	GuildID      string   `json:"guild_id"`
	ChannelID    string   `json:"channel_id"`
	Amount       float64  `json:"amount"`
	Token        string   `json:"token"`
	Each         bool     `json:"each"`
	All          bool     `json:"all"`
	TransferType string   `json:"transfer_type"`
	FullCommand  string   `json:"full_command"`
	Duration     int      `json:"duration"`
	Message      string   `json:"message"`
	Image        string   `json:"image"`
}

type ClaimOnchainTransferRequest struct {
	UserID  string `json:"user_id"`
	ClaimID int    `json:"claim_id"`
	Address string `json:"address"`
}
