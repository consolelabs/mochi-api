package tono

import (
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

// TonoCommandPermissions     godoc
// @Summary     Get Tono command permissions
// @Description Get Tono command permissions
// @Tags        Tono
// @Accept      json
// @Produce     json
// @Param       req   query  request.TonoCommandPermissionsRequest true  "request"
// @Success     200 {object} response.TonoCommandPermissions
// @Router      /tono/command-permissions [get]
func (h *Handler) TonoCommandPermissions(c *gin.Context) {
	req := request.TonoCommandPermissionsRequest{}
	if err := c.ShouldBindQuery(&req); err != nil {
		h.log.Error(err, "[handler.TonoCommandPermissions] ShouldBindQuery() failed")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	data, err := h.entities.GetTonoCommandPermissions(req)
	if err != nil {
		h.log.Error(err, "[handler.TonoCommandPermissions] entities.GetTonoCommandPermissions() failed")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse[any](data, nil, nil, nil))
}
