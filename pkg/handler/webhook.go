package handler

import (
	"encoding/json"
	"net/http"

	"github.com/defipod/mochi/pkg/consts"
	"github.com/defipod/mochi/pkg/logger"

	"github.com/bwmarrin/discordgo"
	"github.com/defipod/mochi/pkg/model"
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

	switch req.Event {
	case request.GUILD_MEMBER_ADD:
		h.handleGuildMemberAdd(c, req.Data)
	case request.MESSAGE_CREATE:
		h.handleMessageCreate(c, req.Data)
	case request.GUILD_CREATE:
		h.handleGuildCreate(c, req.Data)
	case request.MESSAGE_REACTION_ADD:
		h.handleMessageReactionAdd(c, req.Data)
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

// TODO: move logic to entity layer
func (h *Handler) handleInviteTracker(c *gin.Context, invitee *discordgo.Member) {
	var response response.HandleInviteHistoryResponse

	inviter, isVanity, err := h.entities.FindInviter(invitee.GuildID)
	if err != nil {
		h.log.Fields(logger.Fields{"invitee": invitee}).Error(err, "[handler.handleInviteTracker] - failed to find inviter")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	response.IsVanity = isVanity

	if inviter != nil {
		if err := h.entities.CreateUser(request.CreateUserRequest{
			ID:       inviter.User.ID,
			Username: inviter.User.Username,
			Nickname: inviter.Nick,
			GuildID:  inviter.GuildID,
		}); err != nil {
			h.log.Fields(logger.Fields{"inviterID": inviter.User.ID, "inviterUsrName": inviter.User.Username, "inviterNickName": inviter.Nick, "inviterGuildID": inviter.GuildID}).Error(err, "[handler.handleInviteTracker] - failed to create user")
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		response.InviterID = inviter.User.ID
		if inviter.User.Bot {
			response.IsBot = true
		}
	}
	if invitee != nil {
		if err := h.entities.CreateUser(request.CreateUserRequest{
			ID:        invitee.User.ID,
			Username:  invitee.User.Username,
			Nickname:  invitee.Nick,
			GuildID:   invitee.GuildID,
			InvitedBy: inviter.User.ID,
		}); err != nil {
			h.log.Fields(logger.Fields{"inviteeID": invitee.User.ID, "inviteeUsrName": invitee.User.Username, "inviteeNickName": invitee.Nick, "inviteeGuildID": invitee.GuildID}).Error(err, "[handler.handleInviteTracker] - failed to create user")
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		response.InviteeID = invitee.User.ID
	}

	inviteType := model.INVITE_TYPE_NORMAL
	if inviter == nil {
		inviteType = model.INVITE_TYPE_LEFT
	}

	// TODO: Can't find age of user now
	// if time.Now().Unix()-invit < 60*60*24*3 {
	// 	inviteType = model.INVITE_TYPE_FAKE
	// }

	if err := h.entities.CreateInviteHistory(request.CreateInviteHistoryRequest{
		GuildID: invitee.GuildID,
		Inviter: inviter.User.ID,
		Invitee: invitee.User.ID,
		Type:    inviteType,
	}); err != nil {
		h.log.Fields(logger.Fields{"inviteeID": invitee.User.ID, "inviterID": invitee.User.ID, "inviteType": inviteType, "inviteeGuildID": invitee.GuildID}).Error(err, "[handler.handleInviteTracker] - failed to create invite history")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	totalInvites, err := h.entities.CountInviteHistoriesByGuildUser(inviter.GuildID, inviter.User.ID)
	if err != nil {
		h.log.Fields(logger.Fields{"inviterID": invitee.User.ID, "inviterGuildID": inviter.GuildID}).Error(err, "[handler.handleInviteTracker] - failed to count inviter invites")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response.InvitesAmount = int(totalInvites)
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
	var req request.CreateMessageRepostHistRequest

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

	chanID, err := h.entities.CreateRepostReactionEvent(req)
	if err != nil {
		h.log.Fields(logger.Fields{"body": req}).Error(err, "[handler.handleMessageReactionAdd] - failed to create repost reaction event")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":            "OK",
		"repost_channel_id": chanID,
	})
}
