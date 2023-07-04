package handler

import (
	"github.com/defipod/mochi/pkg/entities"
	airdropcampaign "github.com/defipod/mochi/pkg/handler/airdrop-campaign"
	apikey "github.com/defipod/mochi/pkg/handler/api-key"
	"github.com/defipod/mochi/pkg/handler/auth"
	"github.com/defipod/mochi/pkg/handler/community"
	"github.com/defipod/mochi/pkg/handler/config"
	configchannel "github.com/defipod/mochi/pkg/handler/config-channel"
	configcommunity "github.com/defipod/mochi/pkg/handler/config-community"
	configdefi "github.com/defipod/mochi/pkg/handler/config-defi"
	configroles "github.com/defipod/mochi/pkg/handler/config-roles"
	configtwittersales "github.com/defipod/mochi/pkg/handler/config-twitter-sales"
	"github.com/defipod/mochi/pkg/handler/content"
	"github.com/defipod/mochi/pkg/handler/defi"
	"github.com/defipod/mochi/pkg/handler/dex"
	"github.com/defipod/mochi/pkg/handler/emojis"
	"github.com/defipod/mochi/pkg/handler/guild"
	"github.com/defipod/mochi/pkg/handler/healthz"
	"github.com/defipod/mochi/pkg/handler/nft"
	pkpass "github.com/defipod/mochi/pkg/handler/pk-pass"
	"github.com/defipod/mochi/pkg/handler/swap"
	"github.com/defipod/mochi/pkg/handler/telegram"
	"github.com/defipod/mochi/pkg/handler/tip"
	"github.com/defipod/mochi/pkg/handler/user"
	"github.com/defipod/mochi/pkg/handler/vault"
	"github.com/defipod/mochi/pkg/handler/verify"
	"github.com/defipod/mochi/pkg/handler/wallet"
	"github.com/defipod/mochi/pkg/handler/webhook"
	"github.com/defipod/mochi/pkg/handler/widget"
	"github.com/defipod/mochi/pkg/logger"
)

type Handler struct {
	Healthcheck       healthz.IHandler
	Auth              auth.IHandler
	Community         community.IHandler
	Guild             guild.IHandler
	Config            config.IHandler
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
	ConfigTwitterSale configtwittersales.IHandler
	Wallet            wallet.IHandler
	Telegram          telegram.IHandler
	Vault             vault.IHandler
	Swap              swap.IHandler
	ApiKey            apikey.IHandler
	PkPass            pkpass.IHandler
	Emojis            emojis.IHandler
	Dex               dex.IHandler
	Content           content.IHandler
	AirdropCampaign   airdropcampaign.IHandler
}

func New(entities *entities.Entity, logger logger.Logger) *Handler {
	return &Handler{
		Healthcheck:       healthz.New(),
		Auth:              auth.New(entities, logger),
		Community:         community.New(entities, logger),
		Guild:             guild.New(entities, logger),
		Config:            config.New(entities, logger),
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
		ConfigTwitterSale: configtwittersales.New(entities, logger),
		Wallet:            wallet.New(entities, logger),
		Telegram:          telegram.New(entities, logger),
		Vault:             vault.New(entities, logger),
		Swap:              swap.New(entities, logger),
		ApiKey:            apikey.New(entities, logger),
		PkPass:            pkpass.New(entities, logger),
		Emojis:            emojis.New(entities, logger),
		Dex:               dex.New(entities, logger),
		Content:           content.New(entities, logger),
		AirdropCampaign:   airdropcampaign.New(entities, logger),
	}
}
