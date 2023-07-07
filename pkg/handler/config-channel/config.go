package configchannel

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/defipod/mochi/pkg/entities"
	"github.com/defipod/mochi/pkg/logger"
	errs "github.com/defipod/mochi/pkg/model/errors"
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
// @Tags        ConfigChannel
// @Accept      json
// @Produce     json
// @Param       guild_id   query  string true  "Guild ID"
// @Success     200 {object} response.GetGmConfigResponse
// @Router      /config-channels/gm [get]
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
// @Tags        ConfigChannel
// @Accept      json
// @Produce     json
// @Param       Request  body request.UpsertGmConfigRequest true "Upsert GM Config request"
// @Success     200 {object} response.ResponseMessage
// @Router      /config-channels/gm [post]
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
// @Tags        ConfigChannel
// @Accept      json
// @Produce     json
// @Param       guild_id   query  string true  "Guild ID"
// @Success     200 {object} response.GetWelcomeChannelConfigResponse
// @Router      /config-channels/welcome [get]
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
// @Tags        ConfigChannel
// @Accept      json
// @Produce     json
// @Param       Request  body request.UpsertWelcomeConfigRequest true "Upsert welcome channel config request"
// @Success     200 {object} response.GetWelcomeChannelConfigResponse
// @Router      /config-channels/welcome [post]
func (h *Handler) UpsertWelcomeChannelConfig(c *gin.Context) {
	var req request.UpsertWelcomeConfigRequest

	if err := c.BindJSON(&req); err != nil {
		h.log.Fields(logger.Fields{"guildID": req.GuildID, "channelID": req.ChannelID}).Error(err, "[handler.UpsertWelcomeChannelConfig] - failed to read JSON")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if req.GuildID == "" {
		h.log.Info("[handler.UpsertWelcomeChannelConfig] - guild id empty")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("guild_id is required"), nil))
		return
	}
	if req.ChannelID == "" {
		h.log.Info("[handler.UpsertWelcomeChannelConfig] - channel id empty")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("channel_id is required"), nil))
		return
	}

	config, err := h.entities.UpsertWelcomeChannelConfig(req)
	if err != nil {
		h.log.Fields(logger.Fields{"guildID": req.GuildID, "channelID": req.ChannelID}).Error(err, "[handler.UpsertWelcomeChannelConfig] - failed to upsert welcome config")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(config, nil, nil, nil))
}

// DeleteWelcomeChannelConfig     godoc
// @Summary     Delete welcome channel config
// @Description Delete welcome channel config
// @Tags        ConfigChannel
// @Accept      json
// @Produce     json
// @Param       Request  body request.DeleteWelcomeConfigRequest true "Delete welcome channel config request"
// @Success     200 {object} response.ResponseMessage
// @Router      /config-channels/welcome [delete]
func (h *Handler) DeleteWelcomeChannelConfig(c *gin.Context) {
	var req request.DeleteWelcomeConfigRequest

	if err := c.BindJSON(&req); err != nil {
		h.log.Fields(logger.Fields{"guildID": req.GuildID}).Error(err, "[handler.DeleteWelcomeChannelConfig] - failed to read JSON")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if req.GuildID == "" {
		h.log.Info("[handler.DeleteWelcomeChannelConfig] - guild id empty")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("guild_id is required"), nil))
		return
	}

	if err := h.entities.DeleteWelcomeChannelConfig(req); err != nil {
		h.log.Fields(logger.Fields{"guildID": req.GuildID}).Error(err, "[handler.DeleteWelcomeChannelConfig] - failed to delete welcome config")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(response.ResponseMessage{Message: "OK"}, nil, nil, nil))
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

