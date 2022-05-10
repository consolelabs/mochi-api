package handler

import (
	"net/http"

	"github.com/defipod/mochi/pkg/request"
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

func (h *Handler) GetInviteTrackerConfig(c *gin.Context) {
	guildID, exist := c.GetQuery("guild_id")
	if !exist {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "guild_id is required",
		})
		return
	}

	config, err := h.entities.GetInviteTrackerLogChannel(guildID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    config,
		"message": "OK",
	})
}

func (h *Handler) ConfigureInvites(c *gin.Context) {
	var req request.ConfigureInviteRequest
	if err := req.Bind(c); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := req.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := h.entities.CreateOrUpdateInviteTrackerLogChannel(req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": "ok",
	})
}

func (h *Handler) InvitesAggregation(c *gin.Context) {
	guildID := c.Query("guild_id")
	if guildID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "guild_id is required",
		})
		return
	}
	inviterID := c.Query("inviter_id")
	if inviterID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "inviter_id is required",
		})
		return
	}

	aggregation, err := h.entities.GetUserInvitesAggregation(guildID, inviterID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": aggregation,
	})
}
