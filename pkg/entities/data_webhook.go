package entities

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/util"
)

func (e *Entity) NotifyNftCollectionIntegration(req request.SendCollectionIntegrationLogsRequest) error {
	collection, err := e.repo.NFTCollection.GetByAddress(req.CollectionAddress)
	if err != nil {
		e.log.Error(err, "[entity.SendCollectionIntegrationToMochiLogs] repo.NFTCollection.GetByAddress() failed")
		return err
	}
	chain := strings.ToUpper(util.ConvertChainIDToChain(collection.ChainID))
	// send logs to mochi
	err = e.svc.Discord.NotifyCompleteCollectionIntegration(req.GuildID, collection.Name, collection.Symbol, chain, collection.Image)
	if err != nil {
		e.log.Error(err, "[entity.SendCollectionIntegrationToMochiLogs] svc.Discord.NotifyCompleteCollectionIntegration() failed")
		return err
	}

	// reply to orignal command
	_, err = e.discord.ChannelMessageSendEmbedReply(req.ChannelID, &discordgo.MessageEmbed{
		Title:       fmt.Sprintf("Collection %s has been integrated successfully", collection.Name),
		Description: fmt.Sprintf("Symbol: %s\nChain: %s", collection.Symbol, chain),
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: collection.Image,
		},
	}, &discordgo.MessageReference{
		GuildID:   req.GuildID,
		ChannelID: req.ChannelID,
		MessageID: req.MessageID,
	})

	return err
}
