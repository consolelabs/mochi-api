package handler

import (
	"net/http"
	"time"

	"github.com/defipod/mochi/pkg/model"
	"github.com/gin-gonic/gin"
)

func (h *Handler) IndexUsers(c *gin.Context) {
	var body struct {
		ID       string               `json:"id"`
		Username string               `json:"username"`
		Nickname model.JSONNullString `json:"nickname"`
		JoinDate *time.Time           `json:"join_date"`
		GuildID  string               `json:"guild_id"`
	}

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := &model.User{
		ID:       body.ID,
		Username: body.Username,
		GuildUsers: []*model.GuildUser{
			{
				GuildID:  body.GuildID,
				UserID:   body.ID,
				Nickname: body.Nickname,
			},
		},
	}

	if err := h.repo.Users.Create(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}
