package airdropcampaign

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

func (pg *pg) Create(ac *model.AirdropCampaign) (*model.AirdropCampaign, error) {
	return ac, pg.db.Create(ac).Error
}

func (pg *pg) GetById(id int64) (ac *model.AirdropCampaign, err error) {
	return ac, pg.db.First(&ac, id).Error
}

func (pg *pg) List(q ListQuery) (acs []model.AirdropCampaign, total int64, err error) {
	db := pg.db.Model(&model.AirdropCampaign{}).Order("CASE WHEN deadline_at IS NOT NULL and deadline_at > NOW() THEN 0 WHEN deadline_at IS NULL THEN 1 ELSE 2 END, deadline_at ASC")

	db = db.Count(&total).Offset(q.Offset)
	if q.Limit != 0 {
		db = db.Limit(q.Limit)
	}
	return acs, total, db.Find(&acs).Error
}
