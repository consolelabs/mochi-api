package friendtech

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/defipod/mochi/pkg/config"
	"github.com/defipod/mochi/pkg/consts"
	"github.com/defipod/mochi/pkg/response"
	"github.com/defipod/mochi/pkg/service/sentrygo"
	"github.com/defipod/mochi/pkg/util"
)

type FriendTech struct {
	baseUrl string
	sentry  sentrygo.Service
}

var commonHeader = map[string]string{
	"Content-Type": "application/json",
	"Accept":       "application/json",
	"clientData":   fmt.Sprintf("{\"source\": \"%s\"}", consts.ClientID),
	"x-client-id":  consts.ClientID,
}

func NewService(cfg *config.Config, sentry sentrygo.Service) Service {
	baseUrl := cfg.FriendScanAPI
	if baseUrl == "" {
		baseUrl = "https://api.friendscan.io"
	}

	return &FriendTech{
		baseUrl: baseUrl,
		sentry:  sentry,
	}
}

var (
	sentryTags = map[string]string{
		"type": "system",
	}
)

func (n *FriendTech) Search(query string, limit int) (*response.FriendTechKeysResponse, error) {
	if limit == 0 || limit > 200 {
		limit = 200
	}

	url := n.baseUrl + fmt.Sprintf("/accounts?q=%v&limit=%v", query, limit)
	data := response.FriendTechKeysResponse{}
	req := util.SendRequestQuery{
		URL:      url,
		Response: &data,
		Headers:  commonHeader,
	}
	statusCode, err := util.SendRequest(req)
	if err != nil || statusCode != http.StatusOK {
		n.sentry.CaptureErrorEvent(sentrygo.SentryCapturePayload{
			Message: fmt.Sprintf("[API mochi] - FriendTech - Search failed %v", err),
			Tags:    sentryTags,
			Extra: map[string]interface{}{
				"url": url,
			},
		})
		return &response.FriendTechKeysResponse{}, nil
	}
	return &data, nil
}

func (n *FriendTech) GetHistory(accountAddress, interval string) (*response.FriendTechKeyPriceHistoryResponse, error) {
	url := n.baseUrl + fmt.Sprintf("/accounts/%v/historical?interval=%v", accountAddress, interval)
	data := response.FriendTechKeyPriceHistoryResponse{}
	req := util.SendRequestQuery{
		URL:      url,
		Response: &data,
		Headers:  commonHeader,
	}

	statusCode, err := util.SendRequest(req)
	if err != nil || statusCode != http.StatusOK {
		n.sentry.CaptureErrorEvent(sentrygo.SentryCapturePayload{
			Message: fmt.Sprintf("[API mochi] - FriendTech - GetHistory failed %v", err),
			Tags:    sentryTags,
			Extra: map[string]interface{}{
				"url": url,
			},
		})
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
		URL:      url,
		Response: &data,
		Headers:  commonHeader,
	}
	statusCode, err := util.SendRequest(req)
	if err != nil {
		n.sentry.CaptureErrorEvent(sentrygo.SentryCapturePayload{
			Message: fmt.Sprintf("[API mochi] - FriendTech - GetTransaction failed %v", err),
			Tags:    sentryTags,
			Extra: map[string]interface{}{
				"url": url,
			},
		})
		return &response.FriendTechKeyTransactionsResponse{}, err
	}

	if statusCode != http.StatusOK {
		n.sentry.CaptureErrorEvent(sentrygo.SentryCapturePayload{
			Message: fmt.Sprintf("[API mochi] - FriendTech - GetTransaction failed %v", err),
			Tags:    sentryTags,
			Extra: map[string]interface{}{
				"url": url,
			},
		})
		return &response.FriendTechKeyTransactionsResponse{}, errors.New("fetch status code is not 200")
	}

	return &data, nil
}
