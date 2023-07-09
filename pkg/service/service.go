package service

import (
	"fmt"

	"github.com/go-redis/redis/v8"

	"github.com/defipod/mochi/pkg/cache"
	"github.com/defipod/mochi/pkg/config"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/repo/pg"
	"github.com/defipod/mochi/pkg/service/abi"
	"github.com/defipod/mochi/pkg/service/apilayer"
	"github.com/defipod/mochi/pkg/service/binance"
	"github.com/defipod/mochi/pkg/service/birdeye"
	"github.com/defipod/mochi/pkg/service/bluemove"
	"github.com/defipod/mochi/pkg/service/chainexplorer"
	"github.com/defipod/mochi/pkg/service/cloud"
	"github.com/defipod/mochi/pkg/service/coingecko"
	"github.com/defipod/mochi/pkg/service/covalent"
	"github.com/defipod/mochi/pkg/service/discord"
	"github.com/defipod/mochi/pkg/service/indexer"
	"github.com/defipod/mochi/pkg/service/krystal"
	"github.com/defipod/mochi/pkg/service/mochipay"
	"github.com/defipod/mochi/pkg/service/mochiprofile"
	"github.com/defipod/mochi/pkg/service/nghenhan"
	"github.com/defipod/mochi/pkg/service/processor"
	"github.com/defipod/mochi/pkg/service/ronin"
	"github.com/defipod/mochi/pkg/service/skymavis"
	solscan "github.com/defipod/mochi/pkg/service/solscan"
	"github.com/defipod/mochi/pkg/service/sui"
	"github.com/defipod/mochi/pkg/service/swap"
)

type Service struct {
	CoinGecko     coingecko.Service
	Covalent      covalent.Service
	Discord       discord.Service
	Indexer       indexer.Service
	Abi           abi.Service
	Cloud         cloud.Service
	Processor     processor.Service
	Solscan       solscan.Service
	Nghenhan      nghenhan.Service
	Binance       binance.Service
	APILayer      apilayer.Service
	Bluemove      bluemove.Service
	MochiProfile  mochiprofile.Service
	MochiPay      mochipay.Service
	ChainExplorer chainexplorer.Service
	Sui           sui.Service
	Birdeye       birdeye.Service
	Swap          swap.Service
	Skymavis      skymavis.Service
	Ronin         ronin.Service
	Krystal       krystal.Service
}

func NewService(
	cfg config.Config,
	log logger.Logger,
) (*Service, error) {
	s := pg.NewPostgresStore(&cfg)
	repo := pg.NewRepo(s.DB())

	discordSvc, err := discord.NewService(cfg, log, repo)
	if err != nil {
		return nil, fmt.Errorf("failed to init discord: %w", err)
	}

	redisOpt, err := redis.ParseURL(cfg.RedisURL)
	if err != nil {
		log.Fatal(err, "failed to init redis")
	}

	cache, err := cache.NewRedisCache(redisOpt)
	if err != nil {
		log.Fatal(err, "failed to init redis cache")
	}

	roninSvc, err := ronin.New(&cfg, cache)
	if err != nil {
		log.Error(err, "failed to init ronin svc")
	}

	return &Service{
		CoinGecko:     coingecko.NewService(&cfg),
		Covalent:      covalent.NewService(&cfg, log, cache),
		Discord:       discordSvc,
		Indexer:       indexer.NewIndexer(cfg, log),
		Abi:           abi.NewAbi(&cfg),
		Cloud:         cloud.NewCloudClient(&cfg, log),
		Nghenhan:      nghenhan.NewService(),
		Processor:     processor.NewProcessor(&cfg),
		Solscan:       solscan.NewService(&cfg, log, cache),
		Binance:       binance.NewService(&cfg, log, cache),
		APILayer:      apilayer.NewService(&cfg),
		Bluemove:      bluemove.New(&cfg, log),
		MochiProfile:  mochiprofile.NewService(&cfg, log),
		MochiPay:      mochipay.NewService(&cfg, log),
		ChainExplorer: chainexplorer.NewService(cfg, log),
		Sui:           sui.New(&cfg, log, cache),
		Birdeye:       birdeye.NewService(&cfg, log, cache),
		Swap:          swap.New(&cfg, log),
		Skymavis:      skymavis.New(&cfg, cache),
		Ronin:         roninSvc,
		Krystal:       krystal.NewService(&cfg, log, cache),
	}, nil
}
