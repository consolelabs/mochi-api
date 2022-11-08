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

// UpdateUserFeedback     godoc
// @Summary     Update users' feedbacks
// @Description Update users' feedbacks
// @Tags        Feedback
// @Accept      json
// @Produce     json
// @Param       req body request.UpdateUserFeedbackRequest true "request"
// @Success     200 {object} response.UpdateUserFeedbackResponse
// @Router      /feedback [put]
func (h *Handler) UpdateUserFeedback(c *gin.Context) {
	var req request.UpdateUserFeedbackRequest
	if err := c.BindJSON(&req); err != nil {
		h.log.Fields(logger.Fields{"body": req}).Error(err, "[handler.UpdateUserFeedback] - failed to read JSON")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	if req.Status != "none" && req.Status != "confirmed" && req.Status != "completed" {
		h.log.Fields(logger.Fields{"body": req}).Info("[handler.UpdateUserFeedback] - invalid feedback status")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, nil, nil))
		return
	}

	data, err := h.entities.UpdateUserFeedback(&req)
	if err != nil {
		h.log.Fields(logger.Fields{"body": req}).Error(err, "[handler.UpdateUserFeedback] - failed to update feedback")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(data, nil, nil, nil))
}

// GetAllUserFeedback     godoc
// @Summary     Get users' feedbacks
// @Description Get users' feedbacks
// @Tags        Feedback
// @Accept      json
// @Produce     json
// @Param       filter query string true "filter by"
// @Param       value query string true "filtered value"
// @Success     200 {object} response.UserFeedbackResponse
// @Router      /feedback [get]
func (h *Handler) GetAllUserFeedback(c *gin.Context) {
	filter := c.Query("filter")
	if filter != "command" && filter != "status" && filter != "discord_id" {
		filter = "status"
	}
	value := c.Query("value")
	data, err := h.entities.GetAllUserFeedback(filter, value)
	if err != nil {
		h.log.Fields(logger.Fields{"filter": filter, "value": value}).Error(err, "[handler.GetAllUserFeedback] - failed to get feedback")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(data, nil, nil, nil))
}
