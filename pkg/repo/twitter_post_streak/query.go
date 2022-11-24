package twitterpoststreak

type ListQuery struct {
	GuildID       string
	TwitterID     string
	TwitterHandle string
	Sort          string
	Limit         int
	Offset        int
}
