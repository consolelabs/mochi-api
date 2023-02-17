package commonwealth

import "github.com/defipod/mochi/pkg/response"

type Service interface {
	CheckCommunityExist(id string) bool
	GetThreads(communityId string) (*response.CommonwealthThreadResponse, error)
}
