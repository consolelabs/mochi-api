package handler

import (
	"net/http"

	"github.com/defipod/mochi/pkg/request"
	"github.com/gin-gonic/gin"
)

func (h *Handler) HandlerGuildCustomTokenConfig(c *gin.Context) {
	var req request.UpsertCustomTokenConfigRequest

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

	// add to token schemas
	if err := h.entities.UpsertCustomToken(req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// get the Index of the row which has currently been added
	token, err := h.entities.GetTokenBySymbol(req.Symbol, true)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	req.Id = token

	// add to custom token config
	if err := h.entities.UpsertGuildCustomTokenConfig(req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "OK"})
}
