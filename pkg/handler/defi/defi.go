package defi

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/defipod/mochi/pkg/entities"
	"github.com/defipod/mochi/pkg/logger"
	baseerrs "github.com/defipod/mochi/pkg/model/errors"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/response"
)

type Handler struct {
	entities *entities.Entity
	log      logger.Logger
}

func New(entities *entities.Entity, logger logger.Logger) IHandler {
	return &Handler{
		entities: entities,
		log:      logger,
	}
}

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
	var req request.GetMarketChartRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		h.log.Error(err, "[handler.GetHistoricalMarketChart] ShouldBindQuery() failed")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	data, err, statusCode := h.entities.GetHistoricalMarketChart(&req)
	if err != nil {
		h.log.Error(err, "[handler.GetHistoricalMarketChart] - failed to get historical market chart")
		c.JSON(statusCode, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(data, nil, nil, nil))
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

	c.JSON(http.StatusOK, response.CreateResponse(tokens, nil, nil, nil))
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
	coinID := c.Param("id")
	if coinID == "" {
		h.log.Info("[handler.GetCoin] - coin id missing")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("id is required"), nil))
		return
	}

	data, err, statusCode := h.entities.GetCoinData(coinID)
	if err != nil {
		h.log.Error(err, "[handler.GetCoin] - failed to get coin data")
		c.JSON(statusCode, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(data, nil, nil, nil))
}

// SearchCoins     godoc
// @Summary     Search coin
// @Description Search coin
// @Tags        Defi
// @Accept      json
// @Produce     json
// @Param       query   query  string true  "coin query"
// @Success     200 {object} response.SearchCoinResponse
// @Router      /defi/coins [get]
func (h *Handler) SearchCoins(c *gin.Context) {
	query := c.Query("query")
	if query == "" {
		h.log.Info("[handler.SearchCoins] query is required")
		c.JSON(http.StatusBadRequest, gin.H{"error": "query is required"})
		return
	}

	tokens, err := h.entities.SearchCoins(query)
	if err != nil {
		h.log.Error(err, "[handler.SearchCoins] entities.SearchCoins() failed")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(tokens, nil, nil, nil))
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
// @Param       guild_id   query  string false  "Guild ID"
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

	res, err := h.entities.CompareToken(base, target, interval, guildID)
	if err != nil {
		h.log.Fields(logger.Fields{"base": base, "target": target}).Error(err, "[handler.CompareToken] entity.CompareToken failed")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(res, nil, nil, nil))
}

// GetUserWatchlist     godoc
// @Summary     Get user's watchlist
// @Description Get user's watchlist
// @Tags        Defi
// @Accept      json
// @Produce     json
// @Param       req query request.GetUserWatchlistRequest true "request"
// @Success     200 {object} response.GetWatchlistResponse
// @Router      /defi/watchlist [get]
func (h *Handler) GetUserWatchlist(c *gin.Context) {
	var req request.GetUserWatchlistRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		h.log.Error(err, "[handler.AddToWatchlist] ShouldBindQuery() failed")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := h.entities.GetUserWatchlist(req)
	if err != nil {
		h.log.Error(err, "[handler.AddToWatchlist] entity.GetUserWatchlist() failed")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response.CreateResponse(res, nil, nil, nil))
}

// AddToWatchlist     godoc
// @Summary     Add to user's watchlist
// @Description Add to user's watchlist
// @Tags        Defi
// @Accept      json
// @Produce     json
// @Param       req body request.AddToWatchlistRequest true "request"
// @Success     200 {object} response.AddToWatchlistResponse
// @Router      /defi/watchlist [post]
func (h *Handler) AddToWatchlist(c *gin.Context) {
	var req request.AddToWatchlistRequest
	if err := c.BindJSON(&req); err != nil {
		h.log.Error(err, "[handler.AddToWatchlist] Bind() failed")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := h.entities.AddToWatchlist(req)
	if err != nil {
		h.log.Error(err, "[handler.AddToWatchlist] entity.AddToWatchlist() failed")
		c.JSON(baseerrs.GetStatusCode(err), gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// RemoveFromWatchlist     godoc
// @Summary     Remove from user's watchlist
// @Description Remove from user's watchlist
// @Tags        Defi
// @Accept      json
// @Produce     json
// @Param       req query request.RemoveFromWatchlistRequest true "request"
// @Success     200 {object} object
// @Router      /defi/watchlist [delete]
func (h *Handler) RemoveFromWatchlist(c *gin.Context) {
	var req request.RemoveFromWatchlistRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		h.log.Error(err, "[handler.RemoveFromWatchlist] Bind() failed")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := h.entities.RemoveFromWatchlist(req)
	if err != nil {
		h.log.Error(err, "[handler.RemoveFromWatchlist] entity.RemoveFromWatchlist() failed")
		code := http.StatusInternalServerError
		if err == baseerrs.ErrRecordNotFound {
			code = http.StatusNotFound
		}
		c.JSON(code, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response.CreateResponse[any](nil, nil, nil, nil))
}

// GetFiatHistoricalExchangeRates     godoc
// @Summary     Get historical market chart
// @Description Remove from user's watchlist
// @Tags        Defi
// @Accept      json
// @Produce     json
// @Param       req query request.GetFiatHistoricalExchangeRatesRequest true "request"
// @Success     200 {object} response.GetFiatHistoricalExchangeRatesResponse
// @Router      /defi/watchlist [delete]
func (h *Handler) GetFiatHistoricalExchangeRates(c *gin.Context) {
	var req request.GetFiatHistoricalExchangeRatesRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		h.log.Error(err, "[handler.GetFiatHistoricalExchangeRates] ShouldBindQuery() failed")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := h.entities.GetFiatHistoricalExchangeRates(req)
	if err != nil {
		h.log.Error(err, "[handler.GetFiatHistoricalExchangeRates] entity.GetFiatHistoricalExchangeRates() failed")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response.CreateResponse(res, nil, nil, nil))
}

// AddContract   godoc
// @Summary     List All Chain
// @Description List All Chain
// @Tags        Chain
// @Accept      json
// @Produce     json
// @Success     200 {object} response.GetListAllChainsResponse
// @Router      /chains [get]
func (h *Handler) ListAllChain(c *gin.Context) {
	returnChain, err := h.entities.ListAllChain()
	if err != nil {
		h.log.Error(err, "[handler.ListAllChain] - failed to list all chains")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(returnChain, nil, nil, nil))
}
