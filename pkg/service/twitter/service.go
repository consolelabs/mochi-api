package twitter

import (
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/response"
)

type Service interface {
	SendSalesTweet(imageURL string, nft *request.HandleNftWebhookRequest, token *response.IndexerNFTToken) error
}
