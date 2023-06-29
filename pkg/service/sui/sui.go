package sui

import (
	"encoding/json"

	"github.com/defipod/mochi/pkg/cache"
	"github.com/defipod/mochi/pkg/config"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/response"
)

type SuiService struct {
	config *config.Config
	logger logger.Logger
	cache  cache.Cache
}

func New(cfg *config.Config, l logger.Logger, cache cache.Cache) Service {
	return &SuiService{
		config: cfg,
		logger: l,
		cache:  cache,
	}
}

var (
	suiBalanceKey          = "sui-balance"
	suiCoinMetadataKey     = "sui-coin-metadata"
	suiAddressAssetsKey    = "sui-address-asset"
	suiTransactionBlockKey = "sui-transaction-block"
	suiAddressTxnKey       = "sui-address-txn"
)

func (s *SuiService) GetBalance(address string) (*response.SuiAllBalance, error) {
	var res response.SuiAllBalance
	cached, err := s.doCacheBalance(address)
	if err == nil && cached != "" {
		go s.doNetworkBalance(address)
		return &res, json.Unmarshal([]byte(cached), &res)
	}

	// call network
	return s.doNetworkBalance(address)
}

func (s *SuiService) GetCoinMetadata(coinType string) (*response.SuiCoinMetadata, error) {
	var res response.SuiCoinMetadata
	cached, err := s.doCacheCoinMetadata(coinType)
	if err == nil && cached != "" {
		go s.doNetworkCoinMetadata(coinType)
		return &res, json.Unmarshal([]byte(cached), &res)
	}

	// call network
	return s.doNetworkCoinMetadata(coinType)
}

func (s *SuiService) GetAddressAssets(address string) ([]response.WalletAssetData, error) {
	var res []response.WalletAssetData
	cached, err := s.doCacheAddressAssets(address)
	if err == nil && cached != "" {
		go s.doNetworkAddressAssets(address)
		return res, json.Unmarshal([]byte(cached), &res)
	}

	// call network
	return s.doNetworkAddressAssets(address)
}

func (s *SuiService) GetTransactionBlock(address string) (*response.SuiTransactionBlock, error) {
	var res response.SuiTransactionBlock
	cached, err := s.doCacheTransactionBlock(address)
	if err == nil && cached != "" {
		go s.doNetworkTransactionBlock(address)
		return &res, json.Unmarshal([]byte(cached), &res)
	}

	// call network
	return s.doNetworkTransactionBlock(address)
}

func (s *SuiService) GetAddressTxn(address string) ([]response.WalletTransactionData, error) {
	var res []response.WalletTransactionData
	cached, err := s.doCacheAddressTxn(address)
	if err == nil && cached != "" {
		go s.doNetworkAddressTxn(address)
		return res, json.Unmarshal([]byte(cached), &res)
	}

	// call network
	return s.doNetworkAddressTxn(address)
}
