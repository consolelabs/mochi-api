package chain

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	List(discussionID *int64) ([]model.CommonwealthDiscussionSubscription, error)
	Create(sub *model.CommonwealthDiscussionSubscription) error
}
