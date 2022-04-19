package handler

import (
	"github.com/gin-gonic/gin"
)

func (h *Handler) HandleDiscordWebhook(c *gin.Context) {
	// parse request
	// var req request.DiscordWebhookRequest
	// if err := req.Bind(c); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }

	// handle logic
	// switch usecase here
	switch {
	// case isUserJoin():
	// 	h.handleUserJoin(c)
	}

	// response
	// res, err := h.entities.HandleDiscordWebhook(req)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	return
	// }

	// c.JSON(http.StatusOK, res)
}

func (h *Handler) handleUserJoin(c *gin.Context) {
	// handle invite-tracker
	// h.handleInviteTracker(c)

	// handle ...
	// h.handle...(c)
}

// func (h *Handler) handleInviteTracker(c *gin.Context) {

// }
