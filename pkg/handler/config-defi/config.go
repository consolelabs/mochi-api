package configdefi

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/defipod/mochi/pkg/entities"
	"github.com/defipod/mochi/pkg/logger"
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

// UpsertMonikerConfig     godoc
// @Summary     Upsert moniker config
// @Description Upsert moniker config
// @Tags        ConfigDefi
// @Accept      json
// @Produce     json
// @Param       Request  body request.UpsertMonikerConfigRequest true "Upsert moniker config"
// @Success     200 {object} response.ResponseMessage
// @Router      /config-defi/monikers [post]
func (h *Handler) UpsertMonikerConfig(c *gin.Context) {
	var req request.UpsertMonikerConfigRequest
	if err := c.BindJSON(&req); err != nil {
		h.log.Fields(logger.Fields{"request": req}).Error(err, "[handler.UpsertMonikerConfig] - failed to read JSON")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("failed to read JSON"), nil))
		return
	}
	err := h.entities.UpsertMonikerConfig(req)
	if err != nil {
		h.log.Fields(logger.Fields{"request": req}).Error(err, "[handler.UpsertMonikerConfig] - failed to upsert moniker config")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	c.JSON(http.StatusOK, response.CreateResponse(response.ResponseMessage{Message: "OK"}, nil, nil, nil))
}

// GetMonikerByGuildID     godoc
// @Summary     Get moniker configs
// @Description Get moniker configs
// @Tags        ConfigDefi
// @Accept      json
// @Produce     json
// @Param       guild_id   path  string true  "Guild ID"
// @Success     200 {object} response.MonikerConfigResponse
// @Router      /config-defi/monikers/{guild_id} [get]
func (h *Handler) GetMonikerByGuildID(c *gin.Context) {
	guildID := c.Param("guild_id")
	if guildID == "" {
		h.log.Info("[handler.GetMonikerByGuildID] - guild id empty")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("guild_id is required"), nil))
		return
	}
	configs, err := h.entities.GetMonikerByGuildID(guildID)
	if err != nil {
		h.log.Fields(logger.Fields{"guildID": guildID}).Error(err, "[handler.GetMonikerByGuildID] - failed to guild moniker")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	c.JSON(http.StatusOK, response.CreateResponse(configs, nil, nil, nil))
}

// DeleteMonikerConfig     godoc
// @Summary     Delete moniker config
// @Description Delete moniker config
// @Tags        ConfigDefi
// @Accept      json
// @Produce     json
// @Param       Request  body request.DeleteMonikerConfigRequest true "Delete moinker config"
// @Success     200 {object} response.ResponseMessage
// @Router      /config-defi/monikers [delete]
func (h *Handler) DeleteMonikerConfig(c *gin.Context) {
	var req request.DeleteMonikerConfigRequest
	if err := c.BindJSON(&req); err != nil {
		h.log.Fields(logger.Fields{"request": req}).Error(err, "[handler.DeleteMonikerConfig] - failed to read JSON")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("failed to read JSON"), nil))
		return
	}
	err := h.entities.DeleteMonikerConfig(req)
	if err != nil {
		h.log.Fields(logger.Fields{"request": req}).Error(err, "[handler.DeleteMonikerConfig] - failed to delete moniker config")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	c.JSON(http.StatusOK, response.CreateResponse(response.ResponseMessage{Message: "OK"}, nil, nil, nil))
}

