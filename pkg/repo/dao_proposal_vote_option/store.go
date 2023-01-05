package dao_proposal_vote_option

import (
	"gorm.io/gorm"

	"github.com/defipod/mochi/pkg/model"
)

type pg struct {
	db *gorm.DB
}

func NewPG(db *gorm.DB) Store {
	return &pg{db: db}
}

func (pg *pg) GetById(id int64) (model *model.DaoProposalVoteOption, err error) {
	return model, pg.db.Where("id = ?", id).First(&model).Error
}

func (pg *pg) Create(model *model.DaoProposalVoteOption) (*model.DaoProposalVoteOption, error) {
	return model, pg.db.Create(&model).Error
}

func (pg *pg) DeleteAllByProposalID(proposalId int64) error {
	return pg.db.Where("proposal_id = ?", proposalId).Delete(&model.DaoProposalVoteOption{}).Error
}
