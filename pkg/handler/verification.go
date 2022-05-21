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

	res, err := h.entities.NewGuildConfigWalletVerificationMessage(req.GuildConfigWalletVerificationMessage)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status": "ok",
		"data":   res,
	})
}

func (h *Handler) UpdateGuildConfigWalletVerificationMessage(c *gin.Context) {
	var req request.NewGuildConfigWalletVerificationMessageRequest

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := req.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.entities.UpdateGuildConfigWalletVerificationMessage(req.GuildConfigWalletVerificationMessage)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"data":   res,
	})
}

func (h *Handler) DeleteGuildConfigWalletVerificationMessage(c *gin.Context) {

	var guildID = c.Query("guild_id")

	if guildID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "guild_id is required"})
		return
	}

	err := h.entities.DeleteGuildConfigWalletVerificationMessage(guildID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
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
		if err.Error() == "already have a verified wallet" {
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
