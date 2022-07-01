package response

import (
	"time"

	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/util"
)

type IndexerNFTCollectionTickersResponse struct {
	Tickers         TokenTickers `json:"tickers"`
	FloorPrice      float64      `json:"floor_price"`
	Name            string       `json:"name"`
	ContractAddress string       `json:"contract_address"`
	Chain           string       `json:"chain"`
	Platforms       []string     `json:"platforms"`
}

type IndexerGetNFTCollectionsResponse struct {
	util.Pagination
	Data []IndexerNFTCollection `json:"data"`
}

type IndexerNFTCollection struct {
	Address            string       `json:"address"`
	Name               string       `json:"name"`
	Symbol             string       `json:"symbol"`
	ChainId            int          `json:"chain_id"`
	Chain              *model.Chain `json:"chain,omitempty"`
	ERCFormat          string       `json:"erc_format"`
	Supply             uint64       `json:"supply"`
	IsRarityCalculated bool         `json:"is_rarity_calculated"`
	Image              string       `json:"image"`
}

type IndexerGetNFTTokensResponse struct {
	util.Pagination
	Data []IndexerNFTToken `json:"data"`
}

type IndexerNFTToken struct {
	TokenID           string                     `json:"token_id,omitempty"`
	CollectionAddress string                     `json:"collection_address,omitempty"`
	Name              string                     `json:"name,omitempty"`
	Description       string                     `json:"description,omitempty"`
	Amount            string                     `json:"amount,omitempty"`
	Image             string                     `json:"image,omitempty"`
	ImageCDN          string                     `json:"image_cdn,omitempty"`
	ThumbnailCDN      string                     `json:"thumbnail_cdn,omitempty"`
	ImageContentType  string                     `json:"image_content_type"`
	RarityRank        uint64                     `json:"rarity_rank"`
	RarityScore       string                     `json:"rarity_score,omitempty"`
	RarityTier        string                     `json:"rarity_tier"`
	Attributes        []IndexerNFTTokenAttribute `json:"attributes" gorm:"-"`
	Rarity            *IndexerNFTTokenRarity     `json:"rarity"`
	MetadataID        string                     `json:"metadata_id"`
}

type IndexerNFTTokenAttribute struct {
	CollectionAddress string `json:"collection_address"`
	TokenId           string `json:"token_id"`
	TraitType         string `json:"trait_type"`
	Value             string `json:"value"`
	Count             uint64 `json:"count"`
	Rarity            string `json:"rarity"`
	Frequency         string `json:"frequency"`
}

type IndexerNFTTokenRarity struct {
	Rank   uint64 `json:"rank"`
	Score  string `json:"score"`
	Total  uint64 `json:"total"`
	Rarity string `json:"rarity,omitempty"`
}

type IndexerAttribute struct {
	TraitType string `json:"trait_type"`
	Value     string `json:"value"`
	Count     int    `json:"count"`
	Rarity    string `json:"rarity"`
	Frequency string `json:"frequency"`
}

type IndexerContract struct {
	ID              int       `json:"id"`
	LastUpdateTime  time.Time `json:"last_updated_time"`
	LastUpdateBlock int       `json:"last_updated_block"`
	CreationBlock   int       `json:"creation_block"`
	CreatedTime     time.Time `json:"created_time"`
	Address         string    `json:"address"`
	ChainID         int       `json:"chain_id"`
	Type            string    `json:"Type"`
	IsProxy         bool      `json:"is_proxy"`
	LogicAddress    string    `json:"logic_address"`
	Protocol        string    `json:"Protocol"`
	GRPCAddress     string    `json:"GrpcAddress"`
	IsSynced        bool      `json:"is_synced"`
}
