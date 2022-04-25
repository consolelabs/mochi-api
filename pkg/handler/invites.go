package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetInvites(c *gin.Context) {
	memberID := c.Query("member_id")
	guildID := c.Query("guild_id")

	invites, err := h.entities.GetUserGlobalInviteCodes(guildID, memberID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": invites,
	})
}

func (h *Handler) GetInvitesLeaderboard(c *gin.Context) {
	guildID, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "guild_id is required",
		})
		return
	}

	leaderboard, err := h.entities.GetInvitesLeaderboard(guildID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": leaderboard,
	})
}
