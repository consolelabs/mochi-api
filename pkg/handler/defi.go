package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetHistoricalMarketChart(c *gin.Context) {
	data, err, statusCode := h.entities.GetHistoricalMarketChart(c)
	if err != nil {
		c.JSON(statusCode, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": data})
}
