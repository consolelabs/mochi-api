package handler

import (
	"encoding/json"
	"net/http"

	"github.com/defipod/mochi/pkg/consts"
	"github.com/defipod/mochi/pkg/logger"

	"github.com/bwmarrin/discordgo"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/response"
	"github.com/gin-gonic/gin"
)

func (h *Handler) HandleDiscordWebhook(c *gin.Context) {
	var req request.HandleDiscordWebhookRequest
	if err := req.Bind(c); err != nil {
		h.log.Fields(logger.Fields{"body": req}).Error(err, "[handler.HandleDiscordWebhook] - failed to read JSON")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := req.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	h.log.Infof("EVENT: %v", req.Event)
	switch req.Event {
	case request.GUILD_MEMBER_ADD:
		h.handleGuildMemberAdd(c, req.Data)
	case request.MESSAGE_CREATE:
		h.handleMessageCreate(c, req.Data)
	case request.MESSAGE_DELETE:
		h.handleMessageDelete(c, req.Data)
	case request.GUILD_CREATE:
		h.handleGuildCreate(c, req.Data)
	case request.MESSAGE_REACTION_ADD:
		h.handleMessageReactionAdd(c, req.Data)
	case request.MESSAGE_REACTION_REMOVE:
		h.handleMessageReactionRemove(c, req.Data)
	}
}

func (h *Handler) handleGuildMemberAdd(c *gin.Context, data json.RawMessage) {
	var member discordgo.Member
	byteData, err := data.MarshalJSON()
	if err != nil {
		h.log.Info("[handler.handleGuildMemberAdd] - failed to json marshal data")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := discordgo.Unmarshal(byteData, &member); err != nil {
		h.log.Info("[handler.handleGuildMemberAdd] - failed to unmarshal data")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	h.handleInviteTracker(c, &member)
}

func (h *Handler) handleInviteTracker(c *gin.Context, invitee *discordgo.Member) {
	inviter, isVanity, err := h.entities.FindInviter(invitee.GuildID)
	if err != nil {
		h.log.Fields(logger.Fields{"invitee": invitee}).Error(err, "[handler.handleInviteTracker] - failed to find inviter")
	}

	response, err := h.entities.HandleInviteTracker(inviter, invitee)
	if err != nil {
		h.log.Fields(logger.Fields{"inviter": inviter, "invitee": invitee}).Error(err, "[handler.handleInviteTracker] - failed to handle invite tracker")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	response.IsVanity = isVanity

	c.JSON(http.StatusOK, gin.H{
		"data": response,
	})
}

func (h *Handler) handleMessageCreate(c *gin.Context, data json.RawMessage) {
	message := &discordgo.Message{}
	byteData, err := data.MarshalJSON()
	if err != nil {
		h.log.Error(err, "[handler.handleMessageCreate] - failed to json marshal data")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := discordgo.Unmarshal(byteData, &message); err != nil {
		h.log.Error(err, "[handler.handleMessageCreate] - failed to unmarshal data")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	uActivity, err := h.entities.HandleDiscordMessage(message)
	if err != nil {
		h.log.Fields(logger.Fields{"message": message}).Error(err, "[handler.handleMessageCreate] - failed to handle discord message")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// TODO: use response data to send discord message to user
	var resp *response.HandleUserActivityResponse
	switch message.Type {
	case consts.MessageTypeUserPremiumGuildSubscription:
		resp, err = h.entities.BoostXPIncrease(message)
	default:
		if uActivity != nil {
			// break if message was already handled
			break
		}
		resp, err = h.entities.ChatXPIncrease(message)
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "OK", "type": "level_up", "data": resp})
}

func (h *Handler) handleGuildCreate(c *gin.Context, data json.RawMessage) {
	var req struct {
		GuildID string `json:"guild_id"`
	}

	byteData, err := data.MarshalJSON()
	if err != nil {
		h.log.Error(err, "[handler.handleGuildCreate] - failed to json marshal data")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := discordgo.Unmarshal(byteData, &req); err != nil {
		h.log.Error(err, "[handler.handleGuildCreate] - failed to unmarshal data")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err = h.entities.InitGuildDefaultTokenConfigs(req.GuildID); err != nil {
		h.log.Fields(logger.Fields{"guildID": req.GuildID}).Error(err, "[handler.handleGuildCreate] - failed to init default token configs")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err = h.entities.InitGuildDefaultActivityConfigs(req.GuildID); err != nil {
		h.log.Fields(logger.Fields{"guildID": req.GuildID}).Error(err, "[handler.handleGuildCreate] - failed to init default activity configs")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}

func (h *Handler) handleMessageReactionAdd(c *gin.Context, data json.RawMessage) {
	var req request.MessageReactionRequest
	byteData, err := data.MarshalJSON()
	if err != nil {
		h.log.Error(err, "[handler.handleMessageReactionAdd] - failed to json marshal data")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := discordgo.Unmarshal(byteData, &req); err != nil {
		h.log.Error(err, "[handler.handleMessageReactionAdd] - failed to unmarshal data")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.entities.AddMessageReaction(req); err != nil {
		h.log.Fields(logger.Fields{"body": req}).Error(err, "[handler.handleMessageReactionAdd] - failed to create message reaction")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	repostMessage, err := h.entities.CreateRepostReactionEvent(req)
	if err != nil {
		h.log.Fields(logger.Fields{"body": req}).Error(err, "[handler.handleMessageReactionAdd] - failed to create repost reaction event")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if repostMessage == nil {
		c.JSON(http.StatusOK, &response.RepostReactionEventResponse{
			Data: response.RepostReactionEventData{
				Status: "OK",
			},
		})
		return
	}

	c.JSON(http.StatusOK, &response.RepostReactionEventResponse{
		Data: response.RepostReactionEventData{
			Status:          "OK",
			RepostChannelID: repostMessage.RepostChannelID,
			RepostMessageID: repostMessage.RepostMessageID,
		},
	})
}

func (h *Handler) handleMessageReactionRemove(c *gin.Context, data json.RawMessage) {
	var req request.MessageReactionRequest
	byteData, err := data.MarshalJSON()
	if err != nil {
		h.log.Error(err, "[handler.handleMessageReactionRemove] - failed to json marshal data")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := discordgo.Unmarshal(byteData, &req); err != nil {
		h.log.Error(err, "[handler.handleMessageReactionRemove] - failed to unmarshal data")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.entities.RemoveMessageReaction(req); err != nil {
		h.log.Fields(logger.Fields{"body": req}).Error(err, "[handler.handleMessageReactionRemove] - failed to create message reaction")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, &response.RepostReactionEventResponse{
		Data: response.RepostReactionEventData{
			Status: "OK",
		},
	})
}

func (h *Handler) handleMessageDelete(c *gin.Context, data json.RawMessage) {
	message := &discordgo.Message{}
	byteData, err := data.MarshalJSON()
	if err != nil {
		h.log.Error(err, "[handler.handleMessageDelete] - failed to json marshal data")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := discordgo.Unmarshal(byteData, &message); err != nil {
		h.log.Error(err, "[handler.handleMessageDelete] - failed to unmarshal data")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err = h.entities.RemoveAllMessageReactions(message); err != nil {
		h.log.Fields(logger.Fields{"message": message}).Error(err, "[handler.handleMessageDelete] - failed to handle message delete")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"status": "OK",
		},
	})
}

func (h *Handler) WebhookUpvoteTopGG(c *gin.Context) {
	req := request.WebhookUpvoteTopGG{}
	err := c.BindJSON(&req)
	if err != nil {
		h.log.Fields(logger.Fields{"body": req}).Error(err, "[handler.WebhookUpvoteTopGG] - failed to read JSON")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.entities.WebhookUpvoteStreak(req.UserID, consts.TopGGSource)
	if err != nil {
		h.log.Fields(logger.Fields{"body": req}).Error(err, "[handler.WebhookUpvoteTopGG] - failed to add upvote streak")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}

func (h *Handler) WebhookUpvoteDiscordBot(c *gin.Context) {
	req := request.WebhookUpvoteDiscordBot{}
	err := c.BindJSON(&req)
	if err != nil {
		h.log.Fields(logger.Fields{"body": req}).Error(err, "[handler.WebhookUpvoteDiscordBot] - failed to read JSON")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.entities.WebhookUpvoteStreak(req.UserID, consts.DiscordBotListSource)
	if err != nil {
		h.log.Fields(logger.Fields{"body": req}).Error(err, "[handler.WebhookUpvoteDiscordBot] - failed to add upvote streak")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}
