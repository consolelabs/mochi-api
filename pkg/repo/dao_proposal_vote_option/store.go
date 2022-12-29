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
