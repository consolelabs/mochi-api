package configchannel

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

// GetGmConfig     godoc
// @Summary     Get GM config
// @Description Get GM config
// @Tags        Config
// @Accept      json
// @Produce     json
// @Param       guild_id   query  string true  "Guild ID"
// @Success     200 {object} response.GetGmConfigResponse
// @Router      /configs/gm [get]
func (h *Handler) GetGmConfig(c *gin.Context) {
	guildID := c.Query("guild_id")
	if guildID == "" {
		h.log.Info("[handler.GetGmConfig] - guild id empty")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("guild_id is required"), nil))
		return
	}

	config, err := h.entities.GetGmConfig(guildID)
	if err != nil {
		h.log.Fields(logger.Fields{"guildID": guildID}).Error(err, "[handler.GetGmConfig] - failed to get gm config")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(config, nil, nil, nil))
}

// UpsertGmConfig     godoc
// @Summary     Update or insert GM config
// @Description Update or insert GM config
// @Tags        Config
// @Accept      json
// @Produce     json
// @Param       Request  body request.UpsertGmConfigRequest true "Upsert GM Config request"
// @Success     200 {object} response.ResponseMessage
// @Router      /configs/gm [post]
func (h *Handler) UpsertGmConfig(c *gin.Context) {
	var req request.UpsertGmConfigRequest

	if err := c.BindJSON(&req); err != nil {
		h.log.Fields(logger.Fields{"guildID": req.GuildID, "channelID": req.ChannelID}).Error(err, "[handler.UpsertGmConfig] - failed to read JSON")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	if req.GuildID == "" {
		h.log.Info("[handler.UpsertGmConfig] - guild id empty")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("guild_id is required"), nil))
		return
	}
	if req.ChannelID == "" {
		h.log.Info("[handler.UpsertGmConfig] - channel id empty")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("channel_id is required"), nil))
		return
	}

	if err := h.entities.UpsertGmConfig(req); err != nil {
		h.log.Fields(logger.Fields{"guildID": req.GuildID, "channelID": req.ChannelID}).Error(err, "[handler.UpsertGmConfig] - failed to upsert gm config")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(response.ResponseMessage{Message: "OK"}, nil, nil, nil))
}

// GetWelcomeChannelConfig     godoc
// @Summary     Get welcome channel config
// @Description Get welcome channel config
// @Tags        Config
// @Accept      json
// @Produce     json
// @Param       guild_id   query  string true  "Guild ID"
// @Success     200 {object} response.GetWelcomeChannelConfigResponse
// @Router      /configs/welcome [get]
func (h *Handler) GetWelcomeChannelConfig(c *gin.Context) {
	guildID := c.Query("guild_id")
	if guildID == "" {
		h.log.Info("[handler.GetWelcomeChannelConfig] - guild id empty")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("guild_id is required"), nil))
		return
	}

	config, err := h.entities.GetWelcomeChannelConfig(guildID)

	if err != nil {
		h.log.Fields(logger.Fields{"guildID": guildID}).Error(err, "[handler.GetWelcomeChannelConfig] - failed to get welcome config")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(config, nil, nil, nil))
}

// UpsertWelcomeChannelConfig     godoc
// @Summary     Update or insert welcome channel config
// @Description Update or insert welcome channel config
// @Tags        Config
// @Accept      json
// @Produce     json
// @Param       Request  body request.UpsertWelcomeConfigRequest true "Upsert welcome channel config request"
// @Success     200 {object} response.GetWelcomeChannelConfigResponse
// @Router      /configs/welcome [post]
func (h *Handler) UpsertWelcomeChannelConfig(c *gin.Context) {
	var req request.UpsertWelcomeConfigRequest

	if err := c.BindJSON(&req); err != nil {
		h.log.Fields(logger.Fields{"guildID": req.GuildID, "channelID": req.ChannelID}).Error(err, "[handler.GetWelcomeChannelConfig] - failed to read JSON")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if req.GuildID == "" {
		h.log.Info("[handler.GetWelcomeChannelConfig] - guild id empty")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("guild_id is required"), nil))
		return
	}
	if req.ChannelID == "" {
		h.log.Info("[handler.GetWelcomeChannelConfig] - channel id empty")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("channel_id is required"), nil))
		return
	}

	config, err := h.entities.UpsertWelcomeChannelConfig(req)
	if err != nil {
		h.log.Fields(logger.Fields{"guildID": req.GuildID, "channelID": req.ChannelID}).Error(err, "[handler.GetWelcomeChannelConfig] - failed to upsert welcome config")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(config, nil, nil, nil))
}

