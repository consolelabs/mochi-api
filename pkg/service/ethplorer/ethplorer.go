package ethplorer

import (
	"fmt"

	"github.com/defipod/mochi/pkg/util"
)

const (
	baseUrl = "https://api.ethplorer.io"
	apiKey  = "freekey"
)

type ethplorer struct {
}

func NewService() Service {
	return &ethplorer{}
}

func (e *ethplorer) GetTopTokenHolders(address string, limit int) (*TokenHoldersResponse, error) {
	resp := &TokenHoldersResponse{}

	url := fmt.Sprintf("%s/getTopTokenHolders/%s?apiKey=%s&limit=%d", baseUrl, address, apiKey, limit)

	status, err := util.FetchData(url, resp)
	if err != nil {
		return nil, fmt.Errorf("fetch data failed: %w", err)
	}

	if status != 200 {
		return nil, fmt.Errorf("fetch data failed: status %d", status)
	}

	return resp, nil
}

func (e *ethplorer) GetTokenInfo(address string) (*TokenInfoResponse, error) {
	resp := &TokenInfoResponse{}

	url := fmt.Sprintf("%s/getTokenInfo/%s?apiKey=%s", baseUrl, address, apiKey)

	status, err := util.FetchData(url, resp)
	if err != nil {
		return nil, fmt.Errorf("fetch data failed: %w", err)
	}

	if status != 200 {
		return nil, fmt.Errorf("fetch data failed: status %d", status)
	}

	return resp, nil
}
