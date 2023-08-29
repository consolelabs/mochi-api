package github

import "github.com/defipod/mochi/pkg/response"

type Service interface {
	GetContents() []response.RepositoryContents
	GetContentByPath(path string) (*response.RepositoryContent, error)
}
