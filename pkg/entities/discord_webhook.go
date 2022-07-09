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
	data := []*discordgo.MessageEmbed{{
		Author: &discordgo.MessageEmbedAuthor{
			Name:    nftSale.CollectionName,
			IconURL: nftSale.CollectionImage,
		},
		Description: nftSale.TokenName + " sold!",
		Color:       int(e.RarityColors(nftSale.Rarity)),
		Fields: []*discordgo.MessageEmbedField{
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
				Value:  nftSale.Marketplace,
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
			{
				Name:   "HODL",
				Value:  nftSale.Hodl,
				Inline: true,
			},
			{
				Name:   "Gain",
				Value:  nftSale.Gain + " " + strings.ToUpper(nftSale.PaymentToken),
				Inline: true,
			},
			{
				Name:   "Pnl",
				Value:  "$" + nftSale.Pnl + " " + "`+" + nftSale.SubPnl + "%`",
				Inline: true,
			},
		},
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL:    nftSale.TokenImage,
			Width:  10000,
			Height: 5000,
		},
	}}
	resp, _ := e.GetAllNFTSalesTracker()

	for _, saleChannel := range resp {
		if nftSale.CollectionAddress == saleChannel.ContractAddress {
			_, err := e.discord.ChannelMessageSendEmbeds(saleChannel.ChannelID, data)
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
