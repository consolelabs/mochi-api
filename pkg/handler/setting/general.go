package setting

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/response"
)

// GET /profiles/:profile_id/settings/general
// GetUserGeneralSettings get profile's general settings
// @ID getUserGeneralSettings
// @Summary get profile's general settings
// @Description get profile's general settings
// @Tags Settings
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param profile_id path string true "profile ID"
// @Success 200 {object} response.UserGeneralSettingResponse "successful operation"
// @Failure 400 "bad request"
// @Failure 401 "unauthorized"
// @Failure 404 "not found"
// @Failure 500 "internal server error"
// @Router /profiles/{profile_id}/settings/general [get]
func (h *handler) GetUserGeneralSettings(c *gin.Context) {
	logger.Debug("api call ", c.Request.RequestURI)
	defer logger.Debug("api finish ", c.Request.RequestURI)

	var uri request.UserSettingBaseUriRequest
	if err := c.BindUri(&uri); err != nil {
		logger.WithError(err).Error("BindUri() failed")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	payment, privacy, err := h.entities.GetUserGeneralSettings(uri)
	if err != nil {
		logger.WithField("uri", uri).WithError(err).Error("entities.GetUserGeneralSettings() failed")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, errors.New("failed to get user general settings"), nil))
		return
	}

	c.JSON(http.StatusOK, response.ToUserGeneralSettingResponse(payment, privacy))
}

// PUT /profiles/:profile_id/settings/general
// UpdateUserGeneralSettings update profile's general settings
// @ID updateUserGeneralSettings
// @Summary update profile's general settings
// @Description update profile's general settings
// @Tags Settings
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param profile_id path string true "profile ID"
// @Param payload body request.UpdateGeneralSettingsPayloadRequest true "payload"
// @Success 200 {object} response.UserGeneralSettingResponse "successful operation"
// @Failure 400 "bad request"
// @Failure 401 "unauthorized"
// @Failure 404 "not found"
// @Failure 500 "internal server error"
// @Router /profiles/{profile_id}/settings/general [put]
func (h *handler) UpdateUserGeneralSettings(c *gin.Context) {
	logger.Debug("api call ", c.Request.RequestURI)
	defer logger.Debug("api finish ", c.Request.RequestURI)

	var uri request.UserSettingBaseUriRequest
	if err := c.BindUri(&uri); err != nil {
		logger.WithError(err).Error("BindUri() failed")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	var payload request.UpdateGeneralSettingsPayloadRequest
	if err := c.BindJSON(&payload); err != nil {
		logger.WithError(err).Error("BindJSON() failed")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	payment, privacy, err := h.entities.UpdateUserGeneralSettings(uri, payload)
	if err != nil {
		logger.WithFields(logrus.Fields{"uri": uri, "payload": payload}).WithError(err).Error("entities.UpdateUserGeneralSettings() failed")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, errors.New("failed to update user general settings"), nil))
		return
	}

	c.JSON(http.StatusOK, response.ToUserGeneralSettingResponse(payment, privacy))
}
