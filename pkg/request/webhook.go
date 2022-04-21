package request

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
)

type HandleDiscordWebhookRequest struct {
	Event string          `json:"event"`
	Data  json.RawMessage `json:"data"`
}

const (
	GUILD_MEMBER_ADD = "guildMemberAdd"
)

var acceptedEvents = map[string]bool{
	GUILD_MEMBER_ADD: true,
}

func (input *HandleDiscordWebhookRequest) Bind(c *gin.Context) error {
	return c.BindJSON(input)
}

func (input *HandleDiscordWebhookRequest) Validate() error {
	_, ok := acceptedEvents[input.Event]
	if !ok {
		return fmt.Errorf("invalid event")
	}
	if input.Data == nil {
		return fmt.Errorf("data is required")
	}

	return nil
}
