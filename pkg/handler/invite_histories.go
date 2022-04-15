package handler

import (
	"net/http"

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

	inviteHistory := &model.InviteHistory{
		GuildID:   body.GuildID,
		UserID:    body.Invitee,
		InvitedBy: body.Inviter,
	}

	if err := h.repo.InviteHistories.Create(inviteHistory); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := h.repo.GuildUsers.Update(body.GuildID, body.Invitee, "invited_by", body.Inviter); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": inviteHistory,
	})
}

func (h *Handler) CountByGuildUser(c *gin.Context) {
	guildIDstr, _ := c.GetQuery("guild_id")
	inviterStr, _ := c.GetQuery("inviter")

	count, err := h.repo.GuildUsers.CountByGuildUser(guildIDstr, inviterStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": count,
	})
}
