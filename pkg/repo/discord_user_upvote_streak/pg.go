package discord_user_upvote_streak

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

func (pg *pg) GetTopByStreak() ([]model.DiscordUserUpvoteStreak, error) {
	var streaks []model.DiscordUserUpvoteStreak
	return streaks, pg.db.Table("discord_user_upvote_streaks").Order("streak_count DESC").Limit(10).Find(&streaks).Error
}

func (pg *pg) GetTopByTotal() ([]model.DiscordUserUpvoteStreak, error) {
	var streaks []model.DiscordUserUpvoteStreak
	return streaks, pg.db.Table("discord_user_upvote_streaks").Order("total_count DESC").Limit(10).Find(&streaks).Error
}

func (pg *pg) GetGuildTopByStreak(guildId string) ([]model.DiscordUserUpvoteStreak, error) {
	var res []model.DiscordUserUpvoteStreak
	rows, err := pg.db.Raw(`
	select * from discord_user_upvote_streaks where discord_id IN (
		select DISTINCT user_id from guild_users where guild_id=?
	)
	ORDER BY streak_count DESC LIMIT 10
	`, guildId).Rows()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		tmp := model.DiscordUserUpvoteStreak{}
		if err := rows.Scan(&tmp.DiscordID, &tmp.StreakCount, &tmp.TotalCount, &tmp.LastStreakDate, &tmp.CreatedAt, &tmp.UpdatedAt); err != nil {
			return nil, err
		}
		res = append(res, tmp)
	}
	return res, nil
}

func (pg *pg) GetGuildTopByTotal(guildId string) ([]model.DiscordUserUpvoteStreak, error) {
	var res []model.DiscordUserUpvoteStreak
	rows, err := pg.db.Raw(`
	select * from discord_user_upvote_streaks where discord_id IN (
		select DISTINCT user_id from guild_users where guild_id=?
	)
	ORDER BY total_count DESC LIMIT 10
	`, guildId).Rows()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		tmp := model.DiscordUserUpvoteStreak{}
		if err := rows.Scan(&tmp.DiscordID, &tmp.StreakCount, &tmp.TotalCount, &tmp.LastStreakDate, &tmp.CreatedAt, &tmp.UpdatedAt); err != nil {
			return nil, err
		}
		res = append(res, tmp)
	}
	return res, nil
}
