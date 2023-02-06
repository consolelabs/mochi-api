package service

import (
	"fmt"

	"github.com/defipod/mochi/pkg/config"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/service/abi"
	"github.com/defipod/mochi/pkg/service/apilayer"
	"github.com/defipod/mochi/pkg/service/apns"
	"github.com/defipod/mochi/pkg/service/binance"
	"github.com/defipod/mochi/pkg/service/bluemove"
	"github.com/defipod/mochi/pkg/service/cloud"
	"github.com/defipod/mochi/pkg/service/coingecko"
	"github.com/defipod/mochi/pkg/service/covalent"
	"github.com/defipod/mochi/pkg/service/discord"
	"github.com/defipod/mochi/pkg/service/indexer"
	"github.com/defipod/mochi/pkg/service/nghenhan"
	"github.com/defipod/mochi/pkg/service/processor"
	"github.com/defipod/mochi/pkg/service/snapshot"
	solscan "github.com/defipod/mochi/pkg/service/solscan"
	"github.com/defipod/mochi/pkg/service/twitter"
)

// import "github.com/defipod/api/pkg/service/binance"

type Service struct {
	CoinGecko coingecko.Service
	Covalent  covalent.Service
	Discord   discord.Service
	Indexer   indexer.Service
	Abi       abi.Service
	Apns      apns.Service
	Twitter   twitter.Service
	Cloud     cloud.Service
	Processor processor.Service
	Solscan   solscan.Service
	Snapshot  snapshot.Service
	Nghenhan  nghenhan.Service
	Binance   binance.Service
	APILayer  apilayer.Service
	Bluemove  bluemove.Service
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
		CoinGecko: coingecko.NewService(&cfg),
		Covalent:  covalent.NewService(&cfg, log),
		Discord:   discordSvc,
		Indexer:   indexer.NewIndexer(cfg, log),
		Abi:       abi.NewAbi(&cfg),
		Apns:      apns.NewService(&cfg),
		Twitter:   twitter.NewTwitter(&cfg),
		Cloud:     cloud.NewCloudClient(&cfg, log),
		Snapshot:  snapshot.NewService(log),
		Nghenhan:  nghenhan.NewService(),
		Processor: processor.NewProcessor(&cfg),
		Solscan:   solscan.NewService(&cfg),
		Binance:   binance.NewService(),
		APILayer:  apilayer.NewService(&cfg),
		Bluemove:  bluemove.New(&cfg, log),
	}, nil
}
