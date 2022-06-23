package nft_sales_tracker

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	FirstOrCreate(*model.InsertNFTSalesTracker) error
	GetAll() ([]model.NFTSalesTracker, error)
}
