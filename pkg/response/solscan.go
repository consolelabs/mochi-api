package response

type CollectionDataResponse struct {
	Data   CollectionData `json:"data"`
	Code   int64          `json:"code"`
	Status int64          `json:"status"`
}

type CollectionData struct {
	Total       string           `json:"total"`
	Page        int64            `json:"page"`
	Collections []CollectionInfo `json:"collections"`
}

type CollectionInfo struct {
	NftCollectionUrl string  `json:"nft_collection_url"`
	CollectionId     string  `json:"collection_id"`
	CollectionName   string  `json:"collection_name"`
	NumNfts          string  `json:"num_nfts"`
	FloorPrice       float64 `json:"floor_price"`
	Score            int64   `json:"score"`
}

type NftTokenDataResponse struct {
	Data   NftTokenData `json:"data"`
	Code   int64        `json:"code"`
	Status int64        `json:"status"`
}

type NftTokenData struct {
	Page     interface{}    `json:"page"`
	ListNfts []NftTokenInfo `json:"list_nfts"`
}

type NftTokenInfo struct {
	NftAddress             string             `json:"nft_address"`
	NftName                string             `json:"nft_name"`
	NftSymbol              string             `json:"nft_symbol"`
	NftImage               string             `json:"nft_image"`
	NftMintedTime          string             `json:"nft_minted_time"`
	NftMintedTransactionId string             `json:"nft_minted_transaction_id"`
	NftCollectionId        string             `json:"nft_collection_id"`
	NftCollectionName      string             `json:"nft_collection_name"`
	NftAttributes          NftTokenAttributes `json:"nft_attributes"`
	NftCreators            []NftCreator       `json:"nft_creators"`
}

type NftCreator struct {
	Address  string `json:"address"`
	Verified int64  `json:"verified"`
	Share    int64  `json:"share"`
}

type NftTokenAttributes struct {
	Attributes []AttributesSolscan `json:"attributes"`
	Properties Properties          `json:"properties"`
}

type AttributesSolscan struct {
	TraitType string      `json:"trait_type"`
	Value     interface{} `json:"value"`
}

type Properties struct {
	Files    []File              `json:"files"`
	Category string              `json:"category"`
	Creators []PropertiesCreator `json:"creators"`
}

type PropertiesCreator struct {
	Address string `json:"address"`
	Share   int64  `json:"share"`
}

type File struct {
	Url  string `json:"url"`
	Uri  string `json:"uri"`
	Type string `json:"type"`
}
