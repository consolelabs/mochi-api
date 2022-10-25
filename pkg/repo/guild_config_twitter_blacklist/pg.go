package guildconfigtwitterblacklist

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

func (pg *pg) List(q ListQuery) ([]model.GuildConfigTwitterBlacklist, error) {
	db := pg.db
	var list []model.GuildConfigTwitterBlacklist
	if q.GuildID != "" {
		db = db.Where("guild_id = ?", q.GuildID)
	}
	return list, db.Find(&list).Error
}

func (pg *pg) Upsert(model *model.GuildConfigTwitterBlacklist) error {
	tx := pg.db.Begin()
	err := tx.Clauses(clause.OnConflict{DoNothing: true}).Create(model).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (pg *pg) Delete(guildID, twitterID string) error {
	return pg.db.Where("guild_id = ? AND twitter_id = ?", guildID, twitterID).Delete(&model.GuildConfigTwitterBlacklist{}).Error
}
