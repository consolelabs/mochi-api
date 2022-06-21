package response

import "github.com/defipod/mochi/pkg/util"

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
	Address            string `json:"address"`
	Name               string `json:"name"`
	Symbol             string `json:"symbol"`
	ChainId            int    `json:"chain_id"`
	ERCFormat          string `json:"erc_format"`
	Supply             uint64 `json:"supply"`
	IsRarityCalculated bool   `json:"is_rarity_calculated"`
}

type IndexerGetNFTTokensResponse struct {
	util.Pagination
	Data []IndexerNFTToken `json:"data"`
}

type IndexerNFTToken struct {
	TokenID           uint64                     `json:"token_id,omitempty"`
	CollectionAddress string                     `json:"collection_address,omitempty"`
	Name              string                     `json:"name,omitempty"`
	Description       string                     `json:"description,omitempty"`
	Amount            uint64                     `json:"amount,omitempty"`
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
	TokenId           uint64 `json:"token_id"`
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
