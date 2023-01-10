package daovoting

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/defipod/mochi/pkg/entities"
	"github.com/defipod/mochi/pkg/logger"
	errs "github.com/defipod/mochi/pkg/model/errors"
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

// GetProposals     godoc
// @Summary     Get dao proposals
// @Description Get dao proposals
// @Tags        DAO-Voting
// @Accept      json
// @Produce     json
// @Param       user_discord_id   query  string true  "Discord ID"
// @Success     200 {object} response.GetAllDaoProposals
// @Router      /dao-voting/proposals [get]
func (h *Handler) GetProposals(c *gin.Context) {
	userId := c.Query("user_discord_id")
	if userId == "" {
		h.log.Info("[handler.GetProposals] - discord id empty")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("user_discord_id is required"), nil))
		return
	}

	proposals, err := h.entities.GetAllDaoProposalByUserId(userId)
	if err != nil {
		h.log.Fields(logger.Fields{"discord_id": userId}).Error(err, "[handler.GetProposals] - failed to get proposals by discord id")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	c.JSON(http.StatusOK, response.CreateResponse(proposals, nil, nil, nil))
}

// GetProposals     godoc
// @Summary     Get dao votes
// @Description Get dao votes
// @Tags        DAO-Voting
// @Accept      json
// @Produce     json
// @Param       user_discord_id   query  string true  "Discord ID"
// @Param       proposal-id   query  string false  "Proposal ID"
// @Success     200 {object} response.GetAllDaoProposalVotes
// @Router      /dao-voting/user-votes [get]
func (h *Handler) GetUserVotes(c *gin.Context) {
	userId := c.Query("user_discord_id")
	proposalId := c.Param("proposal_id")
	if userId == "" || proposalId == "" {
		h.log.Info("[handler.GetUserVotes] - discord id empty")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("user_discord_id and proposal-id are required"), nil))
		return
	}

	// with proposal and user id, get proposal's votes where user is creator
	votes, err := h.entities.GetDaoProposalVotes(proposalId, userId)
	if err != nil {
		h.log.Fields(logger.Fields{"discord_id": userId}).Error(err, "[handler.GetUserVotes] - failed to get proposal's vote")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	c.JSON(http.StatusOK, response.CreateResponse(votes, nil, nil, nil))

}

