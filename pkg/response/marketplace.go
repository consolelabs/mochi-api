package response

type OpenseaPrimaryAssetContracts struct {
	Address string `json:"address"`
}
type OpenseaCollection struct {
	Editors               interface{}                    `json:"editors"`
	PaymentTokens         interface{}                    `json:"payment_tokens"`
	PrimaryAssetContracts []OpenseaPrimaryAssetContracts `json:"primary_asset_contracts"`
	Traits                interface{}                    `json:"traits"`
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
