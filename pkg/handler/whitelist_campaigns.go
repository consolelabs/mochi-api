package handler

import (
	"net/http"

	"bytes"
	"encoding/csv"
	"github.com/defipod/mochi/pkg/entities"
	"github.com/defipod/mochi/pkg/request"
	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateWhitelistCampaign(c *gin.Context) {
	body := request.CreateWhitelistCampaignRequest{}

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.entities.CreateWhitelistCampaign(body); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, body)
}

func (h *Handler) GetWhitelistCampaigns(c *gin.Context) {
	guildId := c.Query("guild_id")
	if guildId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "guild_id is required"})
		return
	}

	campaigns, err := h.entities.GetWhitelistCampaignsByGuildId(guildId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, campaigns)
}

func (h *Handler) GetWhitelistCampaignById(c *gin.Context) {
	campaignId := c.Param("campaignId")
	if campaignId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "campaign_id is required"})
		return
	}

	campaign, err := h.entities.GetWhitelistCampaign(campaignId)
	if err != nil {
		if err == entities.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, campaign)
}

func (h *Handler) AddWhitelistCampaignUsers(c *gin.Context) {
	var body request.AddWhitelistCampaignUserRequest

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.entities.AddWhitelistCampaignUsers(body); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, body)
}

func (h *Handler) GetWhitelistCampaignUsers(c *gin.Context) {
	campaignId := c.Query("campaign_id")
	if campaignId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "campaign_id is required"})
		return
	}

	wlUsers, err := h.entities.GetWhitelistCampaignUsers(campaignId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, wlUsers)
}

func (h *Handler) GetWhitelistCampaignUserByDiscordId(c *gin.Context) {
	discordId := c.Param("discordId")
	if discordId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "discord_id is required"})
		return
	}

	campaignId := c.Query("campaign_id")
	if campaignId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "campaign_id is required"})
		return
	}

	wlUsers, err := h.entities.GetWhitelistCampaignUser(discordId, campaignId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, wlUsers)
}

func (h *Handler) GetWhitelistCampaignUsersCSV(c *gin.Context) {
	campaignId := c.Query("campaign_id")
	if campaignId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "campaign_id is required"})
		return
	}
	wlUsers, err := h.entities.GetWhitelistCampaignUsers(campaignId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	b := &bytes.Buffer{}
	w := csv.NewWriter(b)

	if err = w.Write([]string{"address", "discord_id", "notes", "whitelist_campaign_id", "created_at"}); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	for _, wl := range wlUsers {
		if err = w.Write([]string{wl.Address, wl.DiscordID, wl.Notes, wl.WhitelistCampaignId, wl.CreatedAt.String()}); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	w.Flush()
	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Disposition", "attachment; filename=wlusers.csv")
	c.Data(http.StatusOK, "text/csv", b.Bytes())
}
