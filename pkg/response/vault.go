package response

import "github.com/defipod/mochi/pkg/model"

type CreateTreasurerRequestResponse struct {
	Request   model.TreasurerRequest `json:"request"`
	Treasurer []model.Treasurer      `json:"treasurer"`
}

type CreateTreasurerSubmissionResponse struct {
	Submission model.TreasurerSubmission `json:"submission"`
	VoteResult VoteResult                `json:"vote_result"`
}

type VoteResult struct {
	TotalApprovedSubmission int64  `json:"total_approved_submission"`
	TotalSubmission         int64  `json:"total_submission"`
	Threshold               string `json:"threshold"`
	Percentage              string `json:"percentage"`
	IsApproved              bool   `json:"is_approved"`
}