// DeleteWelcomeChannelConfig     godoc
// @Summary     Delete welcome channel config
// @Description Delete welcome channel config
// @Tags        Config
// @Accept      json
// @Produce     json
// @Param       Request  body request.DeleteWelcomeConfigRequest true "Delete welcome channel config request"
// @Success     200 {object} response.ResponseMessage
// @Router      /configs/welcome [delete]
func (h *Handler) DeleteWelcomeChannelConfig(c *gin.Context) {
	var req request.DeleteWelcomeConfigRequest

	if err := c.BindJSON(&req); err != nil {
		h.log.Fields(logger.Fields{"guildID": req.GuildID}).Error(err, "[handler.GetWelcomeChannelConfig] - failed to read JSON")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if req.GuildID == "" {
		h.log.Info("[handler.GetWelcomeChannelConfig] - guild id empty")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("guild_id is required"), nil))
		return
	}

	if err := h.entities.DeleteWelcomeChannelConfig(req); err != nil {
		h.log.Fields(logger.Fields{"guildID": req.GuildID}).Error(err, "[handler.GetWelcomeChannelConfig] - failed to delete welcome config")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(response.ResponseMessage{Message: "OK"}, nil, nil, nil))
}

// GetVoteChannelConfig     godoc
// @Summary     Get vote channel config
// @Description Get vote channel config
// @Tags        Config
// @Accept      json
// @Produce     json
// @Param       guild_id   query  string true  "Guild ID"
// @Success     200 {object} response.GetVoteChannelConfigResponse
// @Router      /configs/upvote [get]
func (h *Handler) GetVoteChannelConfig(c *gin.Context) {
	guildID := c.Query("guild_id")
	if guildID == "" {
		h.log.Info("[handler.GetVoteChannelConfig] - guild id empty")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("guild_id is required"), nil))
		return
	}

	config, err := h.entities.GetVoteChannelConfig(guildID)

	if err != nil {
		h.log.Fields(logger.Fields{"guildID": guildID}).Error(err, "[handler.GetVoteChannelConfig] - failed to get vote channel config")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(config, nil, nil, nil))
}

// UpsertVoteChannelConfig     godoc
// @Summary     Update or insert vote channel config
// @Description Update or insert vote channel config
// @Tags        Config
// @Accept      json
// @Produce     json
// @Param       Request  body request.UpsertVoteChannelConfigRequest true "Upsert vote channel config request"
// @Success     200 {object} response.GetVoteChannelConfigResponse
// @Router      /configs/upvote [post]
func (h *Handler) UpsertVoteChannelConfig(c *gin.Context) {
	var req request.UpsertVoteChannelConfigRequest

	if err := c.BindJSON(&req); err != nil {
		h.log.Fields(logger.Fields{"guildID": req.GuildID, "channelID": req.ChannelID}).Error(err, "[handler.UpsertVoteChannelConfig] - failed to read JSON")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	if req.GuildID == "" {
		h.log.Info("[handler.UpsertVoteChannelConfig] - guild id empty")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("guild_id is required"), nil))
		return
	}
	if req.ChannelID == "" {
		h.log.Info("[handler.UpsertVoteChannelConfig] - channel id empty")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("channel_id is required"), nil))
		return
	}

	config, err := h.entities.UpsertVoteChannelConfig(req)
	if err != nil {
		h.log.Fields(logger.Fields{"guildID": req.GuildID, "channelID": req.ChannelID}).Error(err, "[handler.UpsertVoteChannelConfig] - failed to upsert vote channel config")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(config, nil, err, nil))
}

// DeleteVoteChannelConfig     godoc
// @Summary     Delete vote channel config
// @Description Delete vote channel config
// @Tags        Config
// @Accept      json
// @Produce     json
// @Param       Request  body request.DeleteVoteChannelConfigRequest true "Delete vote channel config request"
// @Success     200 {object} response.ResponseMessage
// @Router      /configs/upvote [delete]
func (h *Handler) DeleteVoteChannelConfig(c *gin.Context) {
	var req request.DeleteVoteChannelConfigRequest

	if err := c.BindJSON(&req); err != nil {
		h.log.Fields(logger.Fields{"guildID": req.GuildID}).Error(err, "[handler.DeleteVoteChannelConfig] - failed to read JSON")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	if req.GuildID == "" {
		h.log.Info("[handler.DeleteVoteChannelConfig] - guild id empty")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("guild_id is required"), nil))
		return
	}

	if err := h.entities.DeleteVoteChannelConfig(req); err != nil {
		h.log.Fields(logger.Fields{"guildID": req.GuildID}).Error(err, "[handler.DeleteVoteChannelConfig] - failed to delete vote channel config")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(response.ResponseMessage{Message: "OK"}, nil, nil, nil))
}

