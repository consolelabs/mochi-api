package discord_user_upvote_log

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/defipod/mochi/pkg/model"
)

type pg struct {
	db *gorm.DB
}

func NewPG(db *gorm.DB) Store {
	return &pg{
		db: db,
	}
}

func (pg *pg) UpsertOne(log model.DiscordUserUpvoteLog) error {
	tx := pg.db.Begin()
	// update on conflict
	err := tx.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "discord_id"}, {Name: "source"}},
		UpdateAll: true,
	}).Create(&log).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (pg *pg) GetByDiscordID(discordID string) ([]model.DiscordUserUpvoteLog, error) {
	var logs []model.DiscordUserUpvoteLog
	return logs, pg.db.Table("discord_user_upvote_logs").Where("discord_id = ?", discordID).Find(&logs).Error
}
