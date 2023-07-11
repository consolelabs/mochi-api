package userwalletwatchlistitem

type ListQuery struct {
	ProfileID string
	UserID    string
	IsOwner   *bool
	Address   string
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
