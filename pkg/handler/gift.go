package handler

import (
	"net/http"

	"github.com/defipod/mochi/pkg/request"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GiftXpHandler(c *gin.Context) {
	var req request.GiftXPRequest

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := h.entities.GetUser(req.UserDiscordID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.entities.SendGiftXp(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "cannot create activity log"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": resp})
}
