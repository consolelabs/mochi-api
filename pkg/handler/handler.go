package handler

import (
	"net/http"

	"github.com/defipod/mochi/pkg/config"
	"github.com/defipod/mochi/pkg/entities"
	"github.com/defipod/mochi/pkg/logger"

	"github.com/gin-gonic/gin"
)

// Handler for app
type Handler struct {
	cfg      config.Config
	entities *entities.Entity
}

// New will return an instance of Auth struct
func New(cfg config.Config, l logger.Log, entities *entities.Entity) (*Handler, error) {

	handler := &Handler{
		cfg:      cfg,
		entities: entities,
	}

	return handler, nil
}

// Healthz handler
// Return "OK"
func (h *Handler) Healthz(c *gin.Context) {
	c.Header("Content-Type", "text/plain")
	c.String(http.StatusOK, "OK")
}
