package job

import (
	"github.com/defipod/mochi/pkg/entities"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/model"
	"github.com/schollz/progressbar/v3"
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
	var unnotifiedMessages []model.TwitterSalesMessage
	var bar *progressbar.ProgressBar
	var total int64
	offset := 0
	limit := 100000

	twitterConfigs, err := c.entity.GetAllTwitterConfig()
	if err != nil {
		c.log.Error(err, "failed to get twitter configs")
		return err
	}

	for {
		unnotifiedMessages, total, err = c.entity.GetUnnotifiedSalesMessage(offset, limit)
		if err != nil {
			c.log.Error(err, "failed to get sales messages")
			return err
		}

		// show progress on terminal
		if bar == nil {
			bar = progressbar.Default(total * int64(len(twitterConfigs)))
		}

		for _, t := range twitterConfigs {
			for _, message := range unnotifiedMessages {
				bar.Add(1)
				err := c.entity.GetSvc().Twitter.SendSalesMessageToTwitter(message, t)
				if err != nil {
					c.log.Error(err, "failed to get send tweet")
					return err
				}
				err = c.entity.DeleteSalesMessages(message)
				if err != nil {
					c.log.Error(err, "failed to delete sales message")
					return err
				}
			}
		}

		offset += limit
		if offset >= int(total) {
			break
		}
	}
	return nil
}
