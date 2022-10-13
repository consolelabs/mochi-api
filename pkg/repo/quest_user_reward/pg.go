package questuserreward

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

func (pg *pg) CreateMany(list []model.QuestUserReward) error {
	tx := pg.db.Begin()
	for i, item := range list {
		if err := tx.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "user_id"}, {Name: "quest_id"}, {Name: "start_time"}, {Name: "pass_id"}, {Name: "reward_id"}},
			DoNothing: true,
		}).Create(&item).Error; err != nil {
			tx.Rollback()
			return err
		}
		list[i] = item
	}
	return tx.Commit().Error
}

func (pg *pg) List(q ListQuery) ([]model.QuestUserReward, error) {
	db := pg.db
	var list []model.QuestUserReward
	if q.UserID != nil {
		db = db.Where("user_id = ?", q.UserID)
	}
	if q.StartTime != nil {
		db = db.Where("start_time = ?", q.StartTime)
	}
	return list, db.Find(&list).Error
}
