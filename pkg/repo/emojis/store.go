package emojis

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	ListEmojis(q Query) (model []*model.ProductMetadataEmojis, total int64, err error)
	GetByCode(code string) (model *model.ProductMetadataEmojis, err error)
}
