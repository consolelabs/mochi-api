package producthashtag

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	GetByAlias(alias string) (p *model.ProductHashtag, err error)
}
