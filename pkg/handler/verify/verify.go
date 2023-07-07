package verify

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

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

// NewGuildConfigWalletVerificationMessage     godoc
// @Summary     Config wallet verification message
// @Description Config wallet verification message
// @Tags        Verification
// @Accept      json
// @Produce     json
// @Param       Request  body request.NewGuildConfigWalletVerificationMessageRequest true "New guild config wallet verification message request"
// @Success     200 {object} response.NewGuildConfigWalletVerificationMessageResponse
// @Router      /verify/config [post]
func (h *Handler) NewGuildConfigWalletVerificationMessage(c *gin.Context) {
	var req request.NewGuildConfigWalletVerificationMessageRequest
	if err := c.BindJSON(&req); err != nil {
		h.log.Fields(logger.Fields{"body": req}).Error(err, "[handler.NewGuildConfigWalletVerificationMessage] - failed to read JSON")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	if err := req.Validate(); err != nil {
		h.log.Fields(logger.Fields{"guildID": req.GuildID, "channelID": req.VerifyChannelID}).Error(err, "[handler.NewGuildConfigWalletVerificationMessage] - failed to validate request")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	res, err := h.entities.NewGuildConfigWalletVerificationMessage(req.GuildConfigWalletVerificationMessage)
	if err != nil {
		h.log.Fields(logger.Fields{"message": req.GuildConfigWalletVerificationMessage}).Error(err, "[handler.NewGuildConfigWalletVerificationMessage] - failed to create guild config wallet verification message")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusCreated, response.CreateResponse(res, nil, err, nil))
}

// GetGuildConfigWalletVerificationMessage     godoc
// @Summary     Get guild config wallet verification message
// @Description Get guild config wallet verification message
// @Tags        Verification
// @Accept      json
// @Produce     json
// @Param       guild_id   path  string true  "Guild ID"
// @Success     200 {object} response.NewGuildConfigWalletVerificationMessageResponse
// @Router      /verify/config/{guild_id} [get]
func (h *Handler) GetGuildConfigWalletVerificationMessage(c *gin.Context) {
	guildId := c.Param("guild_id")
	if guildId == "" {
		h.log.Info("[handler.GetGuildConfigWalletVerificationMessage] - guild id empty")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("guild_id is required"), nil))
		return
	}

	res, err := h.entities.GetGuildConfigWalletVerificationMessage(guildId)
	if err != nil && err != gorm.ErrRecordNotFound {
		h.log.Fields(logger.Fields{"guildID": guildId}).Error(err, "[handler.GetGuildConfigWalletVerificationMessage] - failed to get config")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusCreated, response.CreateResponse(res, nil, nil, nil))
}

// UpdateGuildConfigWalletVerificationMessage     godoc
// @Summary     Update guild config wallet verification message
// @Description Update guild config wallet verification message
// @Tags        Verification
// @Accept      json
// @Produce     json
// @Param       Request  body request.NewGuildConfigWalletVerificationMessageRequest true "Update guild config wallet verification message request"
// @Success     200 {object} response.NewGuildConfigWalletVerificationMessageResponse
// @Router      /verify/config [put]
func (h *Handler) UpdateGuildConfigWalletVerificationMessage(c *gin.Context) {
	var req request.NewGuildConfigWalletVerificationMessageRequest

	if err := c.BindJSON(&req); err != nil {
		h.log.Fields(logger.Fields{"body": req}).Error(err, "[handler.UpdateGuildConfigWalletVerificationMessage] - failed to read JSON")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	if err := req.Validate(); err != nil {
		h.log.Fields(logger.Fields{"guildID": req.GuildID, "channelID": req.VerifyChannelID}).Error(err, "[handler.UpdateGuildConfigWalletVerificationMessage] - failed to validate request")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	res, err := h.entities.UpdateGuildConfigWalletVerificationMessage(req.GuildConfigWalletVerificationMessage)
	if err != nil {
		h.log.Fields(logger.Fields{"message": req.GuildConfigWalletVerificationMessage}).Error(err, "[handler.UpdateGuildConfigWalletVerificationMessage] - failed to update guild config wallet verification message")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(res, nil, nil, nil))
}

// DeleteGuildConfigWalletVerificationMessage     godoc
// @Summary     Delete guild config wallet verification message
// @Description Delete guild config wallet verification message
// @Tags        Verification
// @Accept      json
// @Produce     json
// @Param       guild_id   query  string true  "Guild ID"
// @Success     200 {object} response.ResponseStatus
// @Router      /verify/config [delete]
func (h *Handler) DeleteGuildConfigWalletVerificationMessage(c *gin.Context) {
	var guildID = c.Query("guild_id")
	if guildID == "" {
		h.log.Info("[handler.DeleteGuildConfigWalletVerificationMessage] - guild id empty")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("guild_id is required"), nil))
		return
	}

	err := h.entities.DeleteGuildConfigWalletVerificationMessage(guildID)
	if err != nil {
		h.log.Fields(logger.Fields{"guildID": guildID}).Error(err, "[handler.DeleteGuildConfigWalletVerificationMessage] - failed to delete guild config wallet verification message")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(response.ResponseStatus{Status: "OK"}, nil, err, nil))
}

// AssignRole     godoc
// @Summary     Assign verified role if user has verified wallet address
// @Description Assign verified role if user has verified wallet address
// @Tags        Verification
// @Accept      json
// @Produce     json
// @Param       Request  body request.AssignVerifiedRoleRequest true "Assign verified role request"
// @Success     200 {object} response.ResponseStatus
// @Router      /verify/assign-role [post]
func (h *Handler) AssignVerifiedRole(c *gin.Context) {
	var req request.AssignVerifiedRoleRequest
	if err := c.BindJSON(&req); err != nil {
		h.log.Fields(logger.Fields{"body": req}).Error(err, "[handler.AssignVerifiedRole] - failed to read JSON")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	err := h.entities.AssignVerifiedRole(req.UserDiscordID, req.GuildID)
	if err != nil {
		h.log.Fields(logger.Fields{"req": req}).Error(err, "[handler.AssignVerifiedRole] - failed to assign verified role")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusCreated, response.CreateResponse(response.ResponseStatus{Status: "ok"}, nil, nil, nil))
}
