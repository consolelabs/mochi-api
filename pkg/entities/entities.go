package entities

import (
	"errors"

	"github.com/bwmarrin/discordgo"

	"github.com/defipod/mochi/pkg/cache"
	"github.com/defipod/mochi/pkg/discordwallet"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/repo"
)

type Entity struct {
	repo     *repo.Repo
	dcwallet discordwallet.IDiscordWallet
	discord  *discordgo.Session
	cache    cache.Cache
}

func New(
	l logger.Log,
	repo *repo.Repo,
	dcwallet discordwallet.IDiscordWallet,
	discord *discordgo.Session,
	cache cache.Cache,
) *Entity {
	entities := &Entity{
		repo:     repo,
		dcwallet: dcwallet,
		discord:  discord,
		cache:    cache,
	}

	return entities
}

var (
	ErrRecordNotFound = errors.New("not found")
)
