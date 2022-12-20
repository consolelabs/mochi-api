package widget

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/defipod/mochi/pkg/entities"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/response"
)

type Handler struct {
	entities *entities.Entity
	log      logger.Logger
}

func New(entities *entities.Entity, logger logger.Logger) IHandler {
	return &Handler{
		entities: entities,
		log:      logger,
	}
}

// GetUserTokenAlert     godoc
// @Summary     Get user current token alerts
// @Description Get user current token alerts
// @Tags        Widget
// @Accept      json
// @Produce     json
// @Param       discord_id query     string true "Discord ID"
// @Success     200 {object} response.DiscordUserTokenAlertResponse
// @Router      /widget/token-alert [get]
func (h *Handler) GetUserTokenAlert(c *gin.Context) {
	discordID := c.Query("discord_id")
	if discordID == "" {
		h.log.Info("[handler.GetUserTokenAlert] - discord id empty")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("user_id is required"), nil))
		return
	}

	data, err := h.entities.GetUserTokenAlert(discordID)
	if err != nil {
		h.log.Fields(logger.Fields{"discordID": discordID}).Error(err, "[handler.GetUserTokenAlert] - failed to get user token alerts")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, data)
}

// UpsertUserTokenAlert     godoc
// @Summary     Upsert user token alerts
// @Description Upsert user token alerts
// @Tags        Widget
// @Accept      json
// @Produce     json
// @Param       Request  body request.UpsertDiscordUserAlertRequest true "Upsert user token alert"
// @Success     200 {object} response.ResponseMessage
// @Router      /widget/token-alert [post]
func (h *Handler) UpsertUserTokenAlert(c *gin.Context) {
	req := request.UpsertDiscordUserAlertRequest{}
	err := c.BindJSON(&req)
	if err != nil {
		h.log.Fields(logger.Fields{"request": req}).Error(err, "[handler.UpsertUserTokenAlert] - failed to read JSON")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("failed to read JSON"), nil))
		return
	}

	err = h.entities.UpsertUserTokenAlert(&req)
	if err != nil {
		h.log.Fields(logger.Fields{"request": req}).Error(err, "[handler.UpsertUserTokenAlert] - failed to upsert user token alert")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(response.ResponseMessage{Message: "OK"}, nil, nil, nil))
}

// DeleteUserTokenAlert     godoc
// @Summary     Delete user token alerts
// @Description Delete user token alerts
// @Tags        Widget
// @Accept      json
// @Produce     json
// @Param       Request  body request.DeleteDiscordUserAlertRequest true "Delete user token alert"
// @Success     200 {object} response.ResponseMessage
// @Router      /widget/token-alert [delete]
func (h *Handler) DeleteUserTokenAlert(c *gin.Context) {
	req := request.DeleteDiscordUserAlertRequest{}
	err := c.BindJSON(&req)
	if err != nil {
		h.log.Fields(logger.Fields{"request": req}).Error(err, "[handler.DeleteUserTokenAlert] - failed to read JSON")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("failed to read JSON"), nil))
		return
	}

	err = h.entities.DeleteUserTokenAlert(&req)
	if err != nil {
		h.log.Fields(logger.Fields{"request": req}).Error(err, "[handler.DeleteUserTokenAlert] - failed to delete user device")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(response.ResponseMessage{Message: "OK"}, nil, nil, nil))
}

// GetUserDevice     godoc
// @Summary     Get user current device data
// @Description Get user current device data
// @Tags        Widget
// @Accept      json
// @Produce     json
// @Param       device_id query     string true "Device ID"
// @Success     200 {object} response.UserDeviceResponse
// @Router      /widget/device [get]
func (h *Handler) GetUserDevice(c *gin.Context) {
	deviceID := c.Query("device_id")
	if deviceID == "" {
		h.log.Info("[handler.GetUserDevice] - discord id empty")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("user_id is required"), nil))
		return
	}

	data, err := h.entities.GetUserDevice(deviceID)
	if err != nil {
		h.log.Fields(logger.Fields{"deviceID": deviceID}).Error(err, "[handler.GetUserDevice] - failed to get user device")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(data, nil, nil, nil))
}

// UpsertUserDevice     godoc
// @Summary     Upsert user current device data
// @Description Upsert user current device data
// @Tags        Widget
// @Accept      json
// @Produce     json
// @Param       Request  body request.UpsertUserDeviceRequest true "Upsert user device"
// @Success     200 {object} response.ResponseMessage
// @Router      /widget/device [post]
func (h *Handler) UpsertUserDevice(c *gin.Context) {
	req := request.UpsertUserDeviceRequest{}
	err := c.BindJSON(&req)
	if err != nil {
		h.log.Fields(logger.Fields{"request": req}).Error(err, "[handler.UpsertUserDevice] - failed to read JSON")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("failed to read JSON"), nil))
		return
	}

	err = h.entities.UpsertUserDevice(&req)
	if err != nil {
		h.log.Fields(logger.Fields{"request": req}).Error(err, "[handler.UpsertUserDevice] - failed to upsert user device")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(response.ResponseMessage{Message: "OK"}, nil, nil, nil))
}

// DeleteUserDevice     godoc
// @Summary     Delete user current device data
// @Description Delete user current device data
// @Tags        Widget
// @Accept      json
// @Produce     json
// @Param       Request  body request.DeleteUserDeviceRequest true "Delete user device"
// @Success     200 {object} response.ResponseMessage
// @Router      /widget/device [delete]
func (h *Handler) DeleteUserDevice(c *gin.Context) {
	req := request.DeleteUserDeviceRequest{}
	err := c.BindJSON(&req)
	if err != nil {
		h.log.Fields(logger.Fields{"request": req}).Error(err, "[handler.DeleteUserDevice] - failed to read JSON")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("failed to read JSON"), nil))
		return
	}

	err = h.entities.DeleteUserDevice(&req)
	if err != nil {
		h.log.Fields(logger.Fields{"request": req}).Error(err, "[handler.DeleteUserDevice] - failed to upsert user device")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(response.ResponseMessage{Message: "OK"}, nil, nil, nil))
}
