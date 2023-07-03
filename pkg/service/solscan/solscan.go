package solscan

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/defipod/mochi/pkg/cache"
	"github.com/defipod/mochi/pkg/config"
	"github.com/defipod/mochi/pkg/logger"
	resp "github.com/defipod/mochi/pkg/response"
)

type solscan struct {
	config *config.Config
	logger logger.Logger
	cache  cache.Cache
}

func NewService(cfg *config.Config, l logger.Logger, cache cache.Cache) Service {
	return &solscan{
		config: cfg,
		logger: l,
		cache:  cache,
	}
}

var (
	publicSolscanBaseURL        = "https://public-api.solscan.io"
	proSolscanBaseUrl           = "https://pro-api.solscan.io/v1.0"
	solscanTransactionKey       = "solscan-transaction"
	solscanTransactionDetailKey = "solscan-transaction-detail"
	solscanTokenMetadataKey     = "solscan-token-metadata"
	solscanTokenBalanceKey      = "solscan-token-balance"
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
	s.logger.Debug("start Solscan.GetTransactions()")
	defer s.logger.Debug("end Solscan.GetTransactions()")

	var res []TransactionListItem
	// check if data cached

	cached, err := s.doCacheTransaction(address)
	if err == nil && cached != "" {
		s.logger.Infof("hit cache data solscan-service, address: %s", address)
		go s.doNetworkTransaction(address)
		return res, json.Unmarshal([]byte(cached), &res)
	}

	// call network
	go s.doNetworkTransaction(address)
	return nil, nil
}

func (s *solscan) GetTokenBalances(address string) ([]TokenAmountItem, error) {
	s.logger.Debug("start Solscan.GetTokenBalances()")
	defer s.logger.Debug("end Solscan.GetTokenBalances()")

	var res []TokenAmountItem
	// check if data cached

	cached, err := s.doCacheTokenBalance(address)
	if err == nil && cached != "" {
		s.logger.Infof("hit cache data solscan-service, tokenAddress: %s", address)
		go s.doNetworkTokenBalance(address)
		return res, json.Unmarshal([]byte(cached), &res)
	}

	// call network
	return s.doNetworkTokenBalance(address)
}

func (s *solscan) GetTokenMetadata(tokenAddress string) (*TokenMetadataResponse, error) {
	s.logger.Debug("start Solscan.GetTokenMetadata()")
	defer s.logger.Debug("end Solscan.GetTokenMetadata()")

	var res TokenMetadataResponse
	// check if data cached

	cached, err := s.doCacheTokenMetadata(tokenAddress)
	if err == nil && cached != "" {
		s.logger.Infof("hit cache data solscan-service, tokenAddress: %s", tokenAddress)
		go s.doNetworkTokenMetadata(tokenAddress)
		return &res, json.Unmarshal([]byte(cached), &res)
	}

	// call network
	return s.doNetworkTokenMetadata(tokenAddress)
}

func (s *solscan) GetTxDetails(signature string) (*TransactionDetailsResponse, error) {
	s.logger.Debug("start Solscan.GetTxDetails()")
	defer s.logger.Debug("end Solscan.GetTxDetails()")

	var res TransactionDetailsResponse
	// check if data cached

	cached, err := s.doCacheTransactionDetail(signature)
	if err == nil && cached != "" {
		s.logger.Infof("hit cache data solscan-service, signature: %s", signature)
		go s.doNetworkTransactionDetail(signature)
		return &res, json.Unmarshal([]byte(cached), &res)
	}

	// call network
	go s.doNetworkTransactionDetail(signature)
	return nil, nil
}
