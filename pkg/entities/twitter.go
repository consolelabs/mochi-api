package entities

import (
	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/request"
)

func (e *Entity) GetUnnotifiedSalesMessage(offset int, limit int) ([]model.TwitterSalesMessage, int64, error) {
	messages, total, err := e.repo.MochiNFTSales.GetUnnotified(offset, limit)
	if err != nil {
		e.log.Errorf(err, "[e.HandleMochiSalesMessage] failed to get mochi nft sales: %s", err)
		return nil, total, err
	}
	return messages, total, nil
}

func (e *Entity) DeleteSalesMessages(message model.TwitterSalesMessage) error {
	err := e.repo.MochiNFTSales.DeleteOne(message)
	if err != nil {
		e.log.Errorf(err, "[e.HandleMochiSalesMessage] failed to update mochi nft sales: %s", err)
		return err
	}
	return nil
}

func (e *Entity) CreateTwitterPost(req *request.TwitterPost) error {
	err := e.repo.TwitterPost.CreateOne(req)
	if err != nil {
		e.log.Errorf(err, "[e.CreateTwitterPost] failed to create twitter post: %s", err)
		return err
	}
	return nil
}
