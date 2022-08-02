package twitter

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
	"strconv"
	"strings"

	"github.com/ChimeraCoder/anaconda"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/response"
	"github.com/defipod/mochi/pkg/util"
)

type twitter struct {
	twitter *anaconda.TwitterApi
}

func NewTwitter() Service {
	// waiting for api keys
	accessToken := os.Getenv("TWITTER_ACCESS_TOKEN")
	accessTokenScrt := os.Getenv("TWITTER_ACCESS_TOKEN_SECRET")
	consumerKey := os.Getenv("TWITTER_CONSUMER_KEY")
	consumerKeyScrt := os.Getenv("TWITTER_CONSUMER_SECRET")
	t := anaconda.NewTwitterApiWithCredentials(accessToken, accessTokenScrt, consumerKey, consumerKeyScrt)
	return &twitter{
		twitter: t,
	}
}

func (t *twitter) SendSalesTweet(imageURL string, nft *request.HandleNftWebhookRequest, token *response.IndexerNFTToken) error {
	imageFile := "temp.png"
	err := util.DownloadFile(imageURL, imageFile)
	if err != nil {
		return fmt.Errorf("[twitter.SendSalesTweet] cannot download image: %s", err)
	}
	defer deleteFile(imageFile)

	v := url.Values{}
	resizedImageFile, _, _, err := util.CheckAndResizeImg(imageFile)
	if err != nil {
		return fmt.Errorf("[twitter.SendSalesTweet] cannot resize image: %s", err)
	}
	defer deleteFile(resizedImageFile)

	data, err := ioutil.ReadFile(resizedImageFile)
	if err != nil {
		return fmt.Errorf("[twitter.SendSalesTweet] cannot open resized image: %s", err)
	}

	// upload image to twiiter
	mediaResponse, err := t.twitter.UploadMedia(base64.StdEncoding.EncodeToString(data))
	if err != nil {
		return fmt.Errorf("[twitter.SendSalesTweet] cannot upload media to twitter: %s", err)
	}

	// set media to v
	v.Set("media_ids", strconv.FormatInt(mediaResponse.MediaID, 10))

	//currency := util.ConvertChainIDToChain(strconv.Itoa(int(nft.ChainId)))
	price := util.StringWeiToEther(nft.Price.Amount, nft.Price.Token.Decimal)
	marketplaceUrl := util.GetURLMarketPlace(nft.Marketplace) + nft.CollectionAddress
	// rank := strconv.Itoa(int(token.Rarity.Rank))
	// rarity := util.GetTwitterRarityEmoji(token.Rarity.Rarity) + " " + token.Rarity.Rarity

	tweetStatus := fmt.Sprintf("A new sale has been made on %s for %s!\n\nBuyer: %s\n\nSeller: %s\n\nValue: %v %s\n\n→ Check collection at: %s",
		strings.ToTitle(nft.Marketplace), token.Name, util.ShortenAddress(nft.To), util.ShortenAddress(nft.From), util.FormatCryptoPrice(*price), strings.ToUpper(nft.Price.Token.Symbol), marketplaceUrl)
	_, err = t.twitter.PostTweet(tweetStatus, v)
	if err != nil {
		return fmt.Errorf("[twitter.SendSalesTweet] cannot post tweet: %s", err)
	}

	return nil
}

func deleteFile(filePath string) {
	// delete temp image
	_, e := os.Stat(filePath)
	if !os.IsNotExist(e) {
		_ = os.Remove(filePath)
	}
}
