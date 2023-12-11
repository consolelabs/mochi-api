package request

import (
	"errors"

	"github.com/gin-gonic/gin"

	"github.com/defipod/mochi/pkg/util"
)

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
	Message      string   `json:"message"`
	Image        string   `json:"image"`
	ChainID      string   `json:"chain_id"`
}

type TransferV2Request struct {
	Sender         string                 `json:"sender" binding:"required"`
	Recipients     []string               `json:"recipients" binding:"required"`
	Platform       string                 `json:"platform" enums:"discord,telegram,web" binding:"required"`
	GuildID        string                 `json:"guild_id"`
	Amount         float64                `json:"amount" binding:"required"`
	Token          string                 `json:"token" binding:"required"`
	Each           bool                   `json:"each"`
	All            bool                   `json:"all"`
	TransferType   string                 `json:"transfer_type" binding:"required" enums:"transfer,airdrop"`
	Message        string                 `json:"message"`
	ChainID        string                 `json:"chain_id" binding:"required"`
	Metadata       map[string]interface{} `json:"metadata"`
	Moniker        string                 `json:"moniker"`
	OriginalTxId   string                 `json:"original_tx_id"`
	OriginalAmount float64                `json:"original_amount"`
	ChannelId      string                 `json:"channel_id"`
	ChannelName    string                 `json:"channel_name"`
	ChannelUrl     string                 `json:"channel_url"`
	ChannelAvatar  string                 `json:"channel_avatar"`
	ThemeId        int64                  `json:"theme_id"`
}

func (r *TransferV2Request) Bind(c *gin.Context) error {
	// bind payload
	if err := c.BindJSON(r); err != nil {
		return err
	}

	// validate sender format
	if !util.ValidateNumberSeries(r.Sender) {
		return errors.New("invalid sender")
	}

	// check if sender is the authenticated requester
	requester := c.GetString("profile_id")
	isMochi := c.GetBool("is_mochi")
	if !isMochi && requester != r.Sender {
		return errors.New("you can only transfer tokens from your profile")
	}

	// validate recipients
	for _, recipient := range r.Recipients {
		if !util.ValidateNumberSeries(recipient) {
			return errors.New("invalid recipient(s)")
		}
		if recipient == r.Sender {
			return errors.New("you cannot transfer token to your own profile")
		}
		continue
	}

	// validate platforms
	platforms := map[string]bool{"discord": true, "telegram": true, "web": true}
	if r.Platform != "" && !platforms[r.Platform] {
		return errors.New("invalid platforms (available values: discord, telegram, web)")
	}

	// validate chain_id
	if !util.ValidateNumberSeries(r.ChainID) {
		return errors.New("invalid chain ID")
	}

	return nil
}
