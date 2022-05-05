package request

import (
	"fmt"
	"time"

	"github.com/defipod/mochi/pkg/model"
	"github.com/gin-gonic/gin"
)

type NewGuildConfigWalletVerificationMessageRequest struct {
	model.GuildConfigWalletVerificationMessage
	CreatedAt time.Time `json:"-"`
}

func (input *NewGuildConfigWalletVerificationMessageRequest) Validate() error {
	switch true {
	case input.GuildID == "":
		return fmt.Errorf("missing guild_id")
	case input.VerifyChannelID == "":
		return fmt.Errorf("missing verify_channel_id")
	case input.Content == "" && input.EmbeddedMessage == nil:
		return fmt.Errorf("content or embedded_message is required")
	}

	return nil
}

func (input *NewGuildConfigWalletVerificationMessageRequest) Bind(c *gin.Context) (err error) {
	err = c.BindJSON(input)
	if err != nil {
		return err
	}

	return err
}

type GenerateVerificationRequest struct {
	UserDiscordID string `json:"user_discord_id"`
	GuildID       string `json:"guild_id"`
	IsReverify    bool   `json:"is_reverify"`
}

func (input *GenerateVerificationRequest) Validate() error {
	switch true {
	case input.UserDiscordID == "":
		return fmt.Errorf("missing user_discord_id")
	case input.GuildID == "":
		return fmt.Errorf("missing guild_id")
	}

	return nil
}

func (input *GenerateVerificationRequest) Bind(c *gin.Context) (err error) {
	err = c.BindJSON(input)
	if err != nil {
		return err
	}

	return err
}

type VerifyWalletAddressRequest struct {
	WalletAddress string `json:"wallet_address"`
	Code          string `json:"code"`
	Signature     string `json:"signature"`
}

func (input *VerifyWalletAddressRequest) Bind(c *gin.Context) (err error) {
	err = c.BindJSON(input)
	if err != nil {
		return err
	}

	return err
}

func (input *VerifyWalletAddressRequest) Validate() error {
	switch true {
	case input.WalletAddress == "":
		return fmt.Errorf("missing wallet_address")
	case input.Signature == "":
		return fmt.Errorf("missing signature")
	case input.Code == "":
		return fmt.Errorf("no verification code provided")
	}

	return nil
}
