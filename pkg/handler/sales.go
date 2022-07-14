package handler

import (
	"net/http"

	"github.com/defipod/mochi/pkg/request"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetNftSalesHandler(c *gin.Context) {
	addr := c.Query("collection-address")
	platform := c.Query("platform")
	data, err := h.entities.GetNftSales(addr, platform)
	if err != nil || data == nil {
		c.JSON(http.StatusOK, gin.H{"error": "collection not found"})
		return
	}
	c.JSON(http.StatusOK, data)
}

func (h *Handler) WebhookNftSaleHandler(c *gin.Context) {
	var req request.NftSalesRequest
	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	for _, nftSale := range req.NftSales {
		err := h.entities.SendNftSalesToChannel(nftSale)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}
