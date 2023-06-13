package apikey

import (
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

// CreateApiKey     godoc
// @Summary     Create apiKey
// @Description Create apiKey
// @Tags        ApiKey
// @Accept      json
// @Produce     json
// @Param       Authorization header   string true "Authorization"
// @Success     200 {object} response.ProfileApiKeyResponse
// @Router      /api-key/me [post]
func (h *Handler) CreateApiKey(c *gin.Context) {
	profileAccessToken := c.GetString("profile_access_token")

	data, err := h.entities.CreateApiKey(profileAccessToken)
	if err != nil {
		h.log.Fields(logger.Fields{"profileAccessToken": profileAccessToken}).Error(err, "[handler.CreateApiKey] - fail to create apiKey")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	c.JSON(http.StatusOK, response.CreateResponse[any](data, nil, nil, nil))
}

// IntegrateBinanceKey     godoc
// @Summary     Integrate binance key
// @Description Integrate binance key
// @Tags        ApiKey
// @Accept      json
// @Produce     json
// @Success     200 {object} response.ProfileApiKeyResponse
// @Router      /api-key/binance [post]
func (h *Handler) IntegrateBinanceKey(c *gin.Context) {
	var req request.IntegrationBinanceData
	if err := c.BindJSON(&req); err != nil {
		h.log.Error(err, "[handler.IntegrateBinanceKey] ShouldBindQuery() failed")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	data, err := h.entities.IntegrateBinanceData(req)
	if err != nil {
		h.log.Error(err, "[handler.IntegrateBinanceKey] failed to get integrate binance data")
		c.JSON(baseerrs.GetStatusCode(err), response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse[any](data, nil, nil, nil))
}

// UnlinkBinance     godoc
// @Summary     Unlink binance
// @Description Unlink binance
// @Tags        ApiKey
// @Accept      json
// @Produce     json
// @Param       Request  body request.UnlinkBinance true "Unlink Binance request"
// @Success     200 {object} response.UnlinkBinance
// @Router      /api-key/unlink-binance [post]
func (h *Handler) UnlinkBinance(c *gin.Context) {
	var req request.UnlinkBinance
	if err := c.BindJSON(&req); err != nil {
		h.log.Error(err, "[handler.UnlinkBinance] BindJSON() failed")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

  err := h.entities.UnlinkBinance(req)
	if err != nil {
		h.log.Error(err, "[handler.UnlinkBinance] failed to unlink binance")
		c.JSON(baseerrs.GetStatusCode(err), response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse[any](gin.H{"message": "ok"}, nil, nil, nil))
}
