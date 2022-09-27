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

type WebhookUpvoteTopGG struct {
	BotID     string `json:"bot"`
	UserID    string `json:"user"`
	Type      string `json:"type"`
	IsWeekend bool   `json:"isWeekend"`
}
type WebhookUpvoteDiscordBot struct {
	Admin    bool   `json:"admin"`
	Avatar   string `json:"avatar"`
	Username string `json:"username"`
	UserID   string `json:"id"`
}

const (
	GUILD_MEMBER_ADD        = "guildMemberAdd"
	MESSAGE_CREATE          = "messageCreate"
	MESSAGE_DELETE          = "messageDelete"
	SALES_CREATE            = "salesCreate"
	GUILD_CREATE            = "guildCreate"
	MESSAGE_REACTION_ADD    = "messageReactionAdd"
	MESSAGE_REACTION_REMOVE = "messageReactionRemove"
)

var acceptedEvents = map[string]bool{
	GUILD_MEMBER_ADD:        true,
	MESSAGE_CREATE:          true,
	MESSAGE_DELETE:          true,
	SALES_CREATE:            true,
	GUILD_CREATE:            true,
	MESSAGE_REACTION_ADD:    true,
	MESSAGE_REACTION_REMOVE: true,
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
