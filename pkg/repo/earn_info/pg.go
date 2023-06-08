package earninfo

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

func (pg *pg) Create(earnInfo *model.EarnInfo) (*model.EarnInfo, error) {
	return earnInfo, pg.db.Create(earnInfo).Error
}

func (pg *pg) GetById(id int64) (earnInfo *model.EarnInfo, err error) {
	return earnInfo, pg.db.First(&earnInfo, id).Error
}

func (pg *pg) List(q ListQuery) (earnInfos []model.EarnInfo, total int64, err error) {
	db := pg.db.Model(&model.EarnInfo{}).Where("deadline_at > now()").Order("created_at ASC")
	db = db.Count(&total).Offset(q.Offset)
	if q.Limit != 0 {
		db = db.Limit(q.Limit)
	}
	return earnInfos, total, db.Find(&earnInfos).Error
}
