package handler

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/response"
	"github.com/gin-gonic/gin"
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
		c.JSON(http.StatusBadRequest, gin.H{"error": "guild_id is required"})
	}

	config, err := h.entities.GetGmConfig(guildID)

	if err != nil {
		h.log.Fields(logger.Fields{"guildID": guildID}).Error(err, "[handler.GetGmConfig] - failed to get gm config")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response.GetGmConfigResponse{Message: "OK", Data: config})
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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if req.GuildID == "" {
		h.log.Info("[handler.UpsertGmConfig] - guild id empty")
		c.JSON(http.StatusBadRequest, gin.H{"error": "guild_id is required"})
	}
	if req.ChannelID == "" {
		h.log.Info("[handler.UpsertGmConfig] - channel id empty")
		c.JSON(http.StatusBadRequest, gin.H{"error": "channel_id is required"})
	}

	if err := h.entities.UpsertGmConfig(req); err != nil {
		h.log.Fields(logger.Fields{"guildID": req.GuildID, "channelID": req.ChannelID}).Error(err, "[handler.UpsertGmConfig] - failed to upsert gm config")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response.ResponseMessage{Message: "OK"})
}

// GetGmConfig     godoc
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
		c.JSON(http.StatusBadRequest, gin.H{"error": "guild_id is required"})
	}

	config, err := h.entities.GetWelcomeChannelConfig(guildID)

	if err != nil {
		h.log.Fields(logger.Fields{"guildID": guildID}).Error(err, "[handler.GetWelcomeChannelConfig] - failed to get welcome config")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response.GetWelcomeChannelConfigResponse{Message: "OK", Data: config})
}

// UpsertGmConfig     godoc
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
		c.JSON(http.StatusBadRequest, gin.H{"error": "guild_id is required"})
	}
	if req.ChannelID == "" {
		h.log.Info("[handler.GetWelcomeChannelConfig] - channel id empty")
		c.JSON(http.StatusBadRequest, gin.H{"error": "channel_id is required"})
	}

	config, err := h.entities.UpsertWelcomeChannelConfig(req)
	if err != nil {
		h.log.Fields(logger.Fields{"guildID": req.GuildID, "channelID": req.ChannelID}).Error(err, "[handler.GetWelcomeChannelConfig] - failed to upsert welcome config")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response.GetWelcomeChannelConfigResponse{Message: "OK", Data: config})
}

// UpsertGmConfig     godoc
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
		c.JSON(http.StatusBadRequest, gin.H{"error": "guild_id is required"})
	}

	if err := h.entities.DeleteWelcomeChannelConfig(req); err != nil {
		h.log.Fields(logger.Fields{"guildID": req.GuildID}).Error(err, "[handler.GetWelcomeChannelConfig] - failed to delete welcome config")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response.ResponseMessage{Message: "OK"})
}

// GetGmConfig     godoc
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
		c.JSON(http.StatusBadRequest, gin.H{"error": "guild_id is required"})
	}

	config, err := h.entities.GetVoteChannelConfig(guildID)

	if err != nil {
		h.log.Fields(logger.Fields{"guildID": guildID}).Error(err, "[handler.GetVoteChannelConfig] - failed to get vote channel config")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response.GetVoteChannelConfigResponse{Message: "OK", Data: config})
}

// UpsertGmConfig     godoc
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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if req.GuildID == "" {
		h.log.Info("[handler.UpsertVoteChannelConfig] - guild id empty")
		c.JSON(http.StatusBadRequest, gin.H{"error": "guild_id is required"})
	}
	if req.ChannelID == "" {
		h.log.Info("[handler.UpsertVoteChannelConfig] - channel id empty")
		c.JSON(http.StatusBadRequest, gin.H{"error": "channel_id is required"})
	}

	config, err := h.entities.UpsertVoteChannelConfig(req)
	if err != nil {
		h.log.Fields(logger.Fields{"guildID": req.GuildID, "channelID": req.ChannelID}).Error(err, "[handler.UpsertVoteChannelConfig] - failed to upsert vote channel config")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response.GetVoteChannelConfigResponse{Message: "OK", Data: config})
}

