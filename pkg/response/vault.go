package response

import (
	"time"

	"github.com/defipod/mochi/pkg/model"
)

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

type VaultDetailResponse struct {
	WalletAddress     string             `json:"wallet_address"`
	CurrentRequest    []CurrentRequest   `json:"current_request"`
	Balance           []Balance          `json:"balance"`
	MyNft             []MyNft            `json:"my_nft"`
	EstimatedTotal    string             `json:"estimated_total"`
	Treasurer         []model.Treasurer  `json:"treasurer"`
	RecentTransaction []VaultTransaction `json:"recent_transaction"`
}

type VaultTransaction struct {
	Action    string    `json:"action"`
	Target    string    `json:"target"`
	Date      time.Time `json:"date"`
	Amount    string    `json:"amount"`
	Token     string    `json:"token"`
	Threshold string    `json:"threshold"`
	ToAddress string    `json:"to_address"`
}

type CurrentRequest struct {
	TotalApprovedSubmission int64     `json:"total_approved_submission"`
	TotalSubmission         int64     `json:"total_submission"`
	Action                  string    `json:"action"`
	Target                  string    `json:"target"`
	ExpiredDate             time.Time `json:"expired_date"`
	Amount                  string    `json:"amount"`
	Token                   string    `json:"token"`
	Address                 string    `json:"address"`
}

type Balance struct {
	TokenName   string `json:"token_name"`
	Token       string `json:"token"`
	Amount      string `json:"amount"`
	AmountInUsd string `json:"amount_in_usd"`
}

type MyNft struct {
	CollectionName  string `json:"collection_name"`
	CollectionImage string `json:"collection_image"`
	Chain           string `json:"chain"`
	Total           int64  `json:"total"`
	Nft             []Nft  `json:"nft"`
}

type Nft struct {
	Name string `json:"name"`
	Id   string `json:"id"`
}
