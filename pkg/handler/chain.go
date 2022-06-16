package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) ListAllChain(c *gin.Context) {
	returnChain, err := h.entities.ListAllChain()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": returnChain})
}
