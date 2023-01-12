package guild_user_activity_log

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

func (pg *pg) CreateOne(record model.GuildUserActivityLog) error {
	return pg.db.Create(&record).Error
}

func (pg *pg) CreateOneNoGuild(record model.GuildUserActivityLog) error {
	return pg.db.Select("UserID", "ActivityName", "EarnedXP").Create(&record).Error
}

func (pg *pg) CreateBatch(records []model.GuildUserActivityLog) error {
	return pg.db.Create(&records).Error
}