// UpsertGmConfig     godoc
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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if req.GuildID == "" {
		h.log.Info("[handler.DeleteVoteChannelConfig] - guild id empty")
		c.JSON(http.StatusBadRequest, gin.H{"error": "guild_id is required"})
	}

	if err := h.entities.DeleteVoteChannelConfig(req); err != nil {
		h.log.Fields(logger.Fields{"guildID": req.GuildID}).Error(err, "[handler.DeleteVoteChannelConfig] - failed to delete vote channel config")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response.ResponseMessage{Message: "OK"})
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response.GetUpvoteTiersConfig{Message: "OK", Data: tiers})
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
		c.JSON(http.StatusBadRequest, gin.H{"error": "guild_id is required"})
	}

	config, err := h.entities.GetSalesTrackerConfig(guildID)

	if err != nil {
		h.log.Fields(logger.Fields{"guildID": guildID}).Error(err, "[handler.GetSalesTrackerConfig] - failed to get sales tracker config")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response.GetSalesTrackerConfigResponse{Message: "OK", Data: config})
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response.GetGuildTokensResponse{Data: guildTokens})
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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if req.GuildID == "" {
		h.log.Info("[handler.UpsertGuildTokenConfig] - guild id empty")
		c.JSON(http.StatusBadRequest, gin.H{"error": "guild_id is required"})
		return
	}
	if req.Symbol == "" {
		h.log.Info("[handler.UpsertGuildTokenConfig] - symbol empty")
		c.JSON(http.StatusBadRequest, gin.H{"error": "symbol is required"})
		return
	}

	if err := h.entities.UpsertGuildTokenConfig(req); err != nil {
		h.log.Fields(logger.Fields{"guildID": req.GuildID, "symbol": req.Symbol, "active": req.Active}).Error(err, "[handler.UpsertGuildTokenConfig] - failed to upsert guild token config")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response.ResponseMessage{Message: "OK"})
}

// ConfigLevelRole     godoc
// @Summary     Config Level role
// @Description Config level role
// @Tags        Config
// @Accept      json
// @Produce     json
// @Param       Request  body request.ConfigLevelRoleRequest true "Config level role request"
// @Success     200 {object} response.ResponseMessage
// @Router      /configs/level-roles [post]
func (h *Handler) ConfigLevelRole(c *gin.Context) {
	var req request.ConfigLevelRoleRequest
	if err := c.BindJSON(&req); err != nil {
		h.log.Fields(logger.Fields{"guildID": req.GuildID, "roleID": req.RoleID, "level": req.Level}).Error(err, "[handler.ConfigLevelRole] - failed to read JSON")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	if req.GuildID == "" {
		h.log.Info("[handler.ConfigLevelRole] - guild id empty")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("guild_id is required"), nil))
		return
	}
	if req.RoleID == "" {
		h.log.Info("[handler.ConfigLevelRole] - role id empty")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("role_id is required"), nil))
		return
	}
	if req.Level == 0 {
		h.log.Info("[handler.ConfigLevelRole] - level empty")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("invalid level"), nil))
		return
	}

	if err := h.entities.ConfigLevelRole(req); err != nil {
		h.log.Fields(logger.Fields{"guildID": req.GuildID, "roleID": req.RoleID, "level": req.Level}).Error(err, "[handler.ConfigLevelRole] - failed to config level role")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(response.ResponseSucess{
		Success: true,
	}, nil, nil, nil))
}

// GetLevelRoleConfig     godoc
// @Summary     Get level role config
// @Description Get level role config
// @Tags        Config
// @Accept      json
// @Produce     json
// @Param       guild_id   path  string true  "Guild ID"
// @Success     200 {object} response.GetLevelRoleConfigsResponse
// @Router      /configs/level-roles/{guild_id} [get]
func (h *Handler) GetLevelRoleConfigs(c *gin.Context) {
	guildID := c.Param("guild_id")
	if guildID == "" {
		h.log.Info("[handler.GetLevelRoleConfigs] - guild id empty")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("guild_id is required"), nil))
		return
	}

	data, err := h.entities.GetGuildLevelRoleConfigs(guildID)
	if err != nil {
		h.log.Fields(logger.Fields{"guildID": guildID}).Error(err, "[handler.GetLevelRoleConfigs] - failed to get guild level role config")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(data, nil, nil, nil))
}

