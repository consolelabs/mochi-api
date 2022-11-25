package handler

import (
	"errors"
	"net/http"
	"strings"

	"github.com/defipod/mochi/pkg/consts"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/response"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

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

// GetGuildtokens     godoc
// @Summary     Get guild tokens
// @Description Get guild tokens
// @Tags        Config
// @Accept      json
// @Produce     json
// @Param       guild_id   query  string false  "Guild ID"
// @Success     200 {object} response.GetGuildTokensResponse
// @Router      /configs/tokens [get]
func (h *Handler) GetGuildTokens(c *gin.Context) {
	guildID := c.Query("guild_id")
	// if guild id empty, return global default tokens
	guildTokens, err := h.entities.GetGuildTokens(guildID)
	if err != nil {
		h.log.Fields(logger.Fields{"guildID": guildID}).Error(err, "[handler.GetGuildTokens] - failed to get guild tokens")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(guildTokens, nil, nil, nil))
}

// UpsertGuildTokenConfig     godoc
// @Summary     Update or insert guild token config
// @Description Update or insert guild token config
// @Tags        Config
// @Accept      json
// @Produce     json
// @Param       Request  body request.UpsertGuildTokenConfigRequest true "Upsert Guild Token Config request"
// @Success     200 {object} response.ResponseMessage
// @Router      /configs/tokens [post]
func (h *Handler) UpsertGuildTokenConfig(c *gin.Context) {
	var req request.UpsertGuildTokenConfigRequest

	if err := c.BindJSON(&req); err != nil {
		h.log.Fields(logger.Fields{"guildID": req.GuildID, "symbol": req.Symbol, "active": req.Active}).Error(err, "[handler.UpsertGuildTokenConfig] - failed to read JSON")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	if req.GuildID == "" {
		h.log.Info("[handler.UpsertGuildTokenConfig] - guild id empty")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("guild_id is required"), nil))
		return
	}
	if req.Symbol == "" {
		h.log.Info("[handler.UpsertGuildTokenConfig] - symbol empty")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("symbol is required"), nil))
		return
	}

	if err := h.entities.UpsertGuildTokenConfig(req); err != nil {
		h.log.Fields(logger.Fields{"guildID": req.GuildID, "symbol": req.Symbol, "active": req.Active}).Error(err, "[handler.UpsertGuildTokenConfig] - failed to upsert guild token config")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(response.ResponseMessage{Message: "OK"}, nil, nil, nil))
}

// ListGuildNFTRoles     godoc
// @Summary     List guild nft roles
// @Description List guild nft roles
// @Tags        Config
// @Accept      json
// @Produce     json
// @Param       guild_id   query  string true  "Guild ID"
// @Success     200 {object} response.ListGuildGroupNFTRolesResponse
// @Router      /configs/nft-roles [get]
func (h *Handler) ListGuildGroupNFTRoles(c *gin.Context) {
	guildID := c.Query("guild_id")
	if guildID == "" {
		h.log.Info("[handler.ListGuildGroupNFTRoles] - guild id empty")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("guild_id is required"), nil))
		return
	}

	roles, err := h.entities.ListGuildGroupNFTRoles(guildID)
	if err != nil {
		h.log.Fields(logger.Fields{"guildID": guildID}).Error(err, "[handler.ListGuildGroupNFTRoles] - failed to list all nft roles")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(roles, nil, nil, nil))
}

