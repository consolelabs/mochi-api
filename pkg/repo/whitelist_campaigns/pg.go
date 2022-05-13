package whitelist_campaigns

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

func (pg *pg) GetByGuildId(guildId string) ([]model.WhitelistCampaign, error) {
	var campaigns []model.WhitelistCampaign
	err := pg.db.Where("guild_id = ?", guildId).Find(&campaigns).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get campaigns: %w", err)
	}

	return campaigns, nil
}

func (pg *pg) CreateIfNotExists(campaign model.WhitelistCampaign) error {
	tx := pg.db.Begin()
	err := tx.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "name"}, {Name: "guild_id"}},
		DoNothing: true,
	}).Create(&campaign).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (pg *pg) GetByID(id string) (*model.WhitelistCampaign, error) {
	var campaign model.WhitelistCampaign
	return &campaign, pg.db.First(&campaign, "id = ?", id).Error
}
