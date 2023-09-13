package product_changelogs

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	List(ListQuery) ([]model.ProductChangelogs, error)
	Create(changelog *model.ProductChangelogs) error
	GetNewChangelog() ([]model.ProductChangelogs, error)
	InsertBulkProductChangelogSnapshot(changelogSnapshot []model.ProductChangelogSnapshot) error
	DeleteAll() error
}
