package handler

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/defipod/mochi/pkg/entities"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/request"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetGuilds(c *gin.Context) {
	guilds, err := h.entities.GetGuilds()
	if err != nil {
		h.log.Error(err, "[handler.GetGuilds] - failed to get all guilds")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, guilds)
}

func (h *Handler) GetGuild(c *gin.Context) {
	guildID := c.Param("guild_id")

	guild, err := h.entities.GetGuild(guildID)
	if err != nil {
		if err == entities.ErrRecordNotFound {
			h.log.Fields(logger.Fields{"guildID": guildID}).Error(err, "[handler.GetGuild] - guild not exist")
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		h.log.Fields(logger.Fields{"guildID": guildID}).Error(err, "[handler.GetGuild] - failed to get guild")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, guild)
}

func (h *Handler) CreateGuild(c *gin.Context) {
	body := request.CreateGuildRequest{}

	if err := c.BindJSON(&body); err != nil {
		h.log.Fields(logger.Fields{"body": body}).Error(err, "[handler.CreateGuild] - failed to read JSON")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.entities.CreateGuild(body); err != nil {
		h.log.Fields(logger.Fields{"body": body}).Error(err, "[handler.CreateGuild] - failed to creat guild")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, body)
}

func (h *Handler) GetGuildStatsHandler(c *gin.Context) {
	guildID := c.Param("guild_id")

	guildStat, err := h.entities.GetByGuildID(guildID)
	if err != nil {
		if err == entities.ErrRecordNotFound {
			h.log.Fields(logger.Fields{"guildID": guildID}).Error(err, "[handler.GetGuildStatsHandler] - guild not exist")
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		h.log.Fields(logger.Fields{"guildID": guildID}).Error(err, "[handler.GetGuildStatsHandler] - failed to get guild")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, guildStat)
}

func (h *Handler) CreateGuildChannel(c *gin.Context) {
	log := logger.NewLogrusLogger()
	guildID := c.Param("guild_id")
	countType := c.Query("count_type")

	log.Infof("Creating stats channel for counting. GuildId: %v, CountType: %v", guildID, countType)

	err := h.entities.CreateGuildChannel(guildID, countType)
	if err != nil {
		h.log.Fields(logger.Fields{"guildID": guildID, "countType": countType}).Error(err, "[handler.CreateGuildChannel] - failed to create guild channel")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, "ok")
}

func (h *Handler) ListMyGuilds(c *gin.Context) {
	accessToken := c.GetString("discord_access_token")

	guilds, err := h.entities.ListMyDiscordGuilds(accessToken)
	if err != nil {
		h.log.Fields(logger.Fields{"token": accessToken}).Error(err, "[handler.ListMyGuilds] - failed to list discord guilds")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": guilds})
}

func (h *Handler) UpdateGuild(c *gin.Context) {
	guildID := c.Param("guild_id")
	if guildID == "" {
		h.log.Info("[handler.UpdateGuild] - guild id empty")
		c.JSON(http.StatusBadRequest, gin.H{"error": "guild_id is required"})
		return
	}

	var req struct {
		GlobalXP   string `json:"global_xp"`
		LogChannel string `json:"log_channel"`
	}
	if err := c.BindJSON(&req); err != nil {
		h.log.Fields(logger.Fields{"guildID": guildID, "globalXP": req.GlobalXP, "logChannel": req.LogChannel}).Error(err, "[handler.UpdateGuild] - failed to read JSON")
		c.JSON(http.StatusBadRequest, gin.H{"error": "global_xp is required"})
		return
	}

	omit := "log_channel"
	if req.GlobalXP == "" {
		omit = "global_xp"
	}
	globalXP := strings.EqualFold(req.GlobalXP, "true")
	fmt.Println(omit, globalXP, req)

	if err := h.entities.UpdateGuild(omit, guildID, globalXP, req.LogChannel); err != nil {
		h.log.Fields(logger.Fields{"guildID": guildID, "globalXP": req.GlobalXP, "logChannel": req.LogChannel}).Error(err, "[handler.UpdateGuild] - failed to update guild")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "OK"})
}
