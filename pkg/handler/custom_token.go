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

	// set default
	req.Decimals = 18
	req.DiscordBotSupported = true
	req.GuildDefault = false
	req.Active = false

	// get the chainID
	returnChain, isFound, err := h.entities.GetChainIdBySymbol(req.Chain)
	if err != nil {
		h.log.Fields(logger.Fields{"chain": req.Chain}).Error(err, "[handler.HandlerGuildCustomTokenConfig] - failed to get chain data")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if !isFound {
		h.log.Fields(logger.Fields{"chain": req.Chain}).Error(err, "[handler.HandlerGuildCustomTokenConfig] - chain not supported")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Chain is not supported",
		})
		return
	}

	req.ChainID = returnChain.ID

	// check token exists or not
	checkExistToken, err := h.entities.CheckExistToken(req.Symbol)
	if err != nil {
		h.log.Fields(logger.Fields{"symbol": req.Symbol}).Error(err, "[handler.HandlerGuildCustomTokenConfig] - failed to get symbol data")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if !checkExistToken {
		// get the name and coin geck id
		id, name, err := h.entities.GetIDAndName(req.Symbol)
		if err != nil {
			h.log.Fields(logger.Fields{"symbol": req.Symbol}).Error(err, "[handler.HandlerGuildCustomTokenConfig] - failed to get symbol data from coingecko")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "can not get the name and coin geck id"})
			return
		}

		req.CoinGeckoID, req.Name = id, name

		// add to token schemas
		if err := h.entities.CreateCustomToken(req); err != nil {
			h.log.Fields(logger.Fields{"body": req}).Error(err, "[handler.HandlerGuildCustomTokenConfig] - failed to create custom token")
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	// get the Index of the row which has currently been added
	token, err := h.entities.GetTokenBySymbol(req.Symbol, true)
	if err != nil {
		h.log.Fields(logger.Fields{"symbol": req.Symbol}).Error(err, "[handler.HandlerGuildCustomTokenConfig] - failed to get token by symbol")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "cannot get the row index"})
		return
	}

	req.Id = token

	// check token config exists or not
	checkExistTokenConfig, err := h.entities.CheckExistTokenConfig(req.Id, req.GuildID)
	if err != nil {
		h.log.Fields(logger.Fields{"tokenID": req.Id, "guildID": req.GuildID}).Error(err, "[handler.HandlerGuildCustomTokenConfig] - failed to get token config")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if checkExistTokenConfig {
		h.log.Fields(logger.Fields{"tokenID": req.Id, "guildID": req.GuildID}).Error(err, "[handler.HandlerGuildCustomTokenConfig] - token already existed")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Your guild has already added this token."})
		return
	}

	// add to custom token config
	if err := h.entities.CreateGuildCustomTokenConfig(req); err != nil {
		h.log.Fields(logger.Fields{"body": req}).Error(err, "[handler.HandlerGuildCustomTokenConfig] - failed to create custom token config")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "cannot add to token config"})
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
