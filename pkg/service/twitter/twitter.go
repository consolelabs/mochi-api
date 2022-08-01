package twitter

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
	"strconv"

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
	t := anaconda.NewTwitterApiWithCredentials("your-access-token", "your-access-token-secret", "your-consumer-key", "your-consumer-secret")
	return &twitter{
		twitter: t,
	}
}

func (t *twitter) SendSalesTweet(imageURL string, nft *request.HandleNftWebhookRequest, token *response.IndexerNFTToken) error {
	imageFile := "temp.png"
	err := util.DownloadFile(imageURL, imageFile)
	if err != nil {
		return fmt.Errorf("[twitter.SendSalesTweet] cannot download image")
	}
	defer deleteFile(imageFile)

	v := url.Values{}
	resizedImageFile, _, _, err := util.CheckAndResizeImg(imageFile)
	if err != nil {
		return fmt.Errorf("[twitter.SendSalesTweet] cannot resize image")
	}
	defer deleteFile(resizedImageFile)

	data, err := ioutil.ReadFile(resizedImageFile)
	if err != nil {
		return fmt.Errorf("[twitter.SendSalesTweet] cannot open resized image")
	}

	// upload image to twiiter
	mediaResponse, err := t.twitter.UploadMedia(base64.StdEncoding.EncodeToString(data))
	if err != nil {
		return fmt.Errorf("[twitter.SendSalesTweet] cannot upload media to twitter")
	}

	// set media to v
	v.Set("media_ids", strconv.FormatInt(mediaResponse.MediaID, 10))

	currency := util.ConvertChainIDToChain(strconv.Itoa(int(nft.ChainId)))
	price := util.StringWeiToEther(nft.Price.Amount, nft.Price.Token.Decimal)
	rank := strconv.Itoa(int(token.Rarity.Rank))
	rarity := util.GetTwitterRarityEmoji(token.Rarity.Rarity) + " " + token.Rarity.Rarity

	tweetStatus := fmt.Sprintf("%s SOLD for %v %s\n\nRank: %s\n\nRarity: %s\n\nFrom Adress: %s\n\nTo Address: %s\n\n", token.Name, util.FormatCryptoPrice(*price), currency, rank, rarity, util.ShortenAddress(nft.From), util.ShortenAddress(nft.To))
	_, err = t.twitter.PostTweet(tweetStatus, v)
	if err != nil {
		return fmt.Errorf("[twitter.SendSalesTweet] cannot post tweet")
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
