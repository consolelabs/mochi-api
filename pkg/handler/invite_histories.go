package handler

import (
	"net/http"
	"strconv"

	"github.com/defipod/mochi/pkg/model"
	"github.com/gin-gonic/gin"
)

func (h *Handler) IndexInviteHistory(c *gin.Context) {
	var body struct {
		GuildID string `json:"guild_id"`
		Inviter string `json:"inviter"`
		Invitee string `json:"invitee"`
	}

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	guildID, err := strconv.ParseInt(body.GuildID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	inviter, err := strconv.ParseInt(body.Inviter, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	Invitee, err := strconv.ParseInt(body.Invitee, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	inviteHistory := &model.InviteHistory{
		GuildID:   guildID,
		UserID:    Invitee,
		InvitedBy: inviter,
	}

	if err := h.repo.InviteHistories.Create(inviteHistory); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := h.repo.GuildUsers.Update(guildID, Invitee, "invited_by", inviter); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": inviteHistory,
	})
}

func (h *Handler) CountByGuildUser(c *gin.Context) {
	guildIDstr, _ := c.GetQuery("guild_id")
	guildID, err := strconv.ParseInt(guildIDstr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	inviterStr, _ := c.GetQuery("inviter")
	inviter, err := strconv.ParseInt(inviterStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	count, err := h.repo.GuildUsers.CountByGuildUser(guildID, inviter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": count,
	})
}
