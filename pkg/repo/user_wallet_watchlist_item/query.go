package userwalletwatchlistitem

type ListQuery struct {
	UserID  string
	IsOwner *bool
}

type GetOneQuery struct {
	UserID    string
	Query     string
	ForUpdate bool
}

type DeleteQuery struct {
	UserID  string
	Address string
	Alias   string
}
