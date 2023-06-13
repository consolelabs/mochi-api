package service

import (
	"fmt"

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
	"github.com/defipod/mochi/pkg/service/kyber"
	"github.com/defipod/mochi/pkg/service/mochipay"
	"github.com/defipod/mochi/pkg/service/mochiprofile"
	"github.com/defipod/mochi/pkg/service/nghenhan"
	"github.com/defipod/mochi/pkg/service/processor"
	"github.com/defipod/mochi/pkg/service/snapshot"
	solscan "github.com/defipod/mochi/pkg/service/solscan"
	"github.com/defipod/mochi/pkg/service/sui"
	"github.com/defipod/mochi/pkg/service/twitter"
)

// import "github.com/defipod/api/pkg/service/binance"

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
	Kyber         kyber.Service
	Sui           sui.Service
	Birdeye       birdeye.Service
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
		Binance:       binance.NewService(),
		APILayer:      apilayer.NewService(&cfg),
		Bluemove:      bluemove.New(&cfg, log),
		MochiProfile:  mochiprofile.NewService(&cfg, log),
		MochiPay:      mochipay.NewService(&cfg, log),
		ChainExplorer: chainexplorer.NewService(cfg, log),
		Kyber:         kyber.New(&cfg, log),
		Sui:           sui.New(&cfg, log),
		Birdeye:       birdeye.NewService(&cfg, log),
	}, nil
}
