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
	campaigns, err := h.entities.GetWhitelistCampaigns()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, campaigns)
}

func (h *Handler) GetWhitelistCampaign(c *gin.Context) {
	campaignId := c.Param("campaignId")

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
	wlUsers, err := h.entities.GetWhitelistCampaignUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, wlUsers)
}

func (h *Handler) GetWhitelistCampaignUserByAddress(c *gin.Context) {
	campaignId := c.Param("address")
	address := c.Query("campaign_id")
	wlUsers, err := h.entities.GetWhitelistCampaignUser(campaignId, address)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, wlUsers)
}

func (h *Handler) GetWhitelistCampaignUsersCSV(c *gin.Context) {
	wlUsers, err := h.entities.GetWhitelistCampaignUsers()
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
