package data

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/defipod/mochi/pkg/entities"
	"github.com/defipod/mochi/pkg/logger"
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

func (h *Handler) AddGitbookClick(c *gin.Context) {
	url := c.Query("url")
	cmd := c.Query("command")
	action := c.Query("action")
	if url == "" || cmd == "" {
		h.log.Error(nil, "[handler.AddGitbookClick] - url and command are required")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("url and command are required"), nil))
	}

	err := h.entities.AddGitbookClick(url, cmd, action)
	if err != nil {
		h.log.Error(err, "[handler.AddGitbookClick] - faled to add click gitbook info")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	c.Redirect(http.StatusTemporaryRedirect, url)
}

// MetricByProperties   godoc
// @Summary     Metric
// @Description API to get stats of collections, users, servers, ...
// @Tags        Data
// @Accept      json
// @Produce     json
// @Param       q   query  string true  "total_servers | active_users | nft_collections | verified_wallets | supported_tokens | command_usage"
// @Param       guild_id   query  string false  "case active_users"
// @Success     200 {object} response.DataMetric
// @Router      /data/metrics [get]
func (h *Handler) MetricByProperties(c *gin.Context) {
	query := c.Query("q")

	switch query {
	case "nft_collections":
		h.MetricNftCollection(c, query)
	case "active_users":
		h.MetricActiveUsers(c, query, c.Query("guild_id"))
	case "total_servers":
		h.MetricTotalServers(c, query)
	case "verified_wallets":
		h.MetricVerifiedWallets(c, query, c.Query("guild_id"))
	case "supported_tokens":
		h.MetricSupportedTokens(c, query, c.Query("guild_id"))
	default:
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, nil, nil))
	}
}

func (h *Handler) MetricNftCollection(c *gin.Context, query string) {
	totalNftCollection, err := h.entities.TotalNftCollection()
	if err != nil {
		h.log.Fields(logger.Fields{"query": query}).Error(err, "[handler.MetricNftCollection] - failed to get total nft collection")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	c.JSON(http.StatusOK, response.CreateResponse(response.Metric{NftCollections: totalNftCollection}, nil, nil, nil))
}

func (h *Handler) MetricActiveUsers(c *gin.Context, query string, guildId string) {
	totalActiveUsers, err := h.entities.TotalActiveUsers(guildId)
	if err != nil {
		h.log.Fields(logger.Fields{"query": query}).Error(err, "[handler.MetricActiveUsers] - failed to get total active users")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	c.JSON(http.StatusOK, response.CreateResponse(totalActiveUsers, nil, nil, nil))
}

func (h *Handler) MetricTotalServers(c *gin.Context, query string) {
	totalServers, err := h.entities.TotalServers()
	if err != nil {
		h.log.Fields(logger.Fields{"query": query}).Error(err, "[handler.MetricTotalServers] - failed to get total servers")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	c.JSON(http.StatusOK, response.CreateResponse(totalServers, nil, nil, nil))
}

func (h *Handler) MetricVerifiedWallets(c *gin.Context, query string, guildId string) {
	totalVerifiedWallets, err := h.entities.TotalVerifiedWallets(guildId)
	if err != nil {
		h.log.Fields(logger.Fields{"query": query}).Error(err, "[handler.MetricVerifiedWallets] - failed to get total verified wallets")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	c.JSON(http.StatusOK, response.CreateResponse(totalVerifiedWallets, nil, nil, nil))
}

func (h *Handler) MetricSupportedTokens(c *gin.Context, query string, guildId string) {
	totalSupportedTokens, err := h.entities.TotalSupportedTokens(guildId)
	if err != nil {
		h.log.Fields(logger.Fields{"query": query}).Error(err, "[handler.MetricSupportedTokens] - failed to get total supported tokens")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	c.JSON(http.StatusOK, response.CreateResponse(totalSupportedTokens, nil, nil, nil))
}

// MetricProposalUsage   godoc
// @Summary     Metric
// @Description Get proposal usage across Mochi
// @Tags        Data
// @Accept      json
// @Produce     json
// @Param       page   query  string false  "page"
// @Param       size   query  string false  "size"
// @Success     200 {object} response.GuildProposalUsageResponse
// @Router      /data/usage-stats/proposal [get]
func (h *Handler) MetricProposalUsage(c *gin.Context) {
	page := c.Query("page")
	size := c.Query("size")
	if page == "" {
		page = "0"
	}
	if size == "" {
		size = "20"
	}
	res, err := h.entities.GetProposalUsage(page, size)
	if err != nil {
		h.log.Fields(logger.Fields{"page": page, "size": size}).Error(err, "[handler.MetricProposalUsage] - failed to get proposal data")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	c.JSON(http.StatusOK, response.CreateResponse(res, nil, nil, nil))
}

// MetricDaoTracker   godoc
// @Summary     Metric
// @Description Get dao tracker usage across Mochi
// @Tags        Data
// @Accept      json
// @Produce     json
// @Param       page   query  string false  "page"
// @Param       size   query  string false  "size"
// @Success     200 {object} response.DaoTrackerSpaceCountResponse
// @Router      /data/usage-stats/dao-tracker [get]
func (h *Handler) MetricDaoTracker(c *gin.Context) {
	page := c.Query("page")
	size := c.Query("size")
	if page == "" {
		page = "0"
	}
	if size == "" {
		size = "20"
	}
	res, err := h.entities.GetDaoTrackerMetric(page, size)
	if err != nil {
		h.log.Fields(logger.Fields{"page": page, "size": size}).Error(err, "[handler.MetricDaoTracker] - failed to get dao tracker data")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	c.JSON(http.StatusOK, response.CreateResponse(res, nil, nil, nil))
}
