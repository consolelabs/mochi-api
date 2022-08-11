package handler

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/request"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *Handler) GetGmConfig(c *gin.Context) {
	guildID := c.Query("guild_id")
	if guildID == "" {
		h.log.Info("[handler.GetGmConfig] - guild id empty")
		c.JSON(http.StatusBadRequest, gin.H{"error": "guild_id is required"})
	}

	config, err := h.entities.GetGmConfig(guildID)

	if err != nil {
		h.log.Fields(logger.Fields{"guildID": guildID}).Error(err, "[handler.GetGmConfig] - failed to get gm config")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "OK", "data": config})
}

func (h *Handler) UpsertGmConfig(c *gin.Context) {
	var req request.UpsertGmConfigRequest

	if err := c.BindJSON(&req); err != nil {
		h.log.Fields(logger.Fields{"guildID": req.GuildID, "channelID": req.ChannelID}).Error(err, "[handler.UpsertGmConfig] - failed to read JSON")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if req.GuildID == "" {
		h.log.Info("[handler.UpsertGmConfig] - guild id empty")
		c.JSON(http.StatusBadRequest, gin.H{"error": "guild_id is required"})
	}
	if req.ChannelID == "" {
		h.log.Info("[handler.UpsertGmConfig] - channel id empty")
		c.JSON(http.StatusBadRequest, gin.H{"error": "channel_id is required"})
	}

	if err := h.entities.UpsertGmConfig(req); err != nil {
		h.log.Fields(logger.Fields{"guildID": req.GuildID, "channelID": req.ChannelID}).Error(err, "[handler.UpsertGmConfig] - failed to upsert gm config")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "OK"})
}

func (h *Handler) GetSalesTrackerConfig(c *gin.Context) {
	guildID := c.Query("guild_id")
	if guildID == "" {
		h.log.Info("[handler.GetSalesTrackerConfig] - guild id empty")
		c.JSON(http.StatusBadRequest, gin.H{"error": "guild_id is required"})
	}

	config, err := h.entities.GetSalesTrackerConfig(guildID)

	if err != nil {
		h.log.Fields(logger.Fields{"guildID": guildID}).Error(err, "[handler.GetSalesTrackerConfig] - failed to get sales tracker config")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "OK", "data": config})
}

func (h *Handler) GetGuildTokens(c *gin.Context) {
	guildID := c.Query("guild_id")
	if guildID == "" {
		h.log.Info("[handler.GetGuildTokens] - guild id empty")
		c.JSON(http.StatusBadRequest, gin.H{"error": "guild_id is required"})
	}

	guildTokens, err := h.entities.GetGuildTokens(guildID)
	if err != nil {
		h.log.Fields(logger.Fields{"guildID": guildID}).Error(err, "[handler.GetGuildTokens] - failed to get guild tokens")
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
		h.log.Fields(logger.Fields{"guildID": req.GuildID, "symbol": req.Symbol, "active": req.Active}).Error(err, "[handler.UpsertGuildTokenConfig] - failed to read JSON")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if req.GuildID == "" {
		h.log.Info("[handler.UpsertGuildTokenConfig] - guild id empty")
		c.JSON(http.StatusBadRequest, gin.H{"error": "guild_id is required"})
		return
	}
	if req.Symbol == "" {
		h.log.Info("[handler.UpsertGuildTokenConfig] - symbol empty")
		c.JSON(http.StatusBadRequest, gin.H{"error": "symbol is required"})
		return
	}

	if err := h.entities.UpsertGuildTokenConfig(req); err != nil {
		h.log.Fields(logger.Fields{"guildID": req.GuildID, "symbol": req.Symbol, "active": req.Active}).Error(err, "[handler.UpsertGuildTokenConfig] - failed to upsert guild token config")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "OK"})
}

