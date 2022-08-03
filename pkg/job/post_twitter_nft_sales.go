package job

import (
	"github.com/defipod/mochi/pkg/entities"
	"github.com/defipod/mochi/pkg/logger"
)

type postTwitterNftSales struct {
	entity *entities.Entity
	log    logger.Logger
}

func NewPostTwitterNftSales(e *entities.Entity, l logger.Logger) Job {
	return &postTwitterNftSales{
		entity: e,
		log:    l,
	}
}

func (c *postTwitterNftSales) Run() error {
	unnotifiedMessages, err := c.entity.GetUnnotifiedSalesMessage()
	if err != nil {
		c.log.Error(err, "failed to get sales messages")
		return err
	}
	twitterConfigs, err := c.entity.GetAllTwitterConfig()
	if err != nil {
		c.log.Error(err, "failed to get twitter configs")
		return err
	}
	for _, t := range twitterConfigs {
		for _, message := range unnotifiedMessages {
			err := c.entity.GetSvc().Twitter.SendSalesMessageToTwitter(&message, &t)
			if err != nil {
				c.log.Error(err, "failed to get send tweet")
				return err
			}
			err = c.entity.DeleteSalesMessages(&message)
			if err != nil {
				c.log.Error(err, "failed to delete sales message")
				return err
			}
		}
	}
	return nil
}
