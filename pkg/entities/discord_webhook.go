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
	nftSale.TokenName = indexerToken.Name
	nftSale.TokenImage = indexerToken.Image
	nftSale.Rank = indexerToken.Rarity.Rank
	nftSale.Rarity = indexerToken.Rarity.Rarity

	collection, err := e.repo.NFTCollection.GetByAddress(nftSale.CollectionAddress)
	if err != nil {
		return err
	}
	nftSale.CollectionName = collection.Name
	nftSale.CollectionImage = collection.Image

	price := util.ConvertToFloat(nftSale.Price.Amount, nftSale.Price.Token.Decimal)
	gain := util.ConvertToFloat(nftSale.Gain.Amount, nftSale.Gain.Token.Decimal)
	rankDisplay := strconv.Itoa(int(nftSale.Rank))
	rarityDisplay := nftSale.Rarity

	data := []*discordgo.MessageEmbedField{}
	if nftSale.Rarity == "" {
		rarityDisplay = "N/A"
	} else {
		rarityDisplay = e.RarityEmoji(nftSale.Rarity) + " " + nftSale.Rarity
	}
	if nftSale.Rank == 0 {
		rankDisplay = "N/A"
	} else {
		rankDisplay = "<:cup:985137841027821589> " + rankDisplay
	}
	fixed := []*discordgo.MessageEmbedField{
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
			Value:  "[" + nftSale.Marketplace + "](" + util.GetURLMarketPlace(nftSale.Marketplace) + nftSale.CollectionAddress + "/" + nftSale.TokenId + ")",
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
	data = append(data, fixed...)
	if nftSale.Hodl != 0 {
		dataHodl := discordgo.MessageEmbedField{
			Name:   "Hold",
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
			Name:    nftSale.CollectionName,
			IconURL: nftSale.CollectionImage,
		},
		Fields:      data,
		Description: nftSale.TokenName + " sold!",
		Color:       int(e.RarityColors(nftSale.Rarity)),
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: nftSale.TokenImage,
		},
	}}
	resp, _ := e.GetAllNFTSalesTracker()

	for _, saleChannel := range resp {
		//## special case address = *, can be removed if not used
		if saleChannel.ContractAddress == "*" {
			_, err := e.discord.ChannelMessageSendEmbeds(saleChannel.ChannelID, messageSale)
			if err != nil {
				e.log.Errorf(err, "[discord.ChannelMessageSendEmbeds] cannot send message to sale channel. CollectionName: %s, TokenName: %s", nftSale.CollectionName, nftSale.TokenName)
				return fmt.Errorf("cannot send message to sale channel. Error: %v", err)
			}

		}
		//##
		if nftSale.CollectionAddress == saleChannel.ContractAddress {
			_, err := e.discord.ChannelMessageSendEmbeds(saleChannel.ChannelID, messageSale)
			if err != nil {
				e.log.Errorf(err, "[discord.ChannelMessageSendEmbeds] cannot send message to sale channel. CollectionName: %s, TokenName: %s", nftSale.CollectionName, nftSale.TokenName)
				return fmt.Errorf("cannot send message to sale channel. Error: %v", err)
			}
		}
	}

	return nil
}

func (e *Entity) RarityColors(rarity string) int64 {
	switch strings.ToLower(rarity) {
	case "common":
		return 9671571
	case "uncommon":
		return 2282633
	case "rare":
		return 177151
	case "epic":
		return 9962230
	case "legendary":
		return 16744449
	case "mythic":
		return 15542585
	default:
		return 9671571
	}
}

func (e *Entity) RarityEmoji(rarity string) string {
	switch strings.ToLower(rarity) {
	case "common":
		return ":white_circle:"
	case "uncommon":
		return ":green_circle:"
	case "rare":
		return ":blue_circle:"
	case "epic":
		return ":purple_circle:"
	case "legendary":
		return ":orange_circle:"
	case "mythic":
		return ":red_circle:"
	default:
		return ":white_circle:"
	}
}
