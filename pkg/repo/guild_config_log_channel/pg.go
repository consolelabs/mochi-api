package guildconfiglogchannel

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

func (pg pg) Upsert(model *model.GuildConfigLogChannel) (*model.GuildConfigLogChannel, error) {
	tx := pg.db.Begin()

	// update on conflict
	err := tx.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "guild_id"}, {Name: "log_type"}},
		UpdateAll: true,
	}).Create(&model).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	return model, tx.Commit().Error
}

func (pg pg) Get(query Query) (models []model.GuildConfigLogChannel, err error) {
	db := pg.db
	if query.GuildId != "" {
		db = db.Where("guild_id = ?", query.GuildId)
	}

	if query.LogType != "" {
		db = db.Where("log_type = ?", query.LogType)
	}

	return models, db.Find(&models).Error
}
