package handler

import (
	"net/http"

	"github.com/defipod/mochi/pkg/entities"
	"github.com/defipod/mochi/pkg/logger"

	"github.com/gin-gonic/gin"
)

// Handler for app
type Handler struct {
	entities *entities.Entity
	log      logger.Logger
}

// New will return an instance of Auth struct
func New(entities *entities.Entity) (*Handler, error) {
	handler := &Handler{
		entities: entities,
		log:      *entities.GetLogger(),
	}

	return handler, nil
}

// Healthz handler
// Return "OK"
func (h *Handler) Healthz(c *gin.Context) {
	c.Header("Content-Type", "text/plain")
	c.String(http.StatusOK, "OK")
}
