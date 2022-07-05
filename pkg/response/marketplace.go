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
