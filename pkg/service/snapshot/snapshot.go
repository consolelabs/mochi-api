package snapshot

import (
	"context"
	"fmt"

	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/response"
	"github.com/machinebox/graphql"
)

type Snapshot struct {
	client *graphql.Client
	log    logger.Logger
}

func NewService(log logger.Logger) Service {
	return &Snapshot{
		client: graphql.NewClient("https://hub.snapshot.org/graphql"),
		log:    log,
	}
}
func (s *Snapshot) GetSnapshotProposal(id string) (*response.SnapshotProposalDataResponse, error) {
	req := graphql.NewRequest(fmt.Sprintf(`
	query {
		proposal(id:"%s") {
		  id
		  title
		  body
		  choices
		  start
		  end
		  snapshot
		  state
		  author
		  created
		  scores
		  space {
			id
			name
		  }
		}
	}
	`, id))
	ctx := context.Background()

	res := response.SnapshotProposalDataResponse{}
	if err := s.client.Run(ctx, req, &res); err != nil {
		s.log.Error(err, "[service.Snapshot] get proposal failed")
		return nil, err
	}
	return &res, nil
}

func (s *Snapshot) GetSpace(id string) (*response.SnapshotSpaceDataResponse, error) {
	req := graphql.NewRequest(fmt.Sprintf(`
    query {
        space (id: "%s") {
            id
			name
        }
    }
	`, id))

	// set any variables
	ctx := context.Background()

	res := response.SnapshotSpaceDataResponse{}
	if err := s.client.Run(ctx, req, &res); err != nil {
		s.log.Error(err, "[service.Snapshot] get space failed")
		return nil, err
	}
	return &res, nil
}
