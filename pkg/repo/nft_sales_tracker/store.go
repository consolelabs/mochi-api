package nft_sales_tracker

import (
	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/response"
)

type Store interface {
	FirstOrCreate(*model.InsertNFTSalesTracker) error
	GetAll() ([]model.NFTSalesTracker, error)
	GetSalesTrackerByGuildID(guildId string) ([]response.NFTSalesTrackerData, error)
	GetStarTrackerByGuildID(guildId string) (*model.NFTSalesTracker, error)
	DeleteNFTSalesTrackerByContractAddress(contractAddress string) error
}
