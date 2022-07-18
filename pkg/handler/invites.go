package handler

import (
	"net/http"

	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/request"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetInvites(c *gin.Context) {
	memberID := c.Query("member_id")
	guildID := c.Query("guild_id")

	invites, err := h.entities.GetUserGlobalInviteCodes(guildID, memberID)
	if err != nil {
		h.log.Fields(logger.Fields{"memberID": memberID, "guildID": guildID}).Error(err, "[handler.GetInvites] - failed to get user global invite code")
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
		h.log.Info("[handler.GetInvitesLeaderboard] - guild id empty")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "guild_id is required",
		})
		return
	}

	leaderboard, err := h.entities.GetInvitesLeaderboard(guildID)
	if err != nil {
		h.log.Fields(logger.Fields{"guildID": guildID}).Error(err, "[handler.GetInvitesLeaderboard] - failed to get invite leaderboards")
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
		h.log.Info("[handler.GetInviteTrackerConfig] - guild id empty")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "guild_id is required",
		})
		return
	}

	config, err := h.entities.GetInviteTrackerLogChannel(guildID)
	if err != nil {
		h.log.Fields(logger.Fields{"guildID": guildID}).Error(err, "[handler.GetInviteTrackerConfig] - failed to get invite tracker log channel")
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
		h.log.Fields(logger.Fields{"body": req}).Error(err, "[handler.ConfigureInvites] - failed to read JSON")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := req.Validate(); err != nil {
		h.log.Fields(logger.Fields{"logChannel": req.LogChannel, "guildID": req.GuildID}).Error(err, "[handler.ConfigureInvites] - failed to validate request")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := h.entities.CreateOrUpdateInviteTrackerLogChannel(req); err != nil {
		h.log.Fields(logger.Fields{"body": req}).Error(err, "[handler.ConfigureInvites] - failed to upsert invite tracker log channel")
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
		h.log.Info("[handler.InvitesAggregation] - guild id empty")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "guild_id is required",
		})
		return
	}
	inviterID := c.Query("inviter_id")
	if inviterID == "" {
		h.log.Info("[handler.InvitesAggregation] - inviter id empty")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "inviter_id is required",
		})
		return
	}

	aggregation, err := h.entities.GetUserInvitesAggregation(guildID, inviterID)
	if err != nil {
		h.log.Fields(logger.Fields{"guildID": guildID, "inviterID": inviterID}).Error(err, "[handler.InvitesAggregation] - failed to get user invites aggregation")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": aggregation,
	})
}
