package handler

import (
	"net/http"

	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/response"
	"github.com/gin-gonic/gin"
)

// HandleUserFeedback     godoc
// @Summary     Post users' feedbacks
// @Description Post users' feedbacks
// @Tags        Feedback
// @Accept      json
// @Produce     json
// @Param       req body request.UserFeedbackRequest true "request"
// @Success     200 {object} response.ResponseMessage
// @Router      /feedback [post]
func (h *Handler) HandleUserFeedback(c *gin.Context) {
	var req request.UserFeedbackRequest
	if err := c.BindJSON(&req); err != nil {
		h.log.Fields(logger.Fields{"body": req}).Error(err, "[handler.HandleUserFeedback] - failed to read JSON")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	err := h.entities.HandleUserFeedback(&req)
	if err != nil {
		h.log.Fields(logger.Fields{"body": req}).Error(err, "[handler.HandleUserFeedback] - failed to handle feedback")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(response.ResponseMessage{Message: "OK"}, nil, nil, nil))
}
