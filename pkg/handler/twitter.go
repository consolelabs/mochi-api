package handler

import (
	"net/http"

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
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	err = h.entities.CreateTwitterPost(&req)
	if err != nil {
		h.log.Error(err, "[handler.CreateTwitterPost] - failed to create twitter post")
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response.ResponseMessage{Message: "OK"})
}

//Fields(logger.Fields{"address": addr, "platform": platform})
