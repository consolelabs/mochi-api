package handler

import (
	"net/http"
	"strconv"

	"github.com/defipod/mochi/pkg/request"
	"github.com/gin-gonic/gin"
)

func(h *Handler) GiftXpHandler(c *gin.Context) {
	var req request.GiftXpRequest

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	_, err := h.entities.GetUser(req.UserDiscordId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	earnedXp, _ := strconv.Atoi(req.XpAmount)
	resp, err := h.entities.SendGiftXp(req.GuildId, req.UserDiscordId, earnedXp, "gifted")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "cannot create activity log"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": resp})
}