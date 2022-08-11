package handler

import (
	"net/http"

	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/request"
	"github.com/gin-gonic/gin"
)

func (h *Handler) HandlerGuildCustomTokenConfig(c *gin.Context) {
	var req request.UpsertCustomTokenConfigRequest

	// handle input validate
	if err := c.BindJSON(&req); err != nil {
		h.log.Fields(logger.Fields{"body": req}).Error(err, "[handler.HandlerGuildCustomTokenConfig] - failed to read JSON")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if req.GuildID == "" {
		h.log.Info("[handler.HandlerGuildCustomTokenConfig] - guild id empty")
		c.JSON(http.StatusBadRequest, gin.H{"error": "guild_id is required"})
		return
	}
	if req.Symbol == "" {
		h.log.Info("[handler.HandlerGuildCustomTokenConfig] - symbol empty")
		c.JSON(http.StatusBadRequest, gin.H{"error": "symbol is required"})
		return
	}
	if req.Address == "" {
		h.log.Info("[handler.HandlerGuildCustomTokenConfig] - address empty")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Address is required"})
		return
	}
	if req.Chain == "" {
		h.log.Info("[handler.HandlerGuildCustomTokenConfig] - chain empty")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Chain is required"})
		return
	}

	if err := h.entities.CreateCustomToken(req); err != nil {
		h.log.Fields(logger.Fields{"req": req}).Error(err, "[handler.HandlerGuildCustomTokenConfig] - fail to create custom token")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "OK"})
}

func (h *Handler) ListAllCustomToken(c *gin.Context) {
	guildID := c.Param("guild_id")

	// get all token with guildID
	returnToken, err := h.entities.GetAllSupportedToken(guildID)
	if err != nil {
		h.log.Fields(logger.Fields{"guildID": guildID}).Error(err, "[handler.ListAllCustomToken] - failed to get all tokens")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": returnToken})
}
