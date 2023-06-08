package userearn

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

func (pg *pg) UpsertOne(userEarn *model.UserEarn) (*model.UserEarn, error) {
	tx := pg.db.Begin()
	err := tx.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "user_id"}, {Name: "earn_id"}},
		UpdateAll: true,
	}).Create(&userEarn).Error

	if err != nil {
		tx.Rollback()
		return nil, err
	}

	return userEarn, tx.Commit().Error
}

func (pg *pg) Delete(userEarn *model.UserEarn) (*model.UserEarn, error) {
	return userEarn, pg.db.Where("user_id = ? and earn_id = ?", userEarn.UserId, userEarn.EarnId).Delete(userEarn).Error
}

func (pg *pg) GetByUserId(q ListQuery) (userEarns []model.UserEarn, total int64, err error) {
	db := pg.db.Model(&model.UserEarn{}).Where("user_id = ?", q.UserId).Preload("Earn").Order("earn_id ASC")
	db = db.Count(&total).Offset(q.Offset)
	if q.Limit != 0 {
		db = db.Limit(q.Limit)
	}
	return userEarns, total, db.Find(&userEarns).Error
}
