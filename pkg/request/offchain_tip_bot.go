package request

type OffchainTransferRequest struct {
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
	Imgae        string   `json:"image"`
}
type OffchainWithdrawRequest struct {
	Recipient        string  `json:"recipient"`
	RecipientAddress string  `json:"recipient_address"`
	GuildID          string  `json:"guild_id"`
	ChannelID        string  `json:"channel_id"`
	Amount           float64 `json:"amount"`
	Token            string  `json:"token"`
	Each             bool    `json:"each"`
	All              bool    `json:"all"`
	TransferType     string  `json:"transfer_type"`
	FullCommand      string  `json:"full_command"`
	Duration         int     `json:"duration"`
}

type OffchainUpdateTokenFee struct {
	Symbol     string  `json:"symbol"`
	ServiceFee float64 `json:"service_fee"`
}
