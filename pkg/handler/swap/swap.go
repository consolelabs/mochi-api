package swap

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

// GetSwapRoutes     godoc
// @Summary     Get swap route for token pairs
// @Description Get swap route for token pairs
// @Tags        Swap
// @Accept      json
// @Produce     json
// @Param       from   query  string true  "from token symbol"
// @Param       to   query  string true  "to token symbol"
// @Param       amount   query  string true  "from amount value"
// @Success     200 {object} response.KyberSwapRoutes
// @Router      /swap/route [get]
func (h *Handler) GetSwapRoutes(c *gin.Context) {
	req := request.GetSwapRouteRequest{
		From:   c.Query("from"),
		To:     c.Query("to"),
		Amount: c.Query("amount"),
	}

	data, err := h.entities.GetSwapRoutes(&req)
	if err != nil {
		h.log.Fields(logger.Fields{"from": req.From, "to": req.To, "amount": req.Amount}).Error(err, "[handler.GetSwapRoutes] - cannot get data from kyber")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse[any](data, nil, nil, nil))
}
