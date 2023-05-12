package apikey

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/defipod/mochi/pkg/entities"
	"github.com/defipod/mochi/pkg/logger"
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

// CreateApiKey     godoc
// @Summary     Create apiKey
// @Description Create apiKey
// @Tags        ApiKey
// @Accept      json
// @Produce     json
// @Param       Authorization header   string true "Authorization"
// @Success     200 {object} response.ProfileApiKeyResponse
// @Router      /api-key/me [post]
func (h *Handler) CreateApiKey(c *gin.Context) {
	profileAccessToken := c.GetString("profile_access_token")

	data, err := h.entities.CreateApiKey(profileAccessToken)
	if err != nil {
		h.log.Fields(logger.Fields{"profileAccessToken": profileAccessToken}).Error(err, "[handler.CreateApiKey] - fail to create apiKey")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	c.JSON(http.StatusOK, response.CreateResponse[any](data, nil, nil, nil))
}
