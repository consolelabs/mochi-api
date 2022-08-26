package nft_sales_tracker

import (
	"github.com/defipod/mochi/pkg/model"
	"gorm.io/gorm"
)

type pg struct {
	db *gorm.DB
}

func NewPG(db *gorm.DB) Store {
	return &pg{db: db}
}

func (pg *pg) FirstOrCreate(tracker *model.InsertNFTSalesTracker) error {
	return pg.db.Where("contract_address=? AND platform=? AND sales_config_id=?", tracker.ContractAddress, tracker.Platform, tracker.SalesConfigID).FirstOrCreate(tracker).Error
}

func (pg *pg) GetAll() ([]model.NFTSalesTracker, error) {
	data := []model.NFTSalesTracker{}
	err := pg.db.Preload("GuildConfigSalesTracker").Find(&data)
	return data, err.Error
}

func (pg *pg) GetSalesTrackerByGuildID(guildID string) ([]model.NFTSalesTracker, error) {
	trackers := []model.NFTSalesTracker{}
	return trackers, pg.db.
		Preload("GuildConfigSalesTracker").
		Table("nft_sales_trackers").
		Joins("JOIN guild_config_sales_trackers ON nft_sales_trackers.sales_config_id = guild_config_sales_trackers.id").
		Where("guild_id = ?", guildID).
		Find(&trackers).Error
}

func (pg *pg) GetNFTSalesTrackerByContractAndGuildID(guildID, contractAddress string) (*model.NFTSalesTracker, error) {
	var tracker model.NFTSalesTracker
	return &tracker, pg.db.
		Table("nft_sales_trackers").
		Joins("JOIN guild_config_sales_trackers ON nft_sales_trackers.sales_config_id = guild_config_sales_trackers.id").
		Where("guild_id = ? AND contract_address = ?", guildID, contractAddress).
		First(&tracker).Error
}

func (pg *pg) DeleteNFTSalesTracker(salesTrack model.NFTSalesTracker) error {
	return pg.db.Delete(&salesTrack).Error
}
