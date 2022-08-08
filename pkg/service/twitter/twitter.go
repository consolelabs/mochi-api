package twitter

import (
	"fmt"
	"io/ioutil"
	"net/url"
	"os"

	"github.com/ChimeraCoder/anaconda"
	"github.com/defipod/mochi/pkg/config"
	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/util"
)

type twitter struct {
	config  *config.Config
	twitter *anaconda.TwitterApi
}

func NewTwitter(cfg *config.Config) Service {
	t := anaconda.NewTwitterApiWithCredentials(cfg.TwitterAccessToken, cfg.TwitterAccessTokenSecret, cfg.TwitterConsumerKey, cfg.TwitterConsumerSecret)
	return &twitter{
		twitter: t,
		config:  cfg,
	}
}
func (t *twitter) preprocessTwitterImage(image string) ([]byte, string, error) {
	imageFile := "temp.png"
	err := util.DownloadFile(image, imageFile)
	if err != nil {
		return nil, "", fmt.Errorf("[twitter.SendSalesTweet] cannot download image: %s", err)
	}
	// defer deleteFile(imageFile)

	resizedImageFile, _, _, err := util.CheckAndResizeImg(imageFile)
	if err != nil {
		return nil, "", fmt.Errorf("[twitter.SendSalesTweet] cannot resize image: %s", err)
	}

	data, err := ioutil.ReadFile(resizedImageFile)
	if err != nil {
		return nil, "", fmt.Errorf("[twitter.SendSalesTweet] cannot open resized image: %s", err)
	}

	return data, resizedImageFile, nil
}
func (t *twitter) SendSalesMessageToTwitter(message *model.TwitterSalesMessage, twitter *model.GuildConfigTwitterFeed) error {
	twitterApi := anaconda.NewTwitterApiWithCredentials(twitter.TwitterAccessToken, twitter.TwitterAccessTokenSecret, twitter.TwitterConsumerKey, twitter.TwitterConsumerSecret)
	v := url.Values{}

	// **fields format ->
	// marketplace = {Paintswap, Opensea, Optimisim}
	// token name = Cyber Neko #2
	// buyer,seller = 0x23...1234
	// price = 123 FTM
	// edit post UI ->
	tweetStatus := fmt.Sprintf("%s bought for %s from %s by %s\nMarketplace: %s\nTx: %s\nhttps://rarepepe.gg/asset/%s/%s?twitter",
		message.TokenName, message.Price, message.BuyerAddress, message.SellerAddress, message.Marketplace, message.TxURL, message.CollectionAddress, message.TokenID)
	_, err := twitterApi.PostTweet(tweetStatus, v)
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
