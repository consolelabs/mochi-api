package daovoting

import "github.com/gin-gonic/gin"

type IHandler interface {
	GetProposals(c *gin.Context)
	CreateDaoVote(c *gin.Context)
	GetUserVotes(c *gin.Context)
	CreateProposal(c *gin.Context)
}
