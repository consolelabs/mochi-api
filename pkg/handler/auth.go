package handler

import (
	"net/http"

	"github.com/defipod/mochi/pkg/consts"
	_ "github.com/defipod/mochi/pkg/entities"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/response"
	"github.com/gin-gonic/gin"
)

// Login         godoc
// @Summary     Login
// @Description Login
// @Tags        Auth
// @Accept      json
// @Produce     json
// @Param       Request body     request.LoginRequest true "Login request"
// @Success     200     {object} entities.LoginResponse
// @Router      /auth/login [post]
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

// Logout        godoc
// @Summary     Logout
// @Description Logout
// @Tags        Auth
// @Accept      json
// @Produce     json
// @Success     200 {object} response.LogoutResponse
// @Router      /auth/logout [post]
func (h *Handler) Logout(c *gin.Context) {
	c.SetCookie(consts.TokenCookieKey, "", -1, "/", "", false, true)

	c.JSON(http.StatusOK, response.LogoutResponse{
		Status:  "ok",
		Message: "logged out",
	})
}
