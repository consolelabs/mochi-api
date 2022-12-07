package handler

import (
	"net/http"

	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/response"
	"github.com/gin-gonic/gin"
)

func (h *Handler) NotifyNftCollectionIntegration(c *gin.Context) {
	req := request.SendCollectionIntegrationLogsRequest{}
	if err := c.BindJSON(&req); err != nil {
		h.log.Error(err, "[handler.SendCollectionIntegrationLogs] c.BindJSON() failed")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	h.entities.NotifyNftCollectionIntegration(req)
}
