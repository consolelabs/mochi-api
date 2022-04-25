package handler

import (
	"net/http"

	"github.com/defipod/mochi/pkg/request"
	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateGmConfig(c *gin.Context) {
	var req request.CreateGmConfigRequest

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if req.GuildID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "guild_id is required"})
	}
	if req.ChannelID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "channel_id is required"})
	}

	if err := h.entities.CreateGmConfig(req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "OK"})
}
