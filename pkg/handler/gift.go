package handler

import (
	"net/http"

	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/response"
	"github.com/gin-gonic/gin"
)

// GiftXpHandler     godoc
// @Summary     Gift xp handler
// @Description Gift xp handler
// @Tags        Gift
// @Accept      json
// @Produce     json
// @Param       Request  body request.GiftXPRequest true "Gift XP handler request"
// @Success     200 {object} response.GiftXpHandlerResponse
// @Router      /gift/xp [post]
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
	c.JSON(http.StatusOK, response.GiftXpHandlerResponse{Data: resp})
}
