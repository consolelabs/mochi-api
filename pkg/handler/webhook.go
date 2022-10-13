package handler

import (
	"encoding/json"
	"net/http"

	"github.com/bwmarrin/discordgo"
	"github.com/gin-gonic/gin"

	"github.com/defipod/mochi/pkg/consts"
	"github.com/defipod/mochi/pkg/logger"
	baseerrs "github.com/defipod/mochi/pkg/model/errors"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/response"
)

func (h *Handler) HandleDiscordWebhook(c *gin.Context) {
	var req request.HandleDiscordWebhookRequest
	if err := req.Bind(c); err != nil {
		h.log.Fields(logger.Fields{"body": req}).Error(err, "[handler.HandleDiscordWebhook] - failed to read JSON")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	if err := req.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
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
	case request.GUILD_DELETE:
		h.handleGuildDelete(c, req.Data)
	}
}

func (h *Handler) handleGuildMemberAdd(c *gin.Context, data json.RawMessage) {
	var member discordgo.Member
	byteData, err := data.MarshalJSON()
	if err != nil {
		h.log.Info("[handler.handleGuildMemberAdd] - failed to json marshal data")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	if err := discordgo.Unmarshal(byteData, &member); err != nil {
		h.log.Info("[handler.handleGuildMemberAdd] - failed to unmarshal data")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	h.handleInviteTracker(c, &member)
}

func (h *Handler) handleInviteTracker(c *gin.Context, invitee *discordgo.Member) {
	inviter, isVanity, err := h.entities.FindInviter(invitee.GuildID)
	if err != nil {
		h.log.Fields(logger.Fields{"invitee": invitee}).Error(err, "[handler.handleInviteTracker] - failed to find inviter")
	}

	data, err := h.entities.HandleInviteTracker(inviter, invitee)
	if err != nil {
		h.log.Fields(logger.Fields{"inviter": inviter, "invitee": invitee}).Error(err, "[handler.handleInviteTracker] - failed to handle invite tracker")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	data.IsVanity = isVanity

	c.JSON(http.StatusOK, response.CreateResponse(data, nil, nil, nil))
}

func (h *Handler) handleMessageCreate(c *gin.Context, data json.RawMessage) {
	message := &discordgo.Message{}
	byteData, err := data.MarshalJSON()
	if err != nil {
		h.log.Error(err, "[handler.handleMessageCreate] - failed to json marshal data")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	if err := discordgo.Unmarshal(byteData, &message); err != nil {
		h.log.Error(err, "[handler.handleMessageCreate] - failed to unmarshal data")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	uActivity, err := h.entities.HandleDiscordMessage(message)
	if err != nil {
		h.log.Fields(logger.Fields{"message": message}).Error(err, "[handler.handleMessageCreate] - failed to handle discord message")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
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
		h.log.Fields(logger.Fields{"message": message}).Error(err, "[handler.handleMessageCreate] - failed to handle user activity")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(resp, nil, nil, nil))
}

func (h *Handler) handleGuildCreate(c *gin.Context, data json.RawMessage) {
	var req struct {
		GuildID string `json:"guild_id"`
	}

	byteData, err := data.MarshalJSON()
	if err != nil {
		h.log.Error(err, "[handler.handleGuildCreate] - failed to json marshal data")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	if err := discordgo.Unmarshal(byteData, &req); err != nil {
		h.log.Error(err, "[handler.handleGuildCreate] - failed to unmarshal data")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	if err = h.entities.InitGuildDefaultTokenConfigs(req.GuildID); err != nil {
		h.log.Fields(logger.Fields{"guildID": req.GuildID}).Error(err, "[handler.handleGuildCreate] - failed to init default token configs")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	if err = h.entities.InitGuildDefaultActivityConfigs(req.GuildID); err != nil {
		h.log.Fields(logger.Fields{"guildID": req.GuildID}).Error(err, "[handler.handleGuildCreate] - failed to init default activity configs")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(response.ResponseMessage{Message: "OK"}, nil, nil, nil))
}

func (h *Handler) handleMessageReactionAdd(c *gin.Context, data json.RawMessage) {
	var req request.MessageReactionRequest
	byteData, err := data.MarshalJSON()
	if err != nil {
		h.log.Error(err, "[handler.handleMessageReactionAdd] - failed to json marshal data")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	if err := discordgo.Unmarshal(byteData, &req); err != nil {
		h.log.Error(err, "[handler.handleMessageReactionAdd] - failed to unmarshal data")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	if err := h.entities.AddMessageReaction(req); err != nil {
		h.log.Fields(logger.Fields{"body": req}).Error(err, "[handler.handleMessageReactionAdd] - failed to create message reaction")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	// starboard repost conversation
	repostConversation, err := h.entities.CreateRepostConversationReactionEvent(req)
	if err != nil {
		h.log.Fields(logger.Fields{"body": req}).Error(err, "[handler.handleMessageReactionAdd] - failed to create repost reaction event")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	if repostConversation != nil {
		c.JSON(http.StatusOK, response.CreateResponse(response.RepostReactionEventData{
			Status:               "OK",
			RepostChannelID:      repostConversation.RepostChannelID,
			ReactionType:         "conversation",
			OriginStartMessageID: repostConversation.OriginStartMessageID,
			OriginStopMessageID:  repostConversation.OriginStopMessageID,
		}, nil, nil, nil))
		return
	}

	// starboard repost message
	repostMessage, err := h.entities.CreateRepostMessageReactionEvent(req)
	if err != nil {
		h.log.Fields(logger.Fields{"body": req}).Error(err, "[handler.handleMessageReactionAdd] - failed to create repost reaction event")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	if repostMessage == nil {
		c.JSON(http.StatusOK, response.CreateResponse(response.RepostReactionEventData{Status: "OK"}, nil, nil, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(response.RepostReactionEventData{
		Status:          "OK",
		RepostChannelID: repostMessage.RepostChannelID,
		RepostMessageID: repostMessage.RepostMessageID,
	}, nil, nil, nil))
}

func (h *Handler) handleMessageReactionRemove(c *gin.Context, data json.RawMessage) {
	var req request.MessageReactionRequest
	byteData, err := data.MarshalJSON()
	if err != nil {
		h.log.Error(err, "[handler.handleMessageReactionRemove] - failed to json marshal data")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	if err := discordgo.Unmarshal(byteData, &req); err != nil {
		h.log.Error(err, "[handler.handleMessageReactionRemove] - failed to unmarshal data")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	if err := h.entities.RemoveMessageReaction(req); err != nil {
		h.log.Fields(logger.Fields{"body": req}).Error(err, "[handler.handleMessageReactionRemove] - failed to create message reaction")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(response.RepostReactionEventData{Status: "OK"}, nil, nil, nil))
}

func (h *Handler) handleMessageDelete(c *gin.Context, data json.RawMessage) {
	message := &discordgo.Message{}
	byteData, err := data.MarshalJSON()
	if err != nil {
		h.log.Error(err, "[handler.handleMessageDelete] - failed to json marshal data")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	if err := discordgo.Unmarshal(byteData, &message); err != nil {
		h.log.Error(err, "[handler.handleMessageDelete] - failed to unmarshal data")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	if err = h.entities.RemoveAllMessageReactions(message); err != nil {
		h.log.Fields(logger.Fields{"message": message}).Error(err, "[handler.handleMessageDelete] - failed to handle message delete")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(response.ResponseMessage{Message: "OK"}, nil, err, nil))
}

func (h *Handler) WebhookUpvoteTopGG(c *gin.Context) {
	req := request.WebhookUpvoteTopGG{}
	err := c.BindJSON(&req)
	if err != nil {
		h.log.Fields(logger.Fields{"body": req}).Error(err, "[handler.WebhookUpvoteTopGG] - failed to read JSON")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	err = h.entities.WebhookUpvoteStreak(req.UserID, consts.TopGGSource)
	if err != nil {
		h.log.Fields(logger.Fields{"body": req}).Error(err, "[handler.WebhookUpvoteTopGG] - failed to add upvote streak")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(response.ResponseMessage{Message: "OK"}, nil, nil, nil))
}

func (h *Handler) WebhookUpvoteDiscordBot(c *gin.Context) {
	req := request.WebhookUpvoteDiscordBot{}
	err := c.BindJSON(&req)
	if err != nil {
		h.log.Fields(logger.Fields{"body": req}).Error(err, "[handler.WebhookUpvoteDiscordBot] - failed to read JSON")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	err = h.entities.WebhookUpvoteStreak(req.UserID, consts.DiscordBotListSource)
	if err != nil {
		h.log.Fields(logger.Fields{"body": req}).Error(err, "[handler.WebhookUpvoteDiscordBot] - failed to add upvote streak")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(response.ResponseMessage{Message: "OK"}, nil, err, nil))
}

func (h *Handler) handleGuildDelete(c *gin.Context, data json.RawMessage) {
	var req request.HandleGuildDeleteRequest
	byteData, err := data.MarshalJSON()
	if err != nil {
		h.log.Error(err, "[handler.handleGuildDelete] - data.MarshalJSON() failed")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	if err := discordgo.Unmarshal(byteData, &req); err != nil {
		h.log.Error(err, "[handler.handleGuildDelete] - discordgo.Unmarshal() failed")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	err = h.entities.DeactivateGuild(req)
	if err != nil {
		code := http.StatusInternalServerError
		if err == baseerrs.ErrRecordNotFound {
			code = http.StatusNotFound
		}
		c.JSON(code, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	c.JSON(http.StatusOK, response.CreateResponse(response.ResponseMessage{Message: "OK"}, nil, nil, nil))
}