// NewGuildNFTRole     godoc
// @Summary     New guild nft role
// @Description New guild nft role
// @Tags        Config
// @Accept      json
// @Produce     json
// @Param       Request  body request.ConfigGroupNFTRoleRequest true "New NFT role request"
// @Success     200 {object} response.NewGuildGroupNFTRoleResponse
// @Router      /configs/nft-roles [post]
func (h *Handler) NewGuildGroupNFTRole(c *gin.Context) {
	var req request.ConfigGroupNFTRoleRequest
	if err := c.BindJSON(&req); err != nil {
		h.log.Fields(logger.Fields{"body": req}).Error(err, "[handler.NewGuildGroupNFTRole] - failed to read JSON")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	newRole, err := h.entities.NewGuildGroupNFTRoleConfig(req)
	if err != nil {
		h.log.Fields(logger.Fields{"body": req}).Error(err, "[handler.NewGuildGroupNFTRole] - failed to create nft role config")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	c.JSON(http.StatusOK, response.CreateResponse(newRole, nil, nil, nil))
}

// RemoveGuildNFTRole     godoc
// @Summary     Remove guild nft role
// @Description Remove guild nft role
// @Tags        Config
// @Accept      json
// @Produce     json
// @Param       config_ids  query string true "32951e68-9959-4e1d-88ca-22b442e19efe|45d06941-468b-4e5e-8b8f-d20c77c87805"
// @Success     200 {object} response.ResponseMessage
// @Router      /configs/nft-roles [delete]
func (h *Handler) RemoveGuildNFTRole(c *gin.Context) {
	configIDs := c.Query("config_ids")

	if configIDs != "" {
		listConfigIDs := strings.Split(configIDs, "|")
		if err := h.entities.RemoveGuildNFTRoleConfig(listConfigIDs); err != nil {
			h.log.Fields(logger.Fields{"configID": listConfigIDs}).Error(err, "[handler.RemoveGuildNFTRole] - failed to remove nft role config")
			c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
			return
		}
		c.JSON(http.StatusOK, response.CreateResponse(response.ResponseMessage{Message: "OK"}, nil, nil, nil))
	}
}

// RemoveGuildGroupNFTRole     godoc
// @Summary     Remove guild group nft role
// @Description Remove guild group nft role
// @Tags        Config
// @Accept      json
// @Produce     json
// @Param       group_config_id  query string true "Group config ID"
// @Success     200 {object} response.ResponseMessage
// @Router      /configs/nft-roles/group [delete]
func (h *Handler) RemoveGuildGroupNFTRole(c *gin.Context) {
	groupConfigID := c.Query("group_config_id")

	if err := h.entities.RemoveGuildGroupNFTRoleConfig(groupConfigID); err != nil {
		h.log.Fields(logger.Fields{"configID": groupConfigID}).Error(err, "[handler.RemoveGuildGroupNFTRole] - failed to remove nft role config")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	c.JSON(http.StatusOK, response.CreateResponse(response.ResponseMessage{Message: "OK"}, nil, nil, nil))
}

// ConfigResposeReaction     godoc
// @Summary     Config Respost reaction
// @Description Config Respost reaction
// @Tags        Config
// @Accept      json
// @Produce     json
// @Param       Request  body request.ConfigRepostRequest true "Config repost reaction request"
// @Success     200 {object} response.ResponseMessage
// @Router      /configs/repost-reactions [post]
func (h *Handler) ConfigRepostReaction(c *gin.Context) {
	var req request.ConfigRepostRequest
	if err := c.BindJSON(&req); err != nil {
		h.log.Fields(logger.Fields{"guildID": req.GuildID, "emoji": req.Emoji, "quantity": req.Quantity, "channel": req.RepostChannelID}).Error(err, "[handler.ConfigRepostReaction] - failed to read JSON")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	if req.GuildID == "" {
		h.log.Info("[handler.ConfigRepostReaction] - guild id empty")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("guild_id is required"), nil))
		return
	}
	if req.Emoji == "" {
		h.log.Info("[handler.ConfigRepostReaction] - emoji empty")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("emoji is required"), nil))
		return
	}
	if req.Quantity < 1 {
		h.log.Info("[handler.ConfigRepostReaction] - quantity empty")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("quantity is required"), nil))
		return
	}
	if req.RepostChannelID == "" {
		h.log.Info("[handler.ConfigRepostReaction] - channel id empty")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("repost_channel_id is required"), nil))
		return
	}

	if err := h.entities.ConfigRepostReaction(req); err != nil {
		h.log.Fields(logger.Fields{"guildID": req.GuildID, "emoji": req.Emoji, "quantity": req.Quantity, "channel": req.RepostChannelID}).Error(err, "[handler.ConfigRepostReaction] - failed to add config repost reaction")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(response.ResponseMessage{Message: "OK"}, nil, nil, nil))
}

// CreateConfigRepostReactionStartStop     godoc
// @Summary     Config Respost reaction with start stop
// @Description Config Respost reaction with start stop
// @Tags        Config
// @Accept      json
// @Produce     json
// @Param       Request  body request.ConfigRepostReactionStartStop true "Config repost reaction start stop request"
// @Success     200 {object} response.ResponseMessage
// @Router      /configs/repost-reactions/start-stop [post]
func (h *Handler) CreateConfigRepostReactionConversation(c *gin.Context) {
	var req request.ConfigRepostReactionStartStop
	if err := c.BindJSON(&req); err != nil {
		h.log.Fields(logger.Fields{"req": req}).Error(err, "[handler.CreateConfigRepostReactionConversation] - failed to read JSON")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	if req.GuildID == "" {
		h.log.Info("[handler.CreateConfigRepostReactionConversation] - guild id empty")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("guild_id is required"), nil))
		return
	}

	if req.EmojiStart == "" {
		h.log.Info("[handler.CreateConfigRepostReactionConversation] - emoji empty")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("emoji is required"), nil))
		return
	}

	if req.EmojiStop == "" {
		h.log.Info("[handler.CreateConfigRepostReactionConversation] - emoji empty")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("emoji is required"), nil))
		return
	}

	if req.RepostChannelID == "" {
		h.log.Info("[handler.CreateConfigRepostReactionConversation] - channel id empty")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("repost_channel_id is required"), nil))
		return
	}

	if err := h.entities.CreateConfigRepostReactionConversation(req); err != nil {
		h.log.Fields(logger.Fields{"req": req}).Error(err, "[handler.CreateConfigRepostReactionConversation] - failed to add config repost reaction start stop")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(response.ResponseMessage{Message: "OK"}, nil, nil, nil))
}

func (h *Handler) RemoveConfigRepostReactionConversation(c *gin.Context) {
	var req request.ConfigRepostReactionStartStop
	if err := c.BindJSON(&req); err != nil {
		h.log.Fields(logger.Fields{"req": req}).Error(err, "[handler.RemoveConfigRepostReactionConversation] - failed to read JSON")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	if req.GuildID == "" {
		h.log.Info("[handler.RemoveConfigRepostReactionConversation] - guild id empty")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("guild_id is required"), nil))
		return
	}

	if req.EmojiStart == "" {
		h.log.Info("[handler.RemoveConfigRepostReactionConversation] - emoji empty")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("emoji is required"), nil))
		return
	}

	if req.EmojiStop == "" {
		h.log.Info("[handler.RemoveConfigRepostReactionConversation] - emoji empty")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("emoji is required"), nil))
		return
	}
	if err := h.entities.RemoveConfigRepostReactionConversation(req.GuildID, req.EmojiStart, req.EmojiStop); err != nil {
		h.log.Fields(logger.Fields{"req": req}).Error(err, "[handler.RemoveConfigRepostReactionConversation] - failed to add config repost reaction start stop")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(response.ResponseMessage{Message: "OK"}, nil, nil, nil))
}

