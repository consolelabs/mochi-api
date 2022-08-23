package guild_user_activity_log

import (
	"github.com/defipod/mochi/pkg/model"
	"gorm.io/gorm"
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
	return pg.db.Select("UserID", "ActivityName").Create(&record).Error
}
