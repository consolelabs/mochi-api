package airdropcampaign

import (
	"net/http"

	"github.com/gin-gonic/gin"

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

// GetAirdropCampaigns     godoc
// @Summary     Get Airdrop Campaign List
// @Description Get Airdrop Campaign List
// @Tags        Airdrop-campaigns
// @Accept      json
// @Produce     json
// @Param       status   query  string false  "status"
// @Param       keyword   query  string false  "keyword"
// @Param       profile_id   query  string false  "profile id"
// @Param       page   query  string false  "page"
// @Param       size   query  string false  "size"
// @Success     200 {object} response.AirdropCampaignsResponse
// @Router      /earns/airdrop-campaigns [get]
func (h *Handler) GetAirdropCampaigns(c *gin.Context) {
	req := request.GetAirdropCampaignsRequest{}
	if err := c.ShouldBindQuery(&req); err != nil {
		h.log.Error(err, "[handler.GetAirdropCampaigns] ShouldBindQuery() failed")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	if err := req.Validate(); err != nil {
		h.log.Error(err, "[handler.GetAirdropCampaign] validate request failed")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	data, paging, err := h.entities.GetAirdropCampaigns(req)
	if err != nil {
		h.log.Error(err, "[handler.GetAirdropCampaigns] failed to get Airdrop-campaigns list")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse[any](data, paging, nil, nil))
	return
}

// GetAirdropCampaign     godoc
// @Summary     Get Airdrop Campaign By Id
// @Description Get Airdrop Campaign By Id
// @Tags        Airdrop-campaigns
// @Accept      json
// @Produce     json
// @Param       id   path  string true  "airdrop campaign id"
// @Param       profile_id   query  string false  "profile id"
// @Success     200 {object} response.AirdropCampaignResponse
// @Router      /earns/airdrop-campaigns/{id} [get]
func (h *Handler) GetAirdropCampaign(c *gin.Context) {
	req := request.GetAirdropCampaignRequest{}
	if err := c.ShouldBindUri(&req); err != nil {
		h.log.Error(err, "[handler.GetAirdropCampaign] ShouldBindUri() failed")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	if err := c.ShouldBindQuery(&req); err != nil {
		h.log.Error(err, "[handler.GetAirdropCampaign] ShouldBindQuery() failed")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	if err := req.Validate(); err != nil {
		h.log.Error(err, "[handler.GetAirdropCampaign] validate request failed")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	data, err := h.entities.GetAirdropCampaign(req)
	if err != nil {
		h.log.Error(err, "[handler.GetAirdropCampaign] failed to get Airdrop-campaigns by id")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse[any](data, nil, nil, nil))
	return
}

// CreateAirdropCampaign     godoc
// @Summary     Create airdrop campaign
// @Description Create airdrop campaign
// @Tags        Airdrop-campaigns
// @Param       Request  body request.CreateAirdropCampaignRequest true "Create airdrop campaign request"
// @Accept      json
// @Produce     json
// @Success     200 {object} response.AirdropCampaignResponse
// @Router      /earns/airdrop-campaigns [post]
func (h *Handler) CreateAirdropCampaign(c *gin.Context) {
	var req request.CreateAirdropCampaignRequest
	if err := c.BindJSON(&req); err != nil {
		h.log.Error(err, "[handler.CreateAirdropCampaign] BindJSON() failed")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	if err := req.Validate(); err != nil {
		h.log.Error(err, "[handler.CreateAirdropCampaign] validate request failed")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	data, err := h.entities.CreateAirdropCampaign(&req)
	if err != nil {
		h.log.Error(err, "[handler.CreateAirdropCampaign] failed to create airdrop campaign")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse[any](data, nil, nil, nil))
	return
}

// GetAirdropCampaignStats     godoc
// @Summary     Get Airdrop Campaign List
// @Description Get Airdrop Campaign List
// @Tags        Airdrop-campaigns
// @Accept      json
// @Produce     json
// @Param       profile_id   query  string false  "profile_id"
// @Success     200 {object} response.AirdropCampaignStatResponse
// @Router      /earns/airdrop-campaigns/stats [get]
func (h *Handler) GetAirdropCampaignStats(c *gin.Context) {
	var req request.GetAirdropCampaignStatus
	if err := c.BindQuery(&req); err != nil {
		h.log.Error(err, "[handler.GetAirdropCampaignStats] BindQuery() failed")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	if err := req.Validate(); err != nil {
		h.log.Error(err, "[handler.GetAirdropCampaignStats] validate request failed")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	data, err := h.entities.GetAirdropCampaignStats(req)
	if err != nil {
		h.log.Error(err, "[handler.GetAirdropCampaignStats] failed to get Airdrop-campaigns list")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, data)
	return
}

// CreateProfileAirdropCampaign     godoc
// @Summary     Create profile airdrop campaign
// @Description Create profile airdrop campaign
// @Tags        Airdrop-campaigns
// @Param       id   path  string true  "profile id"
// @Param       Request  body request.CreateProfileAirdropCampaignRequest true "Create profile airdrop campaign request"
// @Accept      json
// @Produce     json
// @Success     200 {object} response.ProfileAirdropCampaignResponse
// @Router      /users/{id}/earns/airdrop-campaigns [post]
func (h *Handler) CreateProfileAirdropCampaign(c *gin.Context) {
	var req request.CreateProfileAirdropCampaignRequest
	if err := c.BindJSON(&req); err != nil {
		h.log.Error(err, "[handler.CreateProfileAirdropCampaign] BindJSON() failed")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	req.ProfileId = c.Param("id")

	if err := req.Validate(); err != nil {
		h.log.Error(err, "[handler.CreateProfileAirdropCampaign] validate request failed")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	data, err := h.entities.CreateProfileAirdropCampaign(&req)
	if err != nil {
		h.log.Error(err, "[handler.CreateProfileAirdropCampaign] failed to create profile airdrop campaign")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse[any](data, nil, nil, nil))
	return
}

// GetProfileAirdropCampaigns     godoc
// @Summary     Get user earn list
// @Description Get user earn list
// @Tags        Airdrop-campaigns
// @Param       id   path  string true  "profile id"
// @Param       status   query  string false  "status"
// @Param       is_favorite   query  bool false  "is_favorite"
// @Param       page   query  string false  "page"
// @Param       size   query  string false  "size"
// @Accept      json
// @Produce     json
// @Success     200 {object} response.ProfileAirdropCampaignsResponse
// @Router      /users/{id}/earns/airdrop-campaigns [get]
func (h *Handler) GetProfileAirdropCampaigns(c *gin.Context) {
	req := request.GetProfileAirdropCampaignsRequest{}
	if err := c.ShouldBindQuery(&req); err != nil {
		h.log.Error(err, "[handler.GetProfileAirdropCampaigns] BindJSON() failed")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	req.ProfileId = c.Param("id")

	if err := req.Validate(); err != nil {
		h.log.Error(err, "[handler.GetProfileAirdropCampaigns] validate request failed")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	data, paging, err := h.entities.GetProfileAirdropCampaigns(req)
	if err != nil {
		h.log.Error(err, "[handler.GetProfileAirdropCampaigns] failed to get profile airdrop campaigns")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse[any](data, paging, nil, nil))
	return
}

// DeleteProfileAirdropCampaign     godoc
// @Summary     Delete profile airdrop campaign
// @Description Delete profile airdrop campaign
// @Tags        Airdrop-campaigns
// @Param       id   path  string true  "profile id"
// @Param       airdrop_campaign_id   path  string true  "airdrop campaign id"
// @Accept      json
// @Produce     json
// @Success     200 {object} response.ResponseMessage
// @Router      /users/{id}/earns/airdrop-campaigns/{airdrop_campaign_id} [delete]
func (h *Handler) DeleteProfileAirdropCampaign(c *gin.Context) {
	req := request.DeleteProfileAirdropCampaignRequest{}
	if err := c.ShouldBindUri(&req); err != nil {
		h.log.Error(err, "[handler.DeleteProfileAirdropCampaign] BindUri() failed")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	err := h.entities.RemoveProfileAirdropCampaign(req)
	if err != nil {
		h.log.Error(err, "[handler.DeleteProfileAirdropCampaign] failed to delete user earn")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse[any](gin.H{"message": "ok"}, nil, nil, nil))
	return
}
