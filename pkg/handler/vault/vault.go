package vault

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/defipod/mochi/pkg/entities"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/response"
)

type Handler struct {
	entities *entities.Entity
	log      logger.Logger
}

func New(entities *entities.Entity, logger logger.Logger) IHandler {
	return &Handler{
		entities: entities,
		log:      logger,
	}
}

func (h *Handler) CreateVault(c *gin.Context) {
	var req request.CreateVaultRequest

	if err := c.BindJSON(&req); err != nil {
		h.log.Fields(logger.Fields{"guildID": req.GuildId, "name": req.Name, "threshold": req.Threshold}).Error(err, "[handler.CreateVault] - failed to read JSON")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	vault, err := h.entities.CreateVault(&req)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value") {
			h.log.Fields(logger.Fields{"guildID": req.GuildId, "name": req.Name, "threshold": req.Threshold}).Error(err, "[handler.CreateVault] - failed to create vault")
			c.JSON(http.StatusBadRequest, gin.H{"message": "Vault name is already exist for this server"})
			return
		}
		h.log.Fields(logger.Fields{"guildID": req.GuildId, "name": req.Name, "threshold": req.Threshold}).Error(err, "[handler.CreateVault] - failed to create vault")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse[any](vault, nil, nil, nil))
}
func (h *Handler) GetVault(c *gin.Context) {
	guildId := c.Query("guild_id")
	if guildId == "" {
		h.log.Info("[handler.GetVault] - guildId is empty")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, nil, nil))
		return
	}

	vault, err := h.entities.GetVault(guildId)
	if err != nil {
		h.log.Fields(logger.Fields{"guildID": guildId}).Error(err, "[handler.GetVault] - failed to get vault")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse[any](vault, nil, nil, nil))
}

func (h *Handler) GetVaultInfo(c *gin.Context) {
	vaultInfo, err := h.entities.GetVaultInfo()
	if err != nil {
		h.log.Error(err, "[handler.GetVaultInfo] - failed to get vault info")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse[any](vaultInfo, nil, nil, nil))
}

func (h *Handler) GetVaultConfigChannel(c *gin.Context) {
	guildId := c.Query("guild_id")
	if guildId == "" {
		h.log.Info("[handler.GetVault] - guildId is empty")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, nil, nil))
		return
	}

	vaultInfo, err := h.entities.GetVaultConfigChannel(guildId)
	if err != nil {
		h.log.Error(err, "[handler.GetVaultConfigChannel] - failed to get vault config channel")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse[any](vaultInfo, nil, nil, nil))
}

func (h *Handler) CreateConfigChannel(c *gin.Context) {
	var req request.CreateVaultConfigChannelRequest

	if err := c.BindJSON(&req); err != nil {
		h.log.Fields(logger.Fields{"guildID": req.GuildId, "channelID": req.ChannelId}).Error(err, "[handler.CreateConfigChannel] - failed to read JSON")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	err := h.entities.CreateVaultConfigChannel(&req)
	if err != nil {
		h.log.Fields(logger.Fields{"guildID": req.GuildId, "channelID": req.ChannelId}).Error(err, "[handler.CreateVaultConfigChannel] - failed to create vault config channel")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse[any](response.ResponseMessage{Message: "OK"}, nil, nil, nil))
}

func (h *Handler) CreateConfigThreshold(c *gin.Context) {
	var req request.CreateConfigThresholdRequest

	if err := c.BindJSON(&req); err != nil {
		h.log.Fields(logger.Fields{"guildID": req.GuildId, "name": req.Name, "threshold": req.Threshold}).Error(err, "[handler.CreateConfigThreshold] - failed to read JSON")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	vaultConfigChannel, err := h.entities.CreateConfigThreshold(&req)
	if err != nil {
		h.log.Fields(logger.Fields{"guildID": req.GuildId, "name": req.Name, "threshold": req.Threshold}).Error(err, "[handler.CreateConfigThreshold] - failed to create vault config channel")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse[any](vaultConfigChannel, nil, nil, nil))
}

func (h *Handler) CreateAddTreasurerRequest(c *gin.Context) {
	var req request.CreateAddTreasurerRequest
	if err := c.BindJSON(&req); err != nil {
		h.log.Fields(logger.Fields{"guildID": req.GuildId, "userDiscordId": req.UserDiscordId, "vaultName": req.VaultName, "message": req.Message}).Error(err, "[handler.CreateAddTreasurerRequest] - failed to read JSON")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	treasurerReq, err := h.entities.CreateAddTreasurerRequest(&req)
	if err != nil {
		if strings.Contains(err.Error(), "vault not exist") {
			h.log.Fields(logger.Fields{"guildID": req.GuildId, "userDiscordId": req.UserDiscordId, "vaultName": req.VaultName, "message": req.Message}).Error(err, "[handler.CreateAddTreasurerRequest] - user not found")
			c.JSON(http.StatusBadRequest, gin.H{"error": "This vault is not exist yet"})
			return
		}

		h.log.Fields(logger.Fields{"guildID": req.GuildId, "userDiscordId": req.UserDiscordId, "vaultName": req.VaultName, "message": req.Message}).Error(err, "[handler.CreateAddTreasurerRequest] - failed to create add treasurer req")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse[any](treasurerReq, nil, nil, nil))
}

func (h *Handler) AddTreasurerToVault(c *gin.Context) {
	var req request.AddTreasurerToVaultRequest

	if err := c.BindJSON(&req); err != nil {
		h.log.Fields(logger.Fields{"requestId": req.RequestId}).Error(err, "[handler.AddTreasurerToVault] - failed to read JSON")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	treasurer, err := h.entities.AddTreasurerToVault(&req)
	if err != nil {
		if strings.Contains(err.Error(), "request not exist") {
			h.log.Fields(logger.Fields{"requestId": req.RequestId}).Error(err, "[handler.AddTreasurerToVault] - request not found")
			c.JSON(http.StatusBadRequest, gin.H{"error": "This request is not exist"})
			return
		}
		if strings.Contains(err.Error(), "duplicate key value") {
			h.log.Fields(logger.Fields{"requestId": req.RequestId}).Error(err, "[handler.AddTreasurerToVault] - user already in vault")
			c.JSON(http.StatusBadRequest, gin.H{"error": "This user is already added to this vault"})
			return
		}

		if strings.Contains(err.Error(), "vault not exist") {
			h.log.Fields(logger.Fields{"requestId": req.RequestId}).Error(err, "[handler.AddTreasurerToVault] - user not found")
			c.JSON(http.StatusBadRequest, gin.H{"error": "This vault is not exist yet"})
			return
		}

		h.log.Fields(logger.Fields{"requestId": req.RequestId}).Error(err, "[handler.AddTreasurerToVault] - failed to add treasurer to vault")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse[any](treasurer, nil, nil, nil))
}
