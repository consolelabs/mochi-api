package github

import (
	"fmt"
	"github.com/defipod/mochi/pkg/config"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/response"
)

var (
	githubApi = "https://api.github.com"
	owner     = "consolelabs"
	repo      = "changelog"
)

type Github struct {
	config *config.Config
	logger logger.Logger
}

func NewService(cfg *config.Config, l logger.Logger) Service {
	return &Github{
		config: cfg,
		logger: l,
	}
}

func (g *Github) GetContents() (res []response.RepositoryContents) {
	url := fmt.Sprintf("%s/repos/%s/%s/contents", githubApi, owner, repo)
	code, err := g.fetchGithubApiData(url, &res)
	if err != nil {
		g.logger.Fields(logger.Fields{"endpoint": url, "code": code}).Error(err, "[github.GetContents] util.FetchData() failed")
		return []response.RepositoryContents{}
	}
	return res
}

func (g *Github) GetContentByPath(path string) (*response.RepositoryContent, error) {
	var res *response.RepositoryContent
	code, err := g.fetchGithubApiData(path, &res)
	if err != nil {
		g.logger.Fields(logger.Fields{"endpoint": path, "code": code}).Error(err, "[github.GetContentDetail] util.FetchData() failed")
		return nil, err
	}
	return res, nil
}
