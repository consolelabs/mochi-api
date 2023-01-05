package dao_vote

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

func (pg *pg) GetById(id int64) (model *model.DaoVote, err error) {
	return model, pg.db.Preload("Proposal").First(&model, id).Error
}

func (pg *pg) GetByUserAndProposalID(proposalId int64, userId string) (model *model.DaoVote, err error) {
	return model, pg.db.
		Where("proposal_id = ? AND user_id = ?", proposalId, userId).
		First(&model).Error
}

func (pg *pg) GetByUserId(userId string) (models *[]model.DaoVote, err error) {
	return models, pg.db.Where("user_id = ?", userId).Find(&models).Error
}
func (pg *pg) GetByProposalId(proposalId int64) (models *[]model.DaoVote, err error) {
	return models, pg.db.Where("proposal_id = ?", proposalId).Find(&models).Error
}

func (pg *pg) CreateDaoVote(vote *model.DaoVote) error {
	return pg.db.Create(vote).Error
}
