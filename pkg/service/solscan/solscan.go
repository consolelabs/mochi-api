package solscan

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/defipod/mochi/pkg/config"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/util"
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

var solscanBaseURL = "https://api.solscan.io/collection/id"
var publicSolscanBaseURL = "https://public-api.solscan.io"

func (s *solscan) GetSolanaCollection(collectionId string) (*model.SolanaCollectionMetadata, error) {
	var client = &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("%s?collectionId=%s", solscanBaseURL, collectionId), nil)
	if err != nil {
		return nil, err
	}

	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	resBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	res := &model.SolanaCollectionMetadata{}
	err = json.Unmarshal(resBody, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *solscan) GetTransactions(address string) ([]TransactionListItem, error) {
	var res []TransactionListItem
	url := fmt.Sprintf("%s/account/transactions?account=%s&limit=5", publicSolscanBaseURL, address)
	statusCode, err := util.FetchData(url, &res)
	if err != nil {
		s.logger.Fields(logger.Fields{"url": url, "status": statusCode}).Error(err, "[solscan.getTransactions] util.FetchData() failed")
		return nil, err
	}
	return res, nil
}

func (s *solscan) GetTokenBalances(address string) ([]TokenAmountItem, error) {
	var res []TokenAmountItem
	url := fmt.Sprintf("%s/account/tokens?account=%s", publicSolscanBaseURL, address)
	statusCode, err := util.FetchData(url, &res)
	if err != nil {
		s.logger.Fields(logger.Fields{"url": url, "status": statusCode}).Error(err, "[solscan.getTokenBalances] util.FetchData() failed")
		return nil, err
	}
	return res, nil
}

func (s *solscan) GetTokenMetadata(tokenAddress string) (*TokenMetadataResponse, error) {
	res := &TokenMetadataResponse{}
	url := fmt.Sprintf("%s/token/meta?tokenAddress=%s", publicSolscanBaseURL, tokenAddress)
	statusCode, err := util.FetchData(url, res)
	if err != nil {
		s.logger.Fields(logger.Fields{"url": url, "status": statusCode}).Error(err, "[solscan.GetTokenMetadata] util.FetchData() failed")
		return nil, err
	}
	return res, nil
}

func (s *solscan) GetTxDetails(signature string) (*TransactionDetailsResponse, error) {
	res := &TransactionDetailsResponse{}
	url := fmt.Sprintf("%s/transaction/%s", publicSolscanBaseURL, signature)
	statusCode, err := util.FetchData(url, res)
	if err != nil {
		s.logger.Fields(logger.Fields{"url": url, "status": statusCode}).Error(err, "[solscan.GetTxDetails] util.FetchData() failed")
		return nil, err
	}
	return res, nil
}
