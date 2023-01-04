package request

import "time"

type OffchainTransferRequest struct {
	Sender              string   `json:"sender"`
	Recipients          []string `json:"recipients"`
	RecipientsAddresses []string `json:"recipients_addresses"`
	Platform            string   `json:"platform"`
	GuildID             string   `json:"guild_id"`
	ChannelID           string   `json:"channel_id"`
	Amount              float64  `json:"amount"`
	Token               string   `json:"token"`
	Each                bool     `json:"each"`
	All                 bool     `json:"all"`
	TransferType        string   `json:"transfer_type"`
	FullCommand         string   `json:"full_command"`
	Duration            int      `json:"duration"`
	Message             string   `json:"message"`
	Image               string   `json:"image"`
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

type TipBotDepositRequest struct {
	ChainID       int       `json:"chain_id"`
	FromAddress   string    `json:"from_address"`
	ToAddress     string    `json:"to_address"`
	TokenSymbol   string    `json:"token_symbol"`
	TokenContract string    `json:"token_contract"`
	Amount        float64   `json:"amount"`
	TxHash        string    `json:"tx_hash"`
	BlockNumber   int64     `json:"block_number"`
	SignedAt      time.Time `json:"signed_at"`
}

type GetLatestDepositRequest struct {
	ChainID         string `json:"chain_id" form:"chain_id" binding:"required"`
	ContractAddress string `json:"contract_address" form:"contract_address" binding:"required"`
}

type TipBotGetContractsRequest struct {
	ChainID        string `json:"chain_id" form:"chain_id"`
	IsEVM          *bool  `json:"is_evm" form:"is_evm"`
	SupportDeposit *bool  `json:"support_deposit" form:"support_deposit"`
}
