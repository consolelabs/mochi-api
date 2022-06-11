package handler

import (
	"net/http"

	"github.com/defipod/mochi/pkg/request"
	"github.com/gin-gonic/gin"
)

func (h *Handler) HandlerGuildCustomTokenConfig(c *gin.Context) {
	var req request.UpsertCustomTokenConfigRequest

	// handle input validate
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
	if req.Address == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Address is required"})
		return
	}
	if req.Chain == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Chain is required"})
		return
	}

	// set default
	req.Decimals = 18
	req.DiscordBotSupported = true
	req.GuildDefault = false
	req.Active = false

	// get the name and coin geck id
	id, name, err := h.entities.GetIDAndName(req.Symbol)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "can not get the name and coin geck id"})
		return
	}

	req.CoinGeckoID, req.Name = id, name

	// get the chainID
	chainId, err := h.entities.GetChainIdBySymbol(req.Chain)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	req.ChainID = chainId

	// add to token schemas
	if err := h.entities.CreateCustomToken(req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "cannot add to token schema"})
		return
	}

	// get the Index of the row which has currently been added
	token, err := h.entities.GetTokenBySymbol(req.Symbol, true)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "cannot get the row index"})
		return
	}

	req.Id = token

	// add to custom token config
	if err := h.entities.CreateGuildCustomTokenConfig(req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "cannot add to token config"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "OK"})
}