// RemoveLevelRoleConfig     godoc
// @Summary     Remove level role config
// @Description Remove level role config
// @Tags        Config
// @Accept      json
// @Produce     json
// @Param       guild_id   path  string true  "Guild ID"
// @Success     200 {object} response.ResponseMessage
// @Router      /configs/level-roles/{guild_id} [delete]
func (h *Handler) RemoveLevelRoleConfig(c *gin.Context) {
	guildID := c.Param("guild_id")
	if guildID == "" {
		h.log.Info("[handler.RemoveLevelRoleConfig] - guild id empty")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("guild_id is required"), nil))
		return
	}

	level := c.Query("level")
	if level == "" {
		h.log.Info("[handler.RemoveLevelRoleConfig] - level empty")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("level is required"), nil))
		return
	}

	levelNr, err := strconv.Atoi(level)
	if err != nil {
		h.log.Fields(logger.Fields{"guildID": guildID, "level": level}).Error(err, "[handler.RemoveLevelRoleConfig] - invalid level")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("invalid level"), nil))
		return
	}

	if err := h.entities.RemoveGuildLevelRoleConfig(guildID, levelNr); err != nil {
		h.log.Fields(logger.Fields{"guildID": guildID, "level": level}).Error(err, "[handler.RemoveLevelRoleConfig] - failed to remove guild level role config")
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
		c.JSON(http.StatusBadRequest, gin.H{"error": "guild_id is required"})
		return
	}

	roles, err := h.entities.ListGuildGroupNFTRoles(guildID)
	if err != nil {
		h.log.Fields(logger.Fields{"guildID": guildID}).Error(err, "[handler.ListGuildGroupNFTRoles] - failed to list all nft roles")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response.ListGuildGroupNFTRolesResponse{Data: roles})
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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newRole, err := h.entities.NewGuildGroupNFTRoleConfig(req)
	if err != nil {
		h.log.Fields(logger.Fields{"body": req}).Error(err, "[handler.NewGuildGroupNFTRole] - failed to create nft role config")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response.NewGuildGroupNFTRoleResponse{Message: "OK", Data: newRole})
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
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, response.ResponseMessage{Message: "OK"})
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response.ResponseMessage{Message: "OK"})
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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if req.GuildID == "" {
		h.log.Info("[handler.ConfigRepostReaction] - guild id empty")
		c.JSON(http.StatusBadRequest, gin.H{"error": "guild_id is required"})
		return
	}
	if req.Emoji == "" {
		h.log.Info("[handler.ConfigRepostReaction] - emoji empty")
		c.JSON(http.StatusBadRequest, gin.H{"error": "emoji is required"})
		return
	}
	if req.Quantity < 1 {
		h.log.Info("[handler.ConfigRepostReaction] - quantity empty")
		c.JSON(http.StatusBadRequest, gin.H{"error": "quantity is required"})
		return
	}
	if req.RepostChannelID == "" {
		h.log.Info("[handler.ConfigRepostReaction] - channel id empty")
		c.JSON(http.StatusBadRequest, gin.H{"error": "repost_channel_id is required"})
		return
	}

	if err := h.entities.ConfigRepostReaction(req); err != nil {
		h.log.Fields(logger.Fields{"guildID": req.GuildID, "emoji": req.Emoji, "quantity": req.Quantity, "channel": req.RepostChannelID}).Error(err, "[handler.ConfigRepostReaction] - failed to add config repost reaction")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response.ResponseMessage{Message: "OK"})
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
		c.JSON(http.StatusBadRequest, gin.H{"error": "guild_id is required"})
		return
	}

	data, err := h.entities.GetGuildRepostReactionConfigs(guildID)
	if err != nil {
		h.log.Fields(logger.Fields{"guildID": guildID}).Error(err, "[handler.GetRepostReactionConfigs] - failed to get guild repost reaction config")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response.GetRepostReactionConfigsResponse{Data: data})
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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if req.GuildID == "" {
		h.log.Info("[handler.RemoveRepostReactionConfig] - guild id empty")
		c.JSON(http.StatusBadRequest, gin.H{"error": "guild_id is required"})
		return
	}
	if req.Emoji == "" {
		h.log.Info("[handler.RemoveRepostReactionConfig] - emoji empty")
		c.JSON(http.StatusBadRequest, gin.H{"error": "emoji is required"})
		return
	}

	if err := h.entities.RemoveGuildRepostReactionConfig(req.GuildID, req.Emoji); err != nil {
		h.log.Fields(logger.Fields{"guildID": req.GuildID, "emoji": req.Emoji, "quantity": req.Quantity, "channel": req.RepostChannelID}).Error(err, "[handler.RemoveRepostReactionConfig] - failed to remove repost reaction config")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response.ResponseMessage{Message: "OK"})
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
		c.JSON(http.StatusBadRequest, gin.H{"error": "activity is required"})
		return
	}

	if guildID == "" {
		h.log.Info("[handler.ToggleActivityConfig] - guild id empty")
		c.JSON(http.StatusBadRequest, gin.H{"error": "guild_id is required"})
		return
	}

	config, err := h.entities.ToggleActivityConfig(guildID, activityName)
	if err != nil {
		h.log.Fields(logger.Fields{"guildID": guildID, "activity": activityName}).Error(err, "[handler.ToggleActivityConfig] - failed to toggle activity config")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response.ToggleActivityConfigResponse{Message: "OK", Data: config})
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response.GetAllTwitterConfigResponse{Message: "OK", Data: config})
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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	err = h.entities.CreateTwitterConfig(&cfg)
	if err != nil {
		h.log.Fields(logger.Fields{"body": cfg}).Error(err, "[handler.GetTwitterConfig] - failed to create twitter config")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response.ResponseMessage{Message: "OK"})
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
	if err != nil {
		if strings.Contains(err.Error(), "record not found") {
			h.log.Fields(logger.Fields{"guild_id": guildId}).Info("[handler.GetTwitterHashtagConfig] - hashtag config empty")
			c.JSON(http.StatusOK, gin.H{"data": hashtags})
			return
		}
		h.log.Fields(logger.Fields{"guild_id": guildId}).Error(err, "[handler.GetTwitterHashtagConfig] - failed to get hashtags")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response.GetTwitterHashtagConfigResponse{Data: hashtags})
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
	if err != nil {
		if strings.Contains(err.Error(), "record not found") {
			h.log.Info("[handler.GetTwitterHashtagConfig] - hashtag config empty")
			c.JSON(http.StatusOK, gin.H{"data": hashtags})
			return
		}
		h.log.Error(err, "[handler.GetTwitterHashtagConfig] - failed to get hashtags")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response.GetAllTwitterHashtagConfigResponse{Data: hashtags})
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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response.ResponseMessage{Message: "OK"})
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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.entities.CreateTwitterHashtagConfig(&req)
	if err != nil {
		h.log.Fields(logger.Fields{"body": req}).Error(err, "[handler.CreateTwitterHashtagConfig] - failed to create hashtag")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response.ResponseMessage{Message: "OK"})
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response.GetDefaultTokenResponse{Data: token})
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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.entities.SetDefaultToken(req); err != nil {
		h.log.Fields(logger.Fields{"req": req}).Error(err, "[handler.ConfigDefaultToken] - failed to set default token")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response.ResponseMessage{Message: "OK"})
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response.ResponseMessage{Message: "OK"})
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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.entities.CreateDefaultCollectionSymbol(req); err != nil {
		h.log.Fields(logger.Fields{"req": req}).Error(err, "[handler.CreateDefaultCollectionSymbol] - failed to read JSON")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response.ResponseMessage{Message: "OK"})
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
		c.JSON(http.StatusBadRequest, gin.H{"error": "guild_id is required"})
		return
	}

	config, err := h.entities.GetGuildPruneExclude(guildID)
	if err != nil {
		h.log.Fields(logger.Fields{"guildID": guildID}).Error(err, "[handler.GetGuildPruneExclude] - failed to get prune exclude config")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response.GetGuildPruneExcludeResponse{Message: "OK", Data: config})
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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.entities.UpsertGuildPruneExclude(req); err != nil {
		h.log.Fields(logger.Fields{"req": req}).Error(err, "[handler.UpsertGuildPruneExclude] - failed to upsert guild prune exlude config")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response.ResponseMessage{Message: "OK"})
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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.entities.DeleteGuildPruneExclude(req); err != nil {
		h.log.Fields(logger.Fields{"req": req}).Error(err, "[handler.DeleteGuildPruneExclude] - failed to delete guild prune exlude config")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response.ResponseMessage{Message: "OK"})
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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.entities.EditMessageRepost(&req)
	if err != nil {
		h.log.Fields(logger.Fields{"req": req}).Error(err, "[handler.EditMessageRepost] - fail to edit message repost")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response.ResponseMessage{Message: "OK"})
}
