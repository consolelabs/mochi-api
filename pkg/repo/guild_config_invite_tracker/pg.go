package guildconfiginvitetracker

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

func (pg *pg) Upsert(config *model.GuildConfigInviteTracker) error {
	tx := pg.db.Begin()

	err := tx.Omit(clause.Associations).Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "guild_id"}},
		DoUpdates: clause.Set{
			{
				Column: clause.Column{Name: "channel_id"},
				Value:  config.ChannelID,
			},
		},
	}).Create(&config).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