// GetUpvoteTiersConfig     godoc
// @Summary     Get all upvote tiers
// @Description Get all upvote tiers
// @Tags        Config
// @Accept      json
// @Produce     json
// @Success     200 {object} response.GetUpvoteTiersConfig
// @Router      /configs/upvote-tiers [get]
func (h *Handler) GetUpvoteTiersConfig(c *gin.Context) {
	tiers, err := h.entities.GetUpvoteTiersConfig()
	if err != nil {
		h.log.Error(err, "[handler.GetUpvoteTiersConfig] - failed to get upvote tiers")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(tiers, nil, nil, nil))
}

// GetSalesTrackerConfig     godoc
// @Summary     Get sales tracker config
// @Description Get sales tracker config
// @Tags        Config
// @Accept      json
// @Produce     json
// @Param       guild_id   query  string true  "Guild ID"
// @Success     200 {object} response.GetSalesTrackerConfigResponse
// @Router      /configs/sales-tracker [get]
func (h *Handler) GetSalesTrackerConfig(c *gin.Context) {
	guildID := c.Query("guild_id")
	if guildID == "" {
		h.log.Info("[handler.GetSalesTrackerConfig] - guild id empty")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("guild_id is required"), nil))
		return
	}

	config, err := h.entities.GetSalesTrackerConfig(guildID)
	if err != nil {
		h.log.Fields(logger.Fields{"guildID": guildID}).Error(err, "[handler.GetSalesTrackerConfig] - failed to get sales tracker config")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(config, nil, nil, nil))
}

// GetJoinLeaveChannelConfig     godoc
// @Summary     Get join-leave channel config
// @Description Get join-leave channel config
// @Tags        Config
// @Accept      json
// @Produce     json
// @Param       guild_id   query  string true  "Guild ID"
// @Success     200 {object} response.GetVoteChannelConfigResponse
// @Router      /configs/join-leave [get]
func (h *Handler) GetJoinLeaveChannelConfig(c *gin.Context) {
	guildID := c.Query("guild_id")
	if guildID == "" {
		h.log.Info("[handler.GetJoinLeaveChannelConfig] - guild id empty")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("guild_id is required"), nil))
		return
	}

	config, err := h.entities.GetJoinLeaveChannelConfig(guildID)

	if err != nil {
		h.log.Fields(logger.Fields{"guildID": guildID}).Error(err, "[handler.GetJoinLeaveChannelConfig] - failed to get join-leave channel config")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(config, nil, nil, nil))
}

// UpsertJoinLeaveChannelConfig     godoc
// @Summary     Update or insert join-leave channel config
// @Description Update or insert join-leave channel config
// @Tags        Config
// @Accept      json
// @Produce     json
// @Param       Request  body request.UpsertJoinLeaveChannelConfigRequest true "Upsert join-leave channel config request"
// @Success     200 {object} response.GetVoteChannelConfigResponse
// @Router      /configs/join-leave [post]
func (h *Handler) UpsertJoinLeaveChannelConfig(c *gin.Context) {
	var req request.UpsertJoinLeaveChannelConfigRequest

	if err := c.BindJSON(&req); err != nil {
		h.log.Fields(logger.Fields{"guildID": req.GuildID, "channelID": req.ChannelID}).Error(err, "[handler.UpsertJoinLeaveChannelConfig] - failed to read JSON")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	if req.GuildID == "" {
		h.log.Info("[handler.UpsertJoinLeaveChannelConfig] - guild id empty")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("guild_id is required"), nil))
		return
	}
	if req.ChannelID == "" {
		h.log.Info("[handler.UpsertJoinLeaveChannelConfig] - channel id empty")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("channel_id is required"), nil))
		return
	}

	config, err := h.entities.UpsertJoinLeaveChannelConfig(req)
	if err != nil {
		h.log.Fields(logger.Fields{"guildID": req.GuildID, "channelID": req.ChannelID}).Error(err, "[handler.UpsertJoinLeaveChannelConfig] - failed to upsert join-leave channel config")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(config, nil, err, nil))
}

// DeleteJoinLeaveChannelConfig     godoc
// @Summary     Delete join-leave channel config
// @Description Delete join-leave channel config
// @Tags        Config
// @Accept      json
// @Produce     json
// @Param       Request  body request.DeleteJoinLeaveChannelConfigRequest true "Delete join-leave channel config request"
// @Success     200 {object} response.ResponseMessage
// @Router      /configs/join-leave [delete]
func (h *Handler) DeleteJoinLeaveChannelConfig(c *gin.Context) {
	var req request.DeleteJoinLeaveChannelConfigRequest

	if err := c.BindJSON(&req); err != nil {
		h.log.Fields(logger.Fields{"guildID": req.GuildID}).Error(err, "[handler.DeleteJoinLeaveChannelConfigRequest] - failed to read JSON")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	if req.GuildID == "" {
		h.log.Info("[handler.DeleteJoinLeaveChannelConfigRequest] - guild id empty")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("guild_id is required"), nil))
		return
	}

	if err := h.entities.DeleteJoinLeaveChannelConfig(req); err != nil {
		h.log.Fields(logger.Fields{"guildID": req.GuildID}).Error(err, "[handler.DeleteJoinLeaveChannelConfigRequest] - failed to delete join-leave channel config")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(response.ResponseMessage{Message: "OK"}, nil, nil, nil))
}

