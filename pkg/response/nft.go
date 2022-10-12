package response

import (
	"time"

	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/util"
	"github.com/google/uuid"
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
	PriceChange1d   string          `json:"price_change_1d"`
	PriceChange7d   string          `json:"price_change_7d"`
	PriceChange30d  string          `json:"price_change_30d"`
}

type IndexerGetNFTTokenTickersResponse struct {
	Data IndexerNFTTokenTickersData `json:"data"`
}

type IndexerNFTTokenTickersData struct {
	Tickers           *IndexerTickers `json:"tickers"`
	Name              string          `json:"name"`
	TokenId           string          `json:"token_id"`
	CollectionAddress string          `json:"collection_address"`
	Description       string          `json:"description"`
	Image             string          `json:"image"`
	ImageCDN          string          `json:"image_cdn"`
	RarityRank        uint64          `json:"rarity_rank"`
	RarityScore       string          `json:"rarity_score"`
	RarityTier        string          `json:"rarity_tier"`
	FloorPrice        *IndexerPrice   `json:"floor_price"`
	LastSalePrice     *IndexerPrice   `json:"last_sale_price"`
	PriceChange1d     string          `json:"price_change_1d"`
	PriceChange7d     string          `json:"price_change_7d"`
	PriceChange30d    string          `json:"price_change_30d"`
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
	Data          *IndexerNFTTokenDetailData `json:"data"`
	Suggestions   []CollectionSuggestions    `json:"suggestions"`
	DefaultSymbol *CollectionSuggestions     `json:"default_symbol"`
}

type CollectionSuggestions struct {
	Name    string `json:"name"`
	Symbol  string `json:"symbol"`
	Address string `json:"address"`
	Chain   string `json:"chain"`
	ChainId int64  `json:"chain_id"`
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
	Owner             IndexerNftTokenOwner       `json:"owner"`
	Marketplace       []NftListingMarketplace    `json:"marketplace"`
}

