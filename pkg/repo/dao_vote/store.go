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
	return model, pg.db.First(&model, id).Error
}

func (pg *pg) CreateDaoVote(vote *model.DaoVote) error {
	return pg.db.Create(vote).Error
}
