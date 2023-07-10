package userwalletwatchlistitem

type ListQuery struct {
	ProfileID string
	IsOwner   *bool
}

type GetOneQuery struct {
	ProfileID string
	Query     string
	ForUpdate bool
}

type DeleteQuery struct {
	ProfileID string
	Address   string
	Alias     string
}