type NftListingMarketplace struct {
	ContractAddress      string `json:"contract_address"`
	TokenId              string `json:"token_id"`
	PlatformId           uint64 `json:"platform_id"`
	PlatformName         string `json:"platform_name"`
	ListingStatus        string `json:"listing_status"`
	ListingPrice         string `json:"listing_price"`
	PaymentToken         string `json:"payment_token"`
	PaymentTokenDecimals string `json:"payment_token_decimals"`
	ItemUrl              string `json:"item_url"`
	FloorPrice           string `json:"floor_price"`
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

type IndexerNftTokenOwner struct {
	OwnerAddress      string `json:"owner_address"`
	CollectionAddress string `json:"collection_address"`
	TokenId           string `json:"token_id"`
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

type IndexerGetNFTActivityResponse struct {
	Data []IndexerNFTActivityData `json:"data"`
	util.Pagination
}

type IndexerGetNFTTokenTxHistoryResponse struct {
	Data []IndexerGetNftTokenTxHistory `json:"data"`
}

type IndexerGetNftTokenTxHistory struct {
	From            string    `json:"from"`
	To              string    `json:"to"`
	TokenId         string    `json:"token_id"`
	TransactionHash string    `json:"transaction_hash"`
	ContractAddress string    `json:"contract_address"`
	CreatedTime     time.Time `json:"created_time"`
	ListingStatus   string    `json:"listing_status"`
	EventType       string    `json:"event_type"`
}

type IndexerNFTActivityData struct {
	ID              int          `json:"id,omitempty"`
	PlatformID      int          `json:"platform_id,omitempty"`
	TokenID         string       `json:"token_id,omitempty"`
	ContractAddress string       `json:"contract_address,omitempty"`
	ChainID         int          `json:"chain_id,omitempty"`
	Quantity        string       `json:"quantity,omitempty"`
	PaymentToken    int          `json:"payment_token,omitempty"`
	FromAddress     string       `json:"from_address,omitempty"`
	ToAddress       string       `json:"to_address,omitempty"`
	TransactionHash string       `json:"transaction_hash,omitempty"`
	ListingType     string       `json:"listing_type,omitempty"`
	ListingStatus   string       `json:"listing_status,omitempty"`
	CreatedTime     time.Time    `json:"created_time,omitempty"`
	LastUpdateTime  time.Time    `json:"last_update_time,omitempty"`
	SoldPrice       string       `json:"sold_price,omitempty"`
	ListingPrice    string       `json:"listing_price,omitempty"`
	SoldPriceObj    IndexerPrice `json:"sold_price_obj,omitempty"`
	ListingPriceObj IndexerPrice `json:"listing_price_obj,omitempty"`
	EventType       string       `json:"event_type,omitempty"`
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

type GetNFTCollectionByAddressChain struct {
	ID           uuid.NullUUID `json:"id" gorm:"default:uuid_generate_v4()" swaggertype:"string"`
	Address      string        `json:"address"`
	Name         string        `json:"name"`
	Symbol       string        `json:"symbol"`
	ChainID      string        `json:"chain_id"`
	ERCFormat    string        `json:"erc_format"`
	IsVerified   bool          `json:"is_verified"`
	CreatedAt    time.Time     `json:"created_at"`
	Image        string        `json:"image"`
	Author       string        `json:"author"`
	Description  string        `json:"description"`
	Discord      string        `json:"discord"`
	Twitter      string        `json:"twitter"`
	Website      string        `json:"website"`
	Marketplaces []string      `json:"marketplaces"`
}

type GetNFTCollectionByAddressChainResponse struct {
	Data *GetNFTCollectionByAddressChain `json:"data"`
}

type GetNFTActivityResponse struct {
	Data GetNFTActivityData `json:"data"`
}

type GetNFTActivityData struct {
	Data     []IndexerNFTActivityData `json:"data"`
	Metadata util.Pagination          `json:"metadata"`
}

type NftWatchlistSuggestResponse struct {
	Data *NftWatchlistSuggest `json:"data"`
}

type NftWatchlistSuggest struct {
	Suggestions   []CollectionSuggestions `json:"suggestions"`
	DefaultSymbol *CollectionSuggestions  `json:"default_symbol"`
}

type GetNftWatchlist struct {
	FloorPrice                        float64       `json:"floor_price"`
	Symbol                            string        `json:"symbol"`
	Image                             string        `json:"image"`
	Id                                string        `json:"id"`
	IsPair                            bool          `json:"is_pair"`
	Name                              string        `json:"name"`
	PriceChangePercentage24h          float64       `json:"price_change_percentage_24h"`
	PriceChangePercentage7dInCurrency float64       `json:"price_change_percentage_7d_in_currency"`
	SparkLineIn7d                     SparkLineIn7d `json:"sparkline_in_7d"`
	Token                             IndexerToken  `json:"token"`
}

type SparkLineIn7d struct {
	Price []float64 `json:"price"`
}

type GetNftWatchlistResponse struct {
	Data []GetNftWatchlist `json:"data"`
}

type GetGuildDefaultNftTickerResponse struct {
	Data *model.GuildConfigDefaultCollection `json:"data"`
}

type GetSuggestionNFTCollectionsResponse struct {
	Data []CollectionSuggestions `json:"data"`
}
type IndexerNftCollectionMetadata struct {
	Id                     int64                   `json:"id"`
	CollectionAddress      string                  `json:"collection_address"`
	Name                   string                  `json:"name"`
	Symbol                 string                  `json:"symbol"`
	ChainID                int64                   `json:"chain_id"`
	ERCFormat              string                  `json:"erc_format"`
	Supply                 int64                   `json:"supply"`
	IsRarityCalculated     bool                    `json:"is_rarity_calculated"`
	Image                  string                  `json:"image"`
	Description            string                  `json:"description"`
	ContractScan           string                  `json:"contract_scan"`
	Discord                string                  `json:"discord"`
	Twitter                string                  `json:"twitter"`
	Website                string                  `json:"website"`
	Owners                 int64                   `json:"owners"`
	MarketplaceCollections interface{}             `json:"marketplace_collections"`
	CreatedTime            string                  `json:"created_time"`
	LastUpdatedTime        string                  `json:"last_updated_time"`
	Chain                  interface{}             `json:"chain"`
	Stats                  interface{}             `json:"stats"`
	IsNotifySynced         bool                    `json:"is_notify_synced"`
	Marketplace            []NftListingMarketplace `json:"marketplace"`
}

type IndexerNftCollectionMetadataResponse struct {
	Data IndexerNftCollectionMetadata `json:"data"`
}

type IndexerErrorResponse struct {
	Data  interface{} `json:"data"`
	Error string      `json:"error"`
}
