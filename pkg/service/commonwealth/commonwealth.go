package commonwealth

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/defipod/mochi/pkg/response"
	"github.com/defipod/mochi/pkg/util"
)

type Commonwealth struct {
	baseUrl string
}

func NewService() Service {
	return &Commonwealth{
		baseUrl: "https://commonwealth.im/external/",
	}
}

func (c *Commonwealth) GetCommunities(id string) (*response.ListCommonwealthCommunities, error) {
	res := &response.ListCommonwealthCommunities{}
	url := c.baseUrl + "communities?community_id=" + id
	statusCode, err := util.FetchData(url, res)
	if err != nil || statusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get commnunities")
	}
	return res, nil
}

func (c *Commonwealth) CheckCommunityExist(id string) bool {
	url := c.baseUrl + "communities?community_id=" + id
	client := &http.Client{Timeout: time.Second * 30}
	req, _ := http.NewRequest("GET", url, &io.LimitedReader{})
	req.Header.Set("Content-Type", "application/json")
	res, err := client.Do(req)
	if err != nil || res.StatusCode != http.StatusOK {
		return false
	}
	defer res.Body.Close()
	return true
}

func (c *Commonwealth) GetThreads(communityId string) (*response.CommonwealthThreadResponse, error) {
	res := response.CommonwealthThreadResponse{}
	url := c.baseUrl + fmt.Sprintf("threads?community_id=%s&include_comments=false&count_only=false&limit=10", communityId)
	req := util.SendRequestQuery{
		URL:       url,
		ParseForm: &res,
		Headers:   map[string]string{"Content-Type": "application/json"},
	}
	statusCode, err := util.SendRequest(req)
	if err != nil || statusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get thread")
	}
	return &res, nil
}
