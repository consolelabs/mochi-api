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
