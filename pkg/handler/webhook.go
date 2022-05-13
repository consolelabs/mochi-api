package handler

import (
	"encoding/json"
	"net/http"

	"github.com/bwmarrin/discordgo"
	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (h *Handler) HandleDiscordWebhook(c *gin.Context) {
	var req request.HandleDiscordWebhookRequest
	if err := req.Bind(c); err != nil {
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
	}
}

func (h *Handler) handleGuildMemberAdd(c *gin.Context, data json.RawMessage) {
	var member discordgo.Member
	byteData, err := data.MarshalJSON()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := discordgo.Unmarshal(byteData, &member); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	h.handleInviteTracker(c, &member)
}

func (h *Handler) handleInviteTracker(c *gin.Context, invitee *discordgo.Member) {
	var response response.HandleInviteHistoryResponse

	inviter, isVanity, err := h.entities.FindInviter(invitee.GuildID)
	if err != nil {
		logrus.WithError(err).Errorf("Guild %s: failed to find inviter", invitee.GuildID)
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
			logrus.WithError(err).Errorf("Guild %s: failed to index iviter", invitee.GuildID)
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
			InvitedBy: invitee.User.ID,
		}); err != nil {
			logrus.WithError(err).Errorf("Guild %s: failed to index invitee", invitee.GuildID)
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
		logrus.WithError(err).Errorf("Guild %s: failed to create invite history", invitee.GuildID)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	totalInvites, err := h.entities.CountInviteHistoriesByGuildUser(inviter.GuildID, inviter.User.ID)
	if err != nil {
		logrus.WithError(err).Errorf("Guild %s: failed to count inviter invites", invitee.GuildID)
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := discordgo.Unmarshal(byteData, &message); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err = h.entities.HandleDiscordMessage(message); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// TODO: use response data to send discord message to user
	resp, err := h.entities.HandleUserActivities(&request.HandleUserActivityRequest{
		GuildID:   message.GuildID,
		ChannelID: message.ChannelID,
		UserID:    message.Author.ID,
		Action:    "chat",
		Timestamp: message.Timestamp,
	})
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := discordgo.Unmarshal(byteData, &req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err = h.entities.InitGuildDefaultTokenConfigs(req.GuildID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err = h.entities.InitGuildDefaultActivityConfigs(req.GuildID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}
