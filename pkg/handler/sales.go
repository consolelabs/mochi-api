package handler

import (
	"errors"
	"net/http"

	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/response"
	_ "github.com/defipod/mochi/pkg/response"
	"github.com/gin-gonic/gin"
)

// GetNftsalesHandler     godoc
// @Summary     Get NFT sales
// @Description Get NFT sales
// @Tags        NFT
// @Accept      json
// @Produce     json
// @Param       collection-address   query  string true  "Collection address"
// @Param       platform   query  string true  "Platform"
// @Success     200 {object} response.NftSalesResponse
// @Router      /nfts/sales [get]
func (h *Handler) GetNftSalesHandler(c *gin.Context) {
	addr := c.Query("collection-address")
	platform := c.Query("platform")
	data, err := h.entities.GetNftSales(addr, platform)
	if err != nil || data == nil {
		h.log.Fields(logger.Fields{"address": addr, "platform": platform}).Error(err, "[handler.GetNftSalesHandler] - failed to get NFT sales")
		c.JSON(http.StatusOK, response.CreateResponse[any](nil, nil, errors.New("collection not found"), nil))
		return
	}
	c.JSON(http.StatusOK, data)
}

func (h *Handler) WebhookNftHandler(c *gin.Context) {
	var req request.NotifySaleMarketplaceRequest
	if err := c.Bind(&req); err != nil {
		h.log.Fields(logger.Fields{"body": req}).Error(err, "[handler.WebhookNftSaleHandler] - failed to read JSON")
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	switch req.Event {
	case "sales":
		h.handleNftSales(c, req)
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
func (h *Handler) handleNftSales(c *gin.Context, req request.NotifySaleMarketplaceRequest) {
	err := h.entities.NotifySaleMarketplace(req)
	if err != nil {
		h.log.Fields(logger.Fields{"body": req}).Error(err, "[handler.handleNftSales] - failed to send NFT sales to channel")
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}

// {
// 	"event": "notify_done_sync",
// 	"collection_address": "0x7D1070fdbF0eF8752a9627a79b00221b53F231fA",
// 	"chain": "ftm"
// }
func (h *Handler) handleNofityDoneSync(c *gin.Context, req request.HandleNftWebhookRequest) {
	err := h.entities.SendNftAddedCollection(req)
	if err != nil {
		h.log.Fields(logger.Fields{"body": req}).Error(err, "[handler.handleNofityDoneSync] - failed to send new added collection to channel")
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}
