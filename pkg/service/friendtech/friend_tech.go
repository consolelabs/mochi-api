package friendtech

import (
	"errors"
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

func (n *FriendTech) GetTransactions(subjectAddress string, limit int) (*response.FriendTechKeyTransactionsResponse, error) {
	if limit == 0 || limit > 50 {
		limit = 50
	}

	url := n.baseUrl + fmt.Sprintf("/api/transactions?subjectAddress=%s&limit=%v", subjectAddress, limit)
	data := response.FriendTechKeyTransactionsResponse{}
	req := util.SendRequestQuery{
		URL:       url,
		ParseForm: &data,
		Headers:   map[string]string{"Content-Type": "application/json", "Accept": "application/json"},
	}
	statusCode, err := util.SendRequest(req)
	if err != nil {
		return &response.FriendTechKeyTransactionsResponse{}, err
	}

	if statusCode != http.StatusOK {
		return &response.FriendTechKeyTransactionsResponse{}, errors.New("fetch status code is not 200")
	}

	return &data, nil
}
