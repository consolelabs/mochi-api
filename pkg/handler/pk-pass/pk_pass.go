package pkpass

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

func (h *Handler) GeneratePkPass(c *gin.Context) {
	req := request.GeneratePkPassRequest{}
	if err := c.BindQuery(&req); err != nil {
		h.log.Fields(logger.Fields{"req": req}).Error(err, "[handler.GeneratePkPass] - failed to read JSON")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	data, err := h.entities.PassHandle(req)
	if err != nil {
		h.log.Error(err, "[handler.IntegrateBinanceKey] failed to get integrate binance data")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	c.Header("Content-Type", "application/vnd.apple.pkpass")
	// Set the content disposition header to indicate a file download
	c.Header("Content-Disposition", "attachment; filename=mypkpass.pkpass")

	// Send the pkpass file data as a response
	c.Data(http.StatusOK, "application/octet-stream", data)
	return
}
