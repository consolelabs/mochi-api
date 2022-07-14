package entities

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/util"
)

func (e *Entity) SendNftSalesToChannel(nftSale request.NftSalesRequest) error {
	indexerToken, err := e.indexer.GetNFTDetail(nftSale.CollectionAddress, nftSale.TokenId)
	if err != nil {
		return err
	}

	collection, err := e.repo.NFTCollection.GetByAddress(nftSale.CollectionAddress)
	if err != nil {
		return err
	}

	price := util.StringWeiToEther(nftSale.Price.Amount, nftSale.Price.Token.Decimal)
	gain := util.StringWeiToEther(nftSale.Gain.Amount, nftSale.Gain.Token.Decimal)
	rankDisplay := strconv.Itoa(int(indexerToken.Rarity.Rank))
	rarityDisplay := indexerToken.Rarity.Rarity

	if rarityDisplay == "" {
		rarityDisplay = "N/A"
	} else {
		rarityDisplay = util.RarityEmoji(rarityDisplay) + " " + rarityDisplay
	}
	if indexerToken.Rarity.Rank == 0 {
		rankDisplay = "N/A"
	} else {
		rankDisplay = "<:cup:985137841027821589> " + rankDisplay
	}

	// handle marketplace
	marketplace := strings.ToUpper(string(nftSale.Marketplace[0])) + nftSale.Marketplace[1:]

	data := []*discordgo.MessageEmbedField{
		{
			Name:   "Rarity",
			Value:  rarityDisplay,
			Inline: true,
		},
		{

			Name:   "Rank",
			Value:  rankDisplay,
			Inline: true,
		},
		{
			Name:   "\u200B",
			Value:  "\u200B",
			Inline: true,
		},
		{
			Name:   "Marketplace",
			Value:  "[" + marketplace + "](" + util.GetURLMarketPlace(nftSale.Marketplace) + strings.ToLower(nftSale.CollectionAddress) + ")",
			Inline: true,
		},
		{
			Name:   "Transaction",
			Value:  "[" + util.ShortenAddress(nftSale.Transaction) + "]" + "(https://www.youtube.com/)",
			Inline: true,
		},
		{
			Name:   "\u200B",
			Value:  "\u200B",
			Inline: true,
		},
		{
			Name:   "From",
			Value:  "[" + util.ShortenAddress(nftSale.From) + "]" + "(https://www.youtube.com/)",
			Inline: true,
		},
		{
			Name:   "To",
			Value:  "[" + util.ShortenAddress(nftSale.To) + "]" + "(https://www.youtube.com/)",
			Inline: true,
		},
		{
			Name:   "\u200B",
			Value:  "\u200B",
			Inline: true,
		},
		{
			Name:   "Price",
			Value:  fmt.Sprintf("%.2f", price) + " " + strings.ToUpper(nftSale.Price.Token.Symbol),
			Inline: true,
		},
		{
			Name:   "Sold",
			Value:  fmt.Sprintf("%.2f", price) + " " + strings.ToUpper(nftSale.Price.Token.Symbol),
			Inline: true,
		},
	}
	if nftSale.Hodl != 0 {
		dataHodl := discordgo.MessageEmbedField{
			Name:   "Hodl",
			Value:  strconv.Itoa(util.SecondsToDays(nftSale.Hodl)) + " days",
			Inline: true,
		}
		data = append(data, &dataHodl)
	}

	if nftSale.Gain.Amount != "" {
		dataGain := discordgo.MessageEmbedField{
			Name:   "Gain",
			Value:  fmt.Sprintf("%.2f", gain) + " " + strings.ToUpper(nftSale.Gain.Token.Symbol),
			Inline: true,
		}
		data = append(data, &dataGain)
	}

	if nftSale.Pnl != 0 {
		dataPnl := []*discordgo.MessageEmbedField{
			{
				Name:   "Pnl",
				Value:  "$" + fmt.Sprintf("%v", nftSale.Pnl) + " " + "`+" + fmt.Sprintf("%v", nftSale.SubPnl) + "%`",
				Inline: true,
			},
			{
				Name:   "\u200B",
				Value:  "\u200B",
				Inline: true,
			},
		}
		data = append(data, dataPnl...)
	}

	if !(((nftSale.Pnl != 0) && (nftSale.Hodl != 0) && (nftSale.Gain.Amount != "")) || ((nftSale.Pnl == 0) && (nftSale.Hodl == 0) && (nftSale.Gain.Amount == ""))) {
		dataPnl := discordgo.MessageEmbedField{
			Name:   "\u200B",
			Value:  "\u200B",
			Inline: true,
		}
		data = append(data, &dataPnl)
	}

	messageSale := []*discordgo.MessageEmbed{{
		Author: &discordgo.MessageEmbedAuthor{
			Name:    collection.Name,
			IconURL: collection.Image,
		},
		Fields:      data,
		Description: indexerToken.Name + " sold!",
		Color:       int(util.RarityColors(indexerToken.Rarity.Rarity)),
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: indexerToken.Image,
		},
	}}
	resp, _ := e.GetAllNFTSalesTracker()

	for _, saleChannel := range resp {
		//## special case address = *, can be removed if not used
		if saleChannel.ContractAddress == "*" {
			_, err := e.discord.ChannelMessageSendEmbeds(saleChannel.ChannelID, messageSale)
			if err != nil {
				e.log.Errorf(err, "[discord.ChannelMessageSendEmbeds] cannot send message to sale channel. CollectionName: %s, TokenName: %s", collection.Name, indexerToken.Name)
				return fmt.Errorf("cannot send message to sale channel. Error: %v", err)
			}

		}
		//##
		if nftSale.CollectionAddress == saleChannel.ContractAddress {
			_, err := e.discord.ChannelMessageSendEmbeds(saleChannel.ChannelID, messageSale)
			if err != nil {
				e.log.Errorf(err, "[discord.ChannelMessageSendEmbeds] cannot send message to sale channel. CollectionName: %s, TokenName: %s", collection.Name, indexerToken.Name)
				return fmt.Errorf("cannot send message to sale channel. Error: %v", err)
			}
		}
	}

	return nil
}
