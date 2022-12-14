package community

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/defipod/mochi/pkg/consts"
	"github.com/defipod/mochi/pkg/entities"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/model"
	baseerrs "github.com/defipod/mochi/pkg/model/errors"
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
// @Tags        Feedback
// @Accept      json
// @Produce     json
// @Param       req body request.UserFeedbackRequest true "request"
// @Success     200 {object} response.ResponseMessage
// @Router      /feedback [post]
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
// @Tags        Feedback
// @Accept      json
// @Produce     json
// @Param       req body request.UpdateUserFeedbackRequest true "request"
// @Success     200 {object} response.UpdateUserFeedbackResponse
// @Router      /feedback [put]
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
// @Tags        Feedback
// @Accept      json
// @Produce     json
// @Param       filter query string false "filter by"
// @Param       value query string false "filtered value"
// @Success     200 {object} response.UserFeedbackResponse
// @Router      /feedback [get]
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

// ConfigResposeReaction     godoc
// @Summary     Config Respost reaction
// @Description Config Respost reaction
// @Tags        Config
// @Accept      json
// @Produce     json
// @Param       Request  body request.ConfigRepostRequest true "Config repost reaction request"
// @Success     200 {object} response.ResponseMessage
// @Router      /configs/repost-reactions [post]
func (h *Handler) ConfigRepostReaction(c *gin.Context) {
	var req request.ConfigRepostRequest
	if err := c.BindJSON(&req); err != nil {
		h.log.Fields(logger.Fields{"guildID": req.GuildID, "emoji": req.Emoji, "quantity": req.Quantity, "channel": req.RepostChannelID}).Error(err, "[handler.ConfigRepostReaction] - failed to read JSON")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	if req.GuildID == "" {
		h.log.Info("[handler.ConfigRepostReaction] - guild id empty")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("guild_id is required"), nil))
		return
	}
	if req.Emoji == "" {
		h.log.Info("[handler.ConfigRepostReaction] - emoji empty")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("emoji is required"), nil))
		return
	}
	if req.Quantity < 1 {
		h.log.Info("[handler.ConfigRepostReaction] - quantity empty")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("quantity is required"), nil))
		return
	}
	if req.RepostChannelID == "" {
		h.log.Info("[handler.ConfigRepostReaction] - channel id empty")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("repost_channel_id is required"), nil))
		return
	}

	if err := h.entities.ConfigRepostReaction(req); err != nil {
		h.log.Fields(logger.Fields{"guildID": req.GuildID, "emoji": req.Emoji, "quantity": req.Quantity, "channel": req.RepostChannelID}).Error(err, "[handler.ConfigRepostReaction] - failed to add config repost reaction")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(response.ResponseMessage{Message: "OK"}, nil, nil, nil))
}

// CreateConfigRepostReactionStartStop     godoc
// @Summary     Config Respost reaction with start stop
// @Description Config Respost reaction with start stop
// @Tags        Config
// @Accept      json
// @Produce     json
// @Param       Request  body request.ConfigRepostReactionStartStop true "Config repost reaction start stop request"
// @Success     200 {object} response.ResponseMessage
// @Router      /configs/repost-reactions/start-stop [post]
func (h *Handler) CreateConfigRepostReactionConversation(c *gin.Context) {
	var req request.ConfigRepostReactionStartStop
	if err := c.BindJSON(&req); err != nil {
		h.log.Fields(logger.Fields{"req": req}).Error(err, "[handler.CreateConfigRepostReactionConversation] - failed to read JSON")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	if req.GuildID == "" {
		h.log.Info("[handler.CreateConfigRepostReactionConversation] - guild id empty")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("guild_id is required"), nil))
		return
	}

	if req.EmojiStart == "" {
		h.log.Info("[handler.CreateConfigRepostReactionConversation] - emoji empty")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("emoji is required"), nil))
		return
	}

	if req.EmojiStop == "" {
		h.log.Info("[handler.CreateConfigRepostReactionConversation] - emoji empty")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("emoji is required"), nil))
		return
	}

	if req.RepostChannelID == "" {
		h.log.Info("[handler.CreateConfigRepostReactionConversation] - channel id empty")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("repost_channel_id is required"), nil))
		return
	}

	if err := h.entities.CreateConfigRepostReactionConversation(req); err != nil {
		h.log.Fields(logger.Fields{"req": req}).Error(err, "[handler.CreateConfigRepostReactionConversation] - failed to add config repost reaction start stop")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(response.ResponseMessage{Message: "OK"}, nil, nil, nil))
}

