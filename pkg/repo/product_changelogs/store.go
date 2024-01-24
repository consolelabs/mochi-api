package product_changelogs

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	List(q ListQuery) (changeLogs []model.ProductChangelogs, total int64, err error)
	Create(changelog *model.ProductChangelogs) error
	GetNewChangelog() ([]model.ProductChangelogs, error)
	InsertBulkProductChangelogSnapshot(changelogSnapshot []model.ProductChangelogSnapshot) error
	DeleteAll() error
	GetByVersion(version string) (*model.ProductChangelogs, error)
	GetNextVersion(id int64) (string, error)
	GetPreviousVersion(id int64) (string, error)
}
