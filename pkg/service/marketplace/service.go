package marketplace

type Service interface {
	HandleMarketplaceLink(contractAddress, chain string) string
	GetCollectionFromOpensea(collectionSymbol string) (*openseaGetCollectionResponse, error)
}
