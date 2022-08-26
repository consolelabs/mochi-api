package nft_sales_tracker

import (
	"github.com/defipod/mochi/pkg/model"
)

type Store interface {
	FirstOrCreate(*model.InsertNFTSalesTracker) error
	GetAll() ([]model.NFTSalesTracker, error)
	GetSalesTrackerByGuildID(guildId string) ([]model.NFTSalesTracker, error)
	GetNFTSalesTrackerByContractAndGuildID(guildID, contractAddress string) (*model.NFTSalesTracker, error)
	DeleteNFTSalesTrackerByContractAddress(contractAddress string) error
}
