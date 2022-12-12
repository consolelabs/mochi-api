package guild_config_join_leave_channel

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/defipod/mochi/pkg/model"
)

type pg struct {
	db *gorm.DB
}

func NewPG(db *gorm.DB) Store {
	return &pg{db: db}
}

func (pg *pg) GetByGuildID(guildID string) (*model.GuildConfigJoinLeaveChannel, error) {
	config := &model.GuildConfigJoinLeaveChannel{}
	return config, pg.db.Table("guild_config_join_leave_channels").Where("guild_id = ?", guildID).First(config).Error
}

func (pg *pg) UpsertOne(config *model.GuildConfigJoinLeaveChannel) (*model.GuildConfigJoinLeaveChannel, error) {
	tx := pg.db.Begin()

	// update on conflict
	err := tx.Table("guild_config_join_leave_channels").Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "guild_id"}},
		UpdateAll: true,
	}).Create(&config).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	return config, tx.Commit().Error
}

func (pg *pg) DeleteOne(config *model.GuildConfigJoinLeaveChannel) error {
	return pg.db.Table("guild_config_join_leave_channels").Where("guild_id = ?", config.GuildID).Delete(config).Error
}
