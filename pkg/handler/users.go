package handler

import (
	"net/http"

	"github.com/defipod/mochi/pkg/entities"
	"github.com/defipod/mochi/pkg/request"
	"github.com/gin-gonic/gin"
)

func (h *Handler) IndexUsers(c *gin.Context) {
	var req request.CreateUserRequest

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.entities.CreateUser(req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}

func (h *Handler) GetUser(c *gin.Context) {
	discordID := c.Param("id")
	if discordID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	user, err := h.entities.GetUser(discordID)
	if err != nil {
		if err == entities.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}
