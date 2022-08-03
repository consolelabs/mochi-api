package twitter

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
	"strconv"

	"github.com/ChimeraCoder/anaconda"
	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/util"
)

type twitter struct {
	twitter *anaconda.TwitterApi
}

func NewTwitter() Service {
	accessToken := os.Getenv("TWITTER_ACCESS_TOKEN")
	accessTokenScrt := os.Getenv("TWITTER_ACCESS_TOKEN_SECRET")
	consumerKey := os.Getenv("TWITTER_CONSUMER_KEY")
	consumerKeyScrt := os.Getenv("TWITTER_CONSUMER_SECRET")
	t := anaconda.NewTwitterApiWithCredentials(accessToken, accessTokenScrt, consumerKey, consumerKeyScrt)
	return &twitter{
		twitter: t,
	}
}
func (t *twitter) preprocessTwitterImage(image string) ([]byte, string, error) {
	imageFile := "temp.png"
	err := util.DownloadFile(image, imageFile)
	if err != nil {
		return nil, "", fmt.Errorf("[twitter.SendSalesTweet] cannot download image: %s", err)
	}
	defer deleteFile(imageFile)

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

	data, filePath, err := t.preprocessTwitterImage(message.Image)
	if err != nil {
		return fmt.Errorf("[twitter.SendSalesTweet] cannot process image: %s", err)
	}
	defer deleteFile(filePath)

	// upload image to twitter
	mediaResponse, err := t.twitter.UploadMedia(base64.StdEncoding.EncodeToString(data))
	if err != nil {
		return fmt.Errorf("[twitter.SendSalesTweet] cannot upload media to twitter: %s", err)
	}
	// set media to v
	v.Set("media_ids", strconv.FormatInt(mediaResponse.MediaID, 10))

	// **fields format ->
	// marketplace = {Paintswap, Opensea, Optimisim}
	// token name = Cyber Neko #2
	// buyer,seller = 0x23...1234
	// price = 123 FTM
	// edit post UI ->
	tweetStatus := fmt.Sprintf("A new sale has been made on %s for %s!\n\nBuyer: %s\n\nSeller: %s\n\nValue: %s\n\n→ Check collection at: %s",
		message.Marketplace, message.TokenName, message.BuyerAddress, message.SellerAddress, message.Price, message.MarketplaceURL)
	_, err = twitterApi.PostTweet(tweetStatus, v)
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
