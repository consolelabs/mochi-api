package handler

import (
	"net/http"

	"github.com/defipod/mochi/pkg/consts"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/request"
	"github.com/gin-gonic/gin"
)

func (h *Handler) Login(c *gin.Context) {

	var req request.LoginRequest

	if err := c.BindJSON(&req); err != nil {
		h.log.Fields(logger.Fields{"token": req.AccessToken}).Error(err, "[handler.Login] - failed to read access token")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if req.AccessToken == "" {
		h.log.Info("[handler.Login] - access token empty")
		c.JSON(http.StatusBadRequest, gin.H{"error": "access_token is required"})
		return
	}

	resp, err := h.entities.Login(req.AccessToken)
	if err != nil {
		h.log.Fields(logger.Fields{"token": req.AccessToken}).Error(err, "[handler.Login] - failed to login")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.SetSameSite(http.SameSiteNoneMode)
	c.SetCookie(consts.TokenCookieKey, resp.AccessToken, int(resp.ExpiresAt), "/", "", true, true)
	c.JSON(http.StatusOK, resp)
}

func (h *Handler) Logout(c *gin.Context) {
	c.SetCookie(consts.TokenCookieKey, "", -1, "/", "", false, true)

	c.JSON(200, gin.H{
		"status":  "ok",
		"message": "logged out",
	})
}
