package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetNftSalesHandler(c *gin.Context) {
	addr := c.Query("collection_address")
	platform := c.Query("platform")
	data, err := h.entities.GetNftSales(addr, platform)
	if err != nil || data == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "collection not found"})
		return
	}
	c.JSON(http.StatusOK, data)
}
