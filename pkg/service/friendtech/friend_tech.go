package friendtech

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/defipod/mochi/pkg/config"
	"github.com/defipod/mochi/pkg/consts"
	"github.com/defipod/mochi/pkg/response"
	"github.com/defipod/mochi/pkg/util"
)

type FriendTech struct {
	baseUrl string
}

var commonHeader = map[string]string{
	"Content-Type": "application/json",
	"Accept":       "application/json",
	"clientData":   fmt.Sprintf("{\"source\": \"%s\"}", consts.ClientID),
	"x-client-id":  consts.ClientID,
}

func NewService(cfg *config.Config) Service {
	baseUrl := cfg.FriendScanAPI
	if baseUrl == "" {
		baseUrl = "https://api.friendscan.tech"
	}

	return &FriendTech{
		baseUrl: baseUrl,
	}
}

func (n *FriendTech) Search(query string, limit int) (*response.FriendTechKeysResponse, error) {
	if limit == 0 || limit > 200 {
		limit = 200
	}

	url := n.baseUrl + fmt.Sprintf("/accounts?q=%v&limit=%v", query, limit)
	data := response.FriendTechKeysResponse{}
	req := util.SendRequestQuery{
		URL:       url,
		ParseForm: &data,
		Headers:   commonHeader,
	}
	statusCode, err := util.SendRequest(req)
	if err != nil || statusCode != http.StatusOK {
		return &response.FriendTechKeysResponse{}, nil
	}
	return &data, nil
}

func (n *FriendTech) GetHistory(accountAddress, interval string) (*response.FriendTechKeyPriceHistoryResponse, error) {
	url := n.baseUrl + fmt.Sprintf("/accounts/%v/historical?interval=%v", accountAddress, interval)
	data := response.FriendTechKeyPriceHistoryResponse{}
	req := util.SendRequestQuery{
		URL:       url,
		ParseForm: &data,
		Headers:   commonHeader,
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

	url := n.baseUrl + fmt.Sprintf("/transactions?subjectAddress=%s&limit=%v", subjectAddress, limit)
	data := response.FriendTechKeyTransactionsResponse{}
	req := util.SendRequestQuery{
		URL:       url,
		ParseForm: &data,
		Headers:   commonHeader,
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
