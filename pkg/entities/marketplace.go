package entities

func (e *Entity) HandleMarketplaceLink(contractAddress, chain string) string {
	return e.marketplace.HandleMarketplaceLink(contractAddress, chain)
}
