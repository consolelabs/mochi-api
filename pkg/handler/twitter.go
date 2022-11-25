package handler

import (
	"net/http"

	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/response"
	"github.com/gin-gonic/gin"
)

// CreateTwitterPost     godoc
// @Summary     Create twitter post
// @Description Create twitter post
// @Tags        Twitter
// @Accept      json
// @Produce     json
// @Param       Request  body request.TwitterPost true "Create twitter post request"
// @Success     200 {object} response.ResponseMessage
// @Router      /twitter [post]
func (h *Handler) CreateTwitterPost(c *gin.Context) {
	req := request.TwitterPost{}
	err := c.BindJSON(&req)
	if err != nil {
		h.log.Error(err, "[handler.CreateTwitterPost] - failed to read JSON body")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = h.entities.CreateTwitterPost(&req)
	if err != nil {
		h.log.Error(err, "[handler.CreateTwitterPost] - failed to create twitter post")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response.ResponseMessage{Message: "OK"})
}

// CreateTwitterPost     godoc
// @Summary     Create twitter post
// @Description Create twitter post
// @Tags        Twitter
// @Accept      json
// @Produce     json
// @Param       req  query request.GetTwitterLeaderboardRequest true "Create twitter post request"
// @Success     200 {object} response.GetTwitterLeaderboardResponse
// @Router      /twitter/top [get]
func (h *Handler) GetTwitterLeaderboard(c *gin.Context) {
	req := request.GetTwitterLeaderboardRequest{}
	if err := c.ShouldBindQuery(&req); err != nil {
		h.log.Error(err, "[handler.GetTwitterLeaderboard] BindJSON() failed")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	list, total, err := h.entities.GetTwitterLeaderboard(req)
	if err != nil {
		h.log.Error(err, "[handler.GetTwitterLeaderboard] entity.GetTwitterLeaderboard() failed")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	paging := &response.PaginationResponse{
		Pagination: model.Pagination{Page: req.Page, Size: req.Size},
		Total:      total,
	}
	c.JSON(http.StatusOK, response.CreateResponse(list, paging, nil, nil))
}
