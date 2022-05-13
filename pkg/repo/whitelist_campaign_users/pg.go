package whitelist_campaign_users

import (
	"fmt"

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

func (pg *pg) GetByCampaignId(campaignId string) ([]model.WhitelistCampaignUser, error) {
	var wlUsers []model.WhitelistCampaignUser
	err := pg.db.Where("whitelist_campaign_id = ?", campaignId).Find(&wlUsers).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get campaigns: %w", err)
	}

	return wlUsers, nil
}

func (pg *pg) GetByDiscordIdCampaignId(discordId, campaignId string) (*model.WhitelistCampaignUser, error) {
	var wlUser model.WhitelistCampaignUser
	return &wlUser, pg.db.First(&wlUser, "discord_id = ? and whitelist_campaign_id = ?", discordId, campaignId).Error
}

func (pg *pg) UpsertOne(wlUser model.WhitelistCampaignUser) error {
	tx := pg.db.Begin()

	err := tx.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "address"}, {Name: "discord_id"}, {Name: "whitelist_campaign_id"}},
		UpdateAll: true,
	}).Create(&wlUser).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