// GetDefaultMoniker     godoc
// @Summary     Get default moniker
// @Description Get default moniker
// @Tags        ConfigDefi
// @Accept      json
// @Produce     json
// @Success     200 {object} response.MonikerConfigResponse
// @Router      /config-defi/monikers/default [get]
func (h *Handler) GetDefaultMoniker(c *gin.Context) {
	configs, err := h.entities.GetDefaultMoniker()
	if err != nil {
		h.log.Error(err, "[handler.GetDefaultMoniker] - failed to get default moniker")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	c.JSON(http.StatusOK, response.CreateResponse(configs, nil, nil, nil))
}

// SetGuildDefaultTicker     godoc
// @Summary     Set guild default ticker
// @Description Set guild default ticker
// @Tags        ConfigDefi
// @Accept      json
// @Produce     json
// @Param       Request  body request.GuildConfigDefaultTickerRequest true "Set guild default ticker request"
// @Success     200 {object} response.ResponseMessage
// @Router      /config-defi/default-ticker [post]
func (h *Handler) SetGuildDefaultTicker(c *gin.Context) {
	req := request.GuildConfigDefaultTickerRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		h.log.Error(err, "[handler.SetGuildDefaultTicker] c.ShouldBindJSON failed")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	if err := h.entities.SetGuildDefaultTicker(req); err != nil {
		h.log.Error(err, "[handler.SetGuildDefaultTicker] entity.SetGuildDefaultTicker failed")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	c.JSON(http.StatusCreated, response.CreateResponse(response.ResponseMessage{Message: "OK"}, nil, nil, nil))
}

// GetGuildDefaultTicker     godoc
// @Summary     Get guild default ticker
// @Description Get guild default ticker
// @Tags        ConfigDefi
// @Accept      json
// @Produce     json
// @Param       guild_id   query  string true  "Guild ID"
// @Param       query   query  string true  "Guild ticker query"
// @Success     200 {object} response.GetGuildDefaultTickerResponse
// @Router      /config-defi/default-ticker [get]
func (h *Handler) GetGuildDefaultTicker(c *gin.Context) {
	var req request.GetGuildDefaultTickerRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		h.log.Error(err, "[handler.GetGuildDefaultTicker] ShouldBindQuery() failed")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	res, err := h.entities.GetGuildDefaultTicker(req)
	if err != nil {
		h.log.Error(err, "[handler.GetGuildDefaultTicker] entity.GetGuildDefaultTicker() failed")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	c.JSON(http.StatusOK, response.CreateResponse(res, nil, nil, nil))
}

// GetListGuildDefaultTicker     godoc
// @Summary     Get list default ticker of a guild.
// @Description Get list default ticker of a guild.
// @Tags        ConfigDefi
// @Accept      json
// @Produce     json
// @Param       guild_id   path  string true  "Guild ID"
// @Success     200 {object} response.GetListGuildDefaultTickerResponse
// @Router      /config-defi/default-ticker/{guild_id} [get]
func (h *Handler) GetListGuildDefaultTicker(c *gin.Context) {
	guildID := c.Param("guild_id")
	res, err := h.entities.GetListGuildDefaultTicker(guildID)
	if err != nil {
		h.log.Fields(logger.Fields{"guildID": guildID}).Error(err, "[handler.GetListGuildDefaultTicker] entity.GetListGuildDefaultTicker() failed")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	c.JSON(http.StatusOK, response.CreateResponse(res, nil, nil, nil))
}

// GetGuildtokens     godoc
// @Summary     Get guild tokens
// @Description Get guild tokens
// @Tags        ConfigDefi
// @Accept      json
// @Produce     json
// @Param       guild_id   query  string false  "Guild ID"
// @Success     200 {object} response.GetGuildTokensResponse
// @Router      /config-defi/tokens [get]
func (h *Handler) GetGuildTokens(c *gin.Context) {
	guildID := c.Query("guild_id")
	// if guild id empty, return global default tokens
	guildTokens, err := h.entities.GetGuildTokens(guildID)
	if err != nil {
		h.log.Fields(logger.Fields{"guildID": guildID}).Error(err, "[handler.GetGuildTokens] - failed to get guild tokens")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(guildTokens, nil, nil, nil))
}

// UpsertGuildTokenConfig     godoc
// @Summary     Update or insert guild token config
// @Description Update or insert guild token config
// @Tags        ConfigDefi
// @Accept      json
// @Produce     json
// @Param       Request  body request.UpsertGuildTokenConfigRequest true "Upsert Guild Token Config request"
// @Success     200 {object} response.ResponseMessage
// @Router      /config-defi/tokens [post]
func (h *Handler) UpsertGuildTokenConfig(c *gin.Context) {
	var req request.UpsertGuildTokenConfigRequest

	if err := c.BindJSON(&req); err != nil {
		h.log.Fields(logger.Fields{"guildID": req.GuildID, "symbol": req.Symbol, "active": req.Active}).Error(err, "[handler.UpsertGuildTokenConfig] - failed to read JSON")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	if req.GuildID == "" {
		h.log.Info("[handler.UpsertGuildTokenConfig] - guild id empty")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("guild_id is required"), nil))
		return
	}
	if req.Symbol == "" {
		h.log.Info("[handler.UpsertGuildTokenConfig] - symbol empty")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("symbol is required"), nil))
		return
	}

	if err := h.entities.UpsertGuildTokenConfig(req); err != nil {
		h.log.Fields(logger.Fields{"guildID": req.GuildID, "symbol": req.Symbol, "active": req.Active}).Error(err, "[handler.UpsertGuildTokenConfig] - failed to upsert guild token config")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(response.ResponseMessage{Message: "OK"}, nil, nil, nil))
}

// GetDefaultToken     godoc
// @Summary     Get Default token
// @Description Get Default token
// @Tags        ConfigDefi
// @Accept      json
// @Produce     json
// @Param       guild_id   query  string true  "Guild ID"
// @Success     200 {object} response.GetDefaultTokenResponse
// @Router      /config-defi/tokens/default [get]
func (h *Handler) GetDefaultToken(c *gin.Context) {
	guildID := c.Query("guild_id")
	token, err := h.entities.GetDefaultToken(guildID)
	if err != nil {
		h.log.Fields(logger.Fields{"guild_id": guildID}).Error(err, "[handler.ConfigDefaultToken] - failed to get default token")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(token, nil, nil, nil))
}

// ConfigDefaultToken     godoc
// @Summary     Config Default token
// @Description Config Default token
// @Tags        ConfigDefi
// @Accept      json
// @Produce     json
// @Param       Request  body request.ConfigDefaultTokenRequest true "Config default token request"
// @Success     200 {object} response.ResponseMessage
// @Router      /config-defi/tokens/default [post]
func (h *Handler) ConfigDefaultToken(c *gin.Context) {
	req := request.ConfigDefaultTokenRequest{}
	if err := c.BindJSON(&req); err != nil {
		h.log.Fields(logger.Fields{"req": req}).Error(err, "[handler.ConfigDefaultToken] - failed to read JSON")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	err := h.entities.SetDefaultToken(req)
	if err == baseerrs.ErrRecordNotFound {
		h.log.Fields(logger.Fields{"req": req}).Error(err, "[handler.ConfigDefaultToken] - failed to set default token")
		c.JSON(http.StatusNotFound, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	if err != nil {
		h.log.Fields(logger.Fields{"req": req}).Error(err, "[handler.ConfigDefaultToken] - failed to set default token")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(response.ResponseMessage{Message: "OK"}, nil, nil, nil))
}

// RemoveDefaultToken     godoc
// @Summary     Remove Default token
// @Description Remove Default token
// @Tags        ConfigDefi
// @Accept      json
// @Produce     json
// @Param       guild_id   query  string true  "Guild ID"
// @Success     200 {object} response.ResponseMessage
// @Router      /config-defi/tokens/default [delete]
func (h *Handler) RemoveDefaultToken(c *gin.Context) {
	guildID := c.Query("guild_id")
	if err := h.entities.RemoveDefaultToken(guildID); err != nil {
		h.log.Fields(logger.Fields{"guild_id": guildID}).Error(err, "[handler.RemoveDefaultToken] - failed to remove default token")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(response.ResponseMessage{Message: "OK"}, nil, nil, nil))
}

// CreateDefaultCollectionSymbol     godoc
// @Summary     Create default collection symbol
// @Description Create default collection symbol
// @Tags        ConfigDefi
// @Accept      json
// @Produce     json
// @Param       Request  body request.ConfigDefaultCollection true "Config Default Collection Symbol request"
// @Success     200 {object} response.ResponseMessage
// @Router      /config-defi/default-symbol [post]
func (h *Handler) CreateDefaultCollectionSymbol(c *gin.Context) {
	req := request.ConfigDefaultCollection{}
	if err := c.BindJSON(&req); err != nil {
		h.log.Fields(logger.Fields{"req": req}).Error(err, "[handler.CreateDefaultCollectionSymbol] - failed to read JSON")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	if err := h.entities.CreateDefaultCollectionSymbol(req); err != nil {
		h.log.Fields(logger.Fields{"req": req}).Error(err, "[handler.CreateDefaultCollectionSymbol] - failed to read JSON")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(response.ResponseMessage{Message: "OK"}, nil, nil, nil))
}

// ListAllCustomToken     godoc
// @Summary     List all guild custom token
// @Description List all guild custom token
// @Tags        Guild
// @Accept      json
// @Produce     json
// @Param       guild_id   path  string true  "Guild ID"
// @Success     200 {object} response.ListAllCustomTokenResponse
// @Router      /guilds/{guild_id}/custom-tokens [get]
func (h *Handler) ListAllCustomToken(c *gin.Context) {
	guildID := c.Param("guild_id")

	// get all token with guildID
	returnToken, err := h.entities.GetAllSupportedToken(guildID)
	if err != nil {
		h.log.Fields(logger.Fields{"guildID": guildID}).Error(err, "[handler.ListAllCustomToken] - failed to get all tokens")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response.ListAllCustomTokenResponse{Data: returnToken})
}

// HandlerGuildCustomTokenConfig     godoc
// @Summary     Guild custom token config
// @Description Guild custom token config
// @Tags        ConfigDefi
// @Accept      json
// @Produce     json
// @Param       Request  body request.UpsertCustomTokenConfigRequest true "Custom guild custom token config request"
// @Success     200 {object} response.ResponseMessage
// @Router      /config-defi/custom-tokens [post]
func (h *Handler) HandlerGuildCustomTokenConfig(c *gin.Context) {
	var req request.UpsertCustomTokenConfigRequest

	// handle input validate
	if err := c.BindJSON(&req); err != nil {
		h.log.Fields(logger.Fields{"body": req}).Error(err, "[handler.HandlerGuildCustomTokenConfig] - failed to read JSON")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	if req.GuildID == "" {
		h.log.Info("[handler.HandlerGuildCustomTokenConfig] - guild id empty")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("guild_id is required"), nil))
		return
	}
	if req.Symbol == "" {
		h.log.Info("[handler.HandlerGuildCustomTokenConfig] - symbol empty")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("symbol is required"), nil))
		return
	}
	if req.Address == "" {
		h.log.Info("[handler.HandlerGuildCustomTokenConfig] - address empty")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("address is required"), nil))
		return
	}
	if req.Chain == "" {
		h.log.Info("[handler.HandlerGuildCustomTokenConfig] - chain empty")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Chain is required"})
		return
	}

	if err := h.entities.CreateCustomToken(req); err != nil {
		h.log.Fields(logger.Fields{"req": req}).Error(err, "[handler.HandlerGuildCustomTokenConfig] - fail to create custom token")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(response.ResponseMessage{Message: "OK"}, nil, nil, nil))
}

