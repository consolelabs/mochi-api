package skymavis

import (
	"encoding/json"

	"github.com/defipod/mochi/pkg/cache"
	"github.com/defipod/mochi/pkg/config"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/response"
)

type skymavis struct {
	cfg    *config.Config
	logger logger.Logger
	cache  cache.Cache
}

func New(cfg *config.Config, cache cache.Cache) Service {
	return &skymavis{
		cfg:    cfg,
		logger: logger.NewLogrusLogger(),
		cache:  cache,
	}
}

var (
	farmingKey     = "skymavis-farming"
	nftKey         = "skymavis-nft"
	internalTxsKey = "skymavis-internal-txs"
)

func (s *skymavis) GetAddressFarming(address string) (*response.WalletFarmingResponse, error) {
	s.logger.Debug("start skymavis.GetAddressFarming()")
	defer s.logger.Debug("end skymavis.GetAddressFarming()")

	var data response.WalletFarmingResponse
	// check if data cached

	cached, err := s.doCacheFarming(address)
	if err == nil && cached != "" {
		s.logger.Infof("hit cache data skymavis-service, address: %s", address)
		go s.doNetworkFarming(address)
		return &data, json.Unmarshal([]byte(cached), &data)
	}

	// call network
	return s.doNetworkFarming(address)
}

func (s *skymavis) GetOwnedNfts(address string) (*response.AxieMarketNftResponse, error) {
	s.logger.Debug("start skymavis.GetOwnedNfts()")
	defer s.logger.Debug("end skymavis.GetOwnedNfts()")

	var data response.AxieMarketNftResponse
	// check if data cached

	cached, err := s.doCacheNft(address)
	if err == nil && cached != "" {
		s.logger.Infof("hit cache data skymavis-service, address: %s", address)
		go s.doNetworkNfts(address)
		return &data, json.Unmarshal([]byte(cached), &data)
	}

	// call network
	return s.doNetworkNfts(address)
}

func (s *skymavis) GetInternalTxnsByHash(hash string) (*response.SkymavisTransactionsResponse, error) {
	s.logger.Debug("start skymavis.GetInternalTxnsByHash()")
	defer s.logger.Debug("end skymavis.GetInternalTxnsByHash()")

	var data response.SkymavisTransactionsResponse
	// check if data cached

	cached, err := s.doCacheInternalTxns(hash)
	if err == nil && cached != "" {
		s.logger.Infof("hit cache data skymavis-service, hash: %s", hash)
		go s.doNetworkInternalTxs(hash)
		return &data, json.Unmarshal([]byte(cached), &data)
	}

	// call network
	return s.doNetworkInternalTxs(hash)
}