// GetReposeReactionConfigs     godoc
// @Summary     Get Respost reaction configs
// @Description Get Respost reaction configs
// @Tags        Config
// @Accept      json
// @Produce     json
// @Param       guild_id   path  string true  "Guild ID"
// @Success     200 {object} response.GetRepostReactionConfigsResponse
// @Router      /configs/repost-reactions/{guild_id} [get]
func (h *Handler) GetRepostReactionConfigs(c *gin.Context) {
	guildID := c.Param("guild_id")
	if guildID == "" {
		h.log.Info("[handler.GetRepostReactionConfigs] - guild id empty")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("guild_id is required"), nil))
		return
	}
	reactionType := c.Query("reaction_type")
	if reactionType == "" {
		h.log.Info("[handler.GetRepostReactionConfigs] - type empty")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("type is required"), nil))
		return
	}
	if reactionType != consts.ReactionTypeMessage && reactionType != consts.ReactionTypeConversation {
		h.log.Fields(logger.Fields{"reaction_type": reactionType}).Info("[handler.GetRepostReactionConfigs] - reaction_type is invalid")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("reaction_type is invalid"), nil))
		return
	}

	data, err := h.entities.GetGuildRepostReactionConfigs(guildID, reactionType)
	if err != nil {
		h.log.Fields(logger.Fields{"guildID": guildID}).Error(err, "[handler.GetRepostReactionConfigs] - failed to get guild repost reaction config")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(data, nil, nil, nil))
}

// RemoveReposeReactionConfig     godoc
// @Summary     Remove Respost reaction config
// @Description Remove Respost reaction config
// @Tags        Config
// @Accept      json
// @Produce     json
// @Param       Request  body request.ConfigRepostRequest true "Remove repost reaction config request"
// @Success     200 {object} response.ResponseMessage
// @Router      /configs/repost-reactions [delete]
func (h *Handler) RemoveRepostReactionConfig(c *gin.Context) {
	var req request.ConfigRepostRequest
	if err := c.BindJSON(&req); err != nil {
		h.log.Fields(logger.Fields{"guildID": req.GuildID, "emoji": req.Emoji, "quantity": req.Quantity, "channel": req.RepostChannelID}).Error(err, "[handler.RemoveRepostReactionConfig] - failed to read JSON")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	if req.GuildID == "" {
		h.log.Info("[handler.RemoveRepostReactionConfig] - guild id empty")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("guild_id is required"), nil))
		return
	}
	if req.Emoji == "" {
		h.log.Info("[handler.RemoveRepostReactionConfig] - emoji empty")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("emoji is required"), nil))
		return
	}

	if err := h.entities.RemoveGuildRepostReactionConfig(req.GuildID, req.Emoji); err != nil {
		h.log.Fields(logger.Fields{"guildID": req.GuildID, "emoji": req.Emoji, "quantity": req.Quantity, "channel": req.RepostChannelID}).Error(err, "[handler.RemoveRepostReactionConfig] - failed to remove repost reaction config")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(response.ResponseMessage{Message: "OK"}, nil, nil, nil))
}

// ToggleActivityConfig     godoc
// @Summary     Toggle activity config
// @Description Toggle activity config
// @Tags        Config
// @Accept      json
// @Produce     json
// @Param       activity   path  string true  "Activity name"
// @Param       guild_id   query  string true  "Guild ID"
// @Success     200 {object} response.ToggleActivityConfigResponse
// @Router      /configs/activities/{activity} [post]
func (h *Handler) ToggleActivityConfig(c *gin.Context) {
	var (
		activityName = c.Param("activity")
		guildID      = c.Query("guild_id")
	)

	if activityName == "" {
		h.log.Info("[handler.ToggleActivityConfig] - activity name empty")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("activity is required"), nil))
		return
	}

	if guildID == "" {
		h.log.Info("[handler.ToggleActivityConfig] - guild id empty")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("guild_id is required"), nil))
		return
	}

	config, err := h.entities.ToggleActivityConfig(guildID, activityName)
	if err != nil {
		h.log.Fields(logger.Fields{"guildID": guildID, "activity": activityName}).Error(err, "[handler.ToggleActivityConfig] - failed to toggle activity config")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(config, nil, nil, nil))
}

