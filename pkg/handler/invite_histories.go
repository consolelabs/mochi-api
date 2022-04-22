package handler

import (
	"net/http"

	"github.com/defipod/mochi/pkg/request"
	"github.com/gin-gonic/gin"
)

func (h *Handler) IndexInviteHistory(c *gin.Context) {
	var req request.CreateInviteHistoryRequest

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.entities.CreateInviteHistory(req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}

func (h *Handler) CountByGuildUser(c *gin.Context) {
	guildID, guildIDExist := c.GetQuery("guild_id")
	if !guildIDExist {
		c.JSON(http.StatusBadRequest, gin.H{"error": "guild_id is required"})
		return
	}

	inviter, inviterExist := c.GetQuery("inviter")
	if !inviterExist {
		c.JSON(http.StatusBadRequest, gin.H{"error": "inviter is required"})
		return
	}

	count, err := h.entities.CountInviteHistoriesByGuildUser(guildID, inviter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": count,
	})
}
