package product_changelogs

type ListQuery struct {
	Product string
	Size    int
	Page    int
}

type ProductChangelogSnapshotQuery struct {
	Filename string
	IsPublic bool
}