// GetAllTwitterConfig     godoc
// @Summary     Get all twitter config
// @Description Get all twitter config
// @Tags        Config
// @Accept      json
// @Produce     json
// @Param       guild_id   query  string true  "Guild ID"
// @Success     200 {object} response.GetAllTwitterConfigResponse
// @Router      /configs/twitter [get]
func (h *Handler) GetAllTwitterConfig(c *gin.Context) {
	guildID := c.Query("guild_id")
	config, err := h.entities.GetAllTwitterConfig()
	if err != nil {
		h.log.Fields(logger.Fields{"guildID": guildID}).Error(err, "[handler.GetTwitterConfig] - failed to get twitter config")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	c.JSON(http.StatusOK, response.CreateResponse(config, nil, nil, nil))
}

// CreateTwitterConfig     godoc
// @Summary     Create twitter config
// @Description Create twitter config
// @Tags        Config
// @Accept      json
// @Produce     json
// @Param       Request  body model.GuildConfigTwitterFeed true "Create Twitter config request"
// @Success     200 {object} response.ResponseMessage
// @Router      /configs/twitter [post]
func (h *Handler) CreateTwitterConfig(c *gin.Context) {
	cfg := model.GuildConfigTwitterFeed{}
	err := c.BindJSON(&cfg)
	if err != nil {
		h.log.Fields(logger.Fields{"body": cfg}).Error(err, "[handler.CreateTwitterConfig] - failed to read JSON")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	err = h.entities.CreateTwitterConfig(&cfg)
	if err != nil {
		h.log.Fields(logger.Fields{"body": cfg}).Error(err, "[handler.GetTwitterConfig] - failed to create twitter config")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	c.JSON(http.StatusOK, response.CreateResponse(response.ResponseMessage{Message: "OK"}, nil, nil, nil))
}

// GetTwitterHashtagConfig     godoc
// @Summary     Get twitter hashtag config
// @Description get twitter hashtag config
// @Tags        Config
// @Accept      json
// @Produce     json
// @Param       guild_id   path  string true  "Guild ID"
// @Success     200 {object} response.GetTwitterHashtagConfigResponse
// @Router      /configs/twitter/hashtag/{guild_id} [get]
func (h *Handler) GetTwitterHashtagConfig(c *gin.Context) {
	guildId := c.Param("guild_id")
	hashtags, err := h.entities.GetTwitterHashtagConfig(guildId)
	if err != nil && err != gorm.ErrRecordNotFound {
		h.log.Fields(logger.Fields{"guild_id": guildId}).Error(err, "[handler.GetTwitterHashtagConfig] - failed to get hashtags")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	c.JSON(http.StatusOK, response.CreateResponse(hashtags, nil, nil, nil))
}

// GetAllTwitterHashtagConfig     godoc
// @Summary     Get all twitter hashtag config
// @Description get all twitter hashtag config
// @Tags        Config
// @Accept      json
// @Produce     json
// @Success     200 {object} response.GetAllTwitterHashtagConfigResponse
// @Router      /configs/twitter/hashtag [get]
func (h *Handler) GetAllTwitterHashtagConfig(c *gin.Context) {
	hashtags, err := h.entities.GetAllTwitterHashtagConfig()
	if err != nil && err != gorm.ErrRecordNotFound {
		h.log.Error(err, "[handler.GetTwitterHashtagConfig] - failed to get hashtags")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	c.JSON(http.StatusOK, response.CreateResponse(hashtags, nil, nil, nil))
}

// DeleteTwitterHashtagConfig     godoc
// @Summary     Delete twitter hashtag config
// @Description Delete twitter hashtag config
// @Tags        Config
// @Accept      json
// @Produce     json
// @Param       guild_id   path  string true  "Guild ID"
// @Success     200 {object} response.ResponseMessage
// @Router      /configs/twitter/hashtag/{guild_id} [delete]
func (h *Handler) DeleteTwitterHashtagConfig(c *gin.Context) {
	guildId := c.Param("guild_id")
	err := h.entities.DeleteTwitterHashtagConfig(guildId)
	if err != nil {
		h.log.Fields(logger.Fields{"guild_id": guildId}).Error(err, "[handler.GetTwitterHashtagConfig] - failed to delete")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	c.JSON(http.StatusOK, response.CreateResponse(response.ResponseMessage{Message: "OK"}, nil, nil, nil))
}

// CreateTwitterHashtagConfig     godoc
// @Summary     Create twitter hashtag config
// @Description Create twitter hashtag config
// @Tags        Config
// @Accept      json
// @Produce     json
// @Param       Request  body request.TwitterHashtag true "Create twitter hashtag config request"
// @Success     200 {object} response.ResponseMessage
// @Router      /configs/twitter/hashtag [post]
func (h *Handler) CreateTwitterHashtagConfig(c *gin.Context) {
	req := request.TwitterHashtag{}
	err := c.BindJSON(&req)
	if err != nil {
		h.log.Fields(logger.Fields{"body": req}).Error(err, "[handler.CreateTwitterHashtagConfig] - failed to read JSON")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	err = h.entities.CreateTwitterHashtagConfig(&req)
	if err != nil {
		h.log.Fields(logger.Fields{"body": req}).Error(err, "[handler.CreateTwitterHashtagConfig] - failed to create hashtag")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	c.JSON(http.StatusOK, response.CreateResponse(response.ResponseMessage{Message: "OK"}, nil, nil, nil))
}

// GetDefaultToken     godoc
// @Summary     Get Default token
// @Description Get Default token
// @Tags        Config
// @Accept      json
// @Produce     json
// @Param       guild_id   query  string true  "Guild ID"
// @Success     200 {object} response.GetDefaultTokenResponse
// @Router      /configs/tokens/default [get]
func (h *Handler) GetDefaultToken(c *gin.Context) {
	guildID := c.Query("guild_id")
	token, err := h.entities.GetDefaultToken(guildID)
	if err != nil {
		h.log.Fields(logger.Fields{"guild_id": guildID}).Error(err, "[handler.ConfigDefaultToken] - failed to get default token")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(token, nil, nil, nil))
}

// ConfigDefaultToken     godoc
// @Summary     Config Default token
// @Description Config Default token
// @Tags        Config
// @Accept      json
// @Produce     json
// @Param       Request  body request.ConfigDefaultTokenRequest true "Config default token request"
// @Success     200 {object} response.ResponseMessage
// @Router      /configs/tokens/default [post]
func (h *Handler) ConfigDefaultToken(c *gin.Context) {
	req := request.ConfigDefaultTokenRequest{}
	if err := c.BindJSON(&req); err != nil {
		h.log.Fields(logger.Fields{"req": req}).Error(err, "[handler.ConfigDefaultToken] - failed to read JSON")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	if err := h.entities.SetDefaultToken(req); err != nil {
		h.log.Fields(logger.Fields{"req": req}).Error(err, "[handler.ConfigDefaultToken] - failed to set default token")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(response.ResponseMessage{Message: "OK"}, nil, nil, nil))
}

// RemoveDefaultToken     godoc
// @Summary     Remove Default token
// @Description Remove Default token
// @Tags        Config
// @Accept      json
// @Produce     json
// @Param       guild_id   query  string true  "Guild ID"
// @Success     200 {object} response.ResponseMessage
// @Router      /configs/tokens/default [delete]
func (h *Handler) RemoveDefaultToken(c *gin.Context) {
	guildID := c.Query("guild_id")
	if err := h.entities.RemoveDefaultToken(guildID); err != nil {
		h.log.Fields(logger.Fields{"guild_id": guildID}).Error(err, "[handler.RemoveDefaultToken] - failed to remove default token")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(response.ResponseMessage{Message: "OK"}, nil, nil, nil))
}

// CreateDefaultCollectionSymbol     godoc
// @Summary     Create default collection symbol
// @Description Create default collection symbol
// @Tags        Config
// @Accept      json
// @Produce     json
// @Param       Request  body request.ConfigDefaultCollection true "Config Default Collection Symbol request"
// @Success     200 {object} response.ResponseMessage
// @Router      /configs/default-symbol [post]
func (h *Handler) CreateDefaultCollectionSymbol(c *gin.Context) {
	req := request.ConfigDefaultCollection{}
	if err := c.BindJSON(&req); err != nil {
		h.log.Fields(logger.Fields{"req": req}).Error(err, "[handler.CreateDefaultCollectionSymbol] - failed to read JSON")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	if err := h.entities.CreateDefaultCollectionSymbol(req); err != nil {
		h.log.Fields(logger.Fields{"req": req}).Error(err, "[handler.CreateDefaultCollectionSymbol] - failed to read JSON")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(response.ResponseMessage{Message: "OK"}, nil, nil, nil))
}

// GetGuildPruneExclude     godoc
// @Summary     Get prune exclusion config
// @Description Get prune exclusion config
// @Tags        Config
// @Accept      json
// @Produce     json
// @Param       guild_id   query  string true  "Guild ID"
// @Success     200 {object} response.GetGuildPruneExcludeResponse
// @Router      /configs/whitelist-prune [get]
func (h *Handler) GetGuildPruneExclude(c *gin.Context) {
	guildID := c.Query("guild_id")
	if guildID == "" {
		h.log.Info("[handler.GetGuildPruneExclude] - guild id empty")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("guild_id is required"), nil))
		return
	}

	config, err := h.entities.GetGuildPruneExclude(guildID)
	if err != nil {
		h.log.Fields(logger.Fields{"guildID": guildID}).Error(err, "[handler.GetGuildPruneExclude] - failed to get prune exclude config")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(config, nil, err, nil))
}

// UpsertGuildPruneExclude     godoc
// @Summary     Upsert prune exclude config
// @Description Upsert prune exclude config
// @Tags        Config
// @Accept      json
// @Produce     json
// @Param       Request  body request.UpsertGuildPruneExcludeRequest true "Upsert prune exlude request"
// @Success     200 {object} response.ResponseMessage
// @Router      /configs/whitelist-prune [post]
func (h *Handler) UpsertGuildPruneExclude(c *gin.Context) {
	req := request.UpsertGuildPruneExcludeRequest{}
	if err := c.BindJSON(&req); err != nil {
		h.log.Fields(logger.Fields{"req": req}).Error(err, "[handler.UpsertGuildPruneExclude] - failed to read JSON")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	if err := h.entities.UpsertGuildPruneExclude(req); err != nil {
		h.log.Fields(logger.Fields{"req": req}).Error(err, "[handler.UpsertGuildPruneExclude] - failed to upsert guild prune exlude config")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(response.ResponseMessage{Message: "OK"}, nil, nil, nil))
}

// DeleteGuildPruneExclude     godoc
// @Summary     Delete prune exclude config
// @Description Delete prune exclude config
// @Tags        Config
// @Accept      json
// @Produce     json
// @Param       Request  body request.UpsertGuildPruneExcludeRequest true "Upsert prune exlude request"
// @Success     200 {object} response.ResponseMessage
// @Router      /configs/whitelist-prune [delete]
func (h *Handler) DeleteGuildPruneExclude(c *gin.Context) {
	req := request.UpsertGuildPruneExcludeRequest{}
	if err := c.BindJSON(&req); err != nil {
		h.log.Fields(logger.Fields{"req": req}).Error(err, "[handler.DeleteGuildPruneExclude] - failed to read JSON")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	if err := h.entities.DeleteGuildPruneExclude(req); err != nil {
		h.log.Fields(logger.Fields{"req": req}).Error(err, "[handler.DeleteGuildPruneExclude] - failed to delete guild prune exlude config")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(response.ResponseMessage{Message: "OK"}, nil, nil, nil))
}

// EditMessageRepost     godoc
// @Summary     edit message repost
// @Description edit message repost
// @Tags        Config
// @Accept      json
// @Produce     json
// @Param       Request  body request.EditMessageRepostRequest true "edit message repost request"
// @Success     200 {object} response.ResponseMessage
// @Router      /configs/repost-reactions/message-repost [put]
func (h *Handler) EditMessageRepost(c *gin.Context) {
	req := request.EditMessageRepostRequest{}
	if err := c.BindJSON(&req); err != nil {
		h.log.Fields(logger.Fields{"req": req}).Error(err, "[handler.EditMessageRepost] - failed to read JSON")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	err := h.entities.EditMessageRepost(&req)
	if err != nil {
		h.log.Fields(logger.Fields{"req": req}).Error(err, "[handler.EditMessageRepost] - fail to edit message repost")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(response.ResponseMessage{Message: "OK"}, nil, nil, nil))
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

// CreateBlacklistChannelRepostConfig     godoc
// @Summary     Create blacklist channel repost config
// @Description Create blacklist channel repost config
// @Tags        Config
// @Accept      json
// @Produce     json
// @Param       Request  body request.BalcklistChannelRepostConfigRequest true "Upsert join-leave channel config request"
// @Success     200 {object} response.ResponseMessage
// @Router      /configs/repost-reactions/blacklist-channel [post]
func (h *Handler) CreateBlacklistChannelRepostConfig(c *gin.Context) {
	var req request.BalcklistChannelRepostConfigRequest
	if err := c.BindJSON(&req); err != nil {
		h.log.Fields(logger.Fields{"guildID": req.GuildID, "channelID": req.ChannelID}).Error(err, "[handler.CreateBlacklistChannelRepostConfig] - failed to read JSON")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	if req.GuildID == "" {
		h.log.Info("[handler.CreateBlacklistChannelRepostConfig] - guild id empty")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("guild_id is required"), nil))
		return
	}
	if req.ChannelID == "" {
		h.log.Info("[handler.CreateBlacklistChannelRepostConfig] - channel id empty")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("channel_id is required"), nil))
		return
	}
	if err := h.entities.CreateBlacklistChannelRepostConfig(req); err != nil {
		h.log.Fields(logger.Fields{"guildID": req.GuildID, "channelID": req.ChannelID}).Error(err, "[handler.CreateBlacklistChannelRepostConfig] - failed to create blacklist channel repost config")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(response.ResponseMessage{Message: "OK"}, nil, nil, nil))
}

// GetGuildBlacklistChannelRepostConfig     godoc
// @Summary     Get guild blacklist channel repost config
// @Description Get guild blacklist channel repost config
// @Tags        Config
// @Accept      json
// @Produce     json
// @Param       guild_id   query  string true  "Guild ID"
// @Success     200 {object} response.ResponseMessage
// @Router      /configs/repost-reactions/blacklist-channel [get]
func (h *Handler) GetGuildBlacklistChannelRepostConfig(c *gin.Context) {
	guildID := c.Query("guild_id")
	if guildID == "" {
		h.log.Info("[handler.GetGuildBlacklistChannelRepostConfig] - guild id empty")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("guild_id is required"), nil))
		return
	}

	config, err := h.entities.GetGuildBlacklistChannelRepostConfig(guildID)

	if err != nil {
		h.log.Fields(logger.Fields{"guildID": guildID}).Error(err, "[handler.GetGuildBlacklistChannelRepostConfig] - failed to get blacklist channel repost config")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(config, nil, nil, nil))
}

// DeleteBlacklistChannelRepostConfig     godoc
// @Summary     Delete blacklist channel repost config
// @Description Delete blacklist channel repost config
// @Tags        Config
// @Accept      json
// @Produce     json
// @Param       Request  body request.BalcklistChannelRepostConfigRequest true "Delete blacklist channel repost config request"
// @Success     200 {object} response.ResponseMessage
// @Router      /configs/repost-reactions/blacklist-channel [delete]
func (h *Handler) DeleteBlacklistChannelRepostConfig(c *gin.Context) {
	var req request.BalcklistChannelRepostConfigRequest
	if err := c.BindJSON(&req); err != nil {
		h.log.Fields(logger.Fields{"guildID": req.GuildID, "channelID": req.ChannelID}).Error(err, "[handler.DeleteBlacklistChannelRepostConfig] - failed to read JSON")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	if req.GuildID == "" {
		h.log.Info("[handler.DeleteBlacklistChannelRepostConfig] - guild id empty")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("guild_id is required"), nil))
		return
	}
	if req.ChannelID == "" {
		h.log.Info("[handler.DeleteBlacklistChannelRepostConfig] - channel id empty")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("channel_id is required"), nil))
		return
	}
	if err := h.entities.DeleteBlacklistChannelRepostConfig(req); err != nil {
		h.log.Fields(logger.Fields{"guildID": req.GuildID, "channelID": req.ChannelID}).Error(err, "[handler.DeleteBlacklistChannelRepostConfig] - failed to delete blacklist channel repost config")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(response.ResponseMessage{Message: "OK"}, nil, nil, nil))
}

// AddToTwitterBlackList     godoc
// @Summary     Add an user to twitter watching blacklist
// @Description Add an user to twitter watching blacklist
// @Tags        Config
// @Accept      json
// @Produce     json
// @Param       req body request.AddToTwitterBlackListRequest true "request"
// @Success     200 {string} string "ok"
// @Router      /configs/twitter/blacklist [post]
func (h *Handler) AddToTwitterBlackList(c *gin.Context) {
	var req request.AddToTwitterBlackListRequest
	if err := c.BindJSON(&req); err != nil {
		h.log.Error(err, "[handler.AddToTwitterBlackList] BindJSON() failed")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	err := h.entities.AddToTwitterBlackList(req)
	if err != nil {
		h.log.Fields(logger.Fields{"req": req}).Error(err, "[handler.AddToTwitterBlackList] entity.AddToTwitterBlackList() failed")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	c.JSON(http.StatusOK, response.CreateResponse(response.ResponseMessage{Message: "OK"}, nil, nil, nil))
}

// DeleteFromTwitterBlackList     godoc
// @Summary     Delete an user from twitter watching blacklist
// @Description Delete an user from twitter watching blacklist
// @Tags        Config
// @Accept      json
// @Produce     json
// @Param       req query request.DeleteFromTwitterBlackListRequest true "query"
// @Success     200 {string} string "ok"
// @Router      /configs/twitter/blacklist [delete]
func (h *Handler) DeleteFromTwitterBlackList(c *gin.Context) {
	var req request.DeleteFromTwitterBlackListRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		h.log.Error(err, "[handler.DeleteFromTwitterBlackList] ShouldBindQuery() failed")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	err := h.entities.DeleteFromTwitterBlackList(req)
	if err != nil {
		h.log.Fields(logger.Fields{"req": req}).Error(err, "[handler.DeleteFromTwitterBlackList] entity.DeleteFromTwitterBlackList() failed")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	c.JSON(http.StatusOK, response.CreateResponse(response.ResponseMessage{Message: "OK"}, nil, nil, nil))
}

// GetTwitterBlackList     godoc
// @Summary     Get twitter blacklist
// @Description get twitter blacklist
// @Tags        Config
// @Accept      json
// @Produce     json
// @Param       guild_id   query  string false  "Guild ID"
// @Success     200 {object} response.GetTwitterBlackListResponse
// @Router      /configs/twitter/blacklist [get]
func (h *Handler) GetTwitterBlackList(c *gin.Context) {
	guildID := c.Query("guild_id")
	data, err := h.entities.GetTwitterBlackList(guildID)
	if err != nil {
		h.log.Error(err, "[handler.GetTwitterBlackList] entity.GetTwitterBlackList() failed")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	c.JSON(http.StatusOK, response.CreateResponse(data, nil, nil, nil))
}

// GetUserTokenAlert     godoc
// @Summary     Get user current token alerts
// @Description Get user current token alerts
// @Tags        Config
// @Accept      json
// @Produce     json
// @Param       discord_id query     string true "Discord ID"
// @Success     200 {object} response.DiscordUserTokenAlertResponse
// @Router      /configs/token-alert [get]
func (h *Handler) GetUserTokenAlert(c *gin.Context) {
	discordID := c.Query("discord_id")
	if discordID == "" {
		h.log.Info("[handler.GetUserTokenAlert] - discord id empty")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("user_id is required"), nil))
		return
	}

	data, err := h.entities.GetUserTokenAlert(discordID)
	if err != nil {
		h.log.Fields(logger.Fields{"discordID": discordID}).Error(err, "[handler.GetUserTokenAlert] - failed to get user token alerts")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, data)
}

// UpsertUserTokenAlert     godoc
// @Summary     Upsert user token alerts
// @Description Upsert user token alerts
// @Tags        Config
// @Accept      json
// @Produce     json
// @Param       Request  body request.UpsertDiscordUserAlertRequest true "Upsert user token alert"
// @Success     200 {object} response.ResponseMessage
// @Router      /configs/token-alert [post]
func (h *Handler) UpsertUserTokenAlert(c *gin.Context) {
	req := request.UpsertDiscordUserAlertRequest{}
	err := c.BindJSON(&req)
	if err != nil {
		h.log.Fields(logger.Fields{"request": req}).Error(err, "[handler.UpsertUserTokenAlert] - failed to read JSON")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("failed to read JSON"), nil))
		return
	}

	err = h.entities.UpsertUserTokenAlert(&req)
	if err != nil {
		h.log.Fields(logger.Fields{"request": req}).Error(err, "[handler.UpsertUserTokenAlert] - failed to upsert user token alert")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(response.ResponseMessage{Message: "OK"}, nil, nil, nil))
}

// DeleteUserTokenAlert     godoc
// @Summary     Delete user token alerts
// @Description Delete user token alerts
// @Tags        Config
// @Accept      json
// @Produce     json
// @Param       Request  body request.DeleteDiscordUserAlertRequest true "Delete user token alert"
// @Success     200 {object} response.ResponseMessage
// @Router      /configs/token-alert [delete]
func (h *Handler) DeleteUserTokenAlert(c *gin.Context) {
	req := request.DeleteDiscordUserAlertRequest{}
	err := c.BindJSON(&req)
	if err != nil {
		h.log.Fields(logger.Fields{"request": req}).Error(err, "[handler.DeleteUserTokenAlert] - failed to read JSON")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("failed to read JSON"), nil))
		return
	}

	err = h.entities.DeleteUserTokenAlert(&req)
	if err != nil {
		h.log.Fields(logger.Fields{"request": req}).Error(err, "[handler.DeleteUserTokenAlert] - failed to delete user device")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(response.ResponseMessage{Message: "OK"}, nil, nil, nil))
}

// UpsertMonikerConfig     godoc
// @Summary     Upsert moniker config
// @Description Upsert moniker config
// @Tags        Config
// @Accept      json
// @Produce     json
// @Param       Request  body request.UpsertMonikerConfigRequest true "Upsert moniker config"
// @Success     200 {object} response.ResponseMessage
// @Router      /configs/monikers [post]
func (h *Handler) UpsertMonikerConfig(c *gin.Context) {
	var req request.UpsertMonikerConfigRequest
	if err := c.BindJSON(&req); err != nil {
		h.log.Fields(logger.Fields{"request": req}).Error(err, "[handler.UpsertMonikerConfig] - failed to read JSON")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("failed to read JSON"), nil))
		return
	}
	err := h.entities.UpsertMonikerConfig(req)
	if err != nil {
		h.log.Fields(logger.Fields{"request": req}).Error(err, "[handler.UpsertMonikerConfig] - failed to upsert moniker config")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	c.JSON(http.StatusOK, response.CreateResponse(response.ResponseMessage{Message: "OK"}, nil, nil, nil))
}

// GetMonikerByGuildID     godoc
// @Summary     Get moniker configs
// @Description Get moniker configs
// @Tags        Config
// @Accept      json
// @Produce     json
// @Param       guild_id   path  string true  "Guild ID"
// @Success     200 {object} response.MonikerConfigResponse
// @Router      /configs/monikers/{guild_id} [get]
func (h *Handler) GetMonikerByGuildID(c *gin.Context) {
	guildID := c.Param("guild_id")
	if guildID == "" {
		h.log.Info("[handler.GetMonikerByGuildID] - guild id empty")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("guild_id is required"), nil))
		return
	}
	configs, err := h.entities.GetMonikerByGuildID(guildID)
	if err != nil {
		h.log.Fields(logger.Fields{"guildID": guildID}).Error(err, "[handler.GetMonikerByGuildID] - failed to get user token alerts")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	c.JSON(http.StatusOK, response.CreateResponse(configs, nil, nil, nil))
}

