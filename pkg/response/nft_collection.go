package response

import (
	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/util"
)

type NFTTradingVolume struct {
	CollectionAddress string  `json:"collection_address"`
	CollectionName    string  `json:"collection_name"`
	Symbol            string  `json:"collection_symbol"`
	ChainID           float64 `json:"collection_chain_id"`
	TradingVolume     float64 `json:"trading_volume"`
	Token             string  `json:"token"`
}

type NFTTradingVolumeResponse struct {
	Data []NFTTradingVolume `json:"data"`
}

type NFTNewListed struct {
	Metadata util.Pagination                `json:"metadata"`
	Data     []model.NewListedNFTCollection `json:"data"`
}

type NFTNewListedResponse struct {
	Data *NFTNewListed `json:"data"`
}

type NFTCollectionsResponse struct {
	Data NFTCollectionsData `json:"data"`
}

type NFTCollectionsData struct {
	Metadata util.Pagination       `json:"metadata"`
	Data     []model.NFTCollection `json:"data"`
}

type NFTCollectionCount struct {
	Total int                       `json:"total"`
	Data  []NFTChainCollectionCount `json:"data"`
	// ETHCount int `json:"eth_collection"`
	// FTMCount int `json:"ftm_collection"`
	// OPCount  int `json:"op_collection"`
}

type NFTChainCollectionCount struct {
	Chain model.Chain `json:"chain"`
	Count int         `json:"count"`
}
