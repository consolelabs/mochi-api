package handler

import (
	"net/http"

	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/response"
	"github.com/gin-gonic/gin"
)

// MetricByProperties   godoc
// @Summary     Metric
// @Description API to get stats of collections, users, servers, ...
// @Tags        Metric
// @Accept      json
// @Produce     json
// @Param       q   query  string true  "total_servers | active_users | nft_collections"
// @Param       guild_id   query  string false  "case active_users"
// @Success     200 {object} response.DataMetric
// @Router      /metrics [get]
func (h *Handler) MetricByProperties(c *gin.Context) {
	query := c.Query("q")

	switch query {
	case "nft_collections":
		h.MetricNftCollection(c, query)
	case "active_users":
		h.MetricActiveUsers(c, query, c.Query("guild_id"))
	case "total_servers":
		h.MetricTotalServers(c, query)
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
