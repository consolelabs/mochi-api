package handler

import (
	"net/http"

	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/request"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (h *Handler) NewGuildConfigWalletVerificationMessage(c *gin.Context) {

	var req request.NewGuildConfigWalletVerificationMessageRequest

	if err := c.BindJSON(&req); err != nil {
		h.log.Fields(logger.Fields{"body": req}).Error(err, "[handler.NewGuildConfigWalletVerificationMessage] - failed to read JSON")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := req.Validate(); err != nil {
		h.log.Fields(logger.Fields{"guildID": req.GuildID, "channelID": req.VerifyChannelID}).Error(err, "[handler.NewGuildConfigWalletVerificationMessage] - failed to validate request")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.entities.NewGuildConfigWalletVerificationMessage(req.GuildConfigWalletVerificationMessage)
	if err != nil {
		h.log.Fields(logger.Fields{"message": req.GuildConfigWalletVerificationMessage}).Error(err, "[handler.NewGuildConfigWalletVerificationMessage] - failed to create guild config wallet verification message")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status": "ok",
		"data":   res,
	})
}

func (h *Handler) GetGuildConfigWalletVerificationMessage(c *gin.Context) {
	guildId := c.Param("guild_id")
	if guildId == "" {
		h.log.Info("[handler.GetGuildConfigWalletVerificationMessage] - guild id empty")
		c.JSON(http.StatusBadRequest, gin.H{"error": "guild_id is required"})
		return
	}

	res, err := h.entities.GetGuildConfigWalletVerificationMessage(guildId)
	if err != nil && err != gorm.ErrRecordNotFound {
		h.log.Fields(logger.Fields{"guildID": guildId}).Error(err, "[handler.GetGuildConfigWalletVerificationMessage] - failed to get config")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"data":   res,
	})
}
func (h *Handler) UpdateGuildConfigWalletVerificationMessage(c *gin.Context) {
	var req request.NewGuildConfigWalletVerificationMessageRequest

	if err := c.BindJSON(&req); err != nil {
		h.log.Fields(logger.Fields{"body": req}).Error(err, "[handler.UpdateGuildConfigWalletVerificationMessage] - failed to read JSON")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := req.Validate(); err != nil {
		h.log.Fields(logger.Fields{"guildID": req.GuildID, "channelID": req.VerifyChannelID}).Error(err, "[handler.UpdateGuildConfigWalletVerificationMessage] - failed to validate request")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.entities.UpdateGuildConfigWalletVerificationMessage(req.GuildConfigWalletVerificationMessage)
	if err != nil {
		h.log.Fields(logger.Fields{"message": req.GuildConfigWalletVerificationMessage}).Error(err, "[handler.UpdateGuildConfigWalletVerificationMessage] - failed to update guild config wallet verification message")
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
		h.log.Info("[handler.DeleteGuildConfigWalletVerificationMessage] - guild id empty")
		c.JSON(http.StatusBadRequest, gin.H{"error": "guild_id is required"})
		return
	}

	err := h.entities.DeleteGuildConfigWalletVerificationMessage(guildID)
	if err != nil {
		h.log.Fields(logger.Fields{"guildID": guildID}).Error(err, "[handler.DeleteGuildConfigWalletVerificationMessage] - failed to delete guild config wallet verification message")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func (h *Handler) GenerateVerification(c *gin.Context) {

	var req request.GenerateVerificationRequest

	if err := req.Bind(c); err != nil {
		h.log.Fields(logger.Fields{"body": req}).Error(err, "[handler.GenerateVerification] - failed to read JSON")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := req.Validate(); err != nil {
		h.log.Fields(logger.Fields{"userDiscordID": req.UserDiscordID, "guildID": req.GuildID}).Error(err, "[handler.GenerateVerification] - failed to validate request")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data, statusCode, err := h.entities.GenerateVerification(req)
	if err != nil {
		respData := gin.H{"error": err.Error()}
		if err.Error() == "already have a verified wallet" {
			respData["address"] = data
		}
		h.log.Fields(logger.Fields{"body": req}).Error(err, "[handler.GenerateVerification] - failed to generate verification")
		c.JSON(statusCode, respData)
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok", "code": data})
}

func (h *Handler) VerifyWalletAddress(c *gin.Context) {

	var req request.VerifyWalletAddressRequest

	if err := req.Bind(c); err != nil {
		h.log.Fields(logger.Fields{"body": req}).Error(err, "[handler.VerifyWalletAddress] - failed to read JSON")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := req.Validate(); err != nil {
		h.log.Fields(logger.Fields{"wallet": req.WalletAddress, "signature": req.Signature, "code": req.Code}).Error(err, "[handler.VerifyWalletAddress] - failed to validate request")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	statusCode, err := h.entities.VerifyWalletAddress(req)
	if err != nil {
		h.log.Fields(logger.Fields{"body": req}).Error(err, "[handler.VerifyWalletAddress] - failed to verify wallet address")
		c.JSON(statusCode, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
