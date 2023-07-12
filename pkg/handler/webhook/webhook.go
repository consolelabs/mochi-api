package webhook

import (
	"encoding/json"
	"net/http"

	"github.com/bwmarrin/discordgo"
	"github.com/gin-gonic/gin"

	"github.com/defipod/mochi/pkg/consts"
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

func (h *Handler) HandleDiscordWebhook(c *gin.Context) {
	var req request.HandleDiscordWebhookRequest
	if err := req.Bind(c); err != nil {
		h.log.Fields(logger.Fields{"body": req}).Error(err, "[handler.HandleDiscordWebhook] - failed to read JSON")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	if err := req.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	h.log.Infof("EVENT: %v", req.Event)
	switch req.Event {
	case request.GUILD_MEMBER_ADD:
		h.handleGuildMemberAdd(c, req.Data)
	case request.GUILD_MEMBER_REMOVE:
		h.handleGuildMemberRemove(c, req.Data)
	case request.MESSAGE_CREATE:
		h.handleMessageCreate(c, req)
	case request.MESSAGE_DELETE:
		h.handleMessageDelete(c, req.Data)
	case request.GUILD_CREATE:
		h.handleGuildCreate(c, req.Data)
	case request.MESSAGE_REACTION_ADD:
		h.handleMessageReactionAdd(c, req.Data)
	case request.MESSAGE_REACTION_REMOVE:
		h.handleMessageReactionRemove(c, req.Data)
	case request.GUILD_DELETE:
		h.handleGuildDelete(c, req.Data)
	}

	h.handleAutoTrigger(req.Event, c, req.Data)
}

func (h *Handler) handleGuildMemberAdd(c *gin.Context, data json.RawMessage) {
	var member discordgo.Member
	byteData, err := data.MarshalJSON()
	if err != nil {
		h.log.Info("[handler.handleGuildMemberAdd] - failed to json marshal data")
		// c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		c.JSON(http.StatusOK, response.CreateResponse(response.ResponseMessage{Message: "ok"}, nil, nil, nil))
		return
	}

	if err := discordgo.Unmarshal(byteData, &member); err != nil {
		h.log.Info("[handler.handleGuildMemberAdd] - failed to unmarshal data")
		// c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		c.JSON(http.StatusOK, response.CreateResponse(response.ResponseMessage{Message: "ok"}, nil, nil, nil))
		return
	}
}

func (h *Handler) handleGuildMemberRemove(c *gin.Context, data json.RawMessage) {
	var req request.MemberRemoveWebhookRequest
	byteData, err := data.MarshalJSON()
	if err != nil {
		h.log.Info("[handler.handleGuildMemberRemove] - failed to json marshal data")
		// c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		c.JSON(http.StatusOK, response.CreateResponse(response.ResponseMessage{Message: "ok"}, nil, nil, nil))
		return
	}

	if err := discordgo.Unmarshal(byteData, &req); err != nil {
		h.log.Info("[handler.handleGuildMemberRemove] - failed to unmarshal data")
		// c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		c.JSON(http.StatusOK, response.CreateResponse(response.ResponseMessage{Message: "ok"}, nil, nil, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(response.ResponseMessage{Message: "ok"}, nil, nil, nil))
}

func (h *Handler) handleMessageCreate(c *gin.Context, req request.HandleDiscordWebhookRequest) {
	message := &discordgo.Message{}
	byteData, err := req.Data.MarshalJSON()
	if err != nil {
		h.log.Error(err, "[handler.handleMessageCreate] - failed to json marshal data")
		// c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		c.JSON(http.StatusOK, response.CreateResponse(response.ResponseMessage{Message: "ok"}, nil, nil, nil))
		return
	}

	if err := discordgo.Unmarshal(byteData, &message); err != nil {
		h.log.Error(err, "[handler.handleMessageCreate] - failed to unmarshal data")
		// c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		c.JSON(http.StatusOK, response.CreateResponse(response.ResponseMessage{Message: "ok"}, nil, nil, nil))
		return
	}

	err = h.entities.CreateGuildIfNotExists(message.GuildID)
	if err != nil {
		h.log.Fields(logger.Fields{"message": message}).Error(err, "[handler.handleMessageCreate] entity.CreateGuildIfNotExists() failed")
	}

	_, err = h.entities.HandleDiscordMessage(message, req.ProfileID)
	if err != nil {
		h.log.Fields(logger.Fields{"message": message}).Error(err, "[handler.handleMessageCreate] - failed to handle discord message")
		// c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		c.JSON(http.StatusOK, response.CreateResponse(response.ResponseMessage{Message: "ok"}, nil, nil, nil))
		return
	}

	// TODO: use response data to send discord message to user
	var resp *response.HandleUserActivityResponse
	switch message.Type {
	case consts.MessageTypeUserPremiumGuildSubscription:
		resp, err = h.entities.BoostXPIncrease(message, req.ProfileID)
	default:
		resp, err = h.entities.ChatXPIncrease(message, req.ProfileID)
	}
	if err != nil {
		h.log.Fields(logger.Fields{"message": message}).Error(err, "[handler.handleMessageCreate] - failed to handle user activity")
		// c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		c.JSON(http.StatusOK, response.CreateResponse(response.ResponseMessage{Message: "ok"}, nil, nil, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(resp, nil, nil, nil))
}

func (h *Handler) handleGuildCreate(c *gin.Context, data json.RawMessage) {
	var req struct {
		GuildID string `json:"guild_id"`
	}

	byteData, err := data.MarshalJSON()
	if err != nil {
		h.log.Error(err, "[handler.handleGuildCreate] data.MarshalJSON() failed")
		// c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		c.JSON(http.StatusOK, response.CreateResponse(response.ResponseMessage{Message: "ok"}, nil, nil, nil))
		return
	}

	if err := discordgo.Unmarshal(byteData, &req); err != nil {
		h.log.Error(err, "[handler.handleGuildCreate] discordgo.Unmarshal() failed")
		// c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		c.JSON(http.StatusOK, response.CreateResponse(response.ResponseMessage{Message: "ok"}, nil, nil, nil))
		return
	}

	err = h.entities.HandleGuildCreate(req.GuildID)
	if err != nil {
		h.log.Error(err, "[handler.handleGuildCreate] entity.HandleGuildCreate() failed")
	}

	c.JSON(http.StatusOK, response.CreateResponse(response.ResponseMessage{Message: "OK"}, nil, nil, nil))
}

func (h *Handler) handleMessageReactionAdd(c *gin.Context, data json.RawMessage) {
	var req request.MessageReactionRequest
	byteData, err := data.MarshalJSON()
	if err != nil {
		h.log.Error(err, "[handler.handleMessageReactionAdd] - failed to json marshal data")
		// c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		c.JSON(http.StatusOK, response.CreateResponse(response.ResponseMessage{Message: "ok"}, nil, nil, nil))
		return
	}

	if err := discordgo.Unmarshal(byteData, &req); err != nil {
		h.log.Error(err, "[handler.handleMessageReactionAdd] - failed to unmarshal data")
		// c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		c.JSON(http.StatusOK, response.CreateResponse(response.ResponseMessage{Message: "ok"}, nil, nil, nil))
		return
	}

	if err := h.entities.AddMessageReaction(req); err != nil {
		h.log.Fields(logger.Fields{"body": req}).Error(err, "[handler.handleMessageReactionAdd] - failed to create message reaction")
		// c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		c.JSON(http.StatusOK, response.CreateResponse(response.ResponseMessage{Message: "ok"}, nil, nil, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(response.RepostReactionEventData{
		Status: "OK",
	}, nil, nil, nil))
}

func (h *Handler) handleMessageReactionRemove(c *gin.Context, data json.RawMessage) {
	var req request.MessageReactionRequest
	byteData, err := data.MarshalJSON()
	if err != nil {
		h.log.Error(err, "[handler.handleMessageReactionRemove] - failed to json marshal data")
		// c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		c.JSON(http.StatusOK, response.CreateResponse(response.ResponseMessage{Message: "ok"}, nil, nil, nil))
		return
	}

	if err := discordgo.Unmarshal(byteData, &req); err != nil {
		h.log.Error(err, "[handler.handleMessageReactionRemove] - failed to unmarshal data")
		// c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		c.JSON(http.StatusOK, response.CreateResponse(response.ResponseMessage{Message: "ok"}, nil, nil, nil))
		return
	}

	if err := h.entities.RemoveMessageReaction(req); err != nil {
		h.log.Fields(logger.Fields{"body": req}).Error(err, "[handler.handleMessageReactionRemove] - failed to create message reaction")
		// c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		c.JSON(http.StatusOK, response.CreateResponse(response.ResponseMessage{Message: "ok"}, nil, nil, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(response.RepostReactionEventData{}, nil, nil, nil))
}

func (h *Handler) handleMessageDelete(c *gin.Context, data json.RawMessage) {
	message := &discordgo.Message{}
	byteData, err := data.MarshalJSON()
	if err != nil {
		h.log.Error(err, "[handler.handleMessageDelete] - failed to json marshal data")
		// c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		c.JSON(http.StatusOK, response.CreateResponse(response.ResponseMessage{Message: "ok"}, nil, nil, nil))
		return
	}

	if err := discordgo.Unmarshal(byteData, &message); err != nil {
		h.log.Error(err, "[handler.handleMessageDelete] - failed to unmarshal data")
		// c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		c.JSON(http.StatusOK, response.CreateResponse(response.ResponseMessage{Message: "ok"}, nil, nil, nil))
		return
	}

	if err = h.entities.RemoveAllMessageReactions(message); err != nil {
		h.log.Fields(logger.Fields{"message": message}).Error(err, "[handler.handleMessageDelete] - failed to handle message delete")
		// c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		c.JSON(http.StatusOK, response.CreateResponse(response.ResponseMessage{Message: "ok"}, nil, nil, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(response.ResponseMessage{Message: "OK"}, nil, nil, nil))
}

func (h *Handler) handleGuildDelete(c *gin.Context, data json.RawMessage) {
	var req request.HandleGuildDeleteRequest
	byteData, err := data.MarshalJSON()
	if err != nil {
		h.log.Error(err, "[handler.handleGuildDelete] - data.MarshalJSON() failed")
		// c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		c.JSON(http.StatusOK, response.CreateResponse(response.ResponseMessage{Message: "ok"}, nil, nil, nil))
		return
	}
	if err := discordgo.Unmarshal(byteData, &req); err != nil {
		h.log.Error(err, "[handler.handleGuildDelete] - discordgo.Unmarshal() failed")
		// c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		c.JSON(http.StatusOK, response.CreateResponse(response.ResponseMessage{Message: "ok"}, nil, nil, nil))
		return
	}
	err = h.entities.DeactivateGuild(req)
	if err != nil {
		// code := http.StatusInternalServerError
		// if err == baseerrs.ErrRecordNotFound {
		// 	code = http.StatusNotFound
		// }
		// c.JSON(code, response.CreateResponse[any](nil, nil, err, nil))
		c.JSON(http.StatusOK, response.CreateResponse(response.ResponseMessage{Message: "ok"}, nil, nil, nil))
		return
	}
	c.JSON(http.StatusOK, response.CreateResponse(response.ResponseMessage{Message: "OK"}, nil, nil, nil))
}

func (h *Handler) WebhookNftHandler(c *gin.Context) {
	var req request.NotifySaleMarketplaceRequest
	if err := c.Bind(&req); err != nil {
		h.log.Fields(logger.Fields{"body": req}).Error(err, "[handler.WebhookNftSaleHandler] - failed to read JSON")
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	switch req.Event {
	case "sales":
		h.handleNftSales(c, req)
	}
}

func (h *Handler) handleNftSales(c *gin.Context, req request.NotifySaleMarketplaceRequest) {
	err := h.entities.NotifySaleMarketplace(req)
	if err != nil {
		h.log.Fields(logger.Fields{"body": req}).Error(err, "[handler.handleNftSales] - failed to send NFT sales to channel")
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}

func (h *Handler) NotifyNftCollectionIntegration(c *gin.Context) {
	req := request.NotifyCompleteNftIntegrationRequest{}
	if err := c.BindJSON(&req); err != nil {
		h.log.Error(err, "[handler.NotifyNftCollectionIntegration] c.BindJSON() failed")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	err := h.entities.NotifyNftCollectionIntegration(req)
	if err != nil {
		h.log.Error(err, "[handler.NotifyNftCollectionIntegration] entity.NotifyNftCollectionIntegration() failed")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	c.JSON(http.StatusOK, response.ResponseMessage{Message: "ok"})
}

func (h *Handler) NotifyNftCollectionAdd(c *gin.Context) {
	req := request.NotifyNftCollectionAddRequest{}
	if err := c.BindJSON(&req); err != nil {
		h.log.Error(err, "[handler.NotifyNftCollectionIntegration] c.BindJSON() failed")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	err := h.entities.NotifyNftCollectionAdd(req)
	if err != nil {
		h.log.Error(err, "[handler.NotifyNftCollectionAdd] entity.NotifyNftCollectionAdd() failed")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	c.JSON(http.StatusOK, response.ResponseMessage{Message: "ok"})
}

func (h *Handler) NotifyNftCollectionSync(c *gin.Context) {
	req := request.NotifyCompleteNftSyncRequest{}
	if err := c.BindJSON(&req); err != nil {
		h.log.Error(err, "[handler.NotifyNftCollectionSync] c.BindJSON() failed")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	err := h.entities.NotifyNftCollectionSync(req)
	if err != nil {
		h.log.Error(err, "[handler.NotifyNftCollectionSync] entity.NotifyNftCollectionSync() failed")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	c.JSON(http.StatusOK, response.ResponseMessage{Message: "ok"})
}

func (h *Handler) NotifySaleMarketplace(c *gin.Context) {
	var req request.NotifySaleMarketplaceRequest
	if err := c.Bind(&req); err != nil {
		h.log.Fields(logger.Fields{"body": req}).Error(err, "[handler.WebhookNftSaleHandler] - failed to read JSON")
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	err := h.entities.NotifySaleMarketplace(req)
	if err != nil {
		h.log.Fields(logger.Fields{"body": req}).Error(err, "[handler.handleNftSales] - failed to send NFT sales to channel")
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}

func (h *Handler) handleAutoTrigger(e string, c *gin.Context, data json.RawMessage) {
	h.log.Info("[handler.handleAutoTrigger] - handling auto trigger")
	var req request.AutoTriggerRequest
	byteData, err := data.MarshalJSON()
	if err != nil {
		h.log.Error(err, "[handler.handleMessageReactionAdd] - failed to json marshal data")
		// c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		c.JSON(http.StatusOK, response.CreateResponse(response.ResponseMessage{Message: "ok"}, nil, nil, nil))
		return
	}

	if err := discordgo.Unmarshal(byteData, &req); err != nil {
		h.log.Error(err, "[handler.handleMessageReactionAdd] - failed to unmarshal data")
		// c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		c.JSON(http.StatusOK, response.CreateResponse(response.ResponseMessage{Message: "ok"}, nil, nil, nil))
		return
	}

	err = h.entities.HandleTrigger(req)
	if err != nil {
		h.log.Error(err, "[handler.handleMessageReactionAdd] - failed to handle trigger")
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}
