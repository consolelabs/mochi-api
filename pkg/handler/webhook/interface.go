package webhook

import (
	"encoding/json"

	"github.com/gin-gonic/gin"

	"github.com/defipod/mochi/pkg/request"
)

type IHandler interface {
	HandleDiscordWebhook(c *gin.Context)
	handleGuildMemberAdd(c *gin.Context, data json.RawMessage)
	handleGuildMemberRemove(c *gin.Context, data json.RawMessage)
	handleMessageCreate(c *gin.Context, data json.RawMessage)
	handleMessageDelete(c *gin.Context, data json.RawMessage)
	handleGuildCreate(c *gin.Context, data json.RawMessage)
	handleMessageReactionAdd(c *gin.Context, data json.RawMessage)
	handleMessageReactionRemove(c *gin.Context, data json.RawMessage)
	handleGuildDelete(c *gin.Context, data json.RawMessage)
	WebhookNftHandler(c *gin.Context)
	handleNftSales(c *gin.Context, req request.NotifySaleMarketplaceRequest)
	WebhookUpvoteTopGG(c *gin.Context)
	WebhookUpvoteDiscordBot(c *gin.Context)
	NotifyNftCollectionIntegration(c *gin.Context)
	NotifyNftCollectionSync(c *gin.Context)
	NotifySaleMarketplace(c *gin.Context)
}