// CreateDaoVote      godoc
// @Summary     Create dao vote
// @Description Create dao vote
// @Tags        Dao-voting
// @Accept      json
// @Produce     json
// @Param       Request  body request.CreateDaoVoteRequest true "Create dao vote request"
// @Success     200        {object} response.ResponseMessage
// @Router      /dao-voting/proposals/vote [post]
func (h *Handler) CreateDaoVote(c *gin.Context) {
	var req request.CreateDaoVoteRequest
	if err := c.BindJSON(&req); err != nil {
		h.log.Fields(logger.Fields{"userID": req.UserID, "proposalID": req.ProposalID}).Error(err, "[handler.CreateDaoVote] - failed to read JSON")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	if err := h.entities.CreateDaoVote(req); err != nil {
		h.log.Fields(logger.Fields{"req": req}).Error(err, "[handler.CreateDaoVote] - failed to create vote")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(response.ResponseMessage{Message: "OK"}, nil, nil, nil))
}

// AddContract   godoc
// @Summary     Dao Proposal
// @Description Create dao proposal and then create a discussion channel for users to discuss about the proposal.
// @Tags        DAO Proposal
// @Accept      json
// @Produce     json
// @Param       Request  body request.CreateDaoProposalRequest true "Create dao proposal request"
// @Success     200 {object} response.CreateDaoProposalResponse
// @Router      /dao-voting/proposals [post]
func (h *Handler) CreateProposal(c *gin.Context) {
	var req request.CreateDaoProposalRequest

	if err := c.BindJSON(&req); err != nil {
		h.log.Fields(logger.Fields{"body": req}).Error(err, "[handler.CreateProposal] - failed to read JSON")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	daoProposal, err := h.entities.CreateDaoProposal(&req)
	if err != nil {
		h.log.Fields(logger.Fields{"body": req}).Error(err, "[handler.CreateProposal] - failed to create dao proposal")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(daoProposal, nil, nil, nil))
}

// GetProposalVoteOfUser     godoc
// @Summary     Get dao proposal vote of user
// @Description Get dao proposal vote of user
// @Tags        DAO-Voting
// @Accept      json
// @Produce     json
// @Param       proposal_id   query  string true  "Proposal ID"
// @Param       user_discord_id   query  string true  "Discord ID"
// @Success     200 {object} response.GetVote
// @Router      /dao-voting/votes [get]
func (h *Handler) GetVote(c *gin.Context) {
	userId := c.Query("user_discord_id")
	if userId == "" {
		h.log.Info("[handler.GetVote] - discord id empty")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errs.ErrInvalidDiscordUserID, nil))
		return
	}
	proposalId := c.Query("proposal_id")
	if proposalId == "" {
		h.log.Info("[handler.GetVote] - proposal id empty")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errs.ErrInvalidProposalID, nil))
		return
	}

	vote, err := h.entities.GetDaoProposalVoteOfUser(proposalId, userId)
	if err != nil {
		h.log.Fields(logger.Fields{
			"proposalId": proposalId,
			"discord_id": userId,
		}).Error(err, "[handler.GetVote] - entities.GetDaoProposalVoteOfUser failed")
		c.JSON(errs.GetStatusCode(err), response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	c.JSON(http.StatusOK, response.CreateResponse(vote, nil, nil, nil))
}

// UpdateDaoVote      godoc
// @Summary     Update dao vote
// @Description Update dao vote
// @Tags        Dao-voting
// @Accept      json
// @Produce     json
// @Param       vote_id   path  string true  "DAO Vote ID"
// @Param       Request  body request.UpdateDaoVoteRequest true "Update dao vote request"
// @Success     200        {object} response.UpdateVote
// @Router      /dao-voting/proposals/votes/{vote_id} [put]
func (h *Handler) UpdateDaoVote(c *gin.Context) {
	voteId := c.Param("vote_id")
	if voteId == "" {
		h.log.Info("[handler.UpdateDaoVote] - vote id empty")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errs.ErrInvalidVoteID, nil))
		return
	}

	var req request.UpdateDaoVoteRequest
	if err := c.BindJSON(&req); err != nil {
		h.log.Fields(logger.Fields{"request": req}).Error(err, "[handler.UpdateDaoVote] - failed to read JSON")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	vote, err := h.entities.UpdateDaoVote(voteId, req)
	if err != nil {
		h.log.Fields(logger.Fields{"req": req}).Error(err, "[handler.UpdateDaoVote] - entities.UpdateDaoVote failed")
		c.JSON(errs.GetStatusCode(err), response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	c.JSON(http.StatusOK, response.CreateResponse(vote, nil, nil, nil))
}

// DeleteDaoProposal   godoc
// @Summary     Delete DAO Proposal
// @Description Detele DAO proposal and then remove its discussion channel.
// @Tags        DAO Proposal
// @Accept      json
// @Produce     json
// @Param       Request  body request.DeteteDaoProposalRequest true "Detete dao proposal request"
// @Success     200 {object} response.DeteteDaoProposalResponse
// @Router      /dao-voting/proposals [detete]
func (h *Handler) DeteteProposal(c *gin.Context) {
	proposalId := c.Param("proposal_id")
	if proposalId == "" {
		h.log.Info("[handler.DeteteProposal] - proposal_id empty")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("proposal_id is required"), nil))
		return
	}

	err := h.entities.DeleteDaoProposal(proposalId)
	if err != nil {
		h.log.Fields(logger.Fields{"proposal_id": proposalId}).Error(err, "[handler.DeleteProposal] - failed to delete dao proposal")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(response.ResponseMessage{Message: "OK"}, nil, nil, nil))
}

// TokenHolderStatus     godoc
// @Summary     Get status of token holder for creating proposal and voting
// @Description Check token holder connect wallet yet? And have enough amount based on criteria (has 10 icy, 3 neko, havent connected walelt, …)
// @Tags        DAO-Voting
// @Accept      json
// @Produce     json
// @Param       user-discord-id   query  string true  "Discord ID"
// @Success     200 {object} response.TokenHolderStatus
// @Router      /dao-voting/token-holder/status [get]
func (h *Handler) TokenHolderStatus(c *gin.Context) {
	var query request.TokenHolderStatusRequest
	if err := c.BindQuery(&query); err != nil {
		h.log.Info("[handler.GetProposals] - bind query failed")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("invalid params"), nil))
		return
	}
	resp, err := h.entities.TokenHolderStatus(query)
	if err != nil {
		h.log.Fields(logger.Fields{
			"query": query,
		}).Error(err, "[handler.TokenHolderStatus] - entities.TokenHolderStatus failed")
		c.JSON(errs.GetStatusCode(err), response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	c.JSON(http.StatusOK, resp)
}
