package configrole

import (
	"errors"
	"net/http"
	"strconv"
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

// GetAllRoleReactionConfigs     godoc
// @Summary     Get all role reaction configs
// @Description Get all role reaction configs
// @Tags        Config
// @Accept      json
// @Produce     json
// @Param       guild_id   query  string true  "Guild ID"
// @Success     200 {object} response.DataListRoleReactionResponse
// @Router      /configs/reaction-roles [get]
func (h *Handler) GetAllRoleReactionConfigs(c *gin.Context) {
	guildID := c.Query("guild_id")
	if guildID == "" {
		h.log.Info("[handler.GetAllRoleReactionConfigs] - guild id empty")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("guild_id is required"), nil))
		return
	}

	resp, err := h.entities.ListAllReactionRoles(guildID)
	if err != nil {
		h.log.Fields(logger.Fields{"guildID": guildID}).Error(err, "[handler.GetAllRoleReactionConfigs] - failed to list all reaction roles")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(resp, nil, nil, nil))
}

// AddReactionRoleConfig     godoc
// @Summary     Add reaction role config
// @Description Add reaction role config
// @Tags        Config
// @Accept      json
// @Produce     json
// @Param       Request  body request.RoleReactionUpdateRequest true "Add reaction role config request"
// @Success     200 {object} response.RoleReactionConfigResponse
// @Router      /configs/reaction-roles [post]
func (h *Handler) AddReactionRoleConfig(c *gin.Context) {
	var req request.RoleReactionUpdateRequest

	if err := c.BindJSON(&req); err != nil {
		h.log.Fields(logger.Fields{"body": req}).Error(err, "[handler.AddReactionRoleConfig] - failed to read JSON")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	config, err := h.entities.UpdateConfigByMessageID(req)
	if err != nil {
		h.log.Fields(logger.Fields{"body": req}).Error(err, "[handler.AddReactionRoleConfig] - failed to update config my message id")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(config, nil, nil, nil))
}

// RemoveReactionRoleConfig     godoc
// @Summary     Remove reaction role config
// @Description Remove reaction role config
// @Tags        Config
// @Accept      json
// @Produce     json
// @Param       Request  body request.RoleReactionUpdateRequest true "Remove reaction role config request"
// @Success     200 {object} response.ResponseSucess
// @Router      /configs/reaction-roles [delete]
func (h *Handler) RemoveReactionRoleConfig(c *gin.Context) {
	var req request.RoleReactionUpdateRequest

	if err := c.BindJSON(&req); err != nil {
		h.log.Fields(logger.Fields{"body": req}).Error(err, "[handler.RemoveReactionRoleConfig] - failed to read JSON")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	var err error

	if req.RoleID != "" && req.Reaction != "" {
		err = h.entities.RemoveSpecificRoleReaction(req)
	} else {
		err = h.entities.ClearReactionMessageConfig(req)
	}

	if err != nil {
		h.log.Fields(logger.Fields{"body": req}).Error(err, "[handler.RemoveReactionRoleConfig] - failed to remove reaction config")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(response.ResponseMessage{Message: "OK"}, nil, nil, nil))
}

// FilterConfigByReaction     godoc
// @Summary     Filter config by reaction
// @Description Filter config by reaction
// @Tags        Config
// @Accept      json
// @Produce     json
// @Param       Request  body request.RoleReactionRequest true "Filter config by reaction request"
// @Success     200 {object} response.DataFilterConfigByReaction
// @Router      /configs/reaction-roles/filter [post]
func (h *Handler) FilterConfigByReaction(c *gin.Context) {
	var req request.RoleReactionRequest

	if err := c.BindJSON(&req); err != nil {
		h.log.Fields(logger.Fields{"body": req}).Error(err, "[handler.FilterConfigByReaction] - failed to read JSON")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	config, err := h.entities.GetReactionRoleByMessageID(req.GuildID, req.MessageID, req.Reaction)
	if err != nil {
		h.log.Fields(logger.Fields{"guildID": req.GuildID, "messageID": req.MessageID, "reaction": req.Reaction}).Error(err, "[handler.FilterConfigByReaction] - failed to get reaction role by message id")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(config, nil, nil, nil))
}

// GetDefaultRolesByGuildID     godoc
// @Summary     Get default roles by guild id
// @Description Get default roles by guild id
// @Tags        Config
// @Accept      json
// @Produce     json
// @Param       guild_id   query  string true  "Guild ID"
// @Success     200 {object} response.DefaultRoleResponse
// @Router      /configs/default-roles [get]
func (h *Handler) GetDefaultRolesByGuildID(c *gin.Context) {
	guildID := c.Query("guild_id")
	if guildID == "" {
		h.log.Info("[handler.GetDefaultRolesByGuildID] - guild id empty")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("guild_id is required"), nil))
		return
	}

	data, err := h.entities.GetDefaultRoleByGuildID(guildID)
	if err != nil {
		h.log.Fields(logger.Fields{"guildID": guildID}).Error(err, "[handler.GetDefaultRolesByGuildID] - failed to get default roles")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(data, nil, nil, nil))
}

// CreateDefaultRole     godoc
// @Summary     Create default role
// @Description Create default role
// @Tags        Config
// @Accept      json
// @Produce     json
// @Param       Request  body request.CreateDefaultRoleRequest true "Create default role request"
// @Success     200 {object} response.DefaultRoleResponse
// @Router      /configs/default-roles [post]
func (h *Handler) CreateDefaultRole(c *gin.Context) {
	body := request.CreateDefaultRoleRequest{}

	if err := c.BindJSON(&body); err != nil {
		h.log.Fields(logger.Fields{"body": body}).Error(err, "[handler.CreateDefaultRole] - failed to read JSON")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	if err := h.entities.CreateDefaultRoleConfig(body.GuildID, body.RoleID); err != nil {
		h.log.Fields(logger.Fields{"guildID": body.GuildID, "roleID": body.RoleID}).Error(err, "[handler.CreateDefaultRole] - failed to create default role config")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	defaultRole := response.DefaultRole{
		RoleID:  body.RoleID,
		GuildID: body.GuildID,
	}

	c.JSON(http.StatusOK, response.CreateResponse(defaultRole, nil, nil, nil))
}

// DeleteDefaultRole     godoc
// @Summary     Delete default role
// @Description Delete default role
// @Tags        Config
// @Accept      json
// @Produce     json
// @Param       guild_id   query  string true  "Guild ID"
// @Success     200 {object} response.ResponseSucess
// @Router      /configs/default-roles [delete]
func (h *Handler) DeleteDefaultRoleByGuildID(c *gin.Context) {
	guildID := c.Query("guild_id")
	if guildID == "" {
		h.log.Info("[handler.DeleteDefaultRoleByGuildID] - guild id empty")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("guild_id is required"), nil))
		return
	}
	err := h.entities.DeleteDefaultRoleConfig(guildID)
	if err != nil {
		h.log.Fields(logger.Fields{"guildID": guildID}).Error(err, "[handler.DeleteDefaultRoleByGuildID] - failed to delete default role config")
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

	c.JSON(http.StatusOK, response.CreateResponse(response.ResponseMessage{Message: "OK"}, nil, nil, nil))
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
