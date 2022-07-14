package entities

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/util"
)

func (e *Entity) SendNftSalesToChannel(nftSale request.NftSale) error {
	indexerToken, err := e.indexer.GetNFTDetail(nftSale.CollectionAddress, nftSale.TokenId)
	if err != nil {
		return err
	}
	nftSale.TokenName = indexerToken.Name
	nftSale.TokenImage = indexerToken.Image

	collection, err := e.repo.NFTCollection.GetByAddress(nftSale.CollectionAddress)
	if err != nil {
		return err
	}
	nftSale.CollectionName = collection.Name
	nftSale.CollectionImage = collection.Image

	data := []*discordgo.MessageEmbedField{
		{
			Name:   "Rarity",
			Value:  e.RarityEmoji(nftSale.Rarity) + " " + nftSale.Rarity,
			Inline: true,
		},
		{
			Name:   "Rank",
			Value:  "<:cup:985137841027821589> " + strconv.Itoa(int(nftSale.Rank)),
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
			Value:  nftSale.Price + " " + strings.ToUpper(nftSale.PaymentToken),
			Inline: true,
		},
		{
			Name:   "Bought",
			Value:  nftSale.Bought + " " + strings.ToUpper(nftSale.PaymentToken),
			Inline: true,
		},
		{
			Name:   "Sold",
			Value:  nftSale.Sold + " " + strings.ToUpper(nftSale.PaymentToken),
			Inline: true,
		},
	}

	if nftSale.Hodl != "" {
		dataHodl := discordgo.MessageEmbedField{
			Name:   "HODL",
			Value:  nftSale.Hodl,
			Inline: true,
		}
		data = append(data, &dataHodl)
	}

	if nftSale.Gain != "" {
		dataGain := discordgo.MessageEmbedField{
			Name:   "Gain",
			Value:  nftSale.Gain + " " + strings.ToUpper(nftSale.PaymentToken),
			Inline: true,
		}
		data = append(data, &dataGain)
	}

	if nftSale.Pnl != "" {
		dataPnl := discordgo.MessageEmbedField{
			Name:   "Pnl",
			Value:  "$" + nftSale.Pnl + " " + "`+" + nftSale.SubPnl + "%`",
			Inline: true,
		}
		data = append(data, &dataPnl)
	}

	if !(((nftSale.Pnl != "") && (nftSale.Hodl != "") && (nftSale.Gain != "")) || ((nftSale.Pnl == "") && (nftSale.Hodl == "") && (nftSale.Gain == ""))) {
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
			collectionList, err := e.repo.NFTCollection.ListAll()
			if err != nil {
				e.log.Errorf(err, "[discord.ChannelMessageSendEmbeds] cannot get collection data")
				return fmt.Errorf("cannot get data for address = *")
			}
			// if request address found in database => send sales message
			for _, col := range collectionList {
				if nftSale.CollectionAddress == col.Address {
					_, err := e.discord.ChannelMessageSendEmbeds(saleChannel.ChannelID, messageSale)
					if err != nil {
						e.log.Errorf(err, "[discord.ChannelMessageSendEmbeds] cannot send message to sale channel. CollectionName: %s, TokenName: %s", nftSale.CollectionName, nftSale.TokenName)
						return fmt.Errorf("cannot send message to sale channel. Error: %v", err)
					}
				}
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
