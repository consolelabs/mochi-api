package message

import "github.com/consolelabs/mochi-typeset/typeset"

type OnboardingStart struct {
	Type                    typeset.NotificationType `json:"type"`
	OnboardingStartMetadata OnboardingStartMetadata  `json:"onboarding_start_metadata"`
}

type OnboardingStartMetadata struct {
	UserProfileID string `json:"user_profile_id"`
	Token         string `json:"token"`
	Amount        string `json:"amount"`
	Decimal       int64  `json:"decimal"`
}
