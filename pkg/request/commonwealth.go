package request

import (
	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/response"
)

type NewCommonwealthDiscussionRequest struct {
	ChannelID  string
	Community  model.CommonwealthLatestData
	Discussion response.CommonwealthDiscussion
}
