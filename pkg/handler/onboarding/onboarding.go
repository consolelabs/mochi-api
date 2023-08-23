package onboarding

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

// Start     godoc
// @Summary     User did start onboarding
// @Description User did start onboarding
// @Tags        Onboarding
// @Accept      json
// @Produce     json
// @Param       req body request.OnboardingStartRequest true "onboarding start request"
// @Success     200 {object} response.ResponseDataMessage
// @Router      /onboarding/start [post]
func (h *Handler) Start(c *gin.Context) {
	var req request.OnboardingStartRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.log.Error(err, "[handler.Start] BindJSON() failed")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	if err := h.entities.OnboardingStart(req); err != nil {
		h.log.Error(err, "[handler.Start] entity.OnboardingStart() failed")
		c.JSON(baseerrs.GetStatusCode(err), response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	c.JSON(http.StatusOK, response.CreateResponse[any](response.ResponseMessage{Message: "OK"}, nil, nil, nil))
}
