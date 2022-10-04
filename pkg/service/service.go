package service

import (
	"fmt"

	"github.com/defipod/mochi/pkg/config"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/service/abi"
	"github.com/defipod/mochi/pkg/service/binance"
	blockchainapi "github.com/defipod/mochi/pkg/service/blockchain_api"
	"github.com/defipod/mochi/pkg/service/cloud"
	"github.com/defipod/mochi/pkg/service/coingecko"
	"github.com/defipod/mochi/pkg/service/covalent"
	"github.com/defipod/mochi/pkg/service/discord"
	"github.com/defipod/mochi/pkg/service/indexer"
	"github.com/defipod/mochi/pkg/service/processor"
	"github.com/defipod/mochi/pkg/service/twitter"
)

// import "github.com/defipod/api/pkg/service/binance"

type Service struct {
	CoinGecko     coingecko.Service
	Covalent      covalent.Service
	Discord       discord.Service
	Indexer       indexer.Service
	Abi           abi.Service
	Twitter       twitter.Service
	Cloud         cloud.Service
	Processor     processor.Service
	BlockchainApi blockchainapi.Service
	Binance       binance.Service
}

func NewService(
	cfg config.Config,
	log logger.Logger,
) (*Service, error) {

	discordSvc, err := discord.NewService(cfg, log)
	if err != nil {
		return nil, fmt.Errorf("failed to init discord: %w", err)
	}

	return &Service{
		CoinGecko:     coingecko.NewService(&cfg),
		Covalent:      covalent.NewService(&cfg),
		Discord:       discordSvc,
		Indexer:       indexer.NewIndexer(cfg, log),
		Abi:           abi.NewAbi(&cfg),
		Twitter:       twitter.NewTwitter(&cfg),
		Cloud:         cloud.NewCloudClient(&cfg, log),
		Processor:     processor.NewProcessor(&cfg),
		BlockchainApi: blockchainapi.NewService(&cfg),
		Binance:       binance.NewService(),
	}, nil
}
