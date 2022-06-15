package nft

import "github.com/defipod/mochi/pkg/response"

type Service interface {
	GetNFTIndexer(int) ([]response.IndexerNFTResponse, error)
}
