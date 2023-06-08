package earn

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

// GetEarnInfoList     godoc
// @Summary     Get earn list
// @Description Get earn list
// @Tags        Earn
// @Accept      json
// @Produce     json
// @Param       page   query  string false  "page"
// @Param       size   query  string false  "size"
// @Success     200 {object} response.EarnInfoListResponse
// @Router      /earn [get]
func (h *Handler) GetEarnInfoList(c *gin.Context) {
	req := request.PaginationRequest{}
	if err := c.ShouldBindQuery(&req); err != nil {
		h.log.Error(err, "[handler.GetTwitterLeaderboard] BindJSON() failed")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	data, err := h.entities.GetEarnInfoList(req)
	if err != nil {
		h.log.Error(err, "[handler.GetEarnInfoList] failed to get earn info list")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, data)
	return
}

// CreateEarnInfo     godoc
// @Summary     Create earn info
// @Description Create earn info
// @Tags        Earn
// @Param       Request  body request.CreateEarnInfoRequest true "Create earn info request"
// @Accept      json
// @Produce     json
// @Success     200 {object} response.EarnInfoResponse
// @Router      /earn [post]
func (h *Handler) CreateEarnInfo(c *gin.Context) {
	var req request.CreateEarnInfoRequest
	if err := c.BindJSON(&req); err != nil {
		h.log.Error(err, "[handler.CreateEarnInfo] BindJSON() failed")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	data, err := h.entities.CreateEarnInfo(&req)
	if err != nil {
		h.log.Error(err, "[handler.GetEarnInfoList] failed to get earn info list")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse[any](data, nil, nil, nil))
	return
}

// CreateUserEarn     godoc
// @Summary     Create user earn
// @Description Create user earn
// @Tags        Earn
// @Param       id   path  string true  "user Id"
// @Param       Request  body request.CreateUserEarnRequest true "Create user earn info"
// @Accept      json
// @Produce     json
// @Success     200 {object} response.UserEarnResponse
// @Router      /users/{id}/earn [post]
func (h *Handler) CreateUserEarn(c *gin.Context) {
	var req request.CreateUserEarnRequest
	if err := c.BindJSON(&req); err != nil {
		h.log.Error(err, "[handler.CreateUserEarn] BindQuery() failed")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	userId := c.Param("id")
	if userId == "" {
		h.log.Info("[handler.CreateUserEarn] - user id empty")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("user_id is required"), nil))
		return
	}
	req.UserId = userId

	data, err := h.entities.CreateUserEarn(&req)
	if err != nil {
		h.log.Error(err, "[handler.CreateUserEarn] failed to create user earn")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse[any](data, nil, nil, nil))
	return
}

// GetUserEarnListByUserId     godoc
// @Summary     Get user earn list
// @Description Get user earn list
// @Tags        Earn
// @Param       id   path  string true  "user Id"
// @Param       page   query  string false  "page"
// @Param       size   query  string false  "size"
// @Accept      json
// @Produce     json
// @Success     200 {object} response.UserEarnListResponse
// @Router      /users/{id}/earn [get]
func (h *Handler) GetUserEarnListByUserId(c *gin.Context) {
	req := request.GetUserEarnListByUserIdRequest{}
	if err := c.ShouldBindQuery(&req); err != nil {
		h.log.Error(err, "[handler.GetTwitterLeaderboard] BindJSON() failed")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	req.UserId = c.Param("id")
	if err := req.Validate(); err != nil {
		h.log.Error(err, "[handler.GetUserEarnListByUserId] validate request failed")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	data, err := h.entities.GetUserEarnInfoListByUserId(req)
	if err != nil {
		h.log.Error(err, "[handler.GetUserEarnListByUserId] failed to create user earn")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, data)
	return
}

// DeleteUserEarn     godoc
// @Summary     Delete user earn
// @Description Delete user earn
// @Tags        Earn
// @Param       id   path  string true  "user Id"
// @Param       earn_id   path  string true  "earn Id"
// @Accept      json
// @Produce     json
// @Success     200 {object} response.ResponseMessage
// @Router      /users/{id}/earn/{earn_id} [delete]
func (h *Handler) DeleteUserEarn(c *gin.Context) {
	req := request.DeleteUserEarnRequest{}
	if err := c.ShouldBindUri(&req); err != nil {
		h.log.Error(err, "[handler.DeleteUserEarn] BindUri() failed")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	err := h.entities.RemoveUserEarn(req)
	if err != nil {
		h.log.Error(err, "[handler.DeleteUserEarn] failed to delete user earn")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse[any](gin.H{"message": "ok"}, nil, nil, nil))
	return
}
