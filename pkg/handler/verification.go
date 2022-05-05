package handler

import (
	"net/http"

	"github.com/defipod/mochi/pkg/request"
	"github.com/gin-gonic/gin"
)

func (h *Handler) NewGuildConfigWalletVerificationMessage(c *gin.Context) {

	var req request.NewGuildConfigWalletVerificationMessageRequest

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := req.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.entities.NewGuildConfigWalletVerificationMessage(req.GuildConfigWalletVerificationMessage); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": "ok"})
}

func (h *Handler) GenerateVerification(c *gin.Context) {

	var req request.GenerateVerificationRequest

	if err := req.Bind(c); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := req.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data, statusCode, err := h.entities.GenerateVerification(req)
	if err != nil {
		respData := gin.H{"error": err.Error()}
		if err.Error() == "already have a pod identity" {
			respData["address"] = data
		}
		c.JSON(statusCode, respData)
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok", "code": data})
}

func (h *Handler) VerifyWalletAddress(c *gin.Context) {

	var req request.VerifyWalletAddressRequest

	if err := req.Bind(c); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := req.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	statusCode, err := h.entities.VerifyWalletAddress(req)
	if err != nil {
		c.JSON(statusCode, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
