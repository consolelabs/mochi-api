package handler

import (
	"net/http"

	"bytes"
	"encoding/csv"

	"github.com/defipod/mochi/pkg/entities"
	"github.com/defipod/mochi/pkg/logger"
	_ "github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/request"
	"github.com/gin-gonic/gin"
)

// CreateWhitelistCampaign     godoc
// @Summary     Create whitelist campaign
// @Description Create whitelist campaign
// @Tags        Whitelist campaign
// @Accept      json
// @Produce     json
// @Param       Request  body request.CreateWhitelistCampaignRequest true "Create whitelist campaign request"
// @Success     200 {object} request.CreateWhitelistCampaignRequest
// @Router      /whitelist-campaigns [post]
func (h *Handler) CreateWhitelistCampaign(c *gin.Context) {
	body := request.CreateWhitelistCampaignRequest{}

	if err := c.BindJSON(&body); err != nil {
		h.log.Fields(logger.Fields{"body": body}).Error(err, "[handler.CreateWhitelistCampaign] - failed to read JSON")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.entities.CreateWhitelistCampaign(body); err != nil {
		h.log.Fields(logger.Fields{"body": body}).Error(err, "[handler.CreateWhitelistCampaign] - failed to create whitelist campaign")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, body)
}

// GetWhitelistCampaign     godoc
// @Summary     Get whitelist campaign
// @Description Get whitelist campaign
// @Tags        Whitelist campaign
// @Accept      json
// @Produce     json
// @Param       guild_id   query  string true  "Guild ID"
// @Success     200 {object} []model.WhitelistCampaign
// @Router      /whitelist-campaigns [get]
func (h *Handler) GetWhitelistCampaigns(c *gin.Context) {
	guildId := c.Query("guild_id")
	if guildId == "" {
		h.log.Info("[handler.GetWhitelistCampaigns] - guild id empty")
		c.JSON(http.StatusBadRequest, gin.H{"error": "guild_id is required"})
		return
	}

	campaigns, err := h.entities.GetWhitelistCampaignsByGuildId(guildId)
	if err != nil {
		h.log.Fields(logger.Fields{"guildID": guildId}).Error(err, "[handler.GetWhitelistCampaigns] - failed to get whitelist campaign")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, campaigns)
}

// GetWhitelistCampaignById     godoc
// @Summary     Get whitelist campaign by id
// @Description Get whitelist campaign by id
// @Tags        Whitelist campaign
// @Accept      json
// @Produce     json
// @Param       campaignId   path  string true  "Campaign ID"
// @Success     200 {object} model.WhitelistCampaign
// @Router      /whitelist-campaigns/{campaignId} [get]
func (h *Handler) GetWhitelistCampaignById(c *gin.Context) {
	campaignId := c.Param("campaignId")
	if campaignId == "" {
		h.log.Info("[handler.GetWhitelistCampaignById] - campaign id empty")
		c.JSON(http.StatusBadRequest, gin.H{"error": "campaign_id is required"})
		return
	}

	campaign, err := h.entities.GetWhitelistCampaign(campaignId)
	if err != nil {
		if err == entities.ErrRecordNotFound {
			h.log.Fields(logger.Fields{"campaignID": campaignId}).Error(err, "[handler.GetWhitelistCampaignById] - campaign not exist")
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		h.log.Fields(logger.Fields{"campaignID": campaignId}).Error(err, "[handler.GetWhitelistCampaignById] - failed to get whitelist campaign")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, campaign)
}

// AddWhitelistCampaignUsers     godoc
// @Summary     Add whitelist campaign user
// @Description Add whitelist campaign user
// @Tags        Whitelist campaign
// @Accept      json
// @Produce     json
// @Param       Request  body request.AddWhitelistCampaignUserRequest true "Add whitelist campaign user request"
// @Success     200 {object} request.AddWhitelistCampaignUserRequest
// @Router      /whitelist-campaigns/users [post]
func (h *Handler) AddWhitelistCampaignUsers(c *gin.Context) {
	var body request.AddWhitelistCampaignUserRequest

	if err := c.BindJSON(&body); err != nil {
		h.log.Fields(logger.Fields{"body": body}).Error(err, "[handler.AddWhitelistCampaignUsers] - failed to read JSON")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.entities.AddWhitelistCampaignUsers(body); err != nil {
		h.log.Fields(logger.Fields{"body": body}).Error(err, "[handler.AddWhitelistCampaignUsers] - failed to add whitelist campaign users")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, body)
}

// GetWhitelistCampaignUsers     godoc
// @Summary     Get whitelist campaign user
// @Description Get whitelist campaign user
// @Tags        Whitelist campaign
// @Accept      json
// @Produce     json
// @Param       campaign_id   query  string true  "Campaign ID"
// @Success     200 {object} []model.WhitelistCampaignUser
// @Router      /whitelist-campaigns/users [get]
func (h *Handler) GetWhitelistCampaignUsers(c *gin.Context) {
	campaignId := c.Query("campaign_id")
	if campaignId == "" {
		h.log.Info("[handler.GetWhitelistCampaignUsers] - campaign id empty")
		c.JSON(http.StatusBadRequest, gin.H{"error": "campaign_id is required"})
		return
	}

	wlUsers, err := h.entities.GetWhitelistCampaignUsers(campaignId)
	if err != nil {
		h.log.Fields(logger.Fields{"campaignID": campaignId}).Error(err, "[handler.GetWhitelistCampaignUsers] - failed to get whitelist campaign users")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, wlUsers)
}

// GetWhitelistCampaignUserByDiscordId     godoc
// @Summary     Get whitelist campaign user by discord ID
// @Description Get whitelist campaign user by discord ID
// @Tags        Whitelist campaign
// @Accept      json
// @Produce     json
// @Param       discord_id   path  string true  "Discord ID"
// @Param       campaign_id   query  string true  "Campaign ID"
// @Success     200 {object} model.WhitelistCampaignUser
// @Router      /whitelist-campaigns/users/{discord_id} [get]
func (h *Handler) GetWhitelistCampaignUserByDiscordId(c *gin.Context) {
	discordId := c.Param("discordId")
	if discordId == "" {
		h.log.Info("[handler.GetWhitelistCampaignUserByDiscordId] - discord id empty")
		c.JSON(http.StatusBadRequest, gin.H{"error": "discord_id is required"})
		return
	}

	campaignId := c.Query("campaign_id")
	if campaignId == "" {
		h.log.Info("[handler.GetWhitelistCampaignUserByDiscordId] - campaign id empty")
		c.JSON(http.StatusBadRequest, gin.H{"error": "campaign_id is required"})
		return
	}

	wlUsers, err := h.entities.GetWhitelistCampaignUser(discordId, campaignId)
	if err != nil {
		h.log.Fields(logger.Fields{"discordID": discordId, "campaignID": campaignId}).Error(err, "[handler.GetWhitelistCampaignUserByDiscordId] - failed to get whitelist campaign user")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, wlUsers)
}

// GetWhitelistCampaignUsersCSV     godoc
// @Summary     Get whitelist campaign users csv
// @Description Get whitelist campaign users csv
// @Tags        Whitelist campaign
// @Accept      json
// @Produce     json
// @Param       campaign_id   query  string true  "Campaign ID"
// @Success     200
// @Router      /whitelist-campaigns/users/csv [get]
func (h *Handler) GetWhitelistCampaignUsersCSV(c *gin.Context) {
	campaignId := c.Query("campaign_id")
	if campaignId == "" {
		h.log.Info("[handler.GetWhitelistCampaignUsersCSV] - campaign id empty")
		c.JSON(http.StatusBadRequest, gin.H{"error": "campaign_id is required"})
		return
	}
	wlUsers, err := h.entities.GetWhitelistCampaignUsers(campaignId)
	if err != nil {
		h.log.Fields(logger.Fields{"campaignID": campaignId}).Error(err, "[handler.GetWhitelistCampaignUsersCSV] - failed to get whitelist campaign users")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	b := &bytes.Buffer{}
	w := csv.NewWriter(b)

	if err = w.Write([]string{"address", "discord_id", "notes", "whitelist_campaign_id", "created_at"}); err != nil {
		h.log.Error(err, "[handler.GetWhitelistCampaignUsersCSV] - failed to write to csv")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	for _, wl := range wlUsers {
		if err = w.Write([]string{wl.Address, wl.DiscordID, wl.Notes, wl.WhitelistCampaignId, wl.CreatedAt.String()}); err != nil {
			h.log.Fields(logger.Fields{"address": wl.Address, "discordID": wl.DiscordID, "notes": wl.Notes, "campaignID": wl.WhitelistCampaignId, "createdAt": wl.CreatedAt.String()}).Error(err, "[handler.GetWhitelistCampaignUsersCSV] - failed to write to csv")
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	w.Flush()
	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Disposition", "attachment; filename=wlusers.csv")
	c.Data(http.StatusOK, "text/csv", b.Bytes())
}
