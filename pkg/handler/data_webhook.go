package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/response"
)

func (h *Handler) NotifyNftCollectionIntegration(c *gin.Context) {
	req := request.NotifyCompleteNftIntegrationRequest{}
	if err := c.BindJSON(&req); err != nil {
		h.log.Error(err, "[handler.NotifyNftCollectionIntegration] c.BindJSON() failed")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	err := h.entities.NotifyNftCollectionIntegration(req)
	if err != nil {
		h.log.Error(err, "[handler.NotifyNftCollectionIntegration] entity.NotifyNftCollectionIntegration() failed")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	c.JSON(http.StatusOK, response.ResponseMessage{Message: "ok"})
}

func (h *Handler) NotifyNftCollectionSync(c *gin.Context) {
	req := request.NotifyCompleteNftSyncRequest{}
	if err := c.BindJSON(&req); err != nil {
		h.log.Error(err, "[handler.NotifyNftCollectionSync] c.BindJSON() failed")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	err := h.entities.NotifyNftCollectionSync(req)
	if err != nil {
		h.log.Error(err, "[handler.NotifyNftCollectionSync] entity.NotifyNftCollectionSync() failed")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	c.JSON(http.StatusOK, response.ResponseMessage{Message: "ok"})
}

func (h *Handler) NotifySaleMarketplace(c *gin.Context) {
	var req request.NotifySaleMarketplaceRequest
	if err := c.Bind(&req); err != nil {
		h.log.Fields(logger.Fields{"body": req}).Error(err, "[handler.WebhookNftSaleHandler] - failed to read JSON")
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	err := h.entities.NotifySaleMarketplace(req)
	if err != nil {
		h.log.Fields(logger.Fields{"body": req}).Error(err, "[handler.handleNftSales] - failed to send NFT sales to channel")
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}
