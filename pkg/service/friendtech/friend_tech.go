package friendtech

import (
	"fmt"
	"net/http"

	"github.com/defipod/mochi/pkg/response"
	"github.com/defipod/mochi/pkg/util"
)

type FriendTech struct {
	baseUrl string
}

func NewService() Service {
	return &FriendTech{
		baseUrl: "https://friendscan-api.caliber.build",
	}
}

func (n *FriendTech) Search(query string, limit int) (*response.FriendTechKeysResponse, error) {
	if limit == 0 || limit > 200 {
		limit = 200
	}

	url := n.baseUrl + fmt.Sprintf("/api/accounts?q=%v&limit=%v", query, limit)
	data := response.FriendTechKeysResponse{}
	req := util.SendRequestQuery{
		URL:       url,
		ParseForm: &data,
		Headers:   map[string]string{"Content-Type": "application/json", "Accept": "application/json"},
	}
	statusCode, err := util.SendRequest(req)
	if err != nil || statusCode != http.StatusOK {
		return &response.FriendTechKeysResponse{}, nil
	}
	return &data, nil
}

func (n *FriendTech) GetHistory(accountAddress, interval string) (*response.FriendTechKeyPriceHistoryResponse, error) {
	url := n.baseUrl + fmt.Sprintf("/api/accounts/%v/historical?interval=%v", accountAddress, interval)
	data := response.FriendTechKeyPriceHistoryResponse{}
	req := util.SendRequestQuery{
		URL:       url,
		ParseForm: &data,
		Headers:   map[string]string{"Content-Type": "application/json", "Accept": "application/json"},
	}
	statusCode, err := util.SendRequest(req)
	if err != nil || statusCode != http.StatusOK {
		return &response.FriendTechKeyPriceHistoryResponse{}, nil
	}

	return &data, nil
}