// GetGuildDefaultCurrency     godoc
// @Summary     Get default currency by guild id
// @Description Get default currency by guild id
// @Tags        ConfigDefi
// @Accept      json
// @Produce     json
// @Param       guild_id   query  string true  "Guild ID"
// @Success     200 {object} response.GuildConfigDefaultCurrencyResponse
// @Router      /config-defi/default-currency [get]
func (h *Handler) GetGuildDefaultCurrency(c *gin.Context) {
	guildID := c.Query("guild_id")
	if guildID == "" {
		h.log.Info("[handler.GetGuildDefaultCurrency] - guild id empty")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("guild_id is required"), nil))
		return
	}

	data, err := h.entities.GetGuildDefaultCurrency(guildID)
	if err != nil {
		h.log.Fields(logger.Fields{"guildID": guildID}).Error(err, "[handler.GetGuildDefaultCurrency] - failed to get default currency")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(data, nil, nil, nil))
}

// UpsertGuildDefaultCurrency     godoc
// @Summary     Upsert default currency
// @Description Upsert default currency
// @Tags        ConfigDefi
// @Accept      json
// @Produce     json
// @Param       Request  body request.UpsertGuildDefaultCurrencyRequest true "Upsert default currency config"
// @Success     200 {object} response.ResponseMessage
// @Router      /config-defi/default-currency [post]
func (h *Handler) UpsertGuildDefaultCurrency(c *gin.Context) {
	var req request.UpsertGuildDefaultCurrencyRequest
	if err := c.BindJSON(&req); err != nil {
		h.log.Fields(logger.Fields{"request": req}).Error(err, "[handler.UpsertGuildDefaultCurrency] - failed to read JSON")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("failed to read JSON"), nil))
		return
	}

	err := h.entities.UpsertGuildDefaultCurrency(req)
	if err != nil {
		h.log.Fields(logger.Fields{"request": req}).Error(err, "[handler.UpsertGuildDefaultCurrency] - failed to upsert default currency")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(response.ResponseMessage{Message: "OK"}, nil, nil, nil))
}

