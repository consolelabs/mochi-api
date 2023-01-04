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
// @Param       user-discord-id   query  string true  "Discord ID"
// @Success     200 {object} response.GetAllDaoProposals
// @Router      /dao-voting/proposals [get]
func (h *Handler) GetProposals(c *gin.Context) {
	userId := c.Query("user-discord-id")
	if userId == "" {
		h.log.Info("[handler.GetProposals] - discord id empty")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("user-discord-id is required"), nil))
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
// @Param       user-discord-id   query  string true  "Discord ID"
// @Param       proposal-id   query  string false  "Proposal ID"
// @Success     200 {object} response.GetAllDaoProposalVotes
// @Router      /dao-voting/user-votes [get]
func (h *Handler) GetUserVotes(c *gin.Context) {
	userId := c.Query("user-discord-id")
	proposalId := c.Param("proposal_id")
	if userId == "" || proposalId == "" {
		h.log.Info("[handler.GetUserVotes] - discord id empty")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("user-discord-id and proposal-id are required"), nil))
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

// GetProposals     godoc
// @Summary     Get dao votes
// @Description Get dao votes
// @Tags        DAO-Voting
// @Accept      json
// @Produce     json
// @Param       proposal-id   query  string true  "Proposal ID"
// @Param       user-discord-id   query  string true  "Discord ID"
// @Success     200 {object} response.GetVote
// @Router      /dao-voting/vote [get]
func (h *Handler) GetVote(c *gin.Context) {
	userId := c.Query("user-discord-id")
	if userId == "" {
		h.log.Info("[handler.GetVote] - discord id empty")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errs.ErrInvalidDiscordUserID, nil))
		return
	}
	proposalId := c.Query("proposal-id")
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
