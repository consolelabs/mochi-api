package product_changelogs

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	List(q ListQuery) (changeLogs []model.ProductChangelogs, total int64, err error)
	GetChangelogByFilename(title string) (p model.ProductChangelogs, err error)
	Create(changelog *model.ProductChangelogs) error
	GetNewChangelog() ([]model.ProductChangelogs, error)
	GetChangelogNotConfirmed() ([]model.ProductChangelogs, error)
	UpdateProductChangelogSnapshot(q ProductChangelogSnapshotQuery) error
	InsertBulkProductChangelogSnapshot(changelogSnapshot []model.ProductChangelogSnapshot) error
	DeleteAll() error
}