// DeleteGuildDefaultCurrency     godoc
// @Summary     Delete default currency
// @Description Delete default currency
// @Tags        ConfigDefi
// @Accept      json
// @Produce     json
// @Param       Request  body request.GuildIDRequest true "Delete default currency config"
// @Success     200 {object} response.ResponseMessage
// @Router      /config-defi/default-currency [delete]
func (h *Handler) DeleteGuildDefaultCurrency(c *gin.Context) {
	guildID := c.Query("guild_id")
	if guildID == "" {
		h.log.Fields(logger.Fields{"request": guildID}).Error(nil, "[handler.DeleteGuildDefaultCurrency] - failed to read JSON")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("failed to read JSON"), nil))
		return
	}

	err := h.entities.DeleteGuildDefaultCurrency(guildID)
	if err != nil {
		h.log.Fields(logger.Fields{"request": guildID}).Error(err, "[handler.DeleteGuildDefaultCurrency] - failed to delete default currency")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(response.ResponseMessage{Message: "OK"}, nil, nil, nil))
}

// GetGuildConfigTipRangeByGuildId     godoc
// @Summary     Get config tip range by guild id
// @Description Get config tip range by guild id
// @Tags        ConfigDefi
// @Accept      json
// @Produce     json
// @Param       guild_id   path  string true  "Guild ID"
// @Success     200 {object} response.GuildConfigTipRangeResponse
// @Router      /config-defi/tip-range/{guild_id} [get]
func (h *Handler) GetGuildConfigTipRangeByGuildId(c *gin.Context) {
	guildID := c.Param("guild_id")
	if guildID == "" {
		h.log.Info("[handler.GetGuildConfigTipRangeByGuildId] - guild id empty")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("guild_id is required"), nil))
		return
	}

	data, err := h.entities.GetGuildConfigTipRange(guildID)
	if err != nil {
		h.log.Fields(logger.Fields{"guildID": guildID}).Error(err, "[handler.GetGuildConfigTipRangeByGuildId] - failed to get tip range config")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(data, nil, nil, nil))
}

