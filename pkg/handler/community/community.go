package community

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/defipod/mochi/pkg/entities"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/model"
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

// HandleUserFeedback     godoc
// @Summary     Post users' feedbacks
// @Description Post users' feedbacks
// @Tags        Community
// @Accept      json
// @Produce     json
// @Param       req body request.UserFeedbackRequest true "request"
// @Success     200 {object} response.ResponseMessage
// @Router      /community/feedback [post]
func (h *Handler) HandleUserFeedback(c *gin.Context) {
	var req request.UserFeedbackRequest
	if err := c.BindJSON(&req); err != nil {
		h.log.Fields(logger.Fields{"body": req}).Error(err, "[handler.HandleUserFeedback] - failed to read JSON")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	err := h.entities.HandleUserFeedback(&req)
	if err != nil {
		h.log.Fields(logger.Fields{"body": req}).Error(err, "[handler.HandleUserFeedback] - failed to handle feedback")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(response.ResponseMessage{Message: "OK"}, nil, nil, nil))
}

// UpdateUserFeedback     godoc
// @Summary     Update users' feedbacks
// @Description Update users' feedbacks
// @Tags        Community
// @Accept      json
// @Produce     json
// @Param       req body request.UpdateUserFeedbackRequest true "request"
// @Success     200 {object} response.UpdateUserFeedbackResponse
// @Router      /community/feedback [put]
func (h *Handler) UpdateUserFeedback(c *gin.Context) {
	var req request.UpdateUserFeedbackRequest
	if err := c.BindJSON(&req); err != nil {
		h.log.Fields(logger.Fields{"body": req}).Error(err, "[handler.UpdateUserFeedback] - failed to read JSON")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	if req.Status != "none" && req.Status != "confirmed" && req.Status != "completed" {
		h.log.Fields(logger.Fields{"body": req}).Info("[handler.UpdateUserFeedback] - invalid feedback status")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, nil, nil))
		return
	}

	data, err := h.entities.UpdateUserFeedback(&req)
	if err != nil {
		h.log.Fields(logger.Fields{"body": req}).Error(err, "[handler.UpdateUserFeedback] - failed to update feedback")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(data, nil, nil, nil))
}

// GetAllUserFeedback     godoc
// @Summary     Get users' feedbacks
// @Description Get users' feedbacks
// @Tags        Community
// @Accept      json
// @Produce     json
// @Param       filter query string false "filter by"
// @Param       value query string false "filtered value"
// @Success     200 {object} response.UserFeedbackResponse
// @Router      /community/feedback [get]
func (h *Handler) GetAllUserFeedback(c *gin.Context) {
	filter := c.Query("filter")
	value := c.Query("value")
	page := c.Query("page")
	size := c.Query("size")
	data, err := h.entities.GetAllUserFeedback(filter, value, page, size)
	if err != nil {
		h.log.Fields(logger.Fields{"filter": filter, "value": value}).Error(err, "[handler.GetAllUserFeedback] - failed to get feedback")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(data, nil, nil, nil))
}

// GetUserQuestList     godoc
// @Summary     Get user quest list
// @Description Get user quest list
// @Tags        Community
// @Accept      json
// @Produce     json
// @Param       req query request.GetUserQuestListRequest true "request"
// @Success     200 {object} response.GetUserQuestListResponse
// @Router      /community/quests [get]
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
// @Tags        Community
// @Accept      json
// @Produce     json
// @Param       req body request.ClaimQuestsRewardsRequest true "request"
// @Success     200 {object} response.ClaimQuestsRewardsResponse
// @Router      /community/quests/claim [POST]
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
// @Tags        Community
// @Accept      json
// @Produce     json
// @Param       req body request.UpdateQuestProgressRequest true "request"
// @Success     200 {string} string "ok"
// @Router      /community/quests/progress [POST]
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

