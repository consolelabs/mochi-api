package product_changelogs

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	List(ListQuery) ([]model.ProductChangelogs, error)
}
