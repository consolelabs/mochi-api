package invest

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

// GetInvestList     godoc
// @Summary     Get invest list
// @Description Get invest list
// @Tags        Invest
// @Accept      json
// @Produce     json
// @Param       chainIds   query  string false  "the filterd chain ids, split by comma"
// @Param       platforms   query  string false  "the filterd platforms (aave_v2, aave_v3), split by comma"
// @Param       types   query  string false  "the filterd types (stake, lend), split by comma"
// @Param       address   query  string false  "the filtered token address"
// @Param       status   query  string false  "the filtered status (active, inactive)"
// @Success     200 {object} response.GetInvestListResponse
// @Router      /invests [get]
func (h *Handler) GetInvestList(c *gin.Context) {
	var req request.GetInvestListRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		h.log.Error(err, "[handler.GetInvestList] - c.ShouldBindQuery() failed")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	resp, err := h.entities.GetInvestList(&req)
	if err != nil {
		h.log.Fields(logger.Fields{"request": req}).Error(err, "[handler.GetInvestList] - .entities.GetInvestList() failed")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse[any](resp.Data, nil, nil, nil))
}
