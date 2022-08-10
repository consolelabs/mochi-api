package twitterpost

import (
	"github.com/defipod/mochi/pkg/request"
)

type Store interface {
	CreateOne(*request.TwitterPost) error
}
