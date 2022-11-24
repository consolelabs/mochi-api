package twitterpost

import (
	"github.com/defipod/mochi/pkg/model"
)

type Store interface {
	CreateOne(*model.TwitterPost) error
}
