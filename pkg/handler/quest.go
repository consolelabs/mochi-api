package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/response"
)

// GetUserQuestList     godoc
// @Summary     Get user quest list
// @Description Get user quest list
// @Tags        Quest
// @Accept      json
// @Produce     json
// @Param       req query request.GetUserQuestListRequest true "request"
// @Success     200 {object} response.GetUserQuestListResponse
// @Router      /quests [get]
func (h *Handler) GetUserQuestList(c *gin.Context) {
	var req request.GetUserQuestListRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		h.log.Error(err, "[handler.GetUserQuestList] ShouldBindQuery() failed")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	data, err := h.entities.GetUserQuestList(req)
	if err != nil {
		h.log.Fields(logger.Fields{"req": req}).Error(err, "[handler.GetUserQuestList] entity.GetUserQuestList() failed")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	c.JSON(http.StatusOK, response.CreateResponse(data, nil, nil, nil))
}

// ClaimQuestsRewards     godoc
// @Summary     Claim user quests' rewards
// @Description Claim user quests' rewards
// @Tags        Quest
// @Accept      json
// @Produce     json
// @Param       req body request.ClaimQuestsRewardsRequest true "request"
// @Success     200 {object} response.ClaimQuestsRewardsResponse
// @Router      /quests/claim [POST]
func (h *Handler) ClaimQuestsRewards(c *gin.Context) {
	var req request.ClaimQuestsRewardsRequest
	if err := c.BindJSON(&req); err != nil {
		h.log.Error(err, "[handler.ClaimQuestReward] BindJSON() failed")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	data, err := h.entities.ClaimQuestsRewards(req)
	if err != nil {
		h.log.Error(err, "[handler.ClaimQuestsRewards] entity.ClaimQuestsRewards() failed")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	c.JSON(http.StatusOK, response.CreateResponse(data, nil, nil, nil))
}

// UpdateQuestProgress     godoc
// @Summary     Update user's quest progress
// @Description Update user's quest progress
// @Tags        Quest
// @Accept      json
// @Produce     json
// @Param       req body request.UpdateQuestProgressRequest true "request"
// @Success     200 {string} string "ok"
// @Router      /quests [POST]
func (h *Handler) UpdateQuestProgress(c *gin.Context) {
	var req request.UpdateQuestProgressRequest
	if err := c.BindJSON(&req); err != nil {
		h.log.Error(err, "[handler.UpdateQuestProgress] BindJSON() failed")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	log := &model.QuestUserLog{
		GuildID: req.GuildID,
		UserID:  req.UserID,
		Action:  req.Action,
	}
	err := h.entities.UpdateUserQuestProgress(log)
	if err != nil {
		h.log.Fields(logger.Fields{"log": log}).Error(err, "[handler.UpdateQuestProgress] entity.UpdateUserQuestProgress() failed")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	c.JSON(http.StatusOK, response.CreateResponse(response.ResponseMessage{Message: "OK"}, nil, nil, nil))
}
