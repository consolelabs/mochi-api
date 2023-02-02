package configtwittersale

import (
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

// Get     			godoc
// @Summary     Get twitter sale config
// @Description Get twitter sale config
// @Tags        ConfigTwitterSale
// @Accept      json
// @Produce     json
// @Param       marketplace   query  string true  "marketplace name"
// @Success     200 {object} response.GetSaleTwitterConfigResponse
// @Router      /configs/twitter-sales [get]
func (h *Handler) Get(c *gin.Context) {
	marketplace := c.Query("marketplace")
	configs, err := h.entities.GetSaleBotTwitterConfigs(marketplace)
	if err != nil {
		h.log.Fields(logger.Fields{"marketplace": marketplace}).Error(err, "[handler.ConfigTwitterSale.Get] entity.GetSaleBotTwitterConfigs() failed")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	c.JSON(http.StatusOK, response.CreateResponse(configs, nil, nil, nil))
}

// Create     	godoc
// @Summary     Create twitter sale config
// @Description Create twitter sale config
// @Tags        ConfigTwitterSale
// @Accept      json
// @Produce     json
// @Param       request  body request.CreateTwitterSaleConfigRequest true "req"
// @Success     200 {object} response.CreateTwitterSaleConfigResponse
// @Router      /configs/twitter-sales [post]
func (h *Handler) Create(c *gin.Context) {
	var req request.CreateTwitterSaleConfigRequest
	if err := c.BindJSON(&req); err != nil {
		h.log.Fields(logger.Fields{"req": req}).Error(err, "[handler.ConfigTwitterSale.Create] BindJSON() failed")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	config, err := h.entities.CreateSaleBotTwitterConfig(req)
	if err != nil {
		h.log.Fields(logger.Fields{"req": req}).Error(err, "[handler.ConfigTwitterSale.Create] entity.CreateSaleBotTwitterConfig() failed")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	c.JSON(http.StatusOK, response.CreateResponse[any](config, nil, nil, nil))
}
