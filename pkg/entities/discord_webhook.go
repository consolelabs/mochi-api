package entities

import (
	"fmt"
	"math/big"
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

	// calculate last price, price, pnl, sub pnl
	price := util.StringWeiToEther(nftSale.Price.Amount, nftSale.Price.Token.Decimal)
	lastPrice := util.StringWeiToEther(nftSale.LastPrice.Amount, nftSale.LastPrice.Token.Decimal)
	pnl := new(big.Float)
	pnl = pnl.Sub(price, lastPrice)
	subPnl := new(big.Float).Quo(pnl, lastPrice)
	subPnlPer := subPnl.Mul(subPnl, big.NewFloat(100))

	subPnlDisplay := ""
	if util.FormatCryptoPrice(*lastPrice) != "0" {
		subPnlDisplay = " `" + util.GetChangePnl(pnl) + " " + fmt.Sprintf("%.2f", subPnlPer.Abs(subPnlPer)) + "%`"
	}

	// handle rarity, rank
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
	marketplaceLink := ""
	if strings.ToLower(nftSale.Marketplace) == "opensea" {
		res, err := e.marketplace.GetOpenseaAssetContract(nftSale.CollectionAddress)
		if err != nil {
			e.log.Errorf(err, "[marketplace.GetOpenseaAssetContrace] cannot get opensea data")
			return fmt.Errorf("cannot get opensea data. Error: %v", err)
		}
		marketplaceLink = "[" + marketplace + "](" + util.GetURLMarketPlace(nftSale.Marketplace) + res.Collection.UrlName + ")"
	} else {
		marketplaceLink = "[" + marketplace + "](" + util.GetURLMarketPlace(nftSale.Marketplace) + strings.ToLower(nftSale.CollectionAddress) + ")"
	}

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
			Value:  marketplaceLink,
			Inline: true,
		},
		{
			Name:   "Transaction",
			Value:  "[" + util.ShortenAddress(nftSale.Transaction) + "]" + "(" + util.GetTransactionUrl(nftSale.Marketplace) + strings.ToLower(nftSale.Transaction) + ")",
			Inline: true,
		},
		{
			Name:   "\u200B",
			Value:  "\u200B",
			Inline: true,
		},
		{
			Name:   "From",
			Value:  "[" + util.ShortenAddress(nftSale.From) + "]" + "(" + util.GetWalletUrl(nftSale.Marketplace) + strings.ToLower(nftSale.From) + ")",
			Inline: true,
		},
		{
			Name:   "To",
			Value:  "[" + util.ShortenAddress(nftSale.To) + "]" + "(" + util.GetWalletUrl(nftSale.Marketplace) + strings.ToLower(nftSale.To) + ")",
			Inline: true,
		},
		{
			Name:   "Hodl",
			Value:  strconv.Itoa(util.SecondsToDays(nftSale.Hodl)) + " days",
			Inline: true,
		},
		{
			Name:   "Price",
			Value:  util.FormatCryptoPrice(*price) + " " + strings.ToUpper(nftSale.Price.Token.Symbol),
			Inline: true,
		},
		{
			Name:   "Last Price",
			Value:  util.FormatCryptoPrice(*lastPrice) + " " + strings.ToUpper(nftSale.LastPrice.Token.Symbol),
			Inline: true,
		},
		{
			Name: "PnL",
			// + " " + strings.ToUpper(nftSale.Price.Token.Symbol)
			Value:  util.GetGainEmoji(pnl) + util.FormatCryptoPrice(*pnl) + subPnlDisplay,
			Inline: true,
		},
	}

	// finalize message nft sales
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
