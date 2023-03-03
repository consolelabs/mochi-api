package commonwealth

import "github.com/defipod/mochi/pkg/response"

type Service interface {
	GetCommunities(id string) (*response.ListCommonwealthCommunities, error)
	CheckCommunityExist(id string) bool
	GetThreads(communityId string) (*response.CommonwealthThreadResponse, error)
}
