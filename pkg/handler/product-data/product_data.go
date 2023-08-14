package productdata

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

// ProductBotCommand     godoc
// @Summary     Get product bot commands
// @Description Get product bot commands
// @Tags        ProductMetadata
// @Accept      json
// @Produce     json
// @Param       req   query  request.ProductBotCommandRequest true  "request"
// @Success     200 {object} response.ProductBotCommand
// @Router      /product-metadata/commands [get]
func (h *Handler) ProductBotCommand(c *gin.Context) {
	req := request.ProductBotCommandRequest{}
	if err := c.ShouldBindQuery(&req); err != nil {
		h.log.Error(err, "[handler.ProductBotCommand] ShouldBindQuery() failed")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	commands, err := h.entities.ProductBotCommand(req)
	if err != nil {
		h.log.Error(err, "[handler.ProductBotCommand] entities.ProductBotCommand() failed")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse[any](commands, nil, nil, nil))
}
