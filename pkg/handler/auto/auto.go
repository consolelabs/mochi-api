package auto

import (
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

// GetAutoTriggers     godoc
// @Summary     Get auto triggers
// @Description Get auto triggers
// @Tags        Swap
// @Accept      json
// @Produce     json
// @Param       user_id   query  string true  "user id"
// @Param       guild_id   query  string true  "guild id"
// @Success     200 {object} response.SwapRouteResponseData
// @Router      /swap/route [get]
func (h *Handler) GetAutoTriggers(c *gin.Context) {
	c.JSON(http.StatusOK, response.CreateResponse[any](nil, nil, nil, nil))
}
