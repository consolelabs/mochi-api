package request

type OnboardingStartRequest struct {
	ProfileId string `json:"profile_id" binding:"required"`
	Platform  string `json:"platform" binding:"required,oneof=discord telegram"`
}