func (h *Handler) ConfigLevelRole(c *gin.Context) {
	var req request.ConfigLevelRoleRequest
	if err := c.BindJSON(&req); err != nil {
		h.log.Fields(logger.Fields{"guildID": req.GuildID, "roleID": req.RoleID, "level": req.Level}).Error(err, "[handler.ConfigLevelRole] - failed to read JSON")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if req.GuildID == "" {
		h.log.Info("[handler.ConfigLevelRole] - guild id empty")
		c.JSON(http.StatusBadRequest, gin.H{"error": "guild_id is required"})
		return
	}
	if req.RoleID == "" {
		h.log.Info("[handler.ConfigLevelRole] - role id empty")
		c.JSON(http.StatusBadRequest, gin.H{"error": "role_id is required"})
		return
	}
	if req.Level == 0 {
		h.log.Info("[handler.ConfigLevelRole] - level empty")
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid level"})
		return
	}

	if err := h.entities.ConfigLevelRole(req); err != nil {
		h.log.Fields(logger.Fields{"guildID": req.GuildID, "roleID": req.RoleID, "level": req.Level}).Error(err, "[handler.ConfigLevelRole] - failed to config level role")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "OK"})
}

func (h *Handler) GetLevelRoleConfigs(c *gin.Context) {
	guildID := c.Param("guild_id")
	if guildID == "" {
		h.log.Info("[handler.GetLevelRoleConfigs] - guild id empty")
		c.JSON(http.StatusBadRequest, gin.H{"error": "guild_id is required"})
		return
	}

	data, err := h.entities.GetGuildLevelRoleConfigs(guildID)
	if err != nil {
		h.log.Fields(logger.Fields{"guildID": guildID}).Error(err, "[handler.GetLevelRoleConfigs] - failed to get guild level role config")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": data})
}

func (h *Handler) RemoveLevelRoleConfig(c *gin.Context) {
	guildID := c.Param("guild_id")
	if guildID == "" {
		h.log.Info("[handler.RemoveLevelRoleConfig] - guild id empty")
		c.JSON(http.StatusBadRequest, gin.H{"error": "guild_id is required"})
		return
	}

	level := c.Query("level")
	if level == "" {
		h.log.Info("[handler.RemoveLevelRoleConfig] - level empty")
		c.JSON(http.StatusBadRequest, gin.H{"error": "level is required"})
		return
	}

	levelNr, err := strconv.Atoi(level)
	if err != nil {
		h.log.Fields(logger.Fields{"guildID": guildID, "level": level}).Error(err, "[handler.RemoveLevelRoleConfig] - invalid level")
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid level"})
		return
	}

	if err := h.entities.RemoveGuildLevelRoleConfig(guildID, levelNr); err != nil {
		h.log.Fields(logger.Fields{"guildID": guildID, "level": level}).Error(err, "[handler.RemoveLevelRoleConfig] - failed to remove guild level role config")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "OK"})
}