// DeleteMonikerConfig     godoc
// @Summary     Delete moniker config
// @Description Delete moniker config
// @Tags        Config
// @Accept      json
// @Produce     json
// @Param       Request  body request.DeleteMonikerConfigRequest true "Delete moinker config"
// @Success     200 {object} response.ResponseMessage
// @Router      /configs/monikers [delete]
func (h *Handler) DeleteMonikerConfig(c *gin.Context) {
	var req request.DeleteMonikerConfigRequest
	if err := c.BindJSON(&req); err != nil {
		h.log.Fields(logger.Fields{"request": req}).Error(err, "[handler.DeleteMonikerConfig] - failed to read JSON")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("failed to read JSON"), nil))
		return
	}
	err := h.entities.DeleteMonikerConfig(req)
	if err != nil {
		h.log.Fields(logger.Fields{"request": req}).Error(err, "[handler.DeleteMonikerConfig] - failed to delete moniker config")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	c.JSON(http.StatusOK, response.CreateResponse(response.ResponseMessage{Message: "OK"}, nil, nil, nil))
}

// GetGuildDefaultCurrency     godoc
// @Summary     Get default currency by guild id
// @Description Get default currency by guild id
// @Tags        Config
// @Accept      json
// @Produce     json
// @Param       guild_id   query  string true  "Guild ID"
// @Success     200 {object} response.GuildConfigDefaultCurrencyResponse
// @Router      /configs/default-currency [get]
func (h *Handler) GetGuildDefaultCurrency(c *gin.Context) {
	guildID := c.Query("guild_id")
	if guildID == "" {
		h.log.Info("[handler.GetGuildDefaultCurrency] - guild id empty")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("guild_id is required"), nil))
		return
	}

	data, err := h.entities.GetGuildDefaultCurrency(guildID)
	if err != nil {
		h.log.Fields(logger.Fields{"guildID": guildID}).Error(err, "[handler.GetGuildDefaultCurrency] - failed to get default currency")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(data, nil, nil, nil))
}

