package handler

import (
	"net/http"
	"strconv"

	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/request"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *Handler) GetGmConfig(c *gin.Context) {
	guildID := c.Query("guild_id")
	if guildID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "guild_id is required"})
	}

	config, err := h.entities.GetGmConfig(guildID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "OK", "data": config})
}

func (h *Handler) UpsertGmConfig(c *gin.Context) {
	var req request.UpsertGmConfigRequest

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if req.GuildID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "guild_id is required"})
	}
	if req.ChannelID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "channel_id is required"})
	}

	if err := h.entities.UpsertGmConfig(req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "OK"})
}

<<<<<<< HEAD
<<<<<<< HEAD
func (h *Handler) GetSalesTrackerConfig(c *gin.Context) {
=======
func (h *Handler) GetSTConfig(c *gin.Context) {
>>>>>>> 8098244 (feat: sales tracker config)
=======
func (h *Handler) GetSalesTrackerConfig(c *gin.Context) {
>>>>>>> 4b36907 (feat: sales tracker config)
	guildID := c.Query("guild_id")
	if guildID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "guild_id is required"})
	}

<<<<<<< HEAD
<<<<<<< HEAD
	config, err := h.entities.GetSalesTrackerConfig(guildID)
=======
	config, err := h.entities.GetSTConfig(guildID)
>>>>>>> 8098244 (feat: sales tracker config)
=======
	config, err := h.entities.GetSalesTrackerConfig(guildID)
>>>>>>> 4b36907 (feat: sales tracker config)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "OK", "data": config})
}

func (h *Handler) UpsertSalesTrackerConfig(c *gin.Context) {
<<<<<<< HEAD
<<<<<<< HEAD
	var req request.UpsertSalesTrackerConfigRequest
=======
	var req request.UpsertSTConfigRequest
>>>>>>> 8098244 (feat: sales tracker config)
=======
	var req request.UpsertSalesTrackerConfigRequest
>>>>>>> 665a93b (feat: sales tracker config)

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if req.GuildID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "guild_id is required"})
	}
	if req.ChannelID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "channel_id is required"})
	}

<<<<<<< HEAD
<<<<<<< HEAD
	if err := h.entities.UpsertSalesTrackerConfig(req); err != nil {
=======
	if err := h.entities.UpsertSTConfig(req); err != nil {
>>>>>>> 8098244 (feat: sales tracker config)
=======
	if err := h.entities.UpsertSalesTrackerConfig(req); err != nil {
>>>>>>> 4b36907 (feat: sales tracker config)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "OK"})
}

func (h *Handler) GetGuildTokens(c *gin.Context) {
	guildID := c.Query("guild_id")
	if guildID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "guild_id is required"})
	}

	guildTokens, err := h.entities.GetGuildTokens(guildID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var data []model.Token
	for _, gToken := range guildTokens {
		data = append(data, *gToken.Token)
	}

	c.JSON(http.StatusOK, gin.H{"data": data})
}

func (h *Handler) UpsertGuildTokenConfig(c *gin.Context) {
	var req request.UpsertGuildTokenConfigRequest

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if req.GuildID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "guild_id is required"})
		return
	}
	if req.Symbol == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "symbol is required"})
		return
	}

	if err := h.entities.UpsertGuildTokenConfig(req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "OK"})
}

func (h *Handler) ConfigLevelRole(c *gin.Context) {
	var req request.ConfigLevelRoleRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if req.GuildID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "guild_id is required"})
		return
	}
	if req.RoleID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "role_id is required"})
		return
	}
	if req.Level == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid level"})
		return
	}

	if err := h.entities.ConfigLevelRole(req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "OK"})
}

func (h *Handler) GetLevelRoleConfigs(c *gin.Context) {
	guildID := c.Param("guild_id")
	if guildID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "guild_id is required"})
		return
	}

	data, err := h.entities.GetGuildLevelRoleConfigs(guildID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": data})
}

func (h *Handler) RemoveLevelRoleConfig(c *gin.Context) {
	guildID := c.Param("guild_id")
	if guildID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "guild_id is required"})
		return
	}

	level := c.Query("level")
	if level == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "level is required"})
		return
	}

	levelNr, err := strconv.Atoi(level)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid level"})
		return
	}

	if err := h.entities.RemoveGuildLevelRoleConfig(guildID, levelNr); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "OK"})
}

func (h *Handler) ListGuildNFTRoles(c *gin.Context) {
	guildID := c.Query("guild_id")
	if guildID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "guild_id is required"})
		return
	}

	roles, err := h.entities.ListGuildNFTRoles(guildID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": roles})
}

func (h *Handler) NewGuildNFTRole(c *gin.Context) {
	var req request.ConfigNFTRoleRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := req.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newRole, err := h.entities.NewGuildNFTRoleConfig(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "OK", "data": newRole})
}

func (h *Handler) EditGuildNFTRole(c *gin.Context) {

	var req request.ConfigNFTRoleRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := req.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := uuid.Parse(c.Param("config_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid config_id"})
		return
	}

	req.ID = uuid.NullUUID{UUID: id, Valid: true}

	config, err := h.entities.EditGuildNFTRoleConfig(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "OK", "data": config})
}

func (h *Handler) RemoveGuildNFTRole(c *gin.Context) {

	configID := c.Param("config_id")
	if configID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "config_id is required"})
		return
	}

	if err := h.entities.RemoveGuildNFTRoleConfig(configID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "OK"})
}

func (h *Handler) ConfigRepostReaction(c *gin.Context) {
	var req request.ConfigRepostRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if req.GuildID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "guild_id is required"})
		return
	}
	if req.Emoji == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "emoji is required"})
		return
	}
	if req.Quantity < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "quantity is required"})
		return
	}
	if req.RepostChannelID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "repost_channel_id is required"})
		return
	}

	if err := h.entities.ConfigRepostReaction(req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "OK"})
}

func (h *Handler) GetRepostReactionConfigs(c *gin.Context) {
	guildID := c.Param("guild_id")
	if guildID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "guild_id is required"})
		return
	}

	data, err := h.entities.GetGuildRepostReactionConfigs(guildID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": data})
}

func (h *Handler) RemoveRepostReactionConfig(c *gin.Context) {
	var req request.ConfigRepostRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if req.GuildID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "guild_id is required"})
		return
	}
	if req.Emoji == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "emoji is required"})
		return
	}

	if err := h.entities.RemoveGuildRepostReactionConfig(req.GuildID, req.Emoji); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "OK"})
}

func (h *Handler) ToggleActivityConfig(c *gin.Context) {

	var (
		activityName = c.Param("activity")
		guildID      = c.Query("guild_id")
	)

	if activityName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "activity is required"})
		return
	}

	if guildID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "guild_id is required"})
		return
	}

	config, err := h.entities.ToggleActivityConfig(guildID, activityName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "OK", "data": config})
}
