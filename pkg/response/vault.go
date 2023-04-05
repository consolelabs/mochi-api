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
	TotalApprovedSubmission   int64  `json:"total_approved_submission"`
	TotalSubmission           int64  `json:"total_submission"`
	TotalRejectedSubmisison   int64  `json:"total_rejected_submisison"`
	AllowedRejectedSubmisison int64  `json:"allowed_rejected_submisison"`
	TotalVote                 int64  `json:"total_vote"`
	Threshold                 string `json:"threshold"`
	Percentage                string `json:"percentage"`
	IsApproved                bool   `json:"is_approved"`
}
