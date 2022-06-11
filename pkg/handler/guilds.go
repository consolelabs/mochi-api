package handler

import (
	"net/http"
	"strconv"

	"github.com/defipod/mochi/pkg/entities"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/request"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetGuilds(c *gin.Context) {
	guilds, err := h.entities.GetGuilds()
	if err != nil {
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
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, guild)
}

func (h *Handler) CreateGuild(c *gin.Context) {
	body := request.CreateGuildRequest{}

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.entities.CreateGuild(body); err != nil {
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
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, guildStat)
}

func (h *Handler) CreateGuildChannel(c *gin.Context) {
	log := logger.NewLogrusLogger()
	guildID := c.Param("guild_id")
	countType := c.Query("count_type")
	var coinData []string
	var err error

	log.Infof("Creating stats channel for counting. GuildId: %v, CountType: %v", guildID, countType)
	if countType == "highest_ticker" {
		symbol := c.Query("symbol")
		interval, _ := strconv.Atoi(c.Query("interval"))

		coinData, err = h.entities.GetHighestTicker(symbol, interval)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	}
	err = h.entities.CreateGuildChannel(guildID, countType, coinData...)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, "ok")
}

func (h *Handler) ListMyGuilds(c *gin.Context) {
	accessToken := c.GetString("discord_access_token")

	guilds, err := h.entities.ListMyDiscordGuilds(accessToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": guilds})
}

func (h *Handler) ToggleGlobalXP(c *gin.Context) {
	guildID := c.Param("guild_id")
	if guildID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "guild_id is required"})
		return
	}

	var req struct {
		GlobalXP bool `json:"global_xp"`
	}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "global_xp is required"})
		return
	}

	if err := h.entities.ToggleGuildGlobalXP(guildID, req.GlobalXP); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "OK"})
}
