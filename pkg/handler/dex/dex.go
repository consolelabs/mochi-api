package dex

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/defipod/mochi/pkg/entities"
	"github.com/defipod/mochi/pkg/logger"
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

func (h *Handler) SumarizeBinanceAsset(c *gin.Context) {
	req := request.BinanceRequest{}
	if err := c.ShouldBindQuery(&req); err != nil {
		h.log.Error(err, "[handler.SumarizeBinanceAsset] BindJSON() failed")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	req.Id = c.Param("id")

	res, err := h.entities.SumarizeBinanceAsset(req)
	if err != nil {
		h.log.Error(err, "[handler.SumarizeBinanceAsset] entity.SumarizeBinanceAsset() failed")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(res, nil, nil, nil))
}

func (h *Handler) GetBinanceAssets(c *gin.Context) {
	req := request.GetBinanceAssetsRequest{
		Id:       c.Param("id"),
		Platform: c.Param("platform"),
	}

	if req.Platform != "binance" {
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("Unsupported dex"), nil))
		return
	}

	res, err := h.entities.GetBinanceAssets(req)
	if err != nil {
		h.log.Error(err, "[handler.GetBinanceAssets] entity.GetBinanceAssets() failed")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	c.JSON(http.StatusOK, response.CreateResponse(res, nil, nil, nil))
}

// GetBinanceFutures godoc
// @Summary     Get user's future account balance
// @Description Get user's future account balance
// @Tags        Binance
// @Accept      json
// @Produce     json
// @Param       id   			path  string true  "profile ID"
// @Success     200 {object} response.BinanceFutureAccountBalanceResponse
// @Router      /users/{id}/cexs/binance/futures [get]
func (h *Handler) GetBinanceFutures(c *gin.Context) {
	req := request.GetBinanceFutureRequest{}
	if err := c.ShouldBindUri(&req); err != nil {
		h.log.Error(err, "[handler.GetBinanceFutures] BindJSON() failed")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	res, err := h.entities.GetBinanceFutures(req)
	if err != nil {
		h.log.Error(err, "[handler.GetBinanceFutures] entity.GetBinanceFutures() failed")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	c.JSON(http.StatusOK, response.CreateResponse(res, nil, nil, nil))
}