func (h *Handler) RemoveConfigRepostReactionConversation(c *gin.Context) {
	var req request.ConfigRepostReactionStartStop
	if err := c.BindJSON(&req); err != nil {
		h.log.Fields(logger.Fields{"req": req}).Error(err, "[handler.RemoveConfigRepostReactionConversation] - failed to read JSON")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	if req.GuildID == "" {
		h.log.Info("[handler.RemoveConfigRepostReactionConversation] - guild id empty")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("guild_id is required"), nil))
		return
	}

	if req.EmojiStart == "" {
		h.log.Info("[handler.RemoveConfigRepostReactionConversation] - emoji empty")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("emoji is required"), nil))
		return
	}

	if req.EmojiStop == "" {
		h.log.Info("[handler.RemoveConfigRepostReactionConversation] - emoji empty")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("emoji is required"), nil))
		return
	}
	if err := h.entities.RemoveConfigRepostReactionConversation(req.GuildID, req.EmojiStart, req.EmojiStop); err != nil {
		h.log.Fields(logger.Fields{"req": req}).Error(err, "[handler.RemoveConfigRepostReactionConversation] - failed to add config repost reaction start stop")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(response.ResponseMessage{Message: "OK"}, nil, nil, nil))
}

// GetReposeReactionConfigs     godoc
// @Summary     Get Respost reaction configs
// @Description Get Respost reaction configs
// @Tags        Config
// @Accept      json
// @Produce     json
// @Param       guild_id   path  string true  "Guild ID"
// @Success     200 {object} response.GetRepostReactionConfigsResponse
// @Router      /configs/repost-reactions/{guild_id} [get]
func (h *Handler) GetRepostReactionConfigs(c *gin.Context) {
	guildID := c.Param("guild_id")
	if guildID == "" {
		h.log.Info("[handler.GetRepostReactionConfigs] - guild id empty")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("guild_id is required"), nil))
		return
	}
	reactionType := c.Query("reaction_type")
	if reactionType == "" {
		h.log.Info("[handler.GetRepostReactionConfigs] - type empty")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("type is required"), nil))
		return
	}
	if reactionType != consts.ReactionTypeMessage && reactionType != consts.ReactionTypeConversation {
		h.log.Fields(logger.Fields{"reaction_type": reactionType}).Info("[handler.GetRepostReactionConfigs] - reaction_type is invalid")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("reaction_type is invalid"), nil))
		return
	}

	data, err := h.entities.GetGuildRepostReactionConfigs(guildID, reactionType)
	if err != nil {
		h.log.Fields(logger.Fields{"guildID": guildID}).Error(err, "[handler.GetRepostReactionConfigs] - failed to get guild repost reaction config")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(data, nil, nil, nil))
}

