package solscan

import "github.com/defipod/mochi/pkg/model"

type Service interface {
	GetSolanaCollection(address string) (*model.SolanaCollectionMetadata, error)
}
