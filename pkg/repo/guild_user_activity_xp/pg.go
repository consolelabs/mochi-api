package guild_user_activity_xp

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

func (pg *pg) CreateOne(record model.GuildUserActivityXP) error {
	return pg.db.Create(&record).Error
}
