package dao_guideline_messages

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

func (pg *pg) GetByAuthority(authority model.ProposalAuthorityType) (model *model.DaoGuidelineMessage, err error) {
	return model, pg.db.Where("authority = ?", authority).First(&model).Error
}