// CreateSalesTrackerConfig     godoc
// @Summary     Create sales tracker config
// @Description Create sales tracker config
// @Tags        Config
// @Accept      json
// @Produce     json
// @Param       Request  body request.CreateSalesTrackerConfigRequest true "Create sales tracker config request"
// @Success     200 {object} response.ResponseMessage
// @Router      /configs/sales-tracker [post]
func (h *Handler) CreateSalesTrackerConfig(c *gin.Context) {
	var req request.CreateSalesTrackerConfigRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.log.Fields(logger.Fields{"request": req}).Error(err, "[handler.CreateSalesTrackerConfig] - failed to read JSON request")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	err := h.entities.CreateSalesTrackerConfig(req)
	if err != nil {
		h.log.Fields(logger.Fields{"request": req}).Error(err, "[handler.CreateSalesTrackerConfig] - entities.CreateSalesTrackerConfig failed")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	c.JSON(http.StatusOK, response.CreateResponse(response.ResponseMessage{Message: "OK"}, nil, err, nil))
}

// GetJoinLeaveChannelConfig     godoc
// @Summary     Get join-leave channel config
// @Description Get join-leave channel config
// @Tags        ConfigChannel
// @Accept      json
// @Produce     json
// @Param       guild_id   query  string true  "Guild ID"
// @Success     200 {object} response.GetVoteChannelConfigResponse
// @Router      /config-channels/join-leave [get]
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
// @Tags        ConfigChannel
// @Accept      json
// @Produce     json
// @Param       Request  body request.UpsertJoinLeaveChannelConfigRequest true "Upsert join-leave channel config request"
// @Success     200 {object} response.GetVoteChannelConfigResponse
// @Router      /config-channels/join-leave [post]
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
// @Tags        ConfigChannel
// @Accept      json
// @Produce     json
// @Param       Request  body request.DeleteJoinLeaveChannelConfigRequest true "Delete join-leave channel config request"
// @Success     200 {object} response.ResponseMessage
// @Router      /config-channels/join-leave [delete]
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
// @Tags        ConfigChannel
// @Accept      json
// @Produce     json
// @Param       Request body request.CreateTipConfigNotify true "config notify request"
// @Success     200 {object} response.ResponseMessage
// @Router      /config-channels/tip-notify [post]
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
// @Tags        ConfigChannel
// @Accept      json
// @Produce     json
// @Param       guild_id query string true "guild id"
// @Success     200 {object} response.ListConfigNotifyResponse
// @Router      /config-channels/tip-notify [get]
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
// @Tags        ConfigChannel
// @Accept      json
// @Produce     json
// @Param       id path string true "id of config notify"
// @Success     200 {object} response.ResponseMessage
// @Router      /config-channels/tip-notify/{id} [delete]
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

// GetGuildConfigDaoProposal     godoc
// @Summary     Get dao proposal channel config
// @Description Get dao proposal channel config
// @Tags        ConfigChannel
// @Accept      json
// @Produce     json
// @Param       guild_id   query  string true  "Guild ID"
// @Success     200 {object} response.GetGuildConfigDaoProposal
// @Router      /config-channels/{guild_id}/proposal [get]
func (h *Handler) GetGuildConfigDaoProposal(c *gin.Context) {
	guildId := c.Param("guild_id")
	if guildId == "" {
		h.log.Info("[handler.GetGuildConfigDaoProposal] - guild id empty")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("guild_id is required"), nil))
		return
	}

	config, err := h.entities.GetGuildConfigDaoProposalByGuildID(guildId)
	if err != nil {
		h.log.Error(err, "[handler.GetGuildConfigDaoProposal] - failed to get configs")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(config, nil, nil, nil))
}

