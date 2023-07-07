package configrole

import (
	"errors"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/defipod/mochi/pkg/entities"
	"github.com/defipod/mochi/pkg/job"
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

// GetAllRoleReactionConfigs     godoc
// @Summary     Get all role reaction configs
// @Description Get all role reaction configs
// @Tags        ConfigRole
// @Accept      json
// @Produce     json
// @Param       guild_id   query  string true  "Guild ID"
// @Success     200 {object} response.DataListRoleReactionResponse
// @Router      /config-roles/reaction-roles [get]
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
// @Tags        ConfigRole
// @Accept      json
// @Produce     json
// @Param       Request  body request.RoleReactionUpdateRequest true "Add reaction role config request"
// @Success     200 {object} response.RoleReactionConfigResponse
// @Router      /config-roles/reaction-roles [post]
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
// @Tags        ConfigRole
// @Accept      json
// @Produce     json
// @Param       Request  body request.RoleReactionUpdateRequest true "Remove reaction role config request"
// @Success     200 {object} response.ResponseSucess
// @Router      /config-roles/reaction-roles [delete]
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
// @Tags        ConfigRole
// @Accept      json
// @Produce     json
// @Param       Request  body request.RoleReactionRequest true "Filter config by reaction request"
// @Success     200 {object} response.DataFilterConfigByReaction
// @Router      /config-roles/reaction-roles/filter [post]
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
// @Tags        ConfigRole
// @Accept      json
// @Produce     json
// @Param       guild_id   query  string true  "Guild ID"
// @Success     200 {object} response.DefaultRoleResponse
// @Router      /config-roles/default-roles [get]
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
// @Tags        ConfigRole
// @Accept      json
// @Produce     json
// @Param       Request  body request.CreateDefaultRoleRequest true "Create default role request"
// @Success     200 {object} response.DefaultRoleResponse
// @Router      /config-roles/default-roles [post]
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
// @Tags        ConfigRole
// @Accept      json
// @Produce     json
// @Param       guild_id   query  string true  "Guild ID"
// @Success     200 {object} response.ResponseSucess
// @Router      /config-roles/default-roles [delete]
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
// @Tags        ConfigRole
// @Accept      json
// @Produce     json
// @Param       guild_id   query  string true  "Guild ID"
// @Success     200 {object} response.ListGuildGroupNFTRolesResponse
// @Router      /config-roles/nft-roles [get]
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
// @Tags        ConfigRole
// @Accept      json
// @Produce     json
// @Param       Request  body request.ConfigGroupNFTRoleRequest true "New NFT role request"
// @Success     200 {object} response.NewGuildGroupNFTRoleResponse
// @Router      /config-roles/nft-roles [post]
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
// @Tags        ConfigRole
// @Accept      json
// @Produce     json
// @Param       config_ids  query string true "32951e68-9959-4e1d-88ca-22b442e19efe|45d06941-468b-4e5e-8b8f-d20c77c87805"
// @Success     200 {object} response.ResponseMessage
// @Router      /config-roles/nft-roles [delete]
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
// @Tags        ConfigRole
// @Accept      json
// @Produce     json
// @Param       group_config_id  query string true "Group config ID"
// @Success     200 {object} response.ResponseMessage
// @Router      /config-roles/nft-roles/group [delete]
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
// @Tags        ConfigRole
// @Accept      json
// @Produce     json
// @Param       Request  body request.ConfigLevelRoleRequest true "Config level role request"
// @Success     200 {object} response.ResponseMessage
// @Router      /config-roles/level-roles [post]
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
// @Tags        ConfigRole
// @Accept      json
// @Produce     json
// @Param       guild_id   path  string true  "Guild ID"
// @Success     200 {object} response.GetLevelRoleConfigsResponse
// @Router      /config-roles/level-roles/{guild_id} [get]
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
// @Tags        ConfigRole
// @Accept      json
// @Produce     json
// @Param       guild_id   path  string true  "Guild ID"
// @Success     200 {object} response.ResponseMessage
// @Router      /config-roles/level-roles/{guild_id} [delete]
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

// CreateGuildTokenRole     godoc
// @Summary     Create guild token role config
// @Description Create guild token role config
// @Tags        ConfigRole
// @Accept      json
// @Produce     json
// @Param       Request   body  request.CreateGuildTokenRole true  "Create guild token role config request"
// @Success     200 {object} response.CreateGuildTokenRole
// @Router      /config-roles/token-roles [post]
func (h *Handler) CreateGuildTokenRole(c *gin.Context) {
	req := request.CreateGuildTokenRole{}
	if err := c.ShouldBindJSON(&req); err != nil {
		h.log.Fields(logger.Fields{"req": req}).Error(err, "[handler.CreateGuildTokenRole] - failed to read JSON")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	// list all roles in the guild to check if requested role exists
	trs, err := h.entities.ListGuildTokenRoles(req.GuildID)
	if err != nil {
		h.log.Fields(logger.Fields{"request": req}).Error(err, "[handler.CreateGuildTokenRole] - e.ListGuildTokenRoles failed")
		c.JSON(errs.GetStatusCode(err), response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	// update user roles in the background after the request is done
	defer func() {
		h.log.Fields(logger.Fields{"request": req}).Info("[handler.CreateGuildTokenRole] - start to updateUserRoles...")
		if err := job.NewUpdateUserTokenRolesJob(h.entities, &job.UpdateUserTokenRolesOptions{
			GuildID: req.GuildID,
		}).Run(); err != nil {
			h.log.Fields(logger.Fields{"request": req}).Error(err, "[handler.CreateGuildTokenRole] - failed to run job NewUpdateUserTokenRolesJob")
		}
	}()

	// check if the role already exists, if so, update it
	for _, tr := range trs {
		if tr.RoleID != req.RoleID || (tr.Token != nil && tr.Token.Address != req.Address) {
			continue
		}
		h.log.Fields(logger.Fields{"request": req}).Info("[handler.CreateGuildTokenRole] - role already exists, updating...")
		config, err := h.entities.UpdateGuildTokenRole(tr.ID, request.UpdateGuildTokenRole{
			RoleID:  &req.RoleID,
			Address: &req.Address,
			Chain:   &req.Chain,
			Amount:  &req.Amount,
		})
		if err != nil {
			h.log.Fields(logger.Fields{"request": req}).Error(err, "[handler.CreateGuildTokenRole] - e.UpdateGuildTokenRole failed")
			c.JSON(errs.GetStatusCode(err), response.CreateResponse[any](nil, nil, err, nil))
			return
		}
		c.JSON(http.StatusOK, response.CreateResponse(config, nil, nil, nil))
		return
	}

	config, err := h.entities.CreateGuildTokenRole(req)
	if err != nil {
		h.log.Fields(logger.Fields{"request": req}).Error(err, "[handler.CreateGuildTokenRole] - e.CreateGuildTokenRole failed")
		c.JSON(errs.GetStatusCode(err), response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(config, nil, nil, nil))
}

// ListGuildTokenRole     godoc
// @Summary     Get list token role config of guild
// @Description Get list token role config of guild
// @Tags        ConfigRole
// @Accept      json
// @Produce     json
// @Param       guild_id   path  string true  "Guild ID"
// @Success     200 {object} response.ListGuildTokenRoles
// @Router      /config-roles/token-roles/{guild_id} [get]
func (h *Handler) ListGuildTokenRoles(c *gin.Context) {
	guildID := c.Param("guild_id")
	if guildID == "" {
		h.log.Info("[handler.ListGuildTokenRoles] - guild id empty")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("guild_id is required"), nil))
		return
	}

	data, err := h.entities.ListGuildTokenRoles(guildID)
	if err != nil {
		h.log.Fields(logger.Fields{"guildID": guildID}).Error(err, "[handler.ListGuildTokenRoles] - e.ListGuildTokenRoles failed")
		c.JSON(errs.GetStatusCode(err), response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	currentMinute := time.Now().Minute()
	var nextSync int

	if currentMinute < 20 {
		nextSync = 20 - currentMinute
	} else if currentMinute < 40 {
		nextSync = 40 - currentMinute
	} else {
		nextSync = 60 - currentMinute
	}

	c.JSON(http.StatusOK, response.CreateResponse(data, nil, nil, nil, map[string]interface{}{
		"next_sync": nextSync,
	}))
}

// UpdateGuildTokenRole     godoc
// @Summary     Update guild token role config
// @Description Update guild token role config
// @Tags        ConfigRole
// @Accept      json
// @Produce     json
// @Param       id path  int true  "Config ID"
// @Success     200 {object} response.UpdateGuildTokenRole
// @Router      /config-roles/token-roles/{id} [put]
func (h *Handler) UpdateGuildTokenRole(c *gin.Context) {
	idStr := c.Param("id")
	if idStr == "" {
		h.log.Info("[handler.UpdateGuildTokenRole] - id empty")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("id is required"), nil))
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		h.log.Fields(logger.Fields{
			"id": idStr,
		}).Error(err, "[handler.UpdateGuildTokenRole] - strconv.Atoi failed")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("invalid id"), nil))
		return
	}

	req := request.UpdateGuildTokenRole{}
	if err := c.ShouldBindJSON(&req); err != nil {
		h.log.Fields(logger.Fields{"req": req}).Error(err, "[handler.UpdateGuildTokenRole] - failed to read JSON")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	config, err := h.entities.UpdateGuildTokenRole(id, req)
	if err != nil {
		h.log.Fields(logger.Fields{
			"id": id, "req": req,
		}).Error(err, "[handler.UpdateGuildTokenRole] - e.UpdateGuildTokenRole failed")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(config, nil, nil, nil))
}

// RemoveGuildTokenRole     godoc
// @Summary     Remove guild token role config
// @Description Remove guild token role config
// @Tags        ConfigRole
// @Accept      json
// @Produce     json
// @Param       id  path  int true  "Config ID"
// @Success     200 {object} response.ResponseMessage
// @Router      /config-roles/token-roles/{id} [delete]
func (h *Handler) RemoveGuildTokenRole(c *gin.Context) {
	idStr := c.Param("id")
	if idStr == "" {
		h.log.Info("[handler.RemoveGuildTokenRole] - id empty")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("id is required"), nil))
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		h.log.Fields(logger.Fields{
			"id": idStr,
		}).Error(err, "[handler.RemoveGuildTokenRole] - strconv.Atoi failed")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("invalid id"), nil))
		return
	}

	tr, err := h.entities.GetTokenRole(id)
	if err != nil {
		h.log.Fields(logger.Fields{"id": id}).Error(err, "[handler.RemoveGuildTokenRole] - e.GetTokenRole failed")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("invalid id"), nil))
		return
	}

	// update user roles in the background after the request is done
	defer func() {
		h.log.Fields(logger.Fields{"id": id, "guild_id": tr.GuildID}).Info("[handler.RemoveGuildTokenRole] - start to updateUserRoles...")
		if err := job.NewUpdateUserTokenRolesJob(h.entities, &job.UpdateUserTokenRolesOptions{
			GuildID:       tr.GuildID,
			RolesToRemove: []string{tr.RoleID},
		}).Run(); err != nil {
			h.log.Fields(logger.Fields{"id": id}).Error(err, "[handler.RemoveGuildTokenRole] - failed to run job NewUpdateUserTokenRolesJob")
		}
	}()

	if err := h.entities.RemoveGuildTokenRole(id); err != nil {
		h.log.Fields(logger.Fields{"id": id}).Error(err, "[handler.RemoveGuildTokenRole] - e.RemoveGuildTokenRole failed")
		c.JSON(errs.GetStatusCode(err), response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(response.ResponseMessage{Message: "OK"}, nil, nil, nil))
}

