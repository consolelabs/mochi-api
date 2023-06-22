package entities

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/defipod/mochi/pkg/kafka/message"
	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/response"
	"github.com/defipod/mochi/pkg/service/mochiprofile"
)

func (e *Entity) formatVoteVaultMessage(req *request.CreateTreasurerSubmission, resp *response.CreateTreasurerSubmissionResponse, submitterProfile, changerProfile *mochiprofile.GetProfileResponse, vault *model.Vault, treasurerSubmissions []model.TreasurerSubmission, treasurerReq *model.TreasurerRequest) *message.VaultVoteTreasurer {
	daoVaultTotalTreasurer := make(map[string]string)
	for _, treasurerSubmission := range treasurerSubmissions {
		if treasurerSubmission.Submitter == req.Sumitter {
			continue
		}

		treasurerProfileId, err := e.svc.MochiProfile.GetByDiscordID(treasurerSubmission.Submitter, true)
		if err != nil {
			continue
		}
		if treasurerSubmission.Status == "pending" {
			treasurerSubmission.Status = "waiting"
		}
		daoVaultTotalTreasurer[treasurerProfileId.ID] = treasurerSubmission.Status
	}
	switch req.Type {
	case "add":
		return &message.VaultVoteTreasurer{
			Type: "vault-vote",
			VaultVoteMetadata: message.VaultVoteMetadata{
				TreasurerProfileId:       submitterProfile.ID,
				TreasurerVote:            req.Choice,
				RequestId:                req.RequestId,
				DaoThresholdInPercentage: resp.VoteResult.ThresholdNumber,
				DaoThresholdInNumber:     resp.VoteResult.TotalSubmission - resp.VoteResult.AllowedRejectedSubmisison,
				CurrentApproval:          resp.VoteResult.TotalApprovedSubmission,
				CurrentRejection:         resp.VoteResult.TotalRejectedSubmisison,
				CurrentWaiting:           resp.VoteResult.TotalSubmission - resp.VoteResult.TotalApprovedSubmission - resp.VoteResult.TotalRejectedSubmisison,
				DaoGuild:                 vault.DiscordGuild.Name,
				DaoVault:                 vault.Name,
				// Message:                  "New treasurer request",
				// MessageUrl: "https://mochi.defipod.com/vaults/" + vault.Name + "/treasurers/" + submitterProfile.ID,
				DaoVaultTotalTreasurer: daoVaultTotalTreasurer,
				Action: message.VaultAction{
					Type: "change-treasurer",
					VaultChangeTreasurerActionMetadata: message.VaultChangeTreasurerActionMetadata{
						TreasurerProfileId: changerProfile.ID,
						TreasurerAction:    "add",
					},
				},
			},
		}
	case "remove":
		return &message.VaultVoteTreasurer{
			Type: "vault-vote",
			VaultVoteMetadata: message.VaultVoteMetadata{
				TreasurerProfileId:       submitterProfile.ID,
				TreasurerVote:            req.Choice,
				RequestId:                req.RequestId,
				DaoThresholdInPercentage: resp.VoteResult.ThresholdNumber,
				DaoThresholdInNumber:     resp.VoteResult.TotalSubmission - resp.VoteResult.AllowedRejectedSubmisison,
				CurrentApproval:          resp.VoteResult.TotalApprovedSubmission,
				CurrentRejection:         resp.VoteResult.TotalRejectedSubmisison,
				CurrentWaiting:           resp.VoteResult.TotalSubmission - resp.VoteResult.TotalApprovedSubmission - resp.VoteResult.TotalRejectedSubmisison,
				DaoGuild:                 vault.DiscordGuild.Name,
				DaoVault:                 vault.Name,
				// Message:                  "New treasurer request",
				// MessageUrl: "https://mochi.defipod.com/vaults/" + vault.Name + "/treasurers/" + submitterProfile.ID,
				DaoVaultTotalTreasurer: daoVaultTotalTreasurer,
				Action: message.VaultAction{
					Type: "change-treasurer",
					VaultChangeTreasurerActionMetadata: message.VaultChangeTreasurerActionMetadata{
						TreasurerProfileId: submitterProfile.ID,
						TreasurerAction:    "remove",
					},
				},
			},
		}
	case "transfer":
		token, err := e.svc.MochiPay.GetToken(treasurerReq.Token, treasurerReq.Chain)
		if err != nil {
			return nil
		}

		amountInNumber, _ := strconv.ParseFloat(treasurerReq.Amount, 64)

		return &message.VaultVoteTreasurer{
			Type: "vault-vote",
			VaultVoteMetadata: message.VaultVoteMetadata{
				TreasurerProfileId:       submitterProfile.ID,
				TreasurerVote:            req.Choice,
				RequestId:                req.RequestId,
				DaoThresholdInPercentage: resp.VoteResult.ThresholdNumber,
				DaoThresholdInNumber:     resp.VoteResult.TotalSubmission - resp.VoteResult.AllowedRejectedSubmisison,
				CurrentApproval:          resp.VoteResult.TotalApprovedSubmission,
				CurrentRejection:         resp.VoteResult.TotalRejectedSubmisison,
				CurrentWaiting:           resp.VoteResult.TotalSubmission - resp.VoteResult.TotalApprovedSubmission - resp.VoteResult.TotalRejectedSubmisison,
				DaoGuild:                 vault.DiscordGuild.Name,
				DaoVault:                 vault.Name,
				// Message:                  "New treasurer request",
				// MessageUrl: "https://mochi.defipod.com/vaults/" + vault.Name + "/treasurers/" + submitterProfile.ID,
				DaoVaultTotalTreasurer: daoVaultTotalTreasurer,
				Action: message.VaultAction{
					Type: "transfer",
					VaultTransferActionMetadata: message.VaultTransferActionMetadata{
						TokenAmount:        treasurerReq.Amount,
						TokenDecimal:       token.Decimal,
						TokenAmountInUsd:   fmt.Sprint(token.Price * amountInNumber),
						Token:              strings.ToUpper(treasurerReq.Token),
						RecipientProfileId: changerProfile.ID,
					},
				},
			},
		}
	default:
		return &message.VaultVoteTreasurer{}
	}
}
