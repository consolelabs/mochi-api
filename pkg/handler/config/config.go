package config

import (
	"errors"
	"net/http"

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

// ToggleActivityConfig     godoc
// @Summary     Toggle activity config
// @Description Toggle activity config
// @Tags        Data
// @Accept      json
// @Produce     json
// @Param       activity   path  string true  "Activity name"
// @Param       guild_id   query  string true  "Guild ID"
// @Success     200 {object} response.ToggleActivityConfigResponse
// @Router      /data/configs/activities/{activity} [post]
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

// GetListCommandPermissions     godoc
// @Summary     Get list command permissions
// @Description Get list command permissions
// @Tags        Tono
// @Accept      json
// @Produce     json
// @Param       req   query  request.CommandPermissionsRequest true  "request"
// @Success     200 {object} response.CommandPermissions
// @Router      /config/command-permissions [get]
func (h *Handler) GetListCommandPermissions(c *gin.Context) {
	req := request.CommandPermissionsRequest{}
	if err := c.ShouldBindQuery(&req); err != nil {
		h.log.Error(err, "[handler.GetListCommandPermissions] ShouldBindQuery() failed")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	data, err := h.entities.GetCommandPermissions(req)
	if err != nil {
		h.log.Error(err, "[handler.GetListCommandPermissions] entities.GetCommandPermissions() failed")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse[any](data, nil, nil, nil))
}

// GetInstallBotUrl  godoc
// @Summary     Get bot install url
// @Description Get bot install url
// @Tags        Command Permission
// @Accept      json
// @Produce     json
// @Success     302
// @Router      /config/install-url [get]
func (h *Handler) GetInstallBotUrl(c *gin.Context) {
	url, err := h.entities.GetInstallBotUrl()
	if err != nil {
		h.log.Error(err, "[handler.GetInstallBotUrl] entities.GetCommandPermissions() failed")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.Redirect(http.StatusFound, url)
}
