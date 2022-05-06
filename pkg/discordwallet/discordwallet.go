package discordwallet

import (
	"fmt"

	"github.com/defipod/mochi/pkg/chain"
	"github.com/defipod/mochi/pkg/config"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/repo"
	"github.com/defipod/mochi/pkg/repo/pg"
	"github.com/ethereum/go-ethereum/accounts"

	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
)

type IDiscordWallet interface {
	GetAccountByWalletNumber(i int) (accounts.Account, error)
	GetHDWallet() *hdwallet.Wallet
	FTM() chain.Chain
	Ethereum() chain.Chain
	BSC() chain.Chain
}

type DiscordWallet struct {
	log      logger.Logger
	cfg      config.Config
	repo     *repo.Repo
	hdwallet *hdwallet.Wallet
	ftm      chain.Chain
	ethereum chain.Chain
	bsc      chain.Chain
}

// New will return an instance of DiscordWallet struct
func New(cfg config.Config, l logger.Logger, s repo.Store) (*DiscordWallet, error) {
	r := pg.NewRepo(s.DB())

	wallet, err := hdwallet.NewFromMnemonic(cfg.InDiscordWalletMnemonic)
	if err != nil {
		return nil, err
	}

	ftm, err := chain.NewFTMClient(cfg, wallet)
	if err != nil {
		return nil, err
	}

	ethereum, err := chain.NewEthereumClient(cfg, wallet)
	if err != nil {
		return nil, err
	}

	bsc, err := chain.NewBSCClient(cfg, wallet)
	if err != nil {
		return nil, err
	}

	return &DiscordWallet{
		log:      l,
		cfg:      cfg,
		repo:     r,
		ftm:      ftm,
		ethereum: ethereum,
		bsc:      bsc,
		hdwallet: wallet,
	}, nil
}

func (d *DiscordWallet) GetAccountByWalletNumber(i int) (accounts.Account, error) {
	path := hdwallet.MustParseDerivationPath(fmt.Sprintf("m/44'/60'/0'/0/%d", i))
	return d.hdwallet.Derive(path, false)
}

func (d *DiscordWallet) GetHDWallet() *hdwallet.Wallet {
	return d.hdwallet
}

func (d *DiscordWallet) FTM() chain.Chain {
	return d.ftm
}

func (d *DiscordWallet) Ethereum() chain.Chain {
	return d.ethereum
}

func (d *DiscordWallet) BSC() chain.Chain {
	return d.bsc
}
