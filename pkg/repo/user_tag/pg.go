package user_tag

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

func (pg *pg) GetOne(q GetOneQuery) (model *model.UserTag, err error) {
	db := pg.db
	if q.UserID != "" {
		db = db.Where("user_id", q.UserID)
	}
	if q.GuildID != "" {
		db = db.Where("guild_id", q.GuildID)
	}

	return model, db.First(&model).Error
}

func (pg *pg) UpsertOne(tag model.UserTag) (*model.UserTag, error) {
	tx := pg.db.Begin()

	// update on conflict
	err := tx.Table("user_tags").Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "user_id"}, {Name: "guild_id"}},
		UpdateAll: true,
	}).Create(&tag).Error
	if err != nil {
		tx.Rollback()
		return &tag, err
	}

	return &tag, tx.Commit().Error
}
