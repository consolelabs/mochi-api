package nft_sales_tracker

import (
	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/response"
	"github.com/google/uuid"
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

func (pg *pg) GetSalesTrackerByGuildID(guildID string) (resp []response.NFTSalesTrackerData, err error) {
	rows, err := pg.db.Raw(`
		select 
		nft_sales_trackers.id, 
		contract_address,
		platform,
		sales_config_id,
		guild_config_sales_trackers.guild_id,
		guild_config_sales_trackers.channel_id,
		nft_collections.name,
		chains.id,
		chains.name,
		chains.short_name,
		chains.coin_gecko_id,
		chains.currency
		from (
			nft_sales_trackers
			join guild_config_sales_trackers on nft_sales_trackers.sales_config_id = guild_config_sales_trackers.id
			join nft_collections on contract_address = address
			join chains on chains.id = chain_id::int4
		)
		where guild_id = ?
	`, guildID).Rows()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var ID uuid.NullUUID
		var contract string
		var platform string
		var configID uuid.NullUUID
		var guildID string
		var guildChannelID string
		var name string
		var chainID int
		var chainName string
		var chainShortName string
		var coinGeckoID string
		var currency string

		if err = rows.Scan(&ID, &contract, &platform, &configID, &guildID, &guildChannelID, &name, &chainID, &chainName, &chainShortName, &coinGeckoID, &currency); err != nil {
			return nil, err
		}
		resp = append(resp, response.NFTSalesTrackerData{
			ID:                      ID,
			ContractAddress:         contract,
			Platform:                platform,
			SalesConfigID:           configID,
			GuildConfigSalesTracker: model.GuildConfigSalesTracker{ID: configID, GuildID: guildID, ChannelID: guildChannelID},
			Name:                    name,
			Chain:                   model.Chain{ID: chainID, Name: chainName, ShortName: chainShortName, CoinGeckoID: coinGeckoID, Currency: currency},
		})
	}
	return resp, nil
}

func (pg *pg) DeleteNFTSalesTrackerByContractAddress(contractAddress string) error {
	return pg.db.Where("contract_address = ?", contractAddress).Delete(&model.NFTSalesTracker{}).Error
}

func (pg *pg) GetStarTrackerByGuildID(guildId string) (*model.NFTSalesTracker, error) {
	data := model.NFTSalesTracker{}
	err := pg.db.Table("nft_sales_trackers").Preload("GuildConfigSalesTracker", "guild_id=?", guildId).Where("contract_address='*'").First(&data)
	return &data, err.Error
}
