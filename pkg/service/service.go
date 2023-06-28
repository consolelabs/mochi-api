package service

import (
	"fmt"

	"github.com/go-redis/redis/v8"

	"github.com/defipod/mochi/pkg/cache"
	"github.com/defipod/mochi/pkg/config"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/service/abi"
	"github.com/defipod/mochi/pkg/service/apilayer"
	"github.com/defipod/mochi/pkg/service/apns"
	"github.com/defipod/mochi/pkg/service/binance"
	"github.com/defipod/mochi/pkg/service/birdeye"
	"github.com/defipod/mochi/pkg/service/bluemove"
	"github.com/defipod/mochi/pkg/service/chainexplorer"
	"github.com/defipod/mochi/pkg/service/cloud"
	"github.com/defipod/mochi/pkg/service/coingecko"
	"github.com/defipod/mochi/pkg/service/commonwealth"
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
	"github.com/defipod/mochi/pkg/service/snapshot"
	solscan "github.com/defipod/mochi/pkg/service/solscan"
	"github.com/defipod/mochi/pkg/service/sui"
	"github.com/defipod/mochi/pkg/service/swap"
	"github.com/defipod/mochi/pkg/service/twitter"
)

type Service struct {
	CoinGecko     coingecko.Service
	Covalent      covalent.Service
	Commonwealth  commonwealth.Service
	Discord       discord.Service
	Indexer       indexer.Service
	Abi           abi.Service
	Apns          apns.Service
	Twitter       twitter.Service
	Cloud         cloud.Service
	Processor     processor.Service
	Solscan       solscan.Service
	Snapshot      snapshot.Service
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

	discordSvc, err := discord.NewService(cfg, log)
	if err != nil {
		return nil, fmt.Errorf("failed to init discord: %w", err)
	}

	roninSvc, err := ronin.New(&cfg)
	if err != nil {
		log.Error(err, "failed to init ronin svc")
	}

	redisOpt, err := redis.ParseURL(cfg.RedisURL)
	if err != nil {
		log.Fatal(err, "failed to init redis")
	}

	cache, err := cache.NewRedisCache(redisOpt)
	if err != nil {
		log.Fatal(err, "failed to init redis cache")
	}

	return &Service{
		CoinGecko:     coingecko.NewService(&cfg),
		Covalent:      covalent.NewService(&cfg, log),
		Commonwealth:  commonwealth.NewService(),
		Discord:       discordSvc,
		Indexer:       indexer.NewIndexer(cfg, log),
		Abi:           abi.NewAbi(&cfg),
		Apns:          apns.NewService(&cfg),
		Twitter:       twitter.NewTwitter(&cfg),
		Cloud:         cloud.NewCloudClient(&cfg, log),
		Snapshot:      snapshot.NewService(log),
		Nghenhan:      nghenhan.NewService(),
		Processor:     processor.NewProcessor(&cfg),
		Solscan:       solscan.NewService(&cfg, log),
		Binance:       binance.NewService(&cfg, log, cache),
		APILayer:      apilayer.NewService(&cfg),
		Bluemove:      bluemove.New(&cfg, log),
		MochiProfile:  mochiprofile.NewService(&cfg, log),
		MochiPay:      mochipay.NewService(&cfg, log),
		ChainExplorer: chainexplorer.NewService(cfg, log),
		Sui:           sui.New(&cfg, log),
		Birdeye:       birdeye.NewService(&cfg, log),
		Swap:          swap.New(&cfg, log),
		Skymavis:      skymavis.New(&cfg),
		Ronin:         roninSvc,
		Krystal:       krystal.NewService(&cfg, log, cache),
	}, nil
}
