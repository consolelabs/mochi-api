package response

type OnboardingStartResponse struct {
	Data OnboardingStartData `json:"data"`
}

type OnboardingStartData struct {
	UserAlreadyStarted bool                   `json:"user_already_started"`
	Reward             *OnboardingStartReward `json:"reward"`
}

type OnboardingStartReward struct {
	Token  string `json:"token"`
	Amount string `json:"amount"`
}
