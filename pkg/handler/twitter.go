package handler

import (
	"net/http"

	"github.com/defipod/mochi/pkg/request"
	"github.com/gin-gonic/gin"
)

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
	c.JSON(http.StatusOK, gin.H{"message": "OK"})
}

//Fields(logger.Fields{"address": addr, "platform": platform})
