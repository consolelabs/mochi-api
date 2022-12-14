package healthz

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct{}

func New() IHandler {
	return &Handler{}
}

// Healthz handler
// Return "OK"
func (h *Handler) Healthz(c *gin.Context) {
	c.Header("Content-Type", "text/plain")
	c.String(http.StatusOK, "OK")
}
