package entities

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/consolelabs/mochi-typeset/typeset"
	"github.com/defipod/mochi/pkg/kafka/message"
	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/response"
	"github.com/defipod/mochi/pkg/service/mochiprofile"
	"github.com/defipod/mochi/pkg/util"
)

func (e *Entity) formatVoteVaultMessage(req *request.CreateTreasurerSubmission, resp *response.CreateTreasurerSubmissionResponse, submitterProfile, changerProfile *mochiprofile.GetProfileResponse, vault *model.Vault, treasurerSubmissions []model.TreasurerSubmission, treasurerReq *model.TreasurerRequest) (*message.VaultVoteTreasurer, map[string]string) {
	daoVaultTotalTreasurer := make(map[string]string)
	daoVaultTotalTreasurerProposal := make(map[string]string)
	for _, treasurerSubmission := range treasurerSubmissions {
		treasurerProfileId, err := e.svc.MochiProfile.GetByDiscordID(treasurerSubmission.Submitter, true)
		if err != nil {
			continue
		}

		if treasurerSubmission.Status == "pending" {
			treasurerSubmission.Status = "waiting"
		}

		// if for vote no need, proposal need
		if treasurerSubmission.Submitter == req.Sumitter {
			daoVaultTotalTreasurerProposal[treasurerProfileId.ID] = treasurerSubmission.Status
		} else {
			daoVaultTotalTreasurerProposal[treasurerProfileId.ID] = treasurerSubmission.Status
			daoVaultTotalTreasurer[treasurerProfileId.ID] = treasurerSubmission.Status
		}
	}

	switch req.Type {
	case "add":
		return &message.VaultVoteTreasurer{
			Type: typeset.NOTIFICATION_VAULT_VOTE,
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
				MessageUrl:             treasurerReq.MessageUrl,
				DaoVaultTotalTreasurer: daoVaultTotalTreasurer,
				Action: message.VaultAction{
					Type: "change-treasurer",
					VaultChangeTreasurerActionMetadata: message.VaultChangeTreasurerActionMetadata{
						TreasurerProfileId: changerProfile.ID,
						TreasurerAction:    "add",
					},
				},
			},
		}, daoVaultTotalTreasurerProposal
	case "remove":
		return &message.VaultVoteTreasurer{
			Type: typeset.NOTIFICATION_VAULT_VOTE,
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
				MessageUrl:             treasurerReq.MessageUrl,
				DaoVaultTotalTreasurer: daoVaultTotalTreasurer,
				Action: message.VaultAction{
					Type: "change-treasurer",
					VaultChangeTreasurerActionMetadata: message.VaultChangeTreasurerActionMetadata{
						TreasurerProfileId: submitterProfile.ID,
						TreasurerAction:    "remove",
					},
				},
			},
		}, daoVaultTotalTreasurerProposal
	case "transfer":
		token, err := e.svc.MochiPay.GetToken(treasurerReq.Token, treasurerReq.Chain)
		if err != nil {
			return nil, daoVaultTotalTreasurerProposal
		}

		amountInNumber, _ := strconv.ParseFloat(treasurerReq.Amount, 64)

		return &message.VaultVoteTreasurer{
			Type: typeset.NOTIFICATION_VAULT_VOTE,
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
				MessageUrl:             treasurerReq.MessageUrl,
				DaoVaultTotalTreasurer: daoVaultTotalTreasurer,
				Action: message.VaultAction{
					Type: "transfer",
					VaultTransferActionMetadata: message.VaultTransferActionMetadata{
						TokenAmount:        util.FloatToString(treasurerReq.Amount, token.Decimal),
						TokenDecimal:       token.Decimal,
						TokenAmountInUsd:   fmt.Sprint(token.Price * amountInNumber),
						Token:              strings.ToUpper(treasurerReq.Token),
						RecipientProfileId: changerProfile.ID,
					},
				},
			},
		}, daoVaultTotalTreasurerProposal
	default:
		return &message.VaultVoteTreasurer{}, daoVaultTotalTreasurerProposal
	}
}
