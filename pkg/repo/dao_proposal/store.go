package dao_proposal

import (
	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/response"
)

type Store interface {
	GetById(id int64) (*model.DaoProposal, error)
	GetUsageStatsWithPaging(page int, size int) (models *[]response.ProposalCount, total int64, err error)
	GetAllByCreatorId(userId string) (*[]model.DaoProposal, error)
	GetAllByGuildId(guildId string) (*[]model.DaoProposal, error)
	GetByCreatorIdAndProposalId(proposal int64, userId string) (models []model.DaoProposalWithView, err error)
	Create(model *model.DaoProposal) (*model.DaoProposal, error)
	UpdateDiscussionChannel(id int64, discussionChannelId string) error
	DeleteById(id int64) error
}
