package response

import (
	"time"

	"github.com/defipod/mochi/pkg/model"
)

type CreateTreasurerRequestResponse struct {
	Request              model.VaultRequest     `json:"request"`
	VaultTreasurer       []model.VaultTreasurer `json:"treasurer"`
	IsDecidedAndExecuted bool                   `json:"is_decided_and_executed"`
}

type CreateTreasurerSubmissionResponse struct {
	Submission       model.VaultSubmission   `json:"submission"`
	VoteResult       VoteResult              `json:"vote_result"`
	TotalSubmissions []model.VaultSubmission `json:"total_submissions"`
}

type VoteResult struct {
	TotalApprovedSubmission   int64   `json:"total_approved_submission"`
	TotalSubmission           int64   `json:"total_submission"`
	TotalRejectedSubmisison   int64   `json:"total_rejected_submisison"`
	AllowedRejectedSubmisison int64   `json:"allowed_rejected_submisison"`
	TotalVote                 int64   `json:"total_vote"`
	Threshold                 string  `json:"threshold"`
	ThresholdNumber           float64 `json:"threshold_number"`
	Percentage                string  `json:"percentage"`
	IsApproved                bool    `json:"is_approved"`
}

type VaultDetailResponse struct {
	WalletAddress       string                 `json:"wallet_address"`
	SolanaWalletAddress string                 `json:"solana_wallet_address"`
	CurrentRequest      []CurrentRequest       `json:"current_request"`
	Balance             []Balance              `json:"balance"`
	MyNft               []MyNft                `json:"my_nft"`
	EstimatedTotal      string                 `json:"estimated_total"`
	VaultTreasurer      []model.VaultTreasurer `json:"treasurer"`
	RecentTransaction   []VaultTransaction     `json:"recent_transaction"`
	Threshold           string                 `json:"threshold"`
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
	Token  AssetToken `json:"token"`
	Amount string     `json:"amount"`
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

type GetVaultsResponse struct {
	Data []model.Vault `json:"data"`
}

type GetVaultResponse struct {
	Data model.Vault `json:"data"`
}
