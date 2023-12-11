package github

import (
	"fmt"
	"net/http"

	"github.com/defipod/mochi/pkg/util"
)

func (g *Github) fetchGithubApiData(url string, res interface{}) (int, error) {
	req := util.SendRequestQuery{
		URL:      url,
		Response: &res,
		Headers:  map[string]string{"Authorization": fmt.Sprintf("Bearer %s", g.config.GithubToken), "X-GitHub-Api-Version": "2022-11-28"},
	}

	statusCode, err := util.SendRequest(req)
	if err != nil {
		return http.StatusBadRequest, err
	}

	return statusCode, nil
}
