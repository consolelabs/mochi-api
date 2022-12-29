package dao_proposal

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

func (pg *pg) GetById(id int64) (model *model.DaoProposal, err error) {
	return model, pg.db.First(&model, id).Error
}
