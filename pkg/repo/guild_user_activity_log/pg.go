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

func (pg *pg) CreateBatch(records []model.GuildUserActivityLog) error {
	return pg.db.Create(&records).Error
}

func (pg *pg) UpdateInvalidRecords(userID, profileID string) error {
	return pg.db.Model(&model.GuildUserActivityLog{}).Where("user_id = ? AND profile_id = ?", userID, "").Update("profile_id", profileID).Error
}
