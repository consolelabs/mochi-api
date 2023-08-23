package message

import "github.com/consolelabs/mochi-typeset/typeset"

type OnboardingStart struct {
	Type     typeset.NotificationType `json:"type"`
	Metadata OnboardingStartMetadata  `json:"metadata"`
}

type OnboardingStartMetadata struct {
	Token              string `json:"token"`
	TokenAmount        string `json:"token_amount"`
	TokenDecimal       int64  `json:"token_decimal"`
	RecipientProfileId string `json:"recipient_profile_id"`
}
