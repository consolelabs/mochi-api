package handler

import (
	"fmt"
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

func (h *Handler) WebhookNftHandler(c *gin.Context) {
	var req request.HandleNftWebhookRequest
	if err := c.Bind(&req); err != nil {
		h.log.Fields(logger.Fields{"body": req}).Error(err, "[handler.WebhookNftSaleHandler] - failed to read JSON")
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(req)

	switch req.Event {
	case "sales":
		h.handleNftSales(c, req)
	case "nft_added_collection":
		h.handleNftAddedCollection(c, req)
	}
}

// {
// 			"event": "nft_added_collection",
//             "collection_address": "0x7D1070fdbF0eF8752a9627a79b00221b53F231fA",
//             "token_id": "1",
//             "transaction": "0xf2db72ba2348e2f718a6118c12fead737333a7a5de5c0fa35682c9fc7ef6934f",
//             "from": "0x140dd183e18ba39bd9BE82286ea2d96fdC48117A",
//             "to": "0x140dd183e18ba39bd9BE82286ea2d96fdC48117A",
//             "marketplace": "opensea",
//             "price": {
//                 "token": {
//                     "symbol": "ftm",
//                     "is_native": true,
//                     "address": "0x00000000000000000000000000",
//                     "decimal": 18
//                 },
//                 "amount": "0234566000000000000"
//             },
//             "last_price": {
//                 "token": {
//                     "symbol": "ftm",
//                     "is_native": true,
//                     "address": "0x00000000000000000000000000",
//                     "decimal": 18
//                 },
//                 "amount": "120000000000000000000"
//             },
//             "hold": 440000000
// }
func (h *Handler) handleNftSales(c *gin.Context, req request.HandleNftWebhookRequest) {
	err := h.entities.SendNftSalesToChannel(req)
	if err != nil {
		h.log.Fields(logger.Fields{"body": req}).Error(err, "[handler.handleNftSales] - failed to send NFT sales to channel")
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}

// {
// 	"event": "nft_added_collection",
// 	"collection_address": "0x7D1070fdbF0eF8752a9627a79b00221b53F231fA",
// 	"chain": "ftm"
// }
func (h *Handler) handleNftAddedCollection(c *gin.Context, req request.HandleNftWebhookRequest) {
	err := h.entities.SendNftAddedCollection(req)
	if err != nil {
		h.log.Fields(logger.Fields{"body": req}).Error(err, "[handler.handleNftAddedCollection] - failed to send new added collection to channel")
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}
