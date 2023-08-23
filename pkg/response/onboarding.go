package response

type OnboardingStartResponse struct {
	Data OnboardingStartData `json:"data"`
}

type OnboardingReward struct {
	TokenSymbol string `json:"token_symbol"`
	Amount      int    `json:"amount"`
}

type OnboardingStartData struct {
	DidOnboarding bool             `json:"did_onboarding"`
	Reward        OnboardingReward `json:"reward"`
}