// UpsertGuildConfigTipRange     godoc
// @Summary     Upsert config tip range
// @Description Upsert config tip range
// @Tags        ConfigDefi
// @Accept      json
// @Produce     json
// @Param       Request  body request.UpsertGuildConfigTipRangeRequest true "Upsert config tip range"
// @Success     200 {object} response.GuildConfigTipRangeResponse
// @Router      /config-defi/tip-range [post]
func (h *Handler) UpsertGuildConfigTipRange(c *gin.Context) {
	var req request.UpsertGuildConfigTipRangeRequest
	err := c.BindJSON(&req)
	if err != nil {
		h.log.Fields(logger.Fields{"request": req}).Error(err, "[handler.UpsertGuildConfigTipRange] - failed to read JSON")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("failed to read JSON"), nil))
		return
	}

	err = req.Validate()
	if err != nil {
		h.log.Fields(logger.Fields{"request": req}).Error(err, "[handler.UpsertGuildConfigTipRange] - failed to validate")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	resp, err := h.entities.UpsertGuildConfigTipRange(req)
	if err != nil {
		h.log.Fields(logger.Fields{"request": req}).Error(err, "[handler.UpsertGuildConfigTipRange] - failed to upsert tip range value")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](resp, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(resp, nil, nil, nil))
}

// RemoveGuildConfigTipRange     godoc
// @Summary     Remove config tip range
// @Description Remove config tip range
// @Tags        ConfigDefi
// @Accept      json
// @Produce     json
// @Param       guild_id   path  string true  "Guild ID"
// @Success     200 {object} response.ResponseMessage
// @Router      /config-defi/tip-range/{guild_id} [delete]
func (h *Handler) RemoveGuildConfigTipRange(c *gin.Context) {
	guildID := c.Param("guild_id")
	if guildID == "" {
		h.log.Info("[handler.RemoveGuildConfigTipRange] - guild id empty")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("guild_id is required"), nil))
		return
	}

	if err := h.entities.RemoveGuildConfigTipRange(guildID); err != nil {
		h.log.Fields(logger.Fields{"guild_id": guildID}).Error(err, "[handler.RemoveGuildConfigTipRange] - failed to remove tip range")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(response.ResponseMessage{Message: "OK"}, nil, nil, nil))
}
