package handler

import (
	"fmt"
	"net/http"

	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/request"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetHistoricalMarketChart(c *gin.Context) {
	data, err, statusCode := h.entities.GetHistoricalMarketChart(c)
	if err != nil {
		h.log.Error(err, "[handler.GetHistoricalMarketChart] - failed to get historical market chart")
		c.JSON(statusCode, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": data})
}

func (h *Handler) InDiscordWalletTransfer(c *gin.Context) {
	var req request.TransferRequest
	if err := req.Bind(c); err != nil {
		h.log.Fields(logger.Fields{"body": req}).Error(err, "[handler.InDiscordWalletTransfer] - failed to read JSON")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, errs := h.entities.InDiscordWalletTransfer(req)
	if errs != nil {
		for _, err := range errs {
			fmt.Println("error transfer in dcwallet:", err)
		}
	}

	if len(res) == 0 {
		h.log.Fields(logger.Fields{"body": req}).Info("[handler.InDiscordWalletTransfer] - discord waller transfer returns empty response")
		c.JSON(http.StatusInternalServerError, gin.H{"errors": errs})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":   res,
		"errors": errs,
	})
}

func (h *Handler) InDiscordWalletWithdraw(c *gin.Context) {
	var req request.TransferRequest
	if err := req.Bind(c); err != nil {
		h.log.Fields(logger.Fields{"body": req}).Error(err, "[handler.InDiscordWalletWithdraw] - failed to read JSON")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.entities.InDiscordWalletWithdraw(req)
	if err != nil {
		h.log.Fields(logger.Fields{"body": req}).Error(err, "[handler.InDiscordWalletWithdraw] - failed to withdraw discord wallet")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *Handler) InDiscordWalletBalances(c *gin.Context) {
	guildID := c.Query("guild_id")
	if guildID == "" {
		h.log.Info("[handler.InDiscordWalletBalances] - guild id empty")
		c.JSON(http.StatusBadRequest, gin.H{"error": "guild_id is required"})
		return
	}

	discordID := c.Query("discord_id")
	if discordID == "" {
		h.log.Info("[handler.InDiscordWalletBalances] - discord id empty")
		c.JSON(http.StatusBadRequest, gin.H{"error": "discord_id is required"})
		return
	}

	response, err := h.entities.InDiscordWalletBalances(guildID, discordID)
	if err != nil {
		h.log.Fields(logger.Fields{"guildID": guildID, "discordID": discordID}).Error(err, "[handler.InDiscordWalletBalances] - failed to respond")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "ok", "data": response})
}

func (h *Handler) GetSupportedTokens(c *gin.Context) {
	tokens, err := h.entities.GetSupportedTokens()
	if err != nil {
		h.log.Error(err, "[handler.GetSupportedTokens] - failed to get supported tokens")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": tokens,
	})
}

func (h *Handler) GetCoin(c *gin.Context) {
	data, err, statusCode := h.entities.GetCoinData(c)
	if err != nil {
		h.log.Error(err, "[handler.GetCoin] - failed to get coin data")
		c.JSON(statusCode, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": data})
}

func (h *Handler) SearchCoins(c *gin.Context) {
	data, err, statusCode := h.entities.SearchCoins(c)
	if err != nil {
		h.log.Error(err, "[handler.SearchCoins] - failed to search coin data")
		c.JSON(statusCode, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": data})
}

func (h *Handler) CompareToken(c *gin.Context) {
	base := c.Query("base")
	target := c.Query("target")
	interval := c.Query("interval")

	if base == "" {
		h.log.Info("[handler.CompareToken] base is required")
		c.JSON(http.StatusBadRequest, gin.H{"error": "source symbol is required"})
		return
	}

	if target == "" {
		h.log.Info("[handler.CompareToken] target is required")
		c.JSON(http.StatusBadRequest, gin.H{"error": "target symbol is required"})
		return
	}
	if interval == "" {
		h.log.Info("[handler.CompareToken] interval empty")
		c.JSON(http.StatusBadRequest, gin.H{"error": "interval is required"})
		return
	}

	res, err := h.entities.CompareToken(base, target, interval)
	if err != nil {
		h.log.Fields(logger.Fields{"base": base, "target": target}).Error(err, "[handler.CompareToken] entity.CompareToken failed")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": res})
}
