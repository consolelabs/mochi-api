package discord_user_device

import (
	"github.com/defipod/mochi/pkg/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type pg struct {
	db *gorm.DB
}

func NewPG(db *gorm.DB) Store {
	return &pg{
		db: db,
	}
}

func (pg *pg) UpsertOne(config *model.DiscordUserDevice) error {
	tx := pg.db.Begin()

	// update on conflict
	err := tx.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		UpdateAll: true,
	}).Create(&config).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
func (pg *pg) GetByDeviceID(deviceId string) (*model.DiscordUserDevice, error) {
	device := model.DiscordUserDevice{}
	return &device, pg.db.Where("id=?", deviceId).First(&device).Error
}
func (pg *pg) RemoveByDeviceID(deviceId string) error {
	device := model.DiscordUserDevice{ID: deviceId}
	return pg.db.Where("id=?", deviceId).Delete(&device).Error
}
