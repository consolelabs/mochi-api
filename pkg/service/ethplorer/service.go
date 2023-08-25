package ethplorer

type Service interface {
	GetTokenInfo(address string) (*TokenInfoResponse, error)
	GetTopTokenHolders(address string, limit int) (*TokenHoldersResponse, error)
}
