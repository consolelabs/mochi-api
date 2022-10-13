package handler

import (
	"net/http"

	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/response"
	"github.com/gin-gonic/gin"
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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	data, err := h.entities.GetUserQuestList(req)
	if err != nil {
		h.log.Fields(logger.Fields{"req": req}).Error(err, "[handler.GetUserQuestList] entity.GetUserQuestList() failed")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	c.JSON(http.StatusOK, &response.GetUserQuestListResponse{Data: data})
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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data, err := h.entities.ClaimQuestsRewards(req)
	if err != nil {
		h.log.Error(err, "[handler.ClaimQuestsRewards] entity.ClaimQuestsRewards() failed")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, &response.ClaimQuestsRewardsResponse{Data: data})
}
