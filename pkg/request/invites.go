package request

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type ConfigureInviteRequest struct {
	LogChannel string `json:"log_channel"`
	GuildID    string `json:"guild_id"`
	WebhookURL string `json:"webhook_url"`
}

func (r *ConfigureInviteRequest) Bind(c *gin.Context) error {
	return c.BindJSON(r)
}

func (r *ConfigureInviteRequest) Validate() error {
	if r.LogChannel == "" {
		return fmt.Errorf("log_channel is required")
	}
	if r.GuildID == "" {
		return fmt.Errorf("channel is required")
	}

	return nil
}
