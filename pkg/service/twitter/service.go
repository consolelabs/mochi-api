package twitter

import (
	"github.com/defipod/mochi/pkg/model"
)

type Service interface {
	SendSalesMessageToTwitter(message model.TwitterSalesMessage, twitter model.GuildConfigTwitterFeed) error
}
