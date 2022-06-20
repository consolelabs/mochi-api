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
func (pg *pg) GetAllTracker(addr string, platform string) (*[]model.NFTSalesTracker,error){
	trackers:=[]model.NFTSalesTracker{}
	err:=pg.db.Table("nft_sales_trackers").Preload("SalesConfigID").Joins("JOIN guild_config_sales_trackers ON nft_sales_trackers.sales_config_id = guild_config_sales_trackers.id").Where("contract_address=? AND platform=?",addr,platform).Find(&trackers)
	//err:=pg.db.Model(&model.NFTSalesTracker{}).Joins("JOIN guild_config_sales_trackers ON nft_sales_trackers.sales_config_id = guild_config_sales_trackers.id").Scan(&trackers)
	return &trackers, err.Error
}
func (pg *pg) FirstOrCreate(tracker *model.InsertNFTSalesTracker) error {
	return pg.db.Where("contract_address=? AND platform=? AND sales_config_id=?",tracker.ContractAddress,tracker.Platform,tracker.SalesConfigID).FirstOrCreate(tracker).Error
}