package airdropcampaign

import (
	"strings"

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

func (pg *pg) Upsert(ac *model.AirdropCampaign) (*model.AirdropCampaign, error) {
	tx := pg.db.Begin()

	// update on conflict
	err := tx.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		UpdateAll: true,
	}).Create(&ac).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	return ac, tx.Commit().Error
}

func (pg *pg) GetById(id int64) (ac *model.AirdropCampaign, err error) {
	return ac, pg.db.First(&ac, id).Error
}

func (pg *pg) List(q ListQuery) (acs []model.AirdropCampaign, total int64, err error) {
	db := pg.db.Model(&model.AirdropCampaign{}).Order("CASE WHEN deadline_at IS NOT NULL and deadline_at > NOW() THEN 0 WHEN deadline_at IS NULL THEN 1 ELSE 2 END, deadline_at ASC")

	if q.Status != "" {
		db = db.Where("status = ?", q.Status)
	}

	isSearchById := strings.HasPrefix(q.Keyword, "#")

	if isSearchById {
		searchCampaignId := []string{strings.TrimPrefix(q.Keyword, "#")}
		db = db.Where("id = ?", searchCampaignId)
	}

	if q.Keyword != "" && !isSearchById {
		likeKeyword := "%" + q.Keyword + "%"
		db = db.Where("title ILIKE ? OR detail ILIKE ?", likeKeyword, likeKeyword)
	}

	db = db.Count(&total).Offset(q.Offset)
	if q.Limit != 0 {
		db = db.Limit(q.Limit)
	}
	return acs, total, db.Find(&acs).Error
}

func (pg *pg) CountStat() ([]model.AirdropStatusCount, error) {
	var stats []model.AirdropStatusCount
	err := pg.db.Table("airdrop_campaigns").Select("status, count(*) as count").Group("status").Scan(&stats).Error

	return stats, err
}
