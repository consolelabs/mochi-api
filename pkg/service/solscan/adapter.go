package solscan

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/defipod/mochi/pkg/logger"
)

func (s *solscan) doCacheTransaction(address string) (string, error) {
	return s.cache.GetString(fmt.Sprintf("%s-%s", solscanTransactionKey, strings.ToLower(address)))
}

func (s *solscan) doNetworkTransaction(address string) ([]TransactionListItem, error) {
	res := []TransactionListItem{}
	url := fmt.Sprintf("%s/account/transactions?account=%s&limit=5", publicSolscanBaseURL, address)
	err := s.fetchSolscanData(url, &res)
	if err != nil {
		s.logger.Fields(logger.Fields{"url": url}).Error(err, "[solscan.getTransactions] s.fetchSolscanData() failed")
		return nil, err
	}

	// cache solana-balance-token-data
	// if error occurs -> ignore
	bytes, _ := json.Marshal(&res)
	s.logger.Infof("cache data solscan-service, key: %s", solscanTransactionKey)
	s.cache.Set(fmt.Sprintf("%s-%s", solscanTransactionKey, strings.ToLower(address)), string(bytes), 7*24*time.Hour)

	return res, nil
}

func (s *solscan) doCacheTransactionDetail(signature string) (string, error) {
	return s.cache.GetString(fmt.Sprintf("%s-%s", solscanTransactionDetailKey, signature))
}

func (s *solscan) doNetworkTransactionDetail(signature string) (*TransactionDetailsResponse, error) {
	res := &TransactionDetailsResponse{}
	url := fmt.Sprintf("%s/transaction/%s", publicSolscanBaseURL, signature)
	err := s.fetchSolscanData(url, &res)
	if err != nil {
		s.logger.Fields(logger.Fields{"url": url}).Error(err, "[solscan.getTransactionDetail] s.fetchSolscanData() failed")
		return nil, err
	}

	// cache solana-balance-token-data
	// if error occurs -> ignore
	bytes, _ := json.Marshal(&res)
	s.logger.Infof("cache data solscan-service, key: %s", solscanTransactionDetailKey)
	s.cache.Set(fmt.Sprintf("%s-%s", solscanTransactionDetailKey, signature), string(bytes), 7*24*time.Hour)
	return res, nil
}

func (s *solscan) doCacheTokenMetadata(tokenAddress string) (string, error) {
	return s.cache.GetString(fmt.Sprintf("%s-%s", solscanTokenMetadataKey, tokenAddress))
}

func (s *solscan) doNetworkTokenMetadata(tokenAddress string) (*TokenMetadataResponse, error) {
	res := &TokenMetadataResponse{}
	url := fmt.Sprintf("%s/token/meta?tokenAddress=%s", publicSolscanBaseURL, tokenAddress)
	err := s.fetchSolscanData(url, &res)
	if err != nil {
		s.logger.Fields(logger.Fields{"url": url}).Error(err, "[solscan.getTokenMetadata] s.fetchSolscanData() failed")
		return nil, err
	}

	// cache solana-balance-token-data
	// if error occurs -> ignore
	bytes, _ := json.Marshal(&res)
	s.logger.Infof("cache data solscan-service, key: %s", solscanTokenMetadataKey)
	s.cache.Set(fmt.Sprintf("%s-%s", solscanTokenMetadataKey, tokenAddress), string(bytes), 7*24*time.Hour)
	return res, nil
}

func (s *solscan) doCacheTokenBalance(address string) (string, error) {
	return s.cache.GetString(fmt.Sprintf("%s-%s", solscanTokenBalanceKey, strings.ToLower(address)))
}

func (s *solscan) doNetworkTokenBalance(address string) ([]TokenAmountItem, error) {
	var res []TokenAmountItem
	url := fmt.Sprintf("%s/account/tokens?account=%s", publicSolscanBaseURL, address)
	err := s.fetchSolscanData(url, &res)
	if err != nil {
		s.logger.Fields(logger.Fields{"url": url}).Error(err, "[solscan.getTokenBalances] s.fetchSolscanData() failed")
		return nil, err
	}

	// cache solana-balance-token-data
	// if error occurs -> ignore
	bytes, _ := json.Marshal(&res)
	s.logger.Infof("cache data solscan-service, key: %s", solscanTokenBalanceKey)
	s.cache.Set(fmt.Sprintf("%s-%s", solscanTokenBalanceKey, strings.ToLower(address)), string(bytes), 7*24*time.Hour)
	return res, nil
}
