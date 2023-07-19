package dao_guideline_messages

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	GetByAuthority(authority model.ProposalAuthorityType) (*model.DaoGuidelineMessage, error)
}
