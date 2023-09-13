package friendtech

import "github.com/defipod/mochi/pkg/response"

type Service interface {
	Search(query string, limit int) (*response.FriendTechKeysResponse, error)
}
