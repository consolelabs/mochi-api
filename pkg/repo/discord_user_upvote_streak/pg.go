package discord_user_upvote_streak

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

func (pg *pg) UpsertOne(streak model.DiscordUserUpvoteStreak) error {
	tx := pg.db.Begin()

	// update on conflict
	err := tx.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "discord_id"}},
		UpdateAll: true,
	}).Create(&streak).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (pg *pg) UpsertBatch(streaks []model.DiscordUserUpvoteStreak) error {
	tx := pg.db.Begin()

	// update on conflict
	err := tx.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "discord_id"}},
		UpdateAll: true,
	}).Create(&streaks).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (pg *pg) GetByDiscordID(discordID string) (*model.DiscordUserUpvoteStreak, error) {
	var streak model.DiscordUserUpvoteStreak
	return &streak, pg.db.Table("discord_user_upvote_streaks").Where("discord_id = ?", discordID).First(&streak).Error
}

func (pg *pg) GetAll() ([]model.DiscordUserUpvoteStreak, error) {
	var streaks []model.DiscordUserUpvoteStreak
	return streaks, pg.db.Table("discord_user_upvote_streaks").Find(&streaks).Error
}