// UpsertGuildDefaultCurrency     godoc
// @Summary     Upsert default currency
// @Description Upsert default currency
// @Tags        Config
// @Accept      json
// @Produce     json
// @Param       Request  body request.UpsertGuildDefaultCurrencyRequest true "Upsert default currency config"
// @Success     200 {object} response.ResponseMessage
// @Router      /configs/default-currency [post]
func (h *Handler) UpsertGuildDefaultCurrency(c *gin.Context) {
	var req request.UpsertGuildDefaultCurrencyRequest
	if err := c.BindJSON(&req); err != nil {
		h.log.Fields(logger.Fields{"request": req}).Error(err, "[handler.UpsertGuildDefaultCurrency] - failed to read JSON")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("failed to read JSON"), nil))
		return
	}

	err := h.entities.UpsertGuildDefaultCurrency(req)
	if err != nil {
		h.log.Fields(logger.Fields{"request": req}).Error(err, "[handler.UpsertGuildDefaultCurrency] - failed to upsert default currency")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(response.ResponseMessage{Message: "OK"}, nil, nil, nil))
}

// DeleteGuildDefaultCurrency     godoc
// @Summary     Delete default currency
// @Description Delete default currency
// @Tags        Config
// @Accept      json
// @Produce     json
// @Param       Request  body request.GuildIDRequest true "Delete default currency config"
// @Success     200 {object} response.ResponseMessage
// @Router      /configs/default-currency [delete]
func (h *Handler) DeleteGuildDefaultCurrency(c *gin.Context) {
	var req request.GuildIDRequest
	if err := c.BindJSON(&req); err != nil {
		h.log.Fields(logger.Fields{"request": req}).Error(err, "[handler.DeleteGuildDefaultCurrency] - failed to read JSON")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("failed to read JSON"), nil))
		return
	}

	err := h.entities.DeleteGuildDefaultCurrency(req.GuildID)
	if err != nil {
		h.log.Fields(logger.Fields{"request": req}).Error(err, "[handler.DeleteGuildDefaultCurrency] - failed to delete default currency")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(response.ResponseMessage{Message: "OK"}, nil, nil, nil))
}
