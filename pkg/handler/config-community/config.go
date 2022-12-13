package configcommunity

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

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

// GetAllTwitterConfig     godoc
// @Summary     Get all twitter config
// @Description Get all twitter config
// @Tags        Config
// @Accept      json
// @Produce     json
// @Param       guild_id   query  string true  "Guild ID"
// @Success     200 {object} response.GetAllTwitterConfigResponse
// @Router      /configs/twitter [get]
func (h *Handler) GetAllTwitterConfig(c *gin.Context) {
	guildID := c.Query("guild_id")
	config, err := h.entities.GetAllTwitterConfig()
	if err != nil {
		h.log.Fields(logger.Fields{"guildID": guildID}).Error(err, "[handler.GetTwitterConfig] - failed to get twitter config")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	c.JSON(http.StatusOK, response.CreateResponse(config, nil, nil, nil))
}

// CreateTwitterConfig     godoc
// @Summary     Create twitter config
// @Description Create twitter config
// @Tags        Config
// @Accept      json
// @Produce     json
// @Param       Request  body model.GuildConfigTwitterFeed true "Create Twitter config request"
// @Success     200 {object} response.ResponseMessage
// @Router      /configs/twitter [post]
func (h *Handler) CreateTwitterConfig(c *gin.Context) {
	cfg := model.GuildConfigTwitterFeed{}
	err := c.BindJSON(&cfg)
	if err != nil {
		h.log.Fields(logger.Fields{"body": cfg}).Error(err, "[handler.CreateTwitterConfig] - failed to read JSON")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	err = h.entities.CreateTwitterConfig(&cfg)
	if err != nil {
		h.log.Fields(logger.Fields{"body": cfg}).Error(err, "[handler.GetTwitterConfig] - failed to create twitter config")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	c.JSON(http.StatusOK, response.CreateResponse(response.ResponseMessage{Message: "OK"}, nil, nil, nil))
}

// GetTwitterHashtagConfig     godoc
// @Summary     Get twitter hashtag config
// @Description get twitter hashtag config
// @Tags        Config
// @Accept      json
// @Produce     json
// @Param       guild_id   path  string true  "Guild ID"
// @Success     200 {object} response.GetTwitterHashtagConfigResponse
// @Router      /configs/twitter/hashtag/{guild_id} [get]
func (h *Handler) GetTwitterHashtagConfig(c *gin.Context) {
	guildId := c.Param("guild_id")
	hashtags, err := h.entities.GetTwitterHashtagConfig(guildId)
	if err != nil && err != gorm.ErrRecordNotFound {
		h.log.Fields(logger.Fields{"guild_id": guildId}).Error(err, "[handler.GetTwitterHashtagConfig] - failed to get hashtags")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	c.JSON(http.StatusOK, response.CreateResponse(hashtags, nil, nil, nil))
}

// GetAllTwitterHashtagConfig     godoc
// @Summary     Get all twitter hashtag config
// @Description get all twitter hashtag config
// @Tags        Config
// @Accept      json
// @Produce     json
// @Success     200 {object} response.GetAllTwitterHashtagConfigResponse
// @Router      /configs/twitter/hashtag [get]
func (h *Handler) GetAllTwitterHashtagConfig(c *gin.Context) {
	hashtags, err := h.entities.GetAllTwitterHashtagConfig()
	if err != nil && err != gorm.ErrRecordNotFound {
		h.log.Error(err, "[handler.GetTwitterHashtagConfig] - failed to get hashtags")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	c.JSON(http.StatusOK, response.CreateResponse(hashtags, nil, nil, nil))
}

// DeleteTwitterHashtagConfig     godoc
// @Summary     Delete twitter hashtag config
// @Description Delete twitter hashtag config
// @Tags        Config
// @Accept      json
// @Produce     json
// @Param       guild_id   path  string true  "Guild ID"
// @Success     200 {object} response.ResponseMessage
// @Router      /configs/twitter/hashtag/{guild_id} [delete]
func (h *Handler) DeleteTwitterHashtagConfig(c *gin.Context) {
	guildId := c.Param("guild_id")
	err := h.entities.DeleteTwitterHashtagConfig(guildId)
	if err != nil {
		h.log.Fields(logger.Fields{"guild_id": guildId}).Error(err, "[handler.GetTwitterHashtagConfig] - failed to delete")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	c.JSON(http.StatusOK, response.CreateResponse(response.ResponseMessage{Message: "OK"}, nil, nil, nil))
}

// CreateTwitterHashtagConfig     godoc
// @Summary     Create twitter hashtag config
// @Description Create twitter hashtag config
// @Tags        Config
// @Accept      json
// @Produce     json
// @Param       Request  body request.TwitterHashtag true "Create twitter hashtag config request"
// @Success     200 {object} response.ResponseMessage
// @Router      /configs/twitter/hashtag [post]
func (h *Handler) CreateTwitterHashtagConfig(c *gin.Context) {
	req := request.TwitterHashtag{}
	err := c.BindJSON(&req)
	if err != nil {
		h.log.Fields(logger.Fields{"body": req}).Error(err, "[handler.CreateTwitterHashtagConfig] - failed to read JSON")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	err = h.entities.CreateTwitterHashtagConfig(&req)
	if err != nil {
		h.log.Fields(logger.Fields{"body": req}).Error(err, "[handler.CreateTwitterHashtagConfig] - failed to create hashtag")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	c.JSON(http.StatusOK, response.CreateResponse(response.ResponseMessage{Message: "OK"}, nil, nil, nil))
}

// GetLinkedTelegram     godoc
// @Summary     Get telegram account linked with discord ID
// @Description Get telegram account linked with discord ID
// @Tags        Telegram
// @Accept      json
// @Produce     json
// @Param       telegram_username query string true "request"
// @Success     200 {object} response.GetLinkedTelegramResponse
// @Router      /configs/telegram [get]
func (h *Handler) GetLinkedTelegram(c *gin.Context) {
	telegramUsername := c.Query("telegram_username")
	if telegramUsername == "" {
		h.log.Info("[handler.GetLinkedTelegram] - telegram_username is required")
		c.JSON(http.StatusBadRequest, gin.H{"error": "telegram_username is required"})
		return
	}
	res, err := h.entities.GetByTelegramUsername(telegramUsername)
	if err != nil {
		h.log.Error(err, "[handler.GetLinkedTelegram] entity.GetByTelegramUsername() failed")
		code := http.StatusInternalServerError
		if err == baseerrs.ErrRecordNotFound {
			code = http.StatusNotFound
		}
		c.JSON(code, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// LinkUserTelegramWithDiscord     godoc
// @Summary     Link user's telegram with discord account
// @Description Link user's telegram with discord account
// @Tags        Telegram
// @Accept      json
// @Produce     json
// @Param       req body request.LinkUserTelegramWithDiscordRequest true "request"
// @Success     201 {object} response.LinkUserTelegramWithDiscordResponse
// @Router      /configs/telegram [post]
func (h *Handler) LinkUserTelegramWithDiscord(c *gin.Context) {
	var req request.LinkUserTelegramWithDiscordRequest
	if err := c.Bind(&req); err != nil {
		h.log.Error(err, "[handler.LinkUserTelegramWithDiscord] Bind() failed")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("discord_id and telegramusernam are required"), nil))
		return
	}
	res, err := h.entities.LinkUserTelegramWithDiscord(req)
	if err != nil {
		h.log.Error(err, "[handler.LinkUserTelegramWithDiscord] entity.LinkUserTelegramWithDiscord() failed")
		c.JSON(baseerrs.GetStatusCode(err), gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, res)
}

// AddToTwitterBlackList     godoc
// @Summary     Add an user to twitter watching blacklist
// @Description Add an user to twitter watching blacklist
// @Tags        Config
// @Accept      json
// @Produce     json
// @Param       req body request.AddToTwitterBlackListRequest true "request"
// @Success     200 {string} string "ok"
// @Router      /configs/twitter/blacklist [post]
func (h *Handler) AddToTwitterBlackList(c *gin.Context) {
	var req request.AddToTwitterBlackListRequest
	if err := c.BindJSON(&req); err != nil {
		h.log.Error(err, "[handler.AddToTwitterBlackList] BindJSON() failed")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	err := h.entities.AddToTwitterBlackList(req)
	if err != nil {
		h.log.Fields(logger.Fields{"req": req}).Error(err, "[handler.AddToTwitterBlackList] entity.AddToTwitterBlackList() failed")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	c.JSON(http.StatusOK, response.CreateResponse(response.ResponseMessage{Message: "OK"}, nil, nil, nil))
}

// DeleteFromTwitterBlackList     godoc
// @Summary     Delete an user from twitter watching blacklist
// @Description Delete an user from twitter watching blacklist
// @Tags        Config
// @Accept      json
// @Produce     json
// @Param       req query request.DeleteFromTwitterBlackListRequest true "query"
// @Success     200 {string} string "ok"
// @Router      /configs/twitter/blacklist [delete]
func (h *Handler) DeleteFromTwitterBlackList(c *gin.Context) {
	var req request.DeleteFromTwitterBlackListRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		h.log.Error(err, "[handler.DeleteFromTwitterBlackList] ShouldBindQuery() failed")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	err := h.entities.DeleteFromTwitterBlackList(req)
	if err != nil {
		h.log.Fields(logger.Fields{"req": req}).Error(err, "[handler.DeleteFromTwitterBlackList] entity.DeleteFromTwitterBlackList() failed")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	c.JSON(http.StatusOK, response.CreateResponse(response.ResponseMessage{Message: "OK"}, nil, nil, nil))
}

// GetTwitterBlackList     godoc
// @Summary     Get twitter blacklist
// @Description get twitter blacklist
// @Tags        Config
// @Accept      json
// @Produce     json
// @Param       guild_id   query  string false  "Guild ID"
// @Success     200 {object} response.GetTwitterBlackListResponse
// @Router      /configs/twitter/blacklist [get]
func (h *Handler) GetTwitterBlackList(c *gin.Context) {
	guildID := c.Query("guild_id")
	data, err := h.entities.GetTwitterBlackList(guildID)
	if err != nil {
		h.log.Error(err, "[handler.GetTwitterBlackList] entity.GetTwitterBlackList() failed")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	c.JSON(http.StatusOK, response.CreateResponse(data, nil, nil, nil))
}
