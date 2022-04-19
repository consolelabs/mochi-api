package handler

import (
	"net/http"

	"github.com/bwmarrin/discordgo"
	"github.com/defipod/mochi/pkg/config"
	"github.com/defipod/mochi/pkg/discordwallet"
	"github.com/defipod/mochi/pkg/entities"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/repo"
	"github.com/defipod/mochi/pkg/repo/pg"

	"github.com/gin-gonic/gin"
)

// Handler for app
type Handler struct {
	cfg      config.Config
	repo     *repo.Repo
	dcwallet discordwallet.IDiscordWallet
	entities *entities.Entity
	discord  *discordgo.Session
}

// New will return an instance of Auth struct
func New(cfg config.Config, l logger.Log, s repo.Store, dcwallet *discordwallet.DiscordWallet, entities *entities.Entity) (*Handler, error) {
	r := pg.NewRepo(s.DB())

	discord, err := discordgo.New("Bot " + cfg.DiscordToken)
	if err != nil {
		return nil, err
	}

	handler := &Handler{
		cfg:      cfg,
		repo:     r,
		dcwallet: dcwallet,
		entities: entities,
		discord:  discord,
	}

	return handler, nil
}

// Healthz handler
// Return "OK"
func (h *Handler) Healthz(c *gin.Context) {
	c.Header("Content-Type", "text/plain")
	c.String(http.StatusOK, "OK")
}