// RemoveGuildGroupNFTRole     godoc
// @Summary     Remove guild group nft role
// @Description Remove guild group nft role
// @Tags        Config
// @Accept      json
// @Produce     json
// @Param       group_config_id  query string true "Group config ID"
// @Success     200 {object} response.ResponseMessage
// @Router      /configs/nft-roles/group [delete]
func (h *Handler) RemoveGuildGroupNFTRole(c *gin.Context) {
	groupConfigID := c.Query("group_config_id")

	if err := h.entities.RemoveGuildGroupNFTRoleConfig(groupConfigID); err != nil {
		h.log.Fields(logger.Fields{"configID": groupConfigID}).Error(err, "[handler.RemoveGuildGroupNFTRole] - failed to remove nft role config")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	c.JSON(http.StatusOK, response.CreateResponse(response.ResponseMessage{Message: "OK"}, nil, nil, nil))
}

// EditMessageRepost     godoc
// @Summary     edit message repost
// @Description edit message repost
// @Tags        Community
// @Accept      json
// @Produce     json
// @Param       Request  body request.EditMessageRepostRequest true "edit message repost request"
// @Success     200 {object} response.ResponseMessage
// @Router      /community/repost-reactions/message-repost [put]
func (h *Handler) EditMessageRepost(c *gin.Context) {
	req := request.EditMessageRepostRequest{}
	if err := c.BindJSON(&req); err != nil {
		h.log.Fields(logger.Fields{"req": req}).Error(err, "[handler.EditMessageRepost] - failed to read JSON")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	err := h.entities.EditMessageRepost(&req)
	if err != nil {
		h.log.Fields(logger.Fields{"req": req}).Error(err, "[handler.EditMessageRepost] - fail to edit message repost")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(response.ResponseMessage{Message: "OK"}, nil, nil, nil))
}

// UpsertLevelUpMessage     godoc
// @Summary     Upsert levelup message
// @Description Upsert levelup message
// @Tags        Community
// @Accept      json
// @Produce     json
// @Param       req  query request.UpsertGuildLevelUpMessageRequest true "Upsert levelup message request"
// @Success     200 {object} response.GetGuildLevelUpMessage
// @Router      /community/levelup [post]
func (h *Handler) UpsertLevelUpMessage(c *gin.Context) {
	req := request.UpsertGuildLevelUpMessageRequest{}
	if err := c.BindJSON(&req); err != nil {
		h.log.Error(err, "[handler.UpsertLevelUpMessage] BindJSON() failed")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	config, err := h.entities.UpsertLevelUpMessage(req)
	if err != nil {
		h.log.Error(err, "[handler.UpsertLevelUpMessage] entity.UpsertLevelUpMessage() failed")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	c.JSON(http.StatusOK, response.CreateResponse(config, nil, nil, nil))
}

// GetLevelUpMessage     godoc
// @Summary     Get levelup message
// @Description Get levelup message
// @Tags        Community
// @Accept      json
// @Produce     json
// @Param       guild_id   query  string true  "Guild ID"
// @Success     200 {object} response.GetGuildLevelUpMessage
// @Router      /community/levelup [get]
func (h *Handler) GetLevelUpMessage(c *gin.Context) {
	guildId := c.Query("guild_id")
	if guildId == "" {
		h.log.Info("[handler.GetLevelUpMessage] missing guild id")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("guild_id is required"), nil))
		return
	}
	config, err := h.entities.GetLevelUpMessage(guildId)
	if err != nil {
		h.log.Error(err, "[handler.GetLevelUpMessage] entity.GetLevelUpMessage() failed")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	c.JSON(http.StatusOK, response.CreateResponse(config, nil, nil, nil))
}

// DeleteLevelUpMessage     godoc
// @Summary     Delete levelup message
// @Description Delete levelup message
// @Tags        Community
// @Accept      json
// @Produce     json
// @Param       req  query request.GuildIDRequest true "Delete levelup message request"
// @Success     200 {object} response.ResponseMessage
// @Router      /community/levelup [delete]
func (h *Handler) DeleteLevelUpMessage(c *gin.Context) {
	req := request.GuildIDRequest{}
	if err := c.BindJSON(&req); err != nil {
		h.log.Error(err, "[handler.UpsertLevelUpMessage] BindJSON() failed")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	err := h.entities.DeleteLevelUpMessage(req)
	if err != nil {
		h.log.Error(err, "[handler.UpsertLevelUpMessage] entity.UpsertLevelUpMessage() failed")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	c.JSON(http.StatusOK, response.CreateResponse(response.ResponseMessage{Message: "OK"}, nil, nil, nil))
}

// GetAllAd     godoc
// @Summary     Get all submitted ads
// @Description Get all submitted ads
// @Tags        Community
// @Accept      json
// @Produce     json
// @Success     200 {object} response.GetAllUserSubmittedAdResponse
// @Router      /community/advertise [get]
func (h *Handler) GetAllAd(c *gin.Context) {
	data, _, err := h.entities.GetAllAd()
	if err != nil {
		h.log.Error(err, "[handler.GetAllAd] entity.GetAllAd() failed")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	c.JSON(http.StatusOK, response.CreateResponse(data, nil, nil, nil))
}

// GetAdById     godoc
// @Summary     Get submitted ad
// @Description Get submitted ad
// @Tags        Community
// @Accept      json
// @Produce     json
// @Param       id   query  string true  "ad's id, or 'random'"
// @Success     200 {object} response.GetUserSubmittedAdResponse
// @Router      /community/advertise [get]
func (h *Handler) GetAdById(c *gin.Context) {
	id := c.Param("id")
	data, err := h.entities.GetAdById(id)
	if err != nil {
		h.log.Error(err, "[handler.GetAdById] entity.GetAdById() failed")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	c.JSON(http.StatusOK, response.CreateResponse(data, nil, nil, nil))
}

// CreateAd     godoc
// @Summary     Create ad submission
// @Description Create ad submission
// @Tags        Community
// @Accept      json
// @Produce     json
// @Param       req  query request.InsertUserAd true "Create ad submission"
// @Success     200 {object} response.ResponseMessage
// @Router      /community/advertise [post]
func (h *Handler) CreateAd(c *gin.Context) {
	req := request.InsertUserAd{}
	if err := c.BindJSON(&req); err != nil {
		h.log.Error(err, "[handler.CreateAd] BindJSON() failed")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	err := h.entities.CreateAd(req)
	if err != nil {
		h.log.Error(err, "[handler.CreateAd] entity.CreateAd() failed")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	c.JSON(http.StatusOK, response.CreateResponse(response.ResponseMessage{Message: "OK"}, nil, nil, nil))
}

// InitAdSubmission     godoc
// @Summary     Init ad submission
// @Description Init ad submission
// @Tags        Community
// @Accept      json
// @Produce     json
// @Param       req  query request.InitAdSubmission true "Initiate ad submission"
// @Success     200 {object} response.ResponseMessage
// @Router      /community/advertise/init [post]
func (h *Handler) InitAdSubmission(c *gin.Context) {
	req := request.InitAdSubmission{}
	if err := c.BindJSON(&req); err != nil {
		h.log.Error(err, "[handler.CreateAd] BindJSON() failed")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	err := h.entities.InitAdSubmission(req)
	if err != nil {
		h.log.Error(err, "[handler.CreateAd] entity.CreateAd() failed")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	c.JSON(http.StatusOK, response.CreateResponse(response.ResponseMessage{Message: "OK"}, nil, nil, nil))
}

// DeleteLevelUpMessage     godoc
// @Summary     Delete ad submission
// @Description Delete ad submission
// @Tags        Community
// @Accept      json
// @Produce     json
// @Param       req  query request.DeleteUserAd true "Delete ad submission"
// @Success     200 {object} response.ResponseMessage
// @Router      /community/advertise [delete]
func (h *Handler) DeleteAdById(c *gin.Context) {
	req := request.DeleteUserAd{}
	if err := c.BindJSON(&req); err != nil {
		h.log.Error(err, "[handler.DeleteAdById] BindJSON() failed")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	if err := h.entities.DeleteAdById(req); err != nil {
		h.log.Error(err, "[handler.DeleteAdById] entities.DeleteAdById() failed")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	c.JSON(http.StatusOK, response.CreateResponse(response.ResponseMessage{Message: "OK"}, nil, nil, nil))
}

// UpdateAdById     godoc
// @Summary     Update ad submission
// @Description Update ad submission
// @Tags        Community
// @Accept      json
// @Produce     json
// @Param       req  query request.UpdateUserAd true "Update ad submission"
// @Success     200 {object} response.ResponseMessage
// @Router      /community/advertise [put]
func (h *Handler) UpdateAdById(c *gin.Context) {
	req := request.UpdateUserAd{}
	if err := c.BindJSON(&req); err != nil {
		h.log.Error(err, "[handler.UpdateAdById] BindJSON() failed")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	if req.Status != "approved" && req.Status != "rejected" {
		err := errors.New("invalid request")
		h.log.Fields(logger.Fields{"req": req}).Error(err, "[handler.UpdateAdById] invalid request body")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	if err := h.entities.UpdateAdById(req); err != nil {
		h.log.Error(err, "[handler.UpdateAdById] entities.UpdateAdById() failed")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	c.JSON(http.StatusOK, response.CreateResponse(response.ResponseMessage{Message: "OK"}, nil, nil, nil))

}

// UpsertUserTag     godoc
// @Summary     Upsert user tag
// @Description Upsert user tag
// @Tags        Community
// @Accept      json
// @Produce     json
// @Param       body body request.UpsertUserTag true "Upsert user tag request"
// @Success     200 {object} response.ResponseMessage
// @Router      /community/tagme [post]
func (h *Handler) UpsertUserTag(c *gin.Context) {
	req := request.UpsertUserTag{}
	if err := c.BindJSON(&req); err != nil {
		h.log.Error(err, "[handler.UpsertUserTag] BindJSON() failed")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	config, err := h.entities.UpsertUserTag(req)
	if err != nil {
		h.log.Error(err, "[handler.UpsertTagme] entity.UpsertTagme() failed")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(config, nil, nil, nil))
}

// GetUserTag   godoc
// @Summary     Get tagme
// @Description Get tagme
// @Tags        Community
// @Accept      json
// @Produce     json
// @Param       user_id   query  string true  "User ID"
// @Param       guild_id   query  string false  "Guild ID"
// @Success     200 {object} response.ResponseMessage
// @Router      /community/tagme [get]
func (h *Handler) GetUserTag(c *gin.Context) {
	userID := c.Query("user_id")
	if userID == "" {
		h.log.Info("[handler.GetUserTag] missing user_id")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("user_id is required"), nil))
		return
	}

	guildID := c.Query("guild_id")

	tag, err := h.entities.GetUserTag(userID, guildID)
	if err != nil {
		h.log.Error(err, "[handler.GetUserTag] entity.GetUserTag() failed")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(tag, nil, nil, nil))
}
