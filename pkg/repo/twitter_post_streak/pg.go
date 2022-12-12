package twitterpoststreak

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

func (pg *pg) List(q ListQuery) ([]model.TwitterPostStreak, int64, error) {
	var streaks []model.TwitterPostStreak
	var total int64
	db := pg.db.Table("twitter_post_streaks")
	if q.GuildID != "" {
		db = db.Where("guild_id = ?", q.GuildID)
	}
	if q.TwitterID != "" {
		db = db.Where("twitter_id = ?", q.TwitterID)
	}
	if q.TwitterHandle != "" {
		db = db.Where("twitter_handle = ?", q.TwitterHandle)
	}
	db.Count(&total)
	if q.Sort != "" {
		db = db.Order(q.Sort)
	}
	if q.Offset != 0 {
		db = db.Offset(q.Offset)
	}
	if q.Limit != 0 {
		db = db.Limit(q.Limit)
	}
	return streaks, total, db.Find(&streaks).Error
}

func (pg *pg) UpsertOne(streak *model.TwitterPostStreak) error {
	tx := pg.db.Begin()

	// update on conflict
	err := tx.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "twitter_id"}, {Name: "guild_id"}},
		UpdateAll: true,
	}, clause.OnConflict{
		Columns:   []clause.Column{{Name: "twitter_handle"}, {Name: "guild_id"}},
		UpdateAll: true,
	}).Create(streak).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