// CreateGuildXPRole     godoc
// @Summary     Create guild XP role config
// @Description Create guild XP role config
// @Tags        ConfigRole
// @Accept      json
// @Produce     json
// @Param       Request   body  request.CreateGuildXPRole true  "Create guild XP role config request"
// @Success     200 {object} response.CreateGuildXPRole
// @Router      /config-roles/xp-roles [post]
func (h *Handler) CreateGuildXPRole(c *gin.Context) {
	req := request.CreateGuildXPRole{}
	if err := c.ShouldBindJSON(&req); err != nil {
		h.log.Fields(logger.Fields{"req": req}).Error(err, "[handler.CreateGuildXPRole] - failed to read JSON")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	config, err := h.entities.CreateGuildXPRole(req)
	if err != nil {
		h.log.Fields(logger.Fields{"request": req}).Error(err, "[handler.CreateGuildXPRole] - e.CreateGuildXPRole failed")
		c.JSON(errs.GetStatusCode(err), response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(config, nil, nil, nil))
}

// ListGuildXPRole     godoc
// @Summary     Get list XP role config of guild
// @Description Get list XP role config of guild
// @Tags        ConfigRole
// @Accept      json
// @Produce     json
// @Param       guild_id   query  string true  "Guild ID"
// @Success     200 {object} response.ListGuildXPRoles
// @Router      /config-roles/xp-roles/ [get]
func (h *Handler) ListGuildXPRoles(c *gin.Context) {
	guildID := c.Query("guild_id")
	if guildID == "" {
		h.log.Info("[handler.ListGuildXPRoles] - guild id empty")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("guild_id is required"), nil))
		return
	}

	data, err := h.entities.ListGuildXPRoles(guildID)
	if err != nil {
		h.log.Fields(logger.Fields{"guildID": guildID}).Error(err, "[handler.ListGuildXPRoles] - e.ListGuildXPRoles failed")
		c.JSON(errs.GetStatusCode(err), response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(data, nil, nil, nil))
}

// RemoveGuildXPRole     godoc
// @Summary     Remove guild XP role config
// @Description Remove guild XP role config
// @Tags        ConfigRole
// @Accept      json
// @Produce     json
// @Param       id  path  int true  "Config ID"
// @Success     200 {object} response.ResponseMessage
// @Router      /config-roles/xp-roles/{id} [delete]
func (h *Handler) RemoveGuildXPRole(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		h.log.Fields(logger.Fields{
			"id": idStr,
		}).Error(err, "[handler.RemoveGuildXPRole] - strconv.Atoi failed")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("invalid id"), nil))
		return
	}

	if err := h.entities.RemoveGuildXPRole(id); err != nil {
		h.log.Fields(logger.Fields{"id": id}).Error(err, "[handler.RemoveGuildXPRole] - e.RemoveGuildXPRole failed")
		c.JSON(errs.GetStatusCode(err), response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(response.ResponseMessage{Message: "OK"}, nil, nil, nil))
}

