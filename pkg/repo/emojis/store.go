package emojis

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	ListEmojis(listCode []string) (model []*model.ProductMetadataEmojis, err error)
}
