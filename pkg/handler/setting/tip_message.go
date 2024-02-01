package setting

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/defipod/mochi/pkg/model/errors"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/response"
)

// GET /profiles/:profile_id/settings/tip/default-message
// GetUserTipMessage get profile's default tip message
// @ID getUserTipMessage
// @Summary get profile's default tip message
// @Description get profile's default tip message
// @Tags Settings
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param profile_id path string true "profile ID"
// @Success 200 {object} response.GetUserTipMessageResponse "successful operation"
// @Failure 400 "bad request"
// @Failure 401 "unauthorized"
// @Failure 404 "not found"
// @Failure 500 "internal server error"
// @Router /profiles/{profile_id}/settings/tip/default-message [get]
func (h *handler) GetUserTipMessage(c *gin.Context) {
	logger.Debug("api call ", c.Request.RequestURI)
	defer logger.Debug("api finish ", c.Request.RequestURI)

	var uri request.UserSettingBaseUriRequest
	if err := c.BindUri(&uri); err != nil {
		logger.WithError(err).Error("BindUri() failed")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	msg, err := h.entities.GetUserTipMessage(uri)
	if err != nil {
		logger.WithField("uri", uri).WithError(err).Error("entities.GetUserGeneralSettings() failed")
		c.JSON(errors.GetStatusCode(err), response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.ToUserTipMessageResponse(msg))
}
