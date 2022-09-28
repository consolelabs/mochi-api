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
	// guild
	GUILD_CREATE = "guildCreate"
	GUILD_DELETE = "guildDelete"
	// member
	GUILD_MEMBER_ADD = "guildMemberAdd"
	// message
	MESSAGE_CREATE = "messageCreate"
	MESSAGE_DELETE = "messageDelete"
	// reaction
	MESSAGE_REACTION_ADD    = "messageReactionAdd"
	MESSAGE_REACTION_REMOVE = "messageReactionRemove"
	// sales
	SALES_CREATE = "salesCreate"
)

var acceptedEvents = map[string]bool{
	// guild
	GUILD_CREATE: true,
	GUILD_DELETE: true,
	// member
	GUILD_MEMBER_ADD: true,
	// message
	MESSAGE_DELETE: true,
	MESSAGE_CREATE: true,
	// reaction
	MESSAGE_REACTION_ADD:    true,
	MESSAGE_REACTION_REMOVE: true,
	// sales
	SALES_CREATE: true,
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
