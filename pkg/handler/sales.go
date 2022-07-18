package handler

import (
	"net/http"

	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/request"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetNftSalesHandler(c *gin.Context) {
	addr := c.Query("collection-address")
	platform := c.Query("platform")
	data, err := h.entities.GetNftSales(addr, platform)
	if err != nil || data == nil {
		h.log.Fields(logger.Fields{"address": addr, "platform": platform}).Error(err, "[handler.GetNftSalesHandler] - failed to get NFT sales")
		c.JSON(http.StatusOK, gin.H{"error": "collection not found"})
		return
	}
	c.JSON(http.StatusOK, data)
}

func (h *Handler) WebhookNftSaleHandler(c *gin.Context) {
	var req request.NftSalesRequest
	if err := c.Bind(&req); err != nil {
		h.log.Fields(logger.Fields{"body": req}).Error(err, "[handler.WebhookNftSaleHandler] - failed to read JSON")
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	err := h.entities.SendNftSalesToChannel(req)
	if err != nil {
		h.log.Fields(logger.Fields{"body": req}).Error(err, "[handler.WebhookNftSaleHandler] - failed to send NFT sales to channel")
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}
