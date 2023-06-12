package content

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

func (h *Handler) GetTypeContent(c *gin.Context) {
	contentType := c.Param("type")
	if contentType == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "type of content is required"})
		return
	}

	content, err := h.entities.GetContentByType(contentType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(content, nil, nil, nil))
}
