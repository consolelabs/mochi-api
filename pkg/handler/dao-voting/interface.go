package daovoting

import "github.com/gin-gonic/gin"

type IHandler interface {
	GetProposals(c *gin.Context)
	CreateDaoVote(c *gin.Context)
	UpdateDaoVote(c *gin.Context)
	GetUserVotes(c *gin.Context)
	CreateProposal(c *gin.Context)
	GetVote(c *gin.Context)
	DeteteProposal(c *gin.Context)
}
