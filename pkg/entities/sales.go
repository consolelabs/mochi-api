package entities

import (
	"fmt"

	"github.com/defipod/mochi/pkg/response"
)

func (e *Entity) GetNftSales(addr string, platform string) (*response.NftSalesResponse, error) {
	nft, err := e.indexer.GetNftSales(addr, platform)
	if err != nil {
		err = fmt.Errorf("failed to get sales from indexer: %v", err)
		return nil, err
	}

	return nft, nil
}