func (h *Handler) ListGuildNFTRoles(c *gin.Context) {
	guildID := c.Query("guild_id")
	if guildID == "" {
		h.log.Info("[handler.ListGuildNFTRoles] - guild id empty")
		c.JSON(http.StatusBadRequest, gin.H{"error": "guild_id is required"})
		return
	}

	roles, err := h.entities.ListGuildNFTRoles(guildID)
	if err != nil {
		h.log.Fields(logger.Fields{"guildID": guildID}).Error(err, "[handler.ListGuildNFTRoles] - failed to list all nft roles")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": roles})
}

func (h *Handler) NewGuildNFTRole(c *gin.Context) {
	var req request.ConfigNFTRoleRequest
	if err := c.BindJSON(&req); err != nil {
		h.log.Fields(logger.Fields{"body": req}).Error(err, "[handler.NewGuildNFTRole] - failed to read JSON")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := req.Validate(); err != nil {
		h.log.Fields(logger.Fields{"body": req}).Error(err, "[handler.NewGuildNFTRole] - failed to validate request")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newRole, err := h.entities.NewGuildNFTRoleConfig(req)
	if err != nil {
		h.log.Fields(logger.Fields{"body": req}).Error(err, "[handler.NewGuildNFTRole] - failed to create nft role config")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "OK", "data": newRole})
}

func (h *Handler) EditGuildNFTRole(c *gin.Context) {

	var req request.ConfigNFTRoleRequest
	if err := c.BindJSON(&req); err != nil {
		h.log.Fields(logger.Fields{"body": req}).Error(err, "[handler.EditGuildNFTRole] - failed to read JSON")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := req.Validate(); err != nil {
		h.log.Fields(logger.Fields{"body": req}).Error(err, "[handler.EditGuildNFTRole] - failed to validate request")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := uuid.Parse(c.Param("config_id"))
	if err != nil {
		h.log.Fields(logger.Fields{"configID": c.Param("config_id")}).Error(err, "[handler.EditGuildNFTRole] - failed to read config ID")
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid config_id"})
		return
	}

	req.ID = uuid.NullUUID{UUID: id, Valid: true}

	config, err := h.entities.EditGuildNFTRoleConfig(req)
	if err != nil {
		h.log.Fields(logger.Fields{"body": req}).Error(err, "[handler.EditGuildNFTRole] - failed to edit nft role config")
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
		h.log.Fields(logger.Fields{"configID": configID}).Error(err, "[handler.RemoveGuildNFTRole] - failed to remove nft role config")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "OK"})
}

func (h *Handler) ConfigRepostReaction(c *gin.Context) {
	var req request.ConfigRepostRequest
	if err := c.BindJSON(&req); err != nil {
		h.log.Fields(logger.Fields{"guildID": req.GuildID, "emoji": req.Emoji, "quantity": req.Quantity, "channel": req.RepostChannelID}).Error(err, "[handler.ConfigRepostReaction] - failed to read JSON")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if req.GuildID == "" {
		h.log.Info("[handler.ConfigRepostReaction] - guild id empty")
		c.JSON(http.StatusBadRequest, gin.H{"error": "guild_id is required"})
		return
	}
	if req.Emoji == "" {
		h.log.Info("[handler.ConfigRepostReaction] - emoji empty")
		c.JSON(http.StatusBadRequest, gin.H{"error": "emoji is required"})
		return
	}
	if req.Quantity < 1 {
		h.log.Info("[handler.ConfigRepostReaction] - quantity empty")
		c.JSON(http.StatusBadRequest, gin.H{"error": "quantity is required"})
		return
	}
	if req.RepostChannelID == "" {
		h.log.Info("[handler.ConfigRepostReaction] - channel id empty")
		c.JSON(http.StatusBadRequest, gin.H{"error": "repost_channel_id is required"})
		return
	}

	if err := h.entities.ConfigRepostReaction(req); err != nil {
		h.log.Fields(logger.Fields{"guildID": req.GuildID, "emoji": req.Emoji, "quantity": req.Quantity, "channel": req.RepostChannelID}).Error(err, "[handler.ConfigRepostReaction] - failed to add config repost reaction")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "OK"})
}

func (h *Handler) GetRepostReactionConfigs(c *gin.Context) {
	guildID := c.Param("guild_id")
	if guildID == "" {
		h.log.Info("[handler.GetRepostReactionConfigs] - guild id empty")
		c.JSON(http.StatusBadRequest, gin.H{"error": "guild_id is required"})
		return
	}

	data, err := h.entities.GetGuildRepostReactionConfigs(guildID)
	if err != nil {
		h.log.Fields(logger.Fields{"guildID": guildID}).Error(err, "[handler.GetRepostReactionConfigs] - failed to get guild repost reaction config")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": data})
}

func (h *Handler) RemoveRepostReactionConfig(c *gin.Context) {
	var req request.ConfigRepostRequest
	if err := c.BindJSON(&req); err != nil {
		h.log.Fields(logger.Fields{"guildID": req.GuildID, "emoji": req.Emoji, "quantity": req.Quantity, "channel": req.RepostChannelID}).Error(err, "[handler.RemoveRepostReactionConfig] - failed to read JSON")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if req.GuildID == "" {
		h.log.Info("[handler.RemoveRepostReactionConfig] - guild id empty")
		c.JSON(http.StatusBadRequest, gin.H{"error": "guild_id is required"})
		return
	}
	if req.Emoji == "" {
		h.log.Info("[handler.RemoveRepostReactionConfig] - emoji empty")
		c.JSON(http.StatusBadRequest, gin.H{"error": "emoji is required"})
		return
	}

	if err := h.entities.RemoveGuildRepostReactionConfig(req.GuildID, req.Emoji); err != nil {
		h.log.Fields(logger.Fields{"guildID": req.GuildID, "emoji": req.Emoji, "quantity": req.Quantity, "channel": req.RepostChannelID}).Error(err, "[handler.RemoveRepostReactionConfig] - failed to remove repost reaction config")
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
		h.log.Info("[handler.ToggleActivityConfig] - activity name empty")
		c.JSON(http.StatusBadRequest, gin.H{"error": "activity is required"})
		return
	}

	if guildID == "" {
		h.log.Info("[handler.ToggleActivityConfig] - guild id empty")
		c.JSON(http.StatusBadRequest, gin.H{"error": "guild_id is required"})
		return
	}

	config, err := h.entities.ToggleActivityConfig(guildID, activityName)
	if err != nil {
		h.log.Fields(logger.Fields{"guildID": guildID, "activity": activityName}).Error(err, "[handler.ToggleActivityConfig] - failed to toggle activity config")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "OK", "data": config})
}

func (h *Handler) GetAllTwitterConfig(c *gin.Context) {
	guildID := c.Query("guild_id")
	config, err := h.entities.GetAllTwitterConfig()
	if err != nil {
		h.log.Fields(logger.Fields{"guildID": guildID}).Error(err, "[handler.GetTwitterConfig] - failed to get twitter config")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "OK", "data": config})
}

func (h *Handler) CreateTwitterConfig(c *gin.Context) {
	cfg := model.GuildConfigTwitterFeed{}
	err := c.BindJSON(&cfg)
	if err != nil {
		h.log.Fields(logger.Fields{"body": cfg}).Error(err, "[handler.CreateTwitterConfig] - failed to read JSON")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	err = h.entities.CreateTwitterConfig(&cfg)
	if err != nil {
		h.log.Fields(logger.Fields{"body": cfg}).Error(err, "[handler.GetTwitterConfig] - failed to create twitter config")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "OK"})
}

func (h *Handler) GetTwitterHashtagConfig(c *gin.Context) {
	guildId := c.Param("guild_id")
	hashtags, err := h.entities.GetTwitterHashtagConfig(guildId)
	if err != nil {
		if strings.Contains(err.Error(), "record not found") {
			h.log.Fields(logger.Fields{"guild_id": guildId}).Info("[handler.GetTwitterHashtagConfig] - hashtag config empty")
			c.JSON(http.StatusOK, gin.H{"data": hashtags})
			return
		}
		h.log.Fields(logger.Fields{"guild_id": guildId}).Error(err, "[handler.GetTwitterHashtagConfig] - failed to get hashtags")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": hashtags})
}

func (h *Handler) GetAllTwitterHashtagConfig(c *gin.Context) {
	hashtags, err := h.entities.GetAllTwitterHashtagConfig()
	if err != nil {
		if strings.Contains(err.Error(), "record not found") {
			h.log.Info("[handler.GetTwitterHashtagConfig] - hashtag config empty")
			c.JSON(http.StatusOK, gin.H{"data": hashtags})
			return
		}
		h.log.Error(err, "[handler.GetTwitterHashtagConfig] - failed to get hashtags")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": hashtags})
}

func (h *Handler) DeleteTwitterHashtagConfig(c *gin.Context) {
	guildId := c.Param("guild_id")
	err := h.entities.DeleteTwitterHashtagConfig(guildId)
	if err != nil {
		h.log.Fields(logger.Fields{"guild_id": guildId}).Error(err, "[handler.GetTwitterHashtagConfig] - failed to delete")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "OK"})
}

func (h *Handler) CreateTwitterHashtagConfig(c *gin.Context) {
	req := request.TwitterHashtag{}
	err := c.BindJSON(&req)
	if err != nil {
		h.log.Fields(logger.Fields{"body": req}).Error(err, "[handler.CreateTwitterHashtagConfig] - failed to read JSON")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.entities.CreateTwitterHashtagConfig(&req)
	if err != nil {
		h.log.Fields(logger.Fields{"body": req}).Error(err, "[handler.CreateTwitterHashtagConfig] - failed to create hashtag")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "OK"})
}

func (h *Handler) ConfigDefaultToken(c *gin.Context) {
	req := request.ConfigDefaultTokenRequest{}
	if err := c.BindJSON(&req); err != nil {
		h.log.Fields(logger.Fields{"req": req}).Error(err, "[handler.ConfigDefaultToken] - failed to read JSON")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.entities.SetDefaultToken(req); err != nil {
		h.log.Fields(logger.Fields{"req": req}).Error(err, "[handler.ConfigDefaultToken] - failed to set default token")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "OK"})
}
