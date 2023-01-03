package daovoting

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
func (h *Handler) GetProposals(c *gin.Context) {
	h.entities.Test()
	c.JSON(http.StatusOK, response.CreateResponse("ok", nil, nil, nil))
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
