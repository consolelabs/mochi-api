package profileairdropcampaign

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

func (pg *pg) UpsertOne(profileAirdropCampaign *model.ProfileAirdropCampaign) (*model.ProfileAirdropCampaign, error) {
	tx := pg.db.Begin()
	err := tx.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "profile_id"}, {Name: "airdrop_campaign_id"}},
		UpdateAll: true,
	}).Create(&profileAirdropCampaign).Error

	if err != nil {
		tx.Rollback()
		return nil, err
	}

	return profileAirdropCampaign, tx.Commit().Error
}

func (pg *pg) List(q ListQuery) (pacs []model.ProfileAirdropCampaign, total int64, err error) {
	db := pg.db.Model(&model.ProfileAirdropCampaign{})

	if q.ProfileId != "" {
		db = db.Where("profile_id = ?", q.ProfileId)
	}

	if q.Status != "" {
		db = db.Where("status = ?", q.Status)
	}

	if q.IsFavorite != nil {
		db = db.Where("is_favorite = ?", &q.IsFavorite)
	}

	db = db.Count(&total).Offset(q.Offset)

	if q.Limit != 0 {
		db = db.Limit(q.Limit)
	}
	return pacs, total, db.Preload("AirdropCampaign").Order("airdrop_campaign_id ASC").Find(&pacs).Error
}

func (pg *pg) Delete(pac *model.ProfileAirdropCampaign) (*model.ProfileAirdropCampaign, error) {
	return pac, pg.db.Where("profile_id = ? and airdrop_campaign_id = ?", pac.ProfileId, pac.AirdropCampaignId).Delete(pac).Error
}

func (pg *pg) CountStat(q StatQuery) (stats []model.AirdropStatusCount, err error) {
	db := pg.db.Table("profile_airdrop_campaigns").Select("status, count(*) as count")

	if q.Status != "" {
		db = db.Where("status = ?", q.Status)
	}

	if q.ProfileId != "" {
		db = db.Where("profile_id = ?", q.ProfileId)
	}

	err = db.Group("status").Scan(&stats).Error

	return stats, err
}
