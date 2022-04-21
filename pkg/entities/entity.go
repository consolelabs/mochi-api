package entities

import (
	"errors"

	"github.com/bwmarrin/discordgo"

	"github.com/defipod/mochi/pkg/cache"
	"github.com/defipod/mochi/pkg/discordwallet"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/repo"
	"github.com/defipod/mochi/pkg/service"
)

type Entity struct {
	repo     *repo.Repo
	dcwallet discordwallet.IDiscordWallet
	discord  *discordgo.Session
	cache    cache.Cache
	svc      *service.Service
}

func New(
	l logger.Log,
	repo *repo.Repo,
	dcwallet discordwallet.IDiscordWallet,
	discord *discordgo.Session,
	cache cache.Cache,
	service *service.Service,
) *Entity {
	entities := &Entity{
		repo:     repo,
		dcwallet: dcwallet,
		discord:  discord,
		cache:    cache,
		svc:      service,
	}

	if entities.discord != nil && entities.cache != nil {
		if err := entities.InitInviteTrackerCache(); err != nil {
			return nil
		}
	}

	return entities
}

var (
	ErrRecordNotFound = errors.New("not found")
)
