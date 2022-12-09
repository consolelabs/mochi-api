package response

type OpenseaPrimaryAssetContracts struct {
	Address string `json:"address"`
}
type OpenseaCollection struct {
	Editors               interface{}                    `json:"editors"`
	PaymentTokens         interface{}                    `json:"payment_tokens"`
	PrimaryAssetContracts []OpenseaPrimaryAssetContracts `json:"primary_asset_contracts"`
	Traits                interface{}                    `json:"traits"`
	Stats                 OpenseaCollectionStat          `json:"stats"`
}

type OpenseaCollectionStat struct {
	FloorPrice   float64 `json:"floor_price"`
	AveragePrice float64 `json:"one_day_average_price"`
}

type OpenseaGetCollectionResponse struct {
	Collection OpenseaCollection `json:"collection"`
}

type QuixoticCollectionResponse struct {
	Address         string      `json:"address"`
	Name            string      `json:"name"`
	Symbol          string      `json:"symbol"`
	ContractType    string      `json:"contract_type"`
	ExternalLink    string      `json:"external_link"`
	Description     string      `json:"description"`
	Slug            string      `json:"slug"`
	ImageUrl        string      `json:"image_url"`
	BannerImageUrl  string      `json:"banner_image_url"`
	RoyaltyPerMille int64       `json:"royalty_per_mille"`
	PayoutAddress   string      `json:"payout_address"`
	Verified        bool        `json:"verified"`
	Attributes      interface{} `json:"attributes"`
	Owner           interface{} `json:"owner"`
	FloorPrice      int64       `json:"floor_price"`
	VolumeTraded    int64       `json:"volume_traded"`
}

type OpenseaAssetContractResponse struct {
	Collection OpenseaAssetContract `json:"collection"`
	Address    string               `json:"address"`
}

type OpenseaAssetContract struct {
	// add more fields if needed
	Image   string `json:"image_url"`
	UrlName string `json:"slug"`
}

type PaintswapCollectionResponse struct {
	Collection PaintswapCollection `json:"collection"`
}
type PaintswapCollection struct {
	// add more fields if needed
	Image string                  `json:"poster"`
	Stats PaintswapCollectionStat `json:"stats"`
}

type PaintswapCollectionStat struct {
	FloorPrice   string `json:"floor"`
	AveragePrice string `json:"averagePrice"`
}

type AlchemyContractMetaData struct {
	Name                string `json:"name"`
	Symbol              string `json:"symbol"`
	TotalSupply         string `json:"totalSupply"`
	TokenType           string `json:"tokenType"`
	ContractDeployer    string `json:"contractDeployer"`
	DeployedBlockNumber int64  `json:"deployedBlockNumber"`
}

type AlchemyCollectionResponse struct {
	Address          string                  `json:"address"`
	ContractMetadata AlchemyContractMetaData `json:"contractMetadata"`
}
