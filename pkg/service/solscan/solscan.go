package solscan

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/defipod/mochi/pkg/config"
	"github.com/defipod/mochi/pkg/logger"
	resp "github.com/defipod/mochi/pkg/response"
)

type solscan struct {
	config *config.Config
	logger logger.Logger
}

func NewService(cfg *config.Config, l logger.Logger) Service {
	return &solscan{
		config: cfg,
		logger: l,
	}
}

var (
	publicSolscanBaseURL = "https://public-api.solscan.io"
	proSolscanBaseUrl    = "https://pro-api.solscan.io/v1.0"
)

func (s *solscan) GetCollectionBySolscanId(id string) (*resp.CollectionDataResponse, error) {
	var client = &http.Client{}
	request, err := http.NewRequest("GET", fmt.Sprintf("%s/nft/collection/list?search=%s", proSolscanBaseUrl, id), nil)
	if err != nil {
		return nil, err
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("token", s.config.Solscan.Token)

	response, err := client.Do(request)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	resBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	res := &resp.CollectionDataResponse{}
	err = json.Unmarshal(resBody, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *solscan) GetNftTokenFromCollection(id, page string) (*resp.NftTokenDataResponse, error) {
	var client = &http.Client{}
	request, err := http.NewRequest("GET", fmt.Sprintf("%s/nft/collection/list_nft/%s?page=%s", proSolscanBaseUrl, id, page), nil)
	if err != nil {
		return nil, err
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("token", s.config.Solscan.Token)

	response, err := client.Do(request)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	resBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	res := &resp.NftTokenDataResponse{}
	err = json.Unmarshal(resBody, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *solscan) fetchSolscanData(url string, v any) error {
	var client = &http.Client{}
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("token", s.config.Solscan.Token)
	res, err := client.Do(request)
	if err != nil {
		return err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, v)
	if err != nil {
		return err
	}
	return err
}

func (s *solscan) GetTransactions(address string) ([]TransactionListItem, error) {
	var res []TransactionListItem
	url := fmt.Sprintf("%s/account/transactions?account=%s&limit=5", publicSolscanBaseURL, address)
	err := s.fetchSolscanData(url, &res)
	if err != nil {
		s.logger.Fields(logger.Fields{"url": url}).Error(err, "[solscan.getTransactions] s.fetchSolscanData() failed")
		return nil, err
	}
	return res, nil
}

func (s *solscan) GetTokenBalances(address string) ([]TokenAmountItem, error) {
	var res []TokenAmountItem
	url := fmt.Sprintf("%s/account/tokens?account=%s", publicSolscanBaseURL, address)
	err := s.fetchSolscanData(url, &res)
	if err != nil {
		s.logger.Fields(logger.Fields{"url": url}).Error(err, "[solscan.getTokenBalances] s.fetchSolscanData() failed")
		return nil, err
	}
	return res, nil
}

func (s *solscan) GetTokenMetadata(tokenAddress string) (*TokenMetadataResponse, error) {
	res := &TokenMetadataResponse{}
	url := fmt.Sprintf("%s/token/meta?tokenAddress=%s", publicSolscanBaseURL, tokenAddress)
	err := s.fetchSolscanData(url, &res)
	if err != nil {
		s.logger.Fields(logger.Fields{"url": url}).Error(err, "[solscan.GetTokenMetadata] s.fetchSolscanData() failed")
		return nil, err
	}
	return res, nil
}

func (s *solscan) GetTxDetails(signature string) (*TransactionDetailsResponse, error) {
	res := &TransactionDetailsResponse{}
	url := fmt.Sprintf("%s/transaction/%s", publicSolscanBaseURL, signature)
	err := s.fetchSolscanData(url, &res)
	if err != nil {
		s.logger.Fields(logger.Fields{"url": url}).Error(err, "[solscan.GetTxDetails] s.fetchSolscanData() failed")
		return nil, err
	}
	return res, nil
}
