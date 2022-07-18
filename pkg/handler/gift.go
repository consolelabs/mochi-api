package handler

import (
	"net/http"

	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/request"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GiftXpHandler(c *gin.Context) {
	var req request.GiftXPRequest

	if err := c.BindJSON(&req); err != nil {
		h.log.Fields(logger.Fields{"body": req}).Error(err, "[handler.GiftXpHandler] - failed to read JSON")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := h.entities.GetUser(req.UserDiscordID)
	if err != nil {
		h.log.Fields(logger.Fields{"userID": req.UserDiscordID}).Error(err, "[handler.GiftXpHandler] - failed to get user")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.entities.SendGiftXp(req)
	if err != nil {
		h.log.Fields(logger.Fields{"body": req}).Error(err, "[handler.GiftXpHandler] - failed to send gift xp")
		c.JSON(http.StatusBadRequest, gin.H{"error": "cannot create activity log"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": resp})
}
