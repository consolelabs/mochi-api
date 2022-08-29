package handler

import (
	"fmt"
	"net/http"

	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/response"
	"github.com/gin-gonic/gin"
)

// GetHistoricalMarketChart     godoc
// @Summary     Get historical market chart
// @Description Get historical market chart
// @Tags        Defi
// @Accept      json
// @Produce     json
// @Param       coin_id   path  string true  "Coin ID"
// @Param       day   path  int true  "Day"
// @Param       currency   path  string false  "Currency" default(usd)
// @Success     200 {object} response.GetHistoricalMarketChartResponse
// @Router      /defi/market-chart [get]
func (h *Handler) GetHistoricalMarketChart(c *gin.Context) {
	data, err, statusCode := h.entities.GetHistoricalMarketChart(c)
	if err != nil {
		h.log.Error(err, "[handler.GetHistoricalMarketChart] - failed to get historical market chart")
		c.JSON(statusCode, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response.GetHistoricalMarketChartResponse{Data: data})
}

// InDiscordWalletTransfer     godoc
// @Summary     In Discord Wallet transfer
// @Description In Discord Wallet transfer
// @Tags        Defi
// @Accept      json
// @Produce     json
// @Param       Request  body request.TransferRequest true "In Discord Wallet transfer request"
// @Success     200 {object} response.InDiscordWalletTransferResponseWrapper
// @Router      /defi/transfer [post]
func (h *Handler) InDiscordWalletTransfer(c *gin.Context) {
	var req request.TransferRequest
	if err := req.Bind(c); err != nil {
		h.log.Fields(logger.Fields{"body": req}).Error(err, "[handler.InDiscordWalletTransfer] - failed to read JSON")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, errs := h.entities.InDiscordWalletTransfer(req)
	if len(errs) != 0 {
		for _, err := range errs {
			fmt.Println("error transfer in dcwallet:", err)
		}
	}

	if len(res) == 0 {
		h.log.Fields(logger.Fields{"body": req}).Info("[handler.InDiscordWalletTransfer] - discord waller transfer returns empty response")
		c.JSON(http.StatusInternalServerError, gin.H{"errors": errs})
		return
	}

	c.JSON(http.StatusOK, response.InDiscordWalletTransferResponseWrapper{
		Data:   res,
		Errors: errs,
	})
}

// InDiscordWalletWithdraw     godoc
// @Summary     In Discord Wallet withdraw
// @Description In Discord Wallet withdraw
// @Tags        Defi
// @Accept      json
// @Produce     json
// @Param       Request  body request.TransferRequest true "In Discord Wallet withdraw request"
// @Success     200 {object} response.InDiscordWalletWithdrawResponse
// @Router      /defi/withdraw [post]
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

	c.JSON(http.StatusOK, gin.H{"data": res})
}

// InDiscordWalletBalances     godoc
// @Summary     In Discord Wallet balance
// @Description In Discord Wallet balance
// @Tags        Defi
// @Accept      json
// @Produce     json
// @Param       guild_id   query  string false  "Guild ID"
// @Param       discord_id   path  string true  "Discord ID"
// @Success     200 {object} response.InDiscordWalletBalancesResponse
// @Router      /defi/balances [get]
func (h *Handler) InDiscordWalletBalances(c *gin.Context) {
	guildID := c.Query("guild_id")
	discordID := c.Query("discord_id")
	if discordID == "" {
		h.log.Info("[handler.InDiscordWalletBalances] - discord id empty")
		c.JSON(http.StatusBadRequest, gin.H{"error": "discord_id is required"})
		return
	}

	data, err := h.entities.InDiscordWalletBalances(guildID, discordID)
	if err != nil {
		h.log.Fields(logger.Fields{"guildID": guildID, "discordID": discordID}).Error(err, "[handler.InDiscordWalletBalances] - failed to respond")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response.InDiscordWalletBalancesResponse{Status: "ok", Data: data})
}

// GetSupportedTokens     godoc
// @Summary     Get supported tokens
// @Description Get supported tokens
// @Tags        Defi
// @Accept      json
// @Produce     json
// @Success     200 {object} response.GetSupportedTokensResponse
// @Router      /defi/tokens [get]
func (h *Handler) GetSupportedTokens(c *gin.Context) {
	tokens, err := h.entities.GetSupportedTokens()
	if err != nil {
		h.log.Error(err, "[handler.GetSupportedTokens] - failed to get supported tokens")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response.GetSupportedTokensResponse{
		Data: tokens,
	})
}

// GetCoin     godoc
// @Summary     Get coin
// @Description Get coin
// @Tags        Defi
// @Accept      json
// @Produce     json
// @Param       id   path  string true  "Coin ID"
// @Success     200 {object} response.GetCoinResponseWrapper
// @Router      /defi/coins/{id} [get]
func (h *Handler) GetCoin(c *gin.Context) {
	data, err, statusCode := h.entities.GetCoinData(c)
	if err != nil {
		h.log.Error(err, "[handler.GetCoin] - failed to get coin data")
		c.JSON(statusCode, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response.GetCoinResponseWrapper{Data: data})
}

// SearchCoins     godoc
// @Summary     Search coin
// @Description Search coin
// @Tags        Defi
// @Accept      json
// @Produce     json
// @Param       query   query  string true  "coin query"
// @Success     200 {object} response.SearchCoinsResponse
// @Router      /defi/coins [get]
func (h *Handler) SearchCoins(c *gin.Context) {
	data, err, statusCode := h.entities.SearchCoins(c)
	if err != nil {
		h.log.Error(err, "[handler.SearchCoins] - failed to search coin data")
		c.JSON(statusCode, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response.SearchCoinsResponse{Data: data})
}

// CompareToken     godoc
// @Summary     Compare token
// @Description Compare token
// @Tags        Defi
// @Accept      json
// @Produce     json
// @Param       base   query  string true  "base token"
// @Param       target   query  string true  "target token"
// @Param       interval   query  string true  "compare interval"
// @Param       guild_id   query  string true  "Guild ID"
// @Success     200 {object} response.CompareTokenResponse
// @Router      /defi/coins/compare [get]
func (h *Handler) CompareToken(c *gin.Context) {
	base := c.Query("base")
	target := c.Query("target")
	interval := c.Query("interval")
	guildID := c.Query("guild_id")

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
	if guildID == "" {
		h.log.Info("[handler.CompareToken] guild_id empty")
		c.JSON(http.StatusBadRequest, gin.H{"error": "guild_id is required"})
		return
	}

	res, err := h.entities.CompareToken(base, target, interval, guildID)
	if err != nil {
		h.log.Fields(logger.Fields{"base": base, "target": target}).Error(err, "[handler.CompareToken] entity.CompareToken failed")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response.CompareTokenResponse{Data: res})
}

// SetGuildDefaultTicker     godoc
// @Summary     Set guild default ticker
// @Description Set guild default ticker
// @Tags        Config
// @Accept      json
// @Produce     json
// @Param       Request  body request.GuildConfigDefaultTickerRequest true "Set guild default ticker request"
// @Success     200 {object} response.ResponseMessage
// @Router      /configs/default-ticker [post]
func (h *Handler) SetGuildDefaultTicker(c *gin.Context) {
	req := request.GuildConfigDefaultTickerRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		h.log.Error(err, "[handler.SetGuildDefaultTicker] c.ShouldBindJSON failed")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.entities.SetGuildDefaultTicker(req); err != nil {
		h.log.Error(err, "[handler.SetGuildDefaultTicker] entity.SetGuildDefaultTicker failed")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, response.ResponseMessage{Message: "OK"})
}

// GetGuildDefaultTicker     godoc
// @Summary     Get guild default ticker
// @Description Get guild default ticker
// @Tags        Config
// @Accept      json
// @Produce     json
// @Param       guild_id   query  string true  "Guild ID"
// @Param       query   query  string true  "Guild ticker query"
// @Success     200 {object} response.GetGuildDefaultTickerResponse
// @Router      /configs/default-ticker [get]
func (h *Handler) GetGuildDefaultTicker(c *gin.Context) {
	var q request.GetGuildDefaultTickerQuery
	if err := c.ShouldBindQuery(&q); err != nil {
		h.log.Error(err, "[handler.GetGuildDefaultTicker] ShouldBindQuery failed")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := h.entities.GetGuildDefaultTicker(q)
	if err != nil {
		h.log.Error(err, "[handler.GetGuildDefaultTicker] entity.GetGuildDefaultTicker failed")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}
