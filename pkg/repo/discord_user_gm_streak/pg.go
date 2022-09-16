package discord_user_gm_streak

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

func (pg *pg) UpsertOne(streak model.DiscordUserGMStreak) error {
	tx := pg.db.Begin()

	// update on conflict
	err := tx.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "discord_id"}, {Name: "guild_id"}},
		UpdateAll: true,
	}).Create(&streak).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (pg *pg) UpsertBatch(streaks []model.DiscordUserGMStreak) error {
	tx := pg.db.Begin()

	// update on conflict
	err := tx.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "discord_id"}, {Name: "guild_id"}},
		UpdateAll: true,
	}).Create(&streaks).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (pg *pg) GetByDiscordIDGuildID(discordID, guildID string) (*model.DiscordUserGMStreak, error) {
	var streaks model.DiscordUserGMStreak

	return &streaks, pg.db.Where("discord_id = ? AND guild_id = ?", discordID, guildID).First(&streaks).Error
}

func (pg *pg) GetAll() ([]model.DiscordUserGMStreak, error) {
	var streaks []model.DiscordUserGMStreak
	return streaks, pg.db.Find(&streaks).Error
}
