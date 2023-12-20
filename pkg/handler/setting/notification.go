package setting

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/response"
)

// GET /profiles/:profile_id/settings/notifications
// GetUserNotificationSettings get profile's notification settings
// @ID getUserNotificationSettings
// @Summary get profile's notification settings
// @Description get profile's notification settings
// @Tags Settings
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param profile_id path string true "profile ID"
// @Success 200 {object} response.UserNotificationSettingResponse "successful operation"
// @Failure 400 "bad request"
// @Failure 401 "unauthorized"
// @Failure 404 "not found"
// @Failure 500 "internal server error"
// @Router /profiles/{profile_id}/settings/notifications [get]
func (h *handler) GetUserNotificationSettings(c *gin.Context) {
	logger.Debug("api call ", c.Request.RequestURI)
	defer logger.Debug("api finish ", c.Request.RequestURI)

	var uri request.UserSettingBaseUriRequest
	if err := c.BindUri(&uri); err != nil {
		logger.WithError(err).Error("BindUri() failed")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	data, err := h.entities.GetUserNotificationSettings(uri)
	if err != nil {
		logger.WithField("uri", uri).WithError(err).Error("entities.GetUserNotificationSettings() failed")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("failed to get user notification settings"), nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse[any](data, nil, nil, nil))
}

// PUT /profiles/:profile_id/settings/notifications
// UpdateUserNotificationSettings update profile's notification settings
// @ID updateUserNotificationSettings
// @Summary update profile's notification settings
// @Description update profile's notification settings
// @Tags Settings
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param profile_id path string true "profile ID"
// @Param payload body request.UpdateGeneralNotificationSettingPayloadRequest true "payload"
// @Success 200 {object} response.UserNotificationSettingResponse "successful operation"
// @Failure 400 "bad request"
// @Failure 401 "unauthorized"
// @Failure 404 "not found"
// @Failure 500 "internal server error"
// @Router /profiles/{profile_id}/settings/notifications [put]
func (h *handler) UpdateUserNotificationSettings(c *gin.Context) {
	logger.Debug("api call ", c.Request.RequestURI)
	defer logger.Debug("api finish ", c.Request.RequestURI)

	var uri request.UserSettingBaseUriRequest
	if err := c.BindUri(&uri); err != nil {
		logger.WithError(err).Error("BindUri() failed")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	var payload request.UpdateGeneralNotificationSettingPayloadRequest
	if err := c.BindJSON(&payload); err != nil {
		logger.WithError(err).Error("BindJSON() failed")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	data, err := h.entities.UpdateUserNotificationSettings(uri, payload)
	if err != nil {
		logger.WithFields(logrus.Fields{"uri": uri, "payload": payload}).WithError(err).Error("entities.UpdateUserNotificationSettings() failed")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, errors.New("failed to update user notification settings"), nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse[any](data, nil, nil, nil))
}

// PUT /profiles/:profile_id/settings/notifications/activity/:group/:key
// UpdateUserActivityNotificationSettings update profile's activity notification settings
// @ID updateUserActivityNotificationSettings
// @Summary update profile's activity notification settings
// @Description update profile's activity notification settings
// @Tags Settings
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param profile_id path string true "profile ID"
// @Param group path string true "notification group. available values: `wallet`, `app`, `community`"
// @Param key path string true "notification key"
// @Param payload body request.UpdateActivityNotificationSettingPayloadRequest true "payload"
// @Success 200 {object} response.UserNotificationSettingResponse "successful operation"
// @Failure 400 "bad request"
// @Failure 401 "unauthorized"
// @Failure 404 "not found"
// @Failure 500 "internal server error"
// @Router /profiles/{profile_id}/settings/notifications/activity/{group}/{key} [put]
func (h *handler) UpdateUserActivityNotificationSettings(c *gin.Context) {
	logger.Debug("api call ", c.Request.RequestURI)
	defer logger.Debug("api finish ", c.Request.RequestURI)

	var uri request.UpdateActivityNotificationSettingUriRequest
	if err := c.BindUri(&uri); err != nil {
		logger.WithError(err).Error("BindUri() failed")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	var payload request.UpdateActivityNotificationSettingPayloadRequest
	if err := c.BindJSON(&payload); err != nil {
		logger.WithError(err).Error("BindJSON() failed")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	data, err := h.entities.UpdateUserActivityNotificationSettings(uri, payload)
	if err != nil {
		logger.WithFields(logrus.Fields{"uri": uri, "payload": payload}).WithError(err).Error("entities.UpdateUserActivityNotificationSettings() failed")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, errors.New("failed to update user activity notification settings"), nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse[any](data, nil, nil, nil))
}
