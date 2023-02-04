package snapshot

import "github.com/defipod/mochi/pkg/response"

type Service interface {
	GetSnapshotProposal(id string) (*response.SnapshotProposalDataResponse, error)
	GetSpace(id string) (*response.SnapshotSpaceDataResponse, error)
}
