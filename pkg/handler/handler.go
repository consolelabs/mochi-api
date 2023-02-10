package handler

import (
	"github.com/defipod/mochi/pkg/entities"
	"github.com/defipod/mochi/pkg/handler/auth"
	"github.com/defipod/mochi/pkg/handler/cache"
	"github.com/defipod/mochi/pkg/handler/community"
	"github.com/defipod/mochi/pkg/handler/config"
	configchannel "github.com/defipod/mochi/pkg/handler/config-channel"
	configcommunity "github.com/defipod/mochi/pkg/handler/config-community"
	configdefi "github.com/defipod/mochi/pkg/handler/config-defi"
	configroles "github.com/defipod/mochi/pkg/handler/config-roles"
	configtwittersales "github.com/defipod/mochi/pkg/handler/config-twitter-sales"
	daovoting "github.com/defipod/mochi/pkg/handler/dao-voting"
	"github.com/defipod/mochi/pkg/handler/data"
	"github.com/defipod/mochi/pkg/handler/defi"
	"github.com/defipod/mochi/pkg/handler/guild"
	"github.com/defipod/mochi/pkg/handler/healthz"
	"github.com/defipod/mochi/pkg/handler/nft"
	"github.com/defipod/mochi/pkg/handler/tip"
	"github.com/defipod/mochi/pkg/handler/user"
	"github.com/defipod/mochi/pkg/handler/verify"
	"github.com/defipod/mochi/pkg/handler/wallet"
	"github.com/defipod/mochi/pkg/handler/webhook"
	"github.com/defipod/mochi/pkg/handler/widget"
	"github.com/defipod/mochi/pkg/logger"
)

type Handler struct {
	Healthcheck       healthz.IHandler
	Auth              auth.IHandler
	Cache             cache.IHandler
	Community         community.IHandler
	Guild             guild.IHandler
	Config            config.IHandler
	Data              data.IHandler
	Defi              defi.IHandler
	Nft               nft.IHandler
	User              user.IHandler
	Verify            verify.IHandler
	Webhook           webhook.IHandler
	Tip               tip.IHandler
	Widget            widget.IHandler
	ConfigChannel     configchannel.IHandler
	ConfigCommunity   configcommunity.IHandler
	ConfigDefi        configdefi.IHandler
	ConfigRoles       configroles.IHandler
	DaoVoting         daovoting.IHandler
	ConfigTwitterSale configtwittersales.IHandler
	Wallet            wallet.IHandler
}

func New(entities *entities.Entity, logger logger.Logger) *Handler {
	return &Handler{
		Healthcheck:       healthz.New(),
		Auth:              auth.New(entities, logger),
		Cache:             cache.New(entities, logger),
		Community:         community.New(entities, logger),
		Guild:             guild.New(entities, logger),
		Config:            config.New(entities, logger),
		Data:              data.New(entities, logger),
		Defi:              defi.New(entities, logger),
		Nft:               nft.New(entities, logger),
		User:              user.New(entities, logger),
		Verify:            verify.New(entities, logger),
		Webhook:           webhook.New(entities, logger),
		Tip:               tip.New(entities, logger),
		Widget:            widget.New(entities, logger),
		ConfigChannel:     configchannel.New(entities, logger),
		ConfigCommunity:   configcommunity.New(entities, logger),
		ConfigDefi:        configdefi.New(entities, logger),
		ConfigRoles:       configroles.New(entities, logger),
		DaoVoting:         daovoting.New(entities, logger),
		ConfigTwitterSale: configtwittersales.New(entities, logger),
		Wallet:            wallet.New(entities, logger),
	}
}