// RemoveReposeReactionConfig     godoc
// @Summary     Remove Respost reaction config
// @Description Remove Respost reaction config
// @Tags        Config
// @Accept      json
// @Produce     json
// @Param       Request  body request.ConfigRepostRequest true "Remove repost reaction config request"
// @Success     200 {object} response.ResponseMessage
// @Router      /configs/repost-reactions [delete]
func (h *Handler) RemoveRepostReactionConfig(c *gin.Context) {
	var req request.ConfigRepostRequest
	if err := c.BindJSON(&req); err != nil {
		h.log.Fields(logger.Fields{"guildID": req.GuildID, "emoji": req.Emoji, "quantity": req.Quantity, "channel": req.RepostChannelID}).Error(err, "[handler.RemoveRepostReactionConfig] - failed to read JSON")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	if req.GuildID == "" {
		h.log.Info("[handler.RemoveRepostReactionConfig] - guild id empty")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("guild_id is required"), nil))
		return
	}
	if req.Emoji == "" {
		h.log.Info("[handler.RemoveRepostReactionConfig] - emoji empty")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("emoji is required"), nil))
		return
	}

	err := h.entities.RemoveGuildRepostReactionConfig(req.GuildID, req.Emoji)
	if err == baseerrs.ErrRecordNotFound {
		h.log.Fields(logger.Fields{"guildID": req.GuildID, "emoji": req.Emoji, "quantity": req.Quantity, "channel": req.RepostChannelID}).Error(err, "[handler.RemoveRepostReactionConfig] repost reaction config not found")
		c.JSON(http.StatusNotFound, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	if err != nil {
		h.log.Fields(logger.Fields{"guildID": req.GuildID, "emoji": req.Emoji, "quantity": req.Quantity, "channel": req.RepostChannelID}).Error(err, "[handler.RemoveRepostReactionConfig] - failed to remove repost reaction config")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(response.ResponseMessage{Message: "OK"}, nil, nil, nil))
}

// CreateBlacklistChannelRepostConfig     godoc
// @Summary     Create blacklist channel repost config
// @Description Create blacklist channel repost config
// @Tags        Config
// @Accept      json
// @Produce     json
// @Param       Request  body request.BalcklistChannelRepostConfigRequest true "Upsert join-leave channel config request"
// @Success     200 {object} response.ResponseMessage
// @Router      /configs/repost-reactions/blacklist-channel [post]
func (h *Handler) CreateBlacklistChannelRepostConfig(c *gin.Context) {
	var req request.BalcklistChannelRepostConfigRequest
	if err := c.BindJSON(&req); err != nil {
		h.log.Fields(logger.Fields{"guildID": req.GuildID, "channelID": req.ChannelID}).Error(err, "[handler.CreateBlacklistChannelRepostConfig] - failed to read JSON")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	if req.GuildID == "" {
		h.log.Info("[handler.CreateBlacklistChannelRepostConfig] - guild id empty")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("guild_id is required"), nil))
		return
	}
	if req.ChannelID == "" {
		h.log.Info("[handler.CreateBlacklistChannelRepostConfig] - channel id empty")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("channel_id is required"), nil))
		return
	}
	if err := h.entities.CreateBlacklistChannelRepostConfig(req); err != nil {
		h.log.Fields(logger.Fields{"guildID": req.GuildID, "channelID": req.ChannelID}).Error(err, "[handler.CreateBlacklistChannelRepostConfig] - failed to create blacklist channel repost config")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(response.ResponseMessage{Message: "OK"}, nil, nil, nil))
}

// GetGuildBlacklistChannelRepostConfig     godoc
// @Summary     Get guild blacklist channel repost config
// @Description Get guild blacklist channel repost config
// @Tags        Config
// @Accept      json
// @Produce     json
// @Param       guild_id   query  string true  "Guild ID"
// @Success     200 {object} response.ResponseMessage
// @Router      /configs/repost-reactions/blacklist-channel [get]
func (h *Handler) GetGuildBlacklistChannelRepostConfig(c *gin.Context) {
	guildID := c.Query("guild_id")
	if guildID == "" {
		h.log.Info("[handler.GetGuildBlacklistChannelRepostConfig] - guild id empty")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("guild_id is required"), nil))
		return
	}

	config, err := h.entities.GetGuildBlacklistChannelRepostConfig(guildID)

	if err != nil {
		h.log.Fields(logger.Fields{"guildID": guildID}).Error(err, "[handler.GetGuildBlacklistChannelRepostConfig] - failed to get blacklist channel repost config")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(config, nil, nil, nil))
}

