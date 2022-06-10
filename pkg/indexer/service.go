package indexer

type Indexer interface {
	CreateERC721Contract(CreateERC721ContractRequest) error
}
