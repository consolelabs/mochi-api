package userwalletwatchlistitem

type GetOneQuery struct {
	UserID string
	Query  string
}

type DeleteQuery struct {
	UserID  string
	Address string
	Alias   string
}
