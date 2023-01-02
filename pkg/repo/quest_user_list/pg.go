package questuserlist

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/plugin/dbresolver"

	"github.com/defipod/mochi/pkg/model"
)

type pg struct {
	db *gorm.DB
}

func NewPG(db *gorm.DB) Store {
	return &pg{db: db}
}

func (pg *pg) UpsertMany(list []model.QuestUserList) error {
	tx := pg.db.Begin()
	for i, quest := range list {
		if err := tx.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "user_id"}, {Name: "quest_id"}, {Name: "start_time"}},
			DoUpdates: clause.AssignmentColumns([]string{"current", "is_completed", "is_claimed"}),
		}).Create(&quest).Error; err != nil {
			tx.Rollback()
			return err
		}
		list[i] = quest
	}
	return tx.Commit().Error
}

func (pg *pg) List(q ListQuery) ([]model.QuestUserList, error) {
	db := pg.db
	var list []model.QuestUserList
	if q.UserID != nil {
		db = db.Where("user_id = ?", *q.UserID)
	}
	if q.QuestID != nil {
		db = db.Where("quest_id = ?", *q.QuestID)
	}
	if q.StartTime != nil {
		db = db.Where("start_time = ?", *q.StartTime)
	}
	if q.Routine != nil {
		db = db.Where("routine::TEXT = ?", *q.Routine)
	}
	if q.Action != nil {
		db = db.Where("action::TEXT = ?", *q.Action)
	}
	if q.NotActions != nil && len(q.NotActions) > 0 {
		db = db.Where("action::TEXT NOT IN ?", q.NotActions)
	}
	if q.IsCompleted != nil {
		db = db.Where("is_completed = ?", *q.IsCompleted)
	}
	if q.IsClaimed != nil {
		db = db.Where("is_claimed = ?", *q.IsClaimed)
	}
	db = db.Preload("Quest").Preload("Quest.Rewards").Preload("Quest.Rewards.RewardType")
	return list, db.Clauses(dbresolver.Write).Find(&list).Error
}
