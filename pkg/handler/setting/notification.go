package setting

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/response"
	sliceutils "github.com/defipod/mochi/pkg/util/slice"
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
		logger.WithField("uri", uri).WithError(err).Error("entity.GetUserNotificationSettings() failed")
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
// @Param payload body request.UpdateNotificationSettingPayloadRequest true "payload"
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

	var payload request.UpdateNotificationSettingPayloadRequest
	if err := payload.Bind(c); err != nil {
		logger.WithError(err).Error("Bind() failed")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	// get all notification flags
	notificationFlags, err := h.entities.ListAllNotificationFlags()
	if err != nil {
		logger.WithError(err).Error("entity.ListAllNotificationFlags() failed")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	// validate flags
	systemFlags := sliceutils.Map(notificationFlags, func(s model.NotificationFlag) string {
		return s.Key
	})
	for k := range payload.Flags {
		if !sliceutils.Contains(systemFlags, k) {
			err := errors.New("flags: insufficiant data")
			logger.Error(err.Error())
			c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
			return
		}
	}

	data, err := h.entities.UpdateUserNotificationSettings(uri, payload, notificationFlags)
	if err != nil {
		logger.WithFields(logrus.Fields{"uri": uri, "payload": payload}).WithError(err).Error("entity.UpdateUserNotificationSettings() failed")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, errors.New("failed to update user notification settings"), nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse[any](data, nil, nil, nil))
}
