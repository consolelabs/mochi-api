package guild_config_activity

import (
	"github.com/defipod/mochi/pkg/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type pg struct {
	db *gorm.DB
}

func NewPG(db *gorm.DB) Store {
	return &pg{db: db}
}

func (pg *pg) GetOneByActivityName(guildID, activityName string) (*model.GuildConfigActivity, error) {
	config := &model.GuildConfigActivity{}
	return config, pg.db.Joins("LEFT JOIN activities ON activities.id = guild_config_activities.activity_id").Where("guild_config_activities.guild_id = ? AND lower(activities.name) = lower(?) AND guild_config_activities.active = TRUE", guildID, activityName).Preload("Activity").First(config).Error
}

func (pg *pg) UpsertMany(configs []model.GuildConfigActivity) error {
	tx := pg.db.Begin()

	for _, config := range configs {
		err := tx.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "guild_id"}, {Name: "activity_id"}},
			UpdateAll: true,
		}).Create(&config).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}

func (pg *pg) ForkDefaulActivityConfigs(guildID string) error {
	tx := pg.db.Begin()

	if err := tx.Exec(`
	insert into guild_config_activities(guild_id, activity_id, active)
	select ?, id, true from activities on conflict (guild_id, activity_id) do nothing
	`, guildID).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