// CreateGuildAdminRoles     godoc
// @Summary     Create guild admin role config
// @Description Create guild admin role config
// @Tags        ConfigRole
// @Accept      json
// @Produce     json
// @Param       Request   body  request.CreateGuildAdminRoleRequest true  "Create guild admin role config request"
// @Success     200 {object} response.ResponseMessage
// @Router      /config-roles/admin-roles [post]
func (h *Handler) CreateGuildAdminRoles(c *gin.Context) {
	req := request.CreateGuildAdminRoleRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		h.log.Fields(logger.Fields{"req": req}).Error(err, "[handler.CreateGuildAdminRoles] - failed to read JSON")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	err := h.entities.CreateGuildAdminRoles(req)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("roles existed"), nil))
			return
		}
		h.log.Fields(logger.Fields{"request": req}).Error(err, "[handler.CreateGuildAdminRoles] - e.CreateGuildAdminRoles failed")
		c.JSON(errs.GetStatusCode(err), response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(response.ResponseMessage{Message: "OK"}, nil, nil, nil))
}

// ListGuildAdminRoles     godoc
// @Summary     Get list admin role config of guild
// @Description Get list admin role config of guild
// @Tags        ConfigRole
// @Accept      json
// @Produce     json
// @Param       guild_id   query  string true  "Guild ID"
// @Success     200 {object} response.ListGuildAdminRoles
// @Router      /config-roles/admin-roles [get]
func (h *Handler) ListGuildAdminRoles(c *gin.Context) {
	guildID := c.Query("guild_id")
	if guildID == "" {
		h.log.Info("[handler.ListGuildAdminRoles] - guild id empty")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("guild_id is required"), nil))
		return
	}

	data, err := h.entities.ListGuildAdminRoles(guildID)
	if err != nil {
		h.log.Fields(logger.Fields{"guildID": guildID}).Error(err, "[handler.ListGuildAdminRoles] - e.ListGuildAdminRoles failed")
		c.JSON(errs.GetStatusCode(err), response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(data, nil, nil, nil))
}

// RemoveGuildAdminRole     godoc
// @Summary     Remove guild admin role config
// @Description Remove guild admin role config
// @Tags        ConfigRole
// @Accept      json
// @Produce     json
// @Param       id  path  int true  "Config ID"
// @Success     200 {object} response.ResponseMessage
// @Router      /config-roles/admin-roles/{id} [delete]
func (h *Handler) RemoveGuildAdminRole(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		h.log.Fields(logger.Fields{
			"id": idStr,
		}).Error(err, "[handler.RemoveGuildAdminRole] - strconv.Atoi failed")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("invalid id"), nil))
		return
	}

	if err := h.entities.RemoveGuildAdminRole(id); err != nil {
		h.log.Fields(logger.Fields{"id": id}).Error(err, "[handler.RemoveGuildAdminRole] - e.RemoveGuildAdminRole failed")
		c.JSON(errs.GetStatusCode(err), response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(response.ResponseMessage{Message: "OK"}, nil, nil, nil))
}
