package entities

import (
	"fmt"

	"github.com/defipod/mochi/pkg/response"
)

func (e *Entity) GetNftSales() (*response.NftSalesResponse, error) {
	sales, err := e.indexer.GetNftSales()
	if err != nil {
		err = fmt.Errorf("failed to get sales from indexer: %v", err)
		return nil, err
	}
	return sales, nil
}
