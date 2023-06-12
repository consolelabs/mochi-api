package content

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	GetContentByType(contentType string) (content *model.Content, err error)
}
