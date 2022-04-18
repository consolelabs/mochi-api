package entities

import (
	"errors"

	"github.com/bwmarrin/discordgo"

	"github.com/defipod/mochi/pkg/config"
	"github.com/defipod/mochi/pkg/discordwallet"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/repo"
	"github.com/defipod/mochi/pkg/repo/pg"
)

type Entity struct {
	cfg      config.Config
	repo     *repo.Repo
	dcwallet discordwallet.IDiscordWallet
	discord  *discordgo.Session
}

func New(cfg config.Config, l logger.Log, s repo.Store, dcwallet discordwallet.IDiscordWallet, discord *discordgo.Session) *Entity {
	r := pg.NewRepo(s.DB())
	return &Entity{
		cfg:      cfg,
		repo:     r,
		dcwallet: dcwallet,
		discord:  discord,
	}
}

var (
	ErrRecordNotFound = errors.New("not found")
)
