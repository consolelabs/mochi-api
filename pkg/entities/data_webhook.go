package entities

import (
	"fmt"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/defipod/mochi/pkg/logger"
	nftaddrequesthistory "github.com/defipod/mochi/pkg/repo/nft_add_request_history"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/util"
)

func (e *Entity) NotifyNftCollectionIntegration(req request.NotifyCompleteNftIntegrationRequest) error {
	collection, err := e.repo.NFTCollection.GetByAddress(req.CollectionAddress)
	if err != nil {
		e.log.Error(err, "[entity.SendCollectionIntegrationToMochiLogs] repo.NFTCollection.GetByAddress() failed")
		return err
	}
	history, err := e.repo.NftAddRequestHistory.GetOne(nftaddrequesthistory.GetOneQuery{Address: req.CollectionAddress, ChainID: req.ChainID})
	if err != nil {
		e.log.Fields(logger.Fields{"address": req.CollectionAddress}).Error(err, "[entity.SendCollectionIntegrationToMochiLogs] repo.NftAddRequestHistory.GetOne() failed")
		return err
	}

	chain := strings.ToUpper(util.ConvertChainIDToChain(collection.ChainID))
	// send logs to mochi
	err = e.svc.Discord.NotifyCompleteCollectionIntegration(history.GuildID, collection.Name, collection.Symbol, chain, collection.Image)
	if err != nil {
		e.log.Error(err, "[entity.SendCollectionIntegrationToMochiLogs] svc.Discord.NotifyCompleteCollectionIntegration() failed")
		return err
	}

	// reply to orignal command
	description := fmt.Sprintf("👉 Your collection is being processed. We will let you know when it's ready to use.\n👉 To track other collection in `$nft list`, run `$nft %s <token_id>`", collection.Address)
	_, err = e.discord.ChannelMessageSendEmbedReply(history.ChannelID, &discordgo.MessageEmbed{
		Author: &discordgo.MessageEmbedAuthor{
			IconURL: "https://cdn.discordapp.com/emojis/977508805011181638.png?size=240&quality=lossless",
			Name:    fmt.Sprintf("%s has been added", collection.Name),
		},
		Description: description,
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: collection.Image,
		},
		Timestamp: time.Now().Format("2006-01-02T15:04:05Z07:00"),
	}, &discordgo.MessageReference{
		GuildID:   history.GuildID,
		ChannelID: history.ChannelID,
		MessageID: history.MessageID,
	})

	return err
}

func (e *Entity) NotifyNftCollectionSync(req request.NotifyCompleteNftSyncRequest) error {
	collection, err := e.repo.NFTCollection.GetByAddress(req.CollectionAddress)
	if err != nil {
		e.log.Error(err, "[entity.NotifyNftCollectionSync] repo.NFTCollection.GetByAddress() failed")
		return err
	}
	history, err := e.repo.NftAddRequestHistory.GetOne(nftaddrequesthistory.GetOneQuery{Address: req.CollectionAddress, ChainID: req.ChainID})
	if err != nil {
		e.log.Fields(logger.Fields{"address": req.CollectionAddress}).Error(err, "[entity.NotifyNftCollectionSync] repo.NftAddRequestHistory.GetOne() failed")
		return err
	}

	chain := strings.ToUpper(util.ConvertChainIDToChain(collection.ChainID))
	// send logs to mochi
	err = e.svc.Discord.NotifyCompleteCollectionSync(history.GuildID, collection.Name, collection.Symbol, chain, collection.Image)
	if err != nil {
		e.log.Error(err, "[entity.NotifyNftCollectionSync] svc.Discord.NotifyCompleteCollectionSync() failed")
		return err
	}

	// reply to orignal command
	description := fmt.Sprintf("👉 To check rarity, run `$nft %s <token_id>`.\n👉 To track sales, run `$sales track <channel> %s %s`.", collection.Symbol, collection.Address, collection.ChainID)
	_, err = e.discord.ChannelMessageSendEmbedReply(history.ChannelID, &discordgo.MessageEmbed{
		Author: &discordgo.MessageEmbedAuthor{
			IconURL: "https://cdn.discordapp.com/emojis/977508805011181638.png?size=240&quality=lossless",
			Name:    fmt.Sprintf("%s is ready to use", collection.Name),
		},
		Description: description,
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: collection.Image,
		},
		Timestamp: time.Now().Format("2006-01-02T15:04:05Z07:00"),
	}, &discordgo.MessageReference{
		GuildID:   history.GuildID,
		ChannelID: history.ChannelID,
		MessageID: history.MessageID,
	})

	return err
}