// DeleteVoteChannelConfig     godoc
// @Summary     Delete dao proposal channel config
// @Description Delete dao proposal config
// @Tags        ConfigChannel
// @Accept      json
// @Produce     json
// @Param       Request  body request.DeleteGuildConfigDaoProposal true "Delete dao proposal channel config request"
// @Success     200 {object} response.ResponseMessage
// @Router      /config-channels/proposal [delete]
func (h *Handler) DeleteGuildConfigDaoProposal(c *gin.Context) {
	var req request.DeleteGuildConfigDaoProposal
	if err := c.BindJSON(&req); err != nil {
		h.log.Fields(logger.Fields{"body": req}).Error(err, "[handler.DeleteGuildConfigDaoProposal] - failed to read JSON")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	err := h.entities.DeleteGuildConfigDaoProposalByGuildID(&req)
	if err != nil {
		h.log.Fields(logger.Fields{"body": req}).Error(err, "[handler.DeleteGuildConfigDaoProposalByGuildID] - failed to delete configs")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(response.ResponseMessage{Message: "OK"}, nil, nil, nil))
}

// CreateProposalChannelConfig     godoc
// @Summary     Create proposal channel config
// @Description Create proposal channel config for dao voting
// @Tags        ConfigChannel
// @Accept      json
// @Produce     json
// @Param       Request  body request.CreateProposalChannelConfig true "Create proposal channel config request"
// @Success     200 {object} response.CreateProposalChannelConfigResponse
// @Router      /config-channels/proposal [post]
func (h *Handler) CreateProposalChannelConfig(c *gin.Context) {
	var req request.CreateProposalChannelConfig
	if err := c.ShouldBindJSON(&req); err != nil {
		h.log.Fields(logger.Fields{"request": req}).Error(err, "[handler.CreateProposalChannelConfig] - failed to read JSON request")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	config, err := h.entities.CreateProposalChannelConfig(req)
	if err != nil {
		h.log.Fields(logger.Fields{"request": req}).Error(err, "[handler.CreateProposalChannelConfig] - entities.CreateProposalChannelConfig failed")
		c.JSON(errs.GetStatusCode(err), response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(config, nil, err, nil))
}

// GetGuildConfigDaoTracker     godoc
// @Summary     Get dao tracker channel config
// @Description Get dao tracker channel config
// @Tags        ConfigChannel
// @Accept      json
// @Produce     json
// @Param       guild_id   path  string true  "Guild ID"
// @Success     200 {object} response.GuildConfigDaoTrackerResponse
// @Router      /config-channels/dao-tracker/{guild_id} [get]
func (h *Handler) GetGuildConfigDaoTracker(c *gin.Context) {
	guildId := c.Param("guild_id")
	config, err := h.entities.GetGuildConfigDaoTracker(guildId)
	if err != nil {
		h.log.Error(err, "[handler.GetGuildConfigDaoTracker] - failed to get configs")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(config, nil, nil, nil))
}

// DeleteGuildConfigDaoTracker     godoc
// @Summary     Delete dao tracker channel config
// @Description Delete dao tracker config
// @Tags        ConfigChannel
// @Accept      json
// @Produce     json
// @Param       Request  body request.DeleteGuildConfigDaoTracker true "Delete dao tracker channel config request"
// @Success     200 {object} response.ResponseMessage
// @Router      /config-channels/dao-tracker [delete]
func (h *Handler) DeleteGuildConfigDaoTracker(c *gin.Context) {
	var req request.DeleteGuildConfigDaoTracker
	if err := c.BindJSON(&req); err != nil {
		h.log.Fields(logger.Fields{"request": req}).Error(err, "[handler.DeleteGuildConfigDaoTracker] - failed to read JSON request")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	err := h.entities.DeleteGuildConfigDaoTracker(req)
	if err != nil {
		h.log.Error(err, "[handler.DeleteGuildConfigDaoTracker] - failed to delete configs")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(response.ResponseMessage{Message: "ok"}, nil, nil, nil))
}

// UpsertGuildConfigDaoTracker     godoc
// @Summary     Create tracker channel config
// @Description Create tracker channel config for dao voting
// @Tags        ConfigChannel
// @Accept      json
// @Produce     json
// @Param       Request  body request.UpsertGuildConfigDaoTracer true "Create tracker channel config request"
// @Success     200 {object} response.ResponseMessage
// @Router      /config-channels/dao-tracker [post]
func (h *Handler) UpsertGuildConfigDaoTracker(c *gin.Context) {
	var req request.UpsertGuildConfigDaoTracer
	if err := c.BindJSON(&req); err != nil {
		h.log.Fields(logger.Fields{"request": req}).Error(err, "[handler.UpsertGuildConfigDaoTracker] - failed to read JSON request")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	err := h.entities.UpsertGuildConfigDaoTracker(req)
	if err != nil {
		h.log.Error(err, "[handler.UpsertGuildConfigDaoTracker] - failed to upsert configs")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(response.ResponseMessage{Message: "ok"}, nil, nil, nil))
}

// CreateCommonwealthDiscussionSubscription     godoc
// @Summary     Subscribe commonwealth discussion
// @Description Subscribe commonwealth discussion
// @Tags        Data
// @Accept      json
// @Produce     json
// @Param       Request  body request.CreateCommonwealthDiscussionSubscription true "Create cw discussion subscription request"
// @Success     200 {object} response.CreateCommonwealthDiscussionSubscription
// @Router      /config-channels/dao-tracker/cw-discussion-subs [post]
func (h *Handler) CreateCommonwealthDiscussionSubscription(c *gin.Context) {
	var req request.CreateCommonwealthDiscussionSubscription
	if err := c.BindJSON(&req); err != nil {
		h.log.Fields(logger.Fields{"request": req}).Error(err, "[handler.CommonwealthDiscussionSubscription] - failed to read JSON request")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	sub, err := h.entities.CreateCommonwealthDiscussionSubscription(req)
	if err != nil {
		h.log.Fields(logger.Fields{"req": req}).Error(err, "[handler.CommonwealthDiscussionSubscription] - failed to create cw discussion subscription")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(sub, nil, nil, nil))
}
