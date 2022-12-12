package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/response"
)

func (h *Handler) AddServersUsageStat(c *gin.Context) {
	var req request.UsageInformation
	if err := c.Bind(&req); err != nil {
		h.log.Fields(logger.Fields{"body": req}).Error(err, "[handler.AddServersUsageStat] - failed to read JSON")
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	err := h.entities.AddServersUsageStats(&req)
	if err != nil {
		h.log.Error(err, "[handler.AddServersUsageStat] - failed to add usage stat")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}

func (h *Handler) AddGitbookClick(c *gin.Context) {
	url := c.Query("url")
	cmd := c.Query("command")
	action := c.Query("action")
	if url == "" || cmd == "" {
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("url and command are required"), nil))
	}

	err := h.entities.AddGitbookClick(url, cmd, action)
	if err != nil {
		h.log.Error(err, "[handler.AddGitbookClick] - faled to add click gitbook info")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	c.Redirect(http.StatusTemporaryRedirect, url)
}
