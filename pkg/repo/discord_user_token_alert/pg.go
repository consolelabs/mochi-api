package discord_user_token_alert

import (
	"github.com/defipod/mochi/pkg/model"
	"gorm.io/gorm"
)

type pg struct {
	db *gorm.DB
}

func NewPG(db *gorm.DB) Store {
	return &pg{
		db: db,
	}
}
func (pg *pg) CreateOne(config *model.UpsertDiscordUserTokenAlert) error {
	return pg.db.Table("discord_user_token_alerts").Create(&config).Error
}
func (pg *pg) RemoveOne(config *model.DiscordUserTokenAlert) error {
	return pg.db.Where("id=?", config.ID).Delete(&config).Error
}
func (pg *pg) GetByDiscordID(discordId string) ([]model.DiscordUserTokenAlert, error) {
	configs := []model.DiscordUserTokenAlert{}
	return configs, pg.db.Preload("DiscordUserDevice").Where("discord_id=?", discordId).Find(&configs).Error
}
func (pg *pg) GetByDeviceID(deviceId string) ([]model.DiscordUserTokenAlert, error) {
	configs := []model.DiscordUserTokenAlert{}
	return configs, pg.db.Preload("DiscordUserDevice").Where("device_id=?", deviceId).Find(&configs).Error
}
func (pg *pg) GetAll() ([]model.DiscordUserTokenAlert, error) {
	configs := []model.DiscordUserTokenAlert{}
	return configs, pg.db.Preload("DiscordUserDevice").Find(&configs).Error
}
