package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/response"
)

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
