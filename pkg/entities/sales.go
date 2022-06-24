package entities

import (
	"errors"
	"fmt"

	"github.com/defipod/mochi/pkg/response"
)

func (e *Entity) GetNftSales(addr string, platform string) (*response.NftSales, error) {
	sales, err := e.indexer.GetNftSales()
	if err != nil {
		err = fmt.Errorf("failed to get sales from indexer: %v", err)
		return nil, err
	}

	nft := response.NftSales{}
	for _, ele := range sales.Data {
		if ele.NftCollectionAddress == addr && ele.Platform == platform {
			nft = ele
		}
	}
	if nft == (response.NftSales{}) {
		return nil, errors.New("collection not found")
	}
	return &nft, nil
}
