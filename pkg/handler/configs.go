package handler

import (
	"net/http"

	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/request"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetGmConfig(c *gin.Context) {
	guildID := c.Query("guild_id")
	if guildID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "guild_id is required"})
	}

	config, err := h.entities.GetGmConfig(guildID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "OK", "data": config})
}

func (h *Handler) UpsertGmConfig(c *gin.Context) {
	var req request.UpsertGmConfigRequest

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

	if err := h.entities.UpsertGmConfig(req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "OK"})
}

func (h *Handler) GetGuildTokens(c *gin.Context) {
	guildID := c.Query("guild_id")
	if guildID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "guild_id is required"})
	}

	guildTokens, err := h.entities.GetGuildTokens(guildID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var data []model.Token
	for _, gToken := range guildTokens {
		data = append(data, *gToken.Token)
	}

	c.JSON(http.StatusOK, gin.H{"data": data})
}

func (h *Handler) UpsertGuildTokenConfig(c *gin.Context) {
	var req request.UpsertGuildTokenConfigRequest

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if req.GuildID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "guild_id is required"})
		return
	}
	if req.Symbol == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "symbol is required"})
		return
	}

	if err := h.entities.UpsertGuildTokenConfig(req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "OK"})
}

func (h *Handler) ConfigLevelRole(c *gin.Context) {
	var req request.ConfigLevelRoleRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if req.GuildID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "guild_id is required"})
		return
	}
	if req.RoleID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "role_id is required"})
		return
	}
	if req.Level == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid level"})
		return
	}

	if err := h.entities.ConfigLevelRole(req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "OK"})
}
