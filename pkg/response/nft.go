package response

import (
	"time"

	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/util"
)

type IndexerPrice struct {
	Token  IndexerToken `json:"token"`
	Amount string       `json:"amount"`
}

type IndexerToken struct {
	Symbol   string `json:"symbol"`
	IsNative bool   `json:"is_native"`
	Address  string `json:"address"`
	Decimals int64  `json:"decimals"`
}

type IndexerTickers struct {
	Timestamps []int64        `json:"timestamps"`
	Prices     []IndexerPrice `json:"prices"`
	Times      []string       `json:"times"`
}

type IndexerChain struct {
	Name    string `json:"name"`
	Symbol  string `json:"symbol"`
	ChainId int64  `json:"chain_id"`
	IsEvm   bool   `json:"is_evm"`
}

type IndexerNFTCollectionTickersResponse struct {
	Data *IndexerNFTCollectionTickersData `json:"data"`
}

type IndexerNFTCollectionTickersData struct {
	Tickers         *IndexerTickers `json:"tickers"`
	Name            string          `json:"name"`
	Address         string          `json:"address"`
	Chain           *IndexerChain   `json:"chain"`
	Marketplaces    []string        `json:"marketplaces"`
	Items           int64           `json:"items"`
	Owners          int64           `json:"owners"`
	CollectionImage string          `json:"collection_image"`
	TotalVolume     *IndexerPrice   `json:"total_volume"`
	FloorPrice      *IndexerPrice   `json:"floor_price"`
	LastSalePrice   *IndexerPrice   `json:"last_sale_price"`
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
	Data []IndexerNFTTokenDetailData `json:"data"`
}

type IndexerGetNFTTokenDetailResponse struct {
	Data IndexerNFTTokenDetailData `json:"data"`
}
type IndexerGetNFTTokenDetailResponseWithSuggestions struct {
	Data          IndexerNFTTokenDetailData `json:"data"`
	Suggestions   []CollectionSuggestions   `json:"suggestions"`
	DefaultSymbol *CollectionSuggestions    `json:"default_symbol"`
}

type CollectionSuggestions struct {
	Name    string `json:"name"`
	Symbol  string `json:"symbol"`
	Address string `json:"address"`
	Chain   string `json:"chain"`
}

type IndexerNFTTokenDetailData struct {
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

type NftMetadataAttrIcon struct {
	ID          int    `json:"id"`
	Trait       string `json:"trait_type"`
	DiscordIcon string `json:"discord_icon"`
	UnicodeIcon string `json:"unicode_icon"`
}

type NftMetadataAttrIconResponse struct {
	Data []NftMetadataAttrIcon `json:"data"`
}

type GetGuildDefaultTickerResponse struct {
	Data *model.GuildConfigDefaultTicker `json:"data"`
}

type CreateNFTCollectionResponse struct {
	Data *model.NFTCollection `json:"data"`
}

type GetSupportedChains struct {
	Data []string `json:"data"`
}

type ListAllNFTCollectionsResponse struct {
	Data []model.NFTCollection `json:"data"`
}

type GetDetailNftCollectionResponse struct {
	Data *model.NFTCollectionDetail `json:"data"`
}

type GetAllNFTSalesTrackerResponse struct {
	Data []NFTSalesTrackerResponse `json:"data"`
}

type GetCollectionCountResponse struct {
	Data *NFTCollectionCount `json:"data"`
}

type GetNFTCollectionByAddressChainResponse struct {
	Data *model.NFTCollection `json:"data"`
}
