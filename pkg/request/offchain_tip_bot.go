package request

type OffchainTransferRequest struct {
	Sender     string   `json:"sender"`
	Recipients []string `json:"recipients"`
	Platform   string   `json:"platform"`
	GuildID    string   `json:"guild_id"`
	ChannelID  string   `json:"channel_id"`
	Amount     float64  `json:"amount"`
	// AmountString string   `json:"amount_string"`
	Token        string `json:"token"`
	Each         bool   `json:"each"`
	All          bool   `json:"all"`
	TransferType string `json:"transfer_type"`
	Message      string `json:"message"`
	Image        string `json:"image"`
	ChainID      string `json:"chain_id"`
}

type TransferV2Request struct {
	Sender     string                 `json:"sender"`
	Recipients []string               `json:"recipients"`
	Platform   string                 `json:"platform"`
	GuildID    string                 `json:"guild_id"`
	Amount     float64                `json:"amount"`
	Token      string                 `json:"token"`
	Each       bool                   `json:"each"`
	All        bool                   `json:"all"`
	Message    string                 `json:"message"`
	ChainID    string                 `json:"chain_id"`
	Metadata   map[string]interface{} `json:"metadata"`
}
