package handler

import (
	"net/http"

	responseConverter "github.com/defipod/mochi/pkg/util/response_converter"
	"github.com/gin-gonic/gin"
)

func (h *Handler) Guilds(c *gin.Context) {
	guilds, err := h.repo.DiscordGuilds.Gets()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, responseConverter.ConvertGetGuildsResponse(guilds))
}
