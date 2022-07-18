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

func (h *Handler) TokenCompare(c *gin.Context) {
	sourceSymbol := c.Query("source_symbol")
	targetSymbol := c.Query("target_symbol")
	interval := c.Query("interval")

	if sourceSymbol == "" {
		h.log.Info("[handler.TokenCompare] - source symbol empty")
		c.JSON(http.StatusBadRequest, gin.H{"error": "source symbol is required"})
		return
	}

	if targetSymbol == "" {
		h.log.Info("[handler.TokenCompare] - target symbol empty")
		c.JSON(http.StatusBadRequest, gin.H{"error": "target symbol is required"})
		return
	}
	if interval == "" {
		h.log.Info("[handler.TokenCompare] - interval empty")
		c.JSON(http.StatusBadRequest, gin.H{"error": "interval is required"})
		return
	}

	// get search coin
	sourceSymbol, err, statusCode := h.entities.SearchCoinsBySymbol(sourceSymbol)
	if err != nil {
		h.log.Fields(logger.Fields{"sourceSymbol": sourceSymbol}).Error(err, "[handler.TokenCompare] - failed to search coin by symbol")
		c.JSON(statusCode, gin.H{"error": err.Error()})
		return
	}

	targetSymbol, err, statusCode = h.entities.SearchCoinsBySymbol(targetSymbol)
	if err != nil {
		h.log.Fields(logger.Fields{"targetSymbol": targetSymbol}).Error(err, "[handler.TokenCompare] - failed to search coin by symbol")
		c.JSON(statusCode, gin.H{"error": err.Error()})
		return
	}

	// get all token with guildID
	sourceSymbolInfo, err, statusCode := h.entities.GetHistoryCoinInfo(sourceSymbol, interval)
	if err != nil {
		h.log.Fields(logger.Fields{"sourceSymbol": sourceSymbol}).Error(err, "[handler.TokenCompare] - failed to get history coin info")
		c.JSON(statusCode, gin.H{"error": err.Error()})
		return
	}

	targetSymbolInfo, err, statusCode := h.entities.GetHistoryCoinInfo(targetSymbol, interval)
	if err != nil {
		h.log.Fields(logger.Fields{"targetSymbol": targetSymbol}).Error(err, "[handler.TokenCompare] - failed to get history coin info")
		c.JSON(statusCode, gin.H{"error": err.Error()})
		return
	}

	//check if one of 2 symbol is old
	if len(sourceSymbolInfo) != len(targetSymbolInfo) {
		h.log.Fields(logger.Fields{"sourceInfo": sourceSymbolInfo, "targetInfo": targetSymbolInfo}).Error(err, "[handler.TokenCompare] - token expired")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "One token is expired."})
		return
	}

	tokenCompareReponse, err := h.entities.TokenCompare(sourceSymbolInfo, targetSymbolInfo)
	if err != nil {
		h.log.Fields(logger.Fields{"sourceSymbol": sourceSymbol, "targetSymbol": targetSymbol}).Error(err, "[handler.TokenCompare] - failed to compare token")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": tokenCompareReponse})
}
