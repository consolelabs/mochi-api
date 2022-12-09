package entities

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/util"
)

func (e *Entity) SendNftAddedCollection(nftAddedCollection request.HandleNftWebhookRequest) error {
	collection, err := e.repo.NFTCollection.GetByAddress(nftAddedCollection.CollectionAddress)
	if err != nil {
		e.log.Errorf(err, "[repo.NFTCollection.GetByAddress] cannot get collection. CollectionAddress: %s", nftAddedCollection.CollectionAddress)
		return err
	}

	channelNewCollection := "964780299307343912"
	messageAddedNewCollection := []*discordgo.MessageEmbed{{
		Title:       "New collection: " + collection.Name,
		Description: "We're happy to announce that " + collection.Name + " ranking is available.\n\n" + "You can check your rank using:\n" + "`$nft " + strings.ToLower(collection.Symbol) + " <token_id>`\n\n" + ":warning: Remeber that ranks are calculated using metadata, wrong and bad metadata can impact ranks as well.\n:warning:Ranks are not a financial indicator.\n",
		Color:       0xFCD3C1,
		Timestamp:   time.Now().Format(time.RFC3339),
		Image: &discordgo.MessageEmbedImage{
			URL: collection.Image,
		},
	}}

	_, err = e.discord.ChannelMessageSendEmbeds(channelNewCollection, messageAddedNewCollection)
	if err != nil {
		e.log.Errorf(err, "[discord.ChannelMessageSendEmbeds] cannot send message to new added collection channel. CollectionAddress: %s, Chain: %s", nftAddedCollection.CollectionAddress, nftAddedCollection.ChainId)
		return fmt.Errorf("cannot send message to new added collection channel. Error: %v", err)
	}
	return nil
}

func (e *Entity) SendStealAlert(price float64, address string, marketplace string, token string, image string, name string) error {
	var floor float64 = 0
	var average float64 = 0
	url := ""
	switch marketplace {
	case "opensea":
		//ETH: opensea -> asseet_contract -> slug -> collection/slug -> floor
		res, err := e.marketplace.GetOpenseaAssetContract(address)
		if err != nil {
			return err
		}

		collection, err := e.marketplace.GetCollectionFromOpensea(res.Collection.UrlName)
		if err != nil {
			return err
		}

		floor = collection.Collection.Stats.FloorPrice
		average = collection.Collection.Stats.AveragePrice
		url = fmt.Sprintf("https://opensea.io/assets/ethereum/%s/%s", address, token)

	case "paintswap":
		// FTM: paintswap
		res, err := e.marketplace.GetCollectionFromPaintswap(address)
		if err != nil {
			return err
		}
		floorPrice, _ := util.StringWeiToEther(res.Collection.Stats.FloorPrice, 18).Float64()
		floor = floorPrice
		avgPrice, _ := util.StringWeiToEther(res.Collection.Stats.AveragePrice, 18).Float64()
		average = avgPrice
		url = fmt.Sprintf("https://paintswap.finance/marketplace/assets/%s/%s", address, token)

	case "optimism":
		// OP: quixotic -> collection/slug == collection/address
		res, err := e.marketplace.GetCollectionFromQuixotic(address)
		if err != nil {
			return err
		}
		length := len(strconv.Itoa(int(res.FloorPrice)))
		floorPrice, _ := util.StringWeiToEther(strconv.Itoa(int(res.FloorPrice)), length).Float64()
		floor = floorPrice
		//api does not have average price
		url = fmt.Sprintf("https://quixotic.io/asset/%s/%s", address, token)
	}
	if price < floor {
		err := e.svc.Discord.NotifyStealFloorPrice(price, floor, url, name, image)
		if err != nil {
			return err
		}
	} else if price < average {
		err := e.svc.Discord.NotifyStealAveragePrice(price, average, url, name, image)
		if err != nil {
			return err
		}
	}
	return nil
}