// CreateConfigNotify   godoc
// @Summary     OffChain Tip Bot - Config notify
// @Description API config notify channel for token
// @Tags        OffChain
// @Accept      json
// @Produce     json
// @Param       Request body request.CreateTipConfigNotify true "config notify request"
// @Success     200 {object} response.ResponseMessage
// @Router      /offchain-tip-bot/config-notify [post]
func (h *Handler) CreateConfigNotify(c *gin.Context) {
	req := request.CreateTipConfigNotify{}

	if err := c.BindJSON(&req); err != nil {
		h.log.Fields(logger.Fields{"req": req}).Error(err, "[handler.CreateConfigNotify] - failed to read JSON")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	err := h.entities.CreateConfigNotify(req)
	if err != nil {
		h.log.Error(err, "[handler.CreateConfigNotify] - failed to create config notify")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(response.ResponseMessage{Message: "OK"}, nil, nil, nil))
}

// ListConfigNotify   godoc
// @Summary     OffChain Tip Bot - Config notify
// @Description API get list config notify channel for token
// @Tags        OffChain
// @Accept      json
// @Produce     json
// @Param       guild_id query string true "guild id"
// @Success     200 {object} response.ListConfigNotifyResponse
// @Router      /offchain-tip-bot/config-notify [get]
func (h *Handler) ListConfigNotify(c *gin.Context) {
	guildId := c.Query("guild_id")
	listConfigs, err := h.entities.ListConfigNotify(guildId)
	if err != nil {
		h.log.Error(err, "[handler.ListConfigNotify] - failed to list config notify")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(listConfigs, nil, nil, nil))
}

// DeleteConfigNotify   godoc
// @Summary     OffChain Tip Bot - Config notify
// @Description API delete config notify channel for token
// @Tags        OffChain
// @Accept      json
// @Produce     json
// @Param       id path string true "id of config notify"
// @Success     200 {object} response.ResponseMessage
// @Router      /offchain-tip-bot/config-notify/{id} [delete]
func (h *Handler) DeleteConfigNotify(c *gin.Context) {
	id := c.Param("id")

	err := h.entities.DeleteConfigNotify(id)
	if err != nil {
		h.log.Error(err, "[handler.DeleteConfigNotify] - failed to delete config notify")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(response.ResponseMessage{Message: "OK"}, nil, nil, nil))
}

// GetInviteTrackerConfig     godoc
// @Summary     Get invites tracker config
// @Description Get invites tracker config
// @Tags        Community
// @Accept      json
// @Produce     json
// @Param       guild_id query     string true "Guild ID"
// @Success     200 {object} response.GetInviteTrackerConfigResponse
// @Router      /community/invites/config [get]
func (h *Handler) GetInviteTrackerConfig(c *gin.Context) {
	guildID := c.Query("guild_id")
	if guildID == "" {
		h.log.Info("[handler.GetInviteTrackerConfig] - guild id empty")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("guild_id is required"), nil))
		return
	}

	config, err := h.entities.GetInviteTrackerLogChannel(guildID)
	if err != nil {
		h.log.Fields(logger.Fields{"guildID": guildID}).Error(err, "[handler.GetInviteTrackerConfig] - failed to get invite tracker log channel")

		code := http.StatusInternalServerError
		if err == gorm.ErrRecordNotFound {
			code = http.StatusNotFound
		}
		c.JSON(code, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(config, nil, nil, nil))
}

// ConfigureInvites     godoc
// @Summary     Configure invites
// @Description Configure invites
// @Tags        Community
// @Accept      json
// @Produce     json
// @Param       Request  body request.ConfigureInviteRequest true "Configure Invites request"
// @Success     200 {object} response.ConfigureInvitesResponse
// @Router      /community/invites/config [post]
func (h *Handler) ConfigureInvites(c *gin.Context) {
	var req request.ConfigureInviteRequest
	if err := req.Bind(c); err != nil {
		h.log.Fields(logger.Fields{"body": req}).Error(err, "[handler.ConfigureInvites] - failed to read JSON")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	if err := req.Validate(); err != nil {
		h.log.Fields(logger.Fields{"logChannel": req.LogChannel, "guildID": req.GuildID}).Error(err, "[handler.ConfigureInvites] - failed to validate request")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	if err := h.entities.CreateOrUpdateInviteTrackerLogChannel(req); err != nil {
		h.log.Fields(logger.Fields{"body": req}).Error(err, "[handler.ConfigureInvites] - failed to upsert invite tracker log channel")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(response.ResponseMessage{Message: "OK"}, nil, nil, nil))
}
