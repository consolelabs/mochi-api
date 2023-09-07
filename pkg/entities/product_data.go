package entities

import (
	"encoding/base64"
	"strings"
	"time"

	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/model"
	productbotcommand "github.com/defipod/mochi/pkg/repo/product_bot_command"
	productchangelogs "github.com/defipod/mochi/pkg/repo/product_changelogs"
	productchangelogsview "github.com/defipod/mochi/pkg/repo/product_changelogs_view"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/util"
)

func (e *Entity) ProductBotCommand(req request.ProductBotCommandRequest) ([]model.ProductBotCommand, error) {
	return e.repo.ProductBotCommand.List(productbotcommand.ListQuery{
		Code:  req.Code,
		Scope: req.Scope,
	})
}

func (e *Entity) ProductChangelogs(req request.ProductChangelogsRequest) ([]model.ProductChangelogs, error) {
	return e.repo.ProductChangelogs.List(productchangelogs.ListQuery{
		Product: req.Product,
		Size:    req.Size,
	})
}

func (e *Entity) CreateProductChangelogsView(req request.CreateProductChangelogsViewRequest) (*model.ProductChangelogView, error) {
	productchangelogsview := &model.ProductChangelogView{
		Key:           req.Key,
		ChangelogName: req.ChangelogName,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}
	return productchangelogsview, e.repo.ProductChangelogsView.Create(productchangelogsview)
}

func (e *Entity) GetProductChangelogsView(req request.GetProductChangelogsViewRequest) ([]model.ProductChangelogView, error) {
	return e.repo.ProductChangelogsView.List(productchangelogsview.ListQuery{
		Key:           req.Key,
		ChangelogName: req.ChangelogName,
	})
}

func (e *Entity) CrawlChangelogs() {
	repos := e.svc.Github.GetContents()
	err := e.repo.ProductChangelogs.DeleteAll()
	if err != nil {
		e.log.Error(err, "[entity.CrawlChangelogs()] - cannot delete all")
		return
	}
	for _, repo := range repos {
		// 1. validate file markdown
		if !util.ValidateFileMarkdown(repo.Name) {
			continue
		}

		// 2. get detail content of repo
		repoDetail, err := e.svc.Github.GetContentByPath(repo.URL)
		if err != nil || repoDetail == nil {
			e.log.Fields(logger.Fields{"title": repo.Name}).Error(err, "[entity.CrawlChangelogs()] - cannot get content of repo")
			continue
		}

		// 3. parse content from base64 to string
		rawDecodedText, err := base64.StdEncoding.DecodeString(repoDetail.Content)
		if err != nil {
			e.log.Fields(logger.Fields{"title": repo.Name}).Error(err, "[entity.CrawlChangelogs()] - cannot decode content of repo")
			continue
		}

		// 4. convert content string to model.ProductChangelogs
		changelogs := e.parseChangelogsContent(string(rawDecodedText))
		if changelogs == nil {
			e.log.Fields(logger.Fields{"title": repo.Name}).Error(err, "[entity.CrawlChangelogs()] - cannot parse content of repo")
			continue
		}
		changelogs.GithubUrl = repo.HTMLURL
		changelogs.FileName = repo.Name
		changelogs.IsExpired = false

		// 5. store changelogs
		err = e.repo.ProductChangelogs.Create(changelogs)
		if err != nil {
			e.log.Fields(logger.Fields{"title": repo.Name}).Error(err, "[entity.CrawlChangelogs()] - cannot store repo")
			continue
		}
	}
}

func (e *Entity) parseChangelogsContent(content string) *model.ProductChangelogs {
	var changlogs model.ProductChangelogs
	contentSplit := strings.Split(content, "---")
	if len(contentSplit) < 3 {
		return nil
	}

	// 1. get title, date, product, thumbnail url in contentSplit[1]
	table := strings.TrimSpace(contentSplit[1])
	rows := strings.Split(table, "\n")
	for _, row := range rows {
		cRow := strings.Split(row, ": ")
		if len(cRow) < 2 {
			continue
		}

		switch cRow[0] {
		case "date":
			dateString := strings.TrimSpace(cRow[1])
			date, err := time.Parse("02-01-2006", dateString)
			if err != nil {
				changlogs.CreatedAt = time.Now()
				changlogs.UpdatedAt = time.Now()
			}
			changlogs.CreatedAt = date
			changlogs.UpdatedAt = date
		case "title":
			changlogs.Title = strings.TrimSpace(cRow[1])
		case "product":
			changlogs.Product = strings.TrimSpace(cRow[1])
		case "thumbnai_url":
			changlogs.ThumbnailUrl = strings.TrimSpace(cRow[1])
		}
	}

	// 2. Get content changelogs
	changlogs.Content = strings.TrimSpace(contentSplit[2])

	return &changlogs
}
