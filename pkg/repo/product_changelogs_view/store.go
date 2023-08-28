package product_changelogs_view

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	List(ListQuery) ([]model.ProductChangelogView, error)
	Create(changelogView *model.ProductChangelogView) error
}