// DeleteBlacklistChannelRepostConfig     godoc
// @Summary     Delete blacklist channel repost config
// @Description Delete blacklist channel repost config
// @Tags        Config
// @Accept      json
// @Produce     json
// @Param       Request  body request.BalcklistChannelRepostConfigRequest true "Delete blacklist channel repost config request"
// @Success     200 {object} response.ResponseMessage
// @Router      /configs/repost-reactions/blacklist-channel [delete]
func (h *Handler) DeleteBlacklistChannelRepostConfig(c *gin.Context) {
	var req request.BalcklistChannelRepostConfigRequest
	if err := c.BindJSON(&req); err != nil {
		h.log.Fields(logger.Fields{"guildID": req.GuildID, "channelID": req.ChannelID}).Error(err, "[handler.DeleteBlacklistChannelRepostConfig] - failed to read JSON")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	if req.GuildID == "" {
		h.log.Info("[handler.DeleteBlacklistChannelRepostConfig] - guild id empty")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("guild_id is required"), nil))
		return
	}
	if req.ChannelID == "" {
		h.log.Info("[handler.DeleteBlacklistChannelRepostConfig] - channel id empty")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("channel_id is required"), nil))
		return
	}
	if err := h.entities.DeleteBlacklistChannelRepostConfig(req); err != nil {
		h.log.Fields(logger.Fields{"guildID": req.GuildID, "channelID": req.ChannelID}).Error(err, "[handler.DeleteBlacklistChannelRepostConfig] - failed to delete blacklist channel repost config")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(response.ResponseMessage{Message: "OK"}, nil, nil, nil))
}

// EditMessageRepost     godoc
// @Summary     edit message repost
// @Description edit message repost
// @Tags        Config
// @Accept      json
// @Produce     json
// @Param       Request  body request.EditMessageRepostRequest true "edit message repost request"
// @Success     200 {object} response.ResponseMessage
// @Router      /configs/repost-reactions/message-repost [put]
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

// CreateTwitterPost     godoc
// @Summary     Create twitter post
// @Description Create twitter post
// @Tags        Twitter
// @Accept      json
// @Produce     json
// @Param       Request  body request.TwitterPost true "Create twitter post request"
// @Success     200 {object} response.ResponseMessage
// @Router      /twitter [post]
func (h *Handler) CreateTwitterPost(c *gin.Context) {
	req := request.TwitterPost{}
	err := c.BindJSON(&req)
	if err != nil {
		h.log.Error(err, "[handler.CreateTwitterPost] - failed to read JSON body")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = h.entities.CreateTwitterPost(&req)
	if err != nil {
		h.log.Error(err, "[handler.CreateTwitterPost] - failed to create twitter post")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response.ResponseMessage{Message: "OK"})
}

// CreateTwitterPost     godoc
// @Summary     Create twitter post
// @Description Create twitter post
// @Tags        Twitter
// @Accept      json
// @Produce     json
// @Param       req  query request.GetTwitterLeaderboardRequest true "Create twitter post request"
// @Success     200 {object} response.GetTwitterLeaderboardResponse
// @Router      /twitter/top [get]
func (h *Handler) GetTwitterLeaderboard(c *gin.Context) {
	req := request.GetTwitterLeaderboardRequest{}
	if err := c.ShouldBindQuery(&req); err != nil {
		h.log.Error(err, "[handler.GetTwitterLeaderboard] BindJSON() failed")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	list, total, err := h.entities.GetTwitterLeaderboard(req)
	if err != nil {
		h.log.Error(err, "[handler.GetTwitterLeaderboard] entity.GetTwitterLeaderboard() failed")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	paging := &response.PaginationResponse{
		Pagination: model.Pagination{Page: req.Page, Size: req.Size},
		Total:      total,
	}
	c.JSON(http.StatusOK, response.CreateResponse(list, paging, nil, nil))
}
