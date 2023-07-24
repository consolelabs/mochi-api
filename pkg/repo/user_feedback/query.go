package userfeedback

type FeedbackQuery struct {
	ProfileID string
	DiscordId string
	Command   string
	Status    string
	Sort      string
	Limit     int64
	Offset    int64
}
