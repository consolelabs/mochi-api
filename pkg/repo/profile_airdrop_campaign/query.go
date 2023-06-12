package profileairdropcampaign

type ListQuery struct {
	ProfileId  string
	Status     string
	IsFavorite *bool
	Offset     int
	Limit      int
}
