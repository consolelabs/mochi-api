package service

import (
	"fmt"

	"github.com/defipod/mochi/pkg/config"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/service/abi"
	"github.com/defipod/mochi/pkg/service/coingecko"
	"github.com/defipod/mochi/pkg/service/discord"
	"github.com/defipod/mochi/pkg/service/indexer"
)

// import "github.com/defipod/api/pkg/service/binance"

type Service struct {
	CoinGecko coingecko.Service
	Discord   discord.Service
	Indexer   indexer.Service
	Abi       abi.Service
}

func NewService(
	cfg config.Config,
	log logger.Logger,
) (*Service, error) {

	discordSvc, err := discord.NewService(cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to init discord: %w", err)
	}

	return &Service{
		CoinGecko: coingecko.NewService(),
		Discord:   discordSvc,
		Indexer:   indexer.NewIndexer(cfg, log),
		Abi:       abi.NewAbi(&cfg),
	}, nil
}
